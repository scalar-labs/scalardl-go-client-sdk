package rpc

import (
	"testing"

	"github.com/scalar-labs/dl/v3/crypto"
)

const testKey = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`

func TestContractRegistrationRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request ContractRegistrationRequest = ContractRegistrationRequest{
			ContractId:         "TestContract",
			ContractBinaryName: "com.example.TestContract",
			ContractByteCode:   []byte{0xCA, 0xFE},
			ContractProperties: "{}",
			CertHolderId:       "tester",
			CertVersion:        1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestContractsListingRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request ContractsListingRequest = ContractsListingRequest{
			ContractId:   "TestContract",
			CertHolderId: "tester",
			CertVersion:  1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestContractExecutionRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request ContractExecutionRequest = ContractExecutionRequest{
			ContractId:       "TestContract",
			ContractArgument: `{"foo":"bar"}`,
			CertHolderId:     "tester",
			CertVersion:      1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestLedgerValidationRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request LedgerValidationRequest = LedgerValidationRequest{
			AssetId:      "foo",
			StartAge:     0,
			EndAge:       99,
			CertHolderId: "tester",
			CertVersion:  1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestLedgersValidationRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request LedgersValidationRequest = LedgersValidationRequest{
			AssetId:      "foo",
			CertHolderId: "tester",
			CertVersion:  1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestAssetProofRetrievalRequest_SignWith(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request AssetProofRetrievalRequest = AssetProofRetrievalRequest{
			AssetId:      "foo",
			Age:          0,
			CertHolderId: "tester",
			CertVersion:  1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}

func TestExecutionAbortRequest_SignWithSigner(t *testing.T) {
	var (
		signer  crypto.Signer
		err     error
		request ExecutionAbortRequest = ExecutionAbortRequest{
			Nonce:        "nonce",
			CertHolderId: "tester",
			CertVersion:  1,
		}
	)

	if signer, err = crypto.NewEcdsaSha256Signer([]byte(testKey)); err != nil {
		t.Errorf("should get a Signer")
	}

	if err = request.SignWith(signer); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(request.Signature) == 0 {
		t.Errorf("signature should be filled")
	}
}
