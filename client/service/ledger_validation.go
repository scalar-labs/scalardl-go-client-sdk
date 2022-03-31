package service

import (
	"bytes"
	"context"
	"fmt"

	clientError "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/error"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/crypto"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/json"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/asset"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/model"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// JavaMaxIntValue defines a value according to Java's max value of int.
const JavaMaxIntValue = 2147483647

// ValidateLedger validates the specified asset between the specified ages.
func (s ClientService) ValidateLedger(args ...interface{}) (result model.LedgerValidationResult, err error) {
	if s.clientConfig.ClientMode != "CLIENT" {
		return result, clientError.NewClientError(statuscode.InvalidRequest, "wrong mode specified")
	}

	var (
		assetID  string
		startAge int = 0
		endAge   int = JavaMaxIntValue
		signer   crypto.Signer
		ok       bool
	)

	if len(args) > 0 {
		if assetID, ok = args[0].(string); !ok {
			return result, fmt.Errorf("assetID must be a string")
		}
	}

	if len(args) > 1 {
		if startAge, ok = args[1].(int); !ok {
			return result, fmt.Errorf("startAge must be an integer")
		}
	}

	if len(args) > 2 {
		if endAge, ok = args[2].(int); !ok {
			return result, fmt.Errorf("endAge must be an integer")
		}
	}

	if assetID == "" {
		return result, fmt.Errorf("assetID cannot be empty")
	}

	if endAge < startAge || startAge < 0 || endAge > JavaMaxIntValue {
		return result, fmt.Errorf("invalid ages specified")
	}

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(s.clientConfig.PrivateKey)); err != nil {
		return
	}

	if s.clientConfig.IsAuditorEnabled && s.clientConfig.IsAuditorLinearizableValidationEnabled {
		argument := json.Object{
			"asset_id": assetID,
		}

		argument["start_age"] = startAge
		argument["end_age"] = endAge

		var executed model.ContractExecutionResult

		if executed, err = s.ExecuteContract(
			s.clientConfig.AuditorLinearizableValidationContractID,
			argument,
			nil,
		); err != nil {
			return
		}

		result.Code = statuscode.OK

		if len(executed.Proofs) > 0 {
			result.Proof = executed.Proofs[0]
		}

		if len(executed.AuditorProofs) > 0 {
			result.AuditorProof = executed.AuditorProofs[0]
		}
	} else {
		var request *rpc.LedgerValidationRequest = &rpc.LedgerValidationRequest{
			AssetId:      assetID,
			StartAge:     uint32(startAge),
			EndAge:       uint32(endAge),
			CertHolderId: s.clientConfig.CertHolderID,
			CertVersion:  uint32(s.clientConfig.CertVersion),
		}

		if err = request.SignWith(signer); err != nil {
			return
		}

		var (
			ledgerChan          chan *rpc.LedgerValidationResponse = make(chan *rpc.LedgerValidationResponse)
			auditorChan         chan *rpc.LedgerValidationResponse = make(chan *rpc.LedgerValidationResponse)
			ledgerErrorChan     chan error                         = make(chan error)
			auditorErrorChan    chan error                         = make(chan error)
			responseFromLedger  *rpc.LedgerValidationResponse
			responseFromAuditor *rpc.LedgerValidationResponse
		)

		defer close(ledgerChan)
		defer close(auditorChan)
		defer close(ledgerErrorChan)
		defer close(auditorErrorChan)

		go func(r chan *rpc.LedgerValidationResponse, e chan error) {
			if !s.clientConfig.IsAuditorEnabled {
				r <- nil
				e <- nil
			} else {
				auditor := rpc.NewAuditorClient(s.auditorConnection)
				trailer := metadata.MD{}

				response, err := auditor.ValidateLedger(context.Background(), request, grpc.Trailer(&trailer))

				if err != nil && trailer.Len() > 0 {
					err = getClientErrorFromTrailer(trailer)
				}

				r <- response
				e <- err
			}
		}(auditorChan, auditorErrorChan)

		go func(r chan *rpc.LedgerValidationResponse, e chan error) {
			ledger := rpc.NewLedgerClient(s.ledgerConnection)
			trailer := metadata.MD{}

			response, err := ledger.ValidateLedger(context.Background(), request, grpc.Trailer(&trailer))

			if err != nil && trailer.Len() > 0 {
				err = getClientErrorFromTrailer(trailer)
			}

			r <- response
			e <- err
		}(ledgerChan, ledgerErrorChan)

		var errFromLedger error
		responseFromLedger, errFromLedger = <-ledgerChan, <-ledgerErrorChan

		var errFromAuditor error
		responseFromAuditor, err = <-auditorChan, <-auditorErrorChan

		if errFromLedger != nil {
			err = errFromLedger
			return
		}

		if errFromAuditor != nil {
			err = errFromAuditor
			return
		}

		if !s.clientConfig.IsAuditorEnabled {
			result.Code = statuscode.StatusCode(responseFromLedger.StatusCode)
			result.Proof = asset.FromGRPC(responseFromLedger.Proof)
		} else {
			var (
				p1   asset.Proof           = asset.FromGRPC(responseFromLedger.Proof)
				p2   asset.Proof           = asset.FromGRPC(responseFromAuditor.Proof)
				code statuscode.StatusCode = statuscode.InconsistentStates
			)

			if responseFromLedger.StatusCode == statuscode.OK &&
				responseFromAuditor.StatusCode == statuscode.OK &&
				!p1.Equal(asset.Proof{}) &&
				!p2.Equal(asset.Proof{}) &&
				bytes.Equal(p1.Hash, p2.Hash) {
				code = statuscode.OK
			}

			result.Code = code
			result.Proof = p1
			result.AuditorProof = p2
		}
	}

	return
}
