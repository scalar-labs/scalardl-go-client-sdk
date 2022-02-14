package service

import (
	"crypto/x509"
	"fmt"

	"github.com/scalar-labs/scalardl-go-client-sdk/v3/client/config"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ClientService defines the interface of the client service.
type ClientService struct {
	clientConfig                config.ClientConfig
	signer                      crypto.Signer
	ledgerConnection            *grpc.ClientConn
	ledgerPrivilegedConnection  *grpc.ClientConn
	auditorConnection           *grpc.ClientConn
	auditorPrivilegedConnection *grpc.ClientConn
}

// NewClientService creates ClientService instance.
func NewClientService(c config.ClientConfig) (s ClientService, err error) {
	if err = c.Validate(); err != nil {
		return
	}

	s.clientConfig = c

	if s.signer, err = crypto.NewEcdsaSha256Signer([]byte(c.PrivateKey)); err != nil {
		return
	}

	var opts = make([]grpc.DialOption, 0)

	if c.IsTLSEnabled {
		var certPool *x509.CertPool = x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM([]byte(c.TLSCaRootCert)); !ok {
			err = fmt.Errorf("TLSCaRootCert is not valid")
			return
		}

		var creds credentials.TransportCredentials = credentials.NewClientTLSFromCert(certPool, c.LedgerHost)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	if s.ledgerConnection, err = grpc.Dial(
		fmt.Sprintf("%s:%d", c.LedgerHost, c.LedgerPort),
		opts...,
	); err != nil {
		return
	}

	if s.ledgerPrivilegedConnection, err = grpc.Dial(
		fmt.Sprintf("%s:%d", c.LedgerHost, c.LedgerPrivilegedPort),
		opts...,
	); err != nil {
		return
	}

	if c.IsAuditorEnabled {
		opts = make([]grpc.DialOption, 0)

		if c.IsAuditorTLSEnabled {
			var certPool *x509.CertPool = x509.NewCertPool()
			if ok := certPool.AppendCertsFromPEM([]byte(c.AuditorTLSCaRootCert)); !ok {
				err = fmt.Errorf("AuditorTLSCaRootCert is not valid")
				return
			}

			var creds credentials.TransportCredentials = credentials.NewClientTLSFromCert(certPool, c.AuditorHost)
			opts = append(opts, grpc.WithTransportCredentials(creds))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		if s.auditorConnection, err = grpc.Dial(
			fmt.Sprintf("%s:%d", c.AuditorHost, c.AuditorPort),
			opts...,
		); err != nil {
			return
		}

		if s.auditorPrivilegedConnection, err = grpc.Dial(
			fmt.Sprintf("%s:%d", c.AuditorHost, c.AuditorPrivilegedPort),
			opts...,
		); err != nil {
			return
		}
	}

	return
}

// Close shuts down underlying connections.
func (s ClientService) Close() {
	if s.ledgerConnection != nil {
		s.ledgerConnection.Close()
	}

	if s.ledgerPrivilegedConnection != nil {
		s.ledgerPrivilegedConnection.Close()
	}

	if s.auditorConnection != nil {
		s.auditorConnection.Close()
	}

	if s.auditorPrivilegedConnection != nil {
		s.auditorPrivilegedConnection.Close()
	}
}
