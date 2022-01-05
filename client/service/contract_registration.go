package service

import (
	"context"
	"fmt"

	"github.com/scalar-labs/dl"
	clientError "github.com/scalar-labs/dl/client/error"
	"github.com/scalar-labs/dl/crypto"
	"github.com/scalar-labs/dl/ledger/statuscode"
	"github.com/scalar-labs/dl/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterContract registers contract to Scalar DL networks.
func (s ClientService) RegisterContract(
	id string,
	name string,
	contractBytes []byte,
	properties dl.JSONObject,
) (err error) {
	if s.clientConfig.ClientMode != "CLIENT" {
		return clientError.NewClientError(statuscode.InvalidRequest, "wrong mode specified")
	}

	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if contractBytes == nil {
		return fmt.Errorf("contractBytes cannot be nil")
	}

	var (
		signer  crypto.Signer
		trailer metadata.MD
		request = &rpc.ContractRegistrationRequest{
			ContractId:         id,
			ContractBinaryName: name,
			ContractByteCode:   contractBytes,
			CertHolderId:       s.clientConfig.CertHolderID,
			CertVersion:        uint32(s.clientConfig.CertVersion),
		}
	)

	if properties != nil {
		request.ContractProperties = properties.String()
	}

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(s.clientConfig.PrivateKey)); err != nil {
		return
	}

	if err = request.SignWith(signer); err != nil {
		return
	}

	if s.clientConfig.IsAuditorEnabled {
		var auditor = rpc.NewAuditorClient(s.auditorConnection)
		if _, err := auditor.RegisterContract(context.Background(), request, grpc.Trailer(&trailer)); err != nil {
			if trailer.Len() > 0 {
				err = getClientErrorFromTrailer(trailer)
			}

			return err
		}
	}

	trailer = metadata.MD{}
	var ledger = rpc.NewLedgerClient(s.ledgerConnection)
	if _, err := ledger.RegisterContract(context.Background(), request, grpc.Trailer(&trailer)); err != nil {
		if trailer.Len() > 0 {
			err = getClientErrorFromTrailer(trailer)
		}

		return err
	}

	return
}
