package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/google/uuid"
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

// ExecuteContract executes a registered contract.
func (s ClientService) ExecuteContract(
	id string,
	argument json.Object,
	functionArgument json.Object,
) (result model.ContractExecutionResult, err error) {
	if s.clientConfig.ClientMode != "CLIENT" {
		return result, clientError.NewClientError(statuscode.InvalidRequest, "wrong mode specified")
	}

	if id == "" {
		return result, fmt.Errorf("id cannot be empty")
	}

	if argument == nil {
		return result, fmt.Errorf("argument cannot be nil")
	}

	if _, ok := argument["nonce"]; !ok {
		argument["nonce"] = uuid.NewString()
	}

	if argument["nonce"] == "" {
		argument["nonce"] = uuid.NewString()
	}

	var (
		signer  crypto.Signer
		request = &rpc.ContractExecutionRequest{
			ContractId:       id,
			CertHolderId:     s.clientConfig.CertHolderID,
			CertVersion:      uint32(s.clientConfig.CertVersion),
			ContractArgument: argument.String(),
		}
	)

	if functionArgument != nil {
		request.FunctionArgument = functionArgument.String()
	}

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(s.clientConfig.PrivateKey)); err != nil {
		return
	}

	if err = request.SignWith(signer); err != nil {
		return
	}

	var (
		auditor             rpc.AuditorClient
		ledger              = rpc.NewLedgerClient(s.ledgerConnection)
		trailer             metadata.MD
		responseFromLedger  *rpc.ContractExecutionResponse
		responseFromAuditor *rpc.ContractExecutionResponse
	)

	// send the ordering request to Auditor first to get Auditor's signature.
	if s.clientConfig.IsAuditorEnabled {
		auditor = rpc.NewAuditorClient(s.auditorConnection)
		trailer = metadata.MD{}

		var ordered *rpc.ExecutionOrderingResponse
		if ordered, err = auditor.OrderExecution(context.Background(), request, grpc.Trailer(&trailer)); err != nil {
			if trailer.Len() > 0 {
				err = getClientErrorFromTrailer(trailer)
			}

			return
		}

		request.AuditorSignature = ordered.GetSignature()
	}

	trailer = metadata.MD{}
	if responseFromLedger, err = ledger.ExecuteContract(context.Background(), request, grpc.Trailer(&trailer)); err != nil {
		if trailer.Len() > 0 {
			err = getClientErrorFromTrailer(trailer)
		}

		return
	}

	// send the execution validation reqeust to Auditor.k
	if s.clientConfig.IsAuditorEnabled {
		trailer = metadata.MD{}

		if responseFromAuditor, err = auditor.ValidateExecution(
			context.Background(),
			&rpc.ExecutionValidationRequest{
				Request: request,
				Proofs:  responseFromLedger.GetProofs(),
			},
			grpc.Trailer(&trailer),
		); err != nil {
			if trailer.Len() > 0 {
				err = getClientErrorFromTrailer(trailer)
			}

			return
		}

		if responseFromLedger.GetResult() != responseFromAuditor.GetResult() ||
			len(responseFromLedger.GetProofs()) != len(responseFromAuditor.GetProofs()) {
			return result, clientError.NewClientError(
				statuscode.InconsistentStates,
				"The results from Ledger and Auditor don't match",
			)
		}

		var proofsFromLedger map[string]*rpc.AssetProof = make(map[string]*rpc.AssetProof)

		for _, p1 := range responseFromLedger.GetProofs() {
			proofsFromLedger[p1.GetAssetId()] = p1
		}

		for _, p2 := range responseFromAuditor.GetProofs() {
			p1, ok := proofsFromLedger[p2.GetAssetId()]

			if !ok || p1.GetAge() != p2.GetAge() || !bytes.Equal(p1.GetHash(), p2.GetHash()) {
				return result, clientError.NewClientError(
					statuscode.InconsistentStates,
					"The results from Ledger and Auditor don't match",
				)
			}
		}

		for _, p := range responseFromAuditor.GetProofs() {
			result.AuditorProofs = append(result.AuditorProofs, asset.FromGRPC(p))
		}
	}

	result.Result, _ = json.FromJSON(responseFromLedger.Result)

	for _, p := range responseFromLedger.GetProofs() {
		result.Proofs = append(result.Proofs, asset.FromGRPC(p))
	}

	return
}
