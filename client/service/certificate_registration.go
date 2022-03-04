package service

import (
	"context"

	clientError "github.com/scalar-labs/scalardl-go-client-sdk/v3/client/error"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RegisterCertificate registers the certificate in the client config to Ledger and Auditor.
func (s ClientService) RegisterCertificate() (err error) {
	if s.clientConfig.ClientMode != "CLIENT" {
		return clientError.NewClientError(
			statuscode.InvalidRequest,
			"wrong mode specified",
		)
	}

	var (
		trailer metadata.MD
		request = &rpc.CertificateRegistrationRequest{
			CertHolderId: s.clientConfig.CertHolderID,
			CertVersion:  (uint32)(s.clientConfig.CertVersion),
			CertPem:      s.clientConfig.Cert,
		}
	)

	if s.clientConfig.IsAuditorEnabled {
		var privileged = rpc.NewAuditorPrivilegedClient(s.auditorPrivilegedConnection)

		if _, err = privileged.RegisterCert(
			context.Background(),
			request,
			grpc.Trailer(&trailer),
		); err != nil {
			if trailer.Len() > 0 {
				err = getClientErrorFromTrailer(trailer)
			}
		}
	}

	var priviledged = rpc.NewLedgerPrivilegedClient(s.ledgerPrivilegedConnection)

	trailer = metadata.MD{}

	if _, err = priviledged.RegisterCert(
		context.Background(),
		request,
		grpc.Trailer(&trailer),
	); err != nil {
		if trailer.Len() > 0 {
			err = getClientErrorFromTrailer(trailer)
		}
	}

	return
}
