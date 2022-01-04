package crypto

import "testing"

func TestNewEcdsaSha256Signer(t *testing.T) {
	if _, err := NewEcdsaSha256Signer([]byte("not a key")); err == nil {
		t.Errorf("should get an error")
	}

	var key string = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`

	if _, err := NewEcdsaSha256Signer([]byte(key)); err != nil {
		t.Errorf("should get an correct signer instance")
	}
}

func TestEcdsaSha256Signer_Sign(t *testing.T) {
	var (
		privateKey string = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`
		s         Signer
		err       error
		signature []byte
	)

	if s, err = NewEcdsaSha256Signer([]byte(privateKey)); err != nil {
		t.Errorf("should get an correct signer instance")
	}

	if signature, err = s.Sign([]byte("hello world!")); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(signature) < 2 || len(signature) < (int(signature[1])-2) {
		t.Errorf("signature is in an incorrect length")
	}

	if int(signature[0]) != 48 {
		t.Errorf("signature has an incorrect first-byte")
	}
}

func TestEcdsaSha256Signer_Verify(t *testing.T) {
	var (
		privateKey string = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`
		publicKey string = `-----BEGIN CERTIFICATE-----
MIICizCCAjKgAwIBAgIUMEUDTdWsQpftFkqs6bCd6U++4nEwCgYIKoZIzj0EAwIw
bzELMAkGA1UEBhMCSlAxDjAMBgNVBAgTBVRva3lvMQ4wDAYDVQQHEwVUb2t5bzEf
MB0GA1UEChMWU2FtcGxlIEludGVybWVkaWF0ZSBDQTEfMB0GA1UEAxMWU2FtcGxl
IEludGVybWVkaWF0ZSBDQTAeFw0xODA5MTAwODA3MDBaFw0yMTA5MDkwODA3MDBa
MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJ
bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNC
AAQEa6Gq6bKHsFU2pw0oBBCKkMaihSlRG97Z07rqlAKCO1J+7uUlXbRdhZ2uCjRj
d5cSG8rSWxRE703Ses+JBZPgo4HVMIHSMA4GA1UdDwEB/wQEAwIFoDATBgNVHSUE
DDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRDd2MS9Ndo68PJ
y9K/RNY6syZW0zAfBgNVHSMEGDAWgBR+Y+v8yByDNp39G7trYrTfZ0UjJzAxBggr
BgEFBQcBAQQlMCMwIQYIKwYBBQUHMAGGFWh0dHA6Ly9sb2NhbGhvc3Q6ODg4OTAq
BgNVHR8EIzAhMB+gHaAbhhlodHRwOi8vbG9jYWxob3N0Ojg4ODgvY3JsMAoGCCqG
SM49BAMCA0cAMEQCIC/Bo4oNU6yHFLJeme5ApxoNdyu3rWyiqWPxJmJAr9L0AiBl
Gc/v+yh4dHIDhCrimajTQAYOG9n0kajULI70Gg7TNw==
-----END CERTIFICATE-----
`
		s         Signer
		v         Verifier
		err       error
		signature []byte
	)

	if s, err = NewEcdsaSha256Signer([]byte(privateKey)); err != nil {
		t.Errorf("should get an correct signer instance")
	}

	if signature, err = s.Sign([]byte("hello world!")); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(signature) < 2 || len(signature) < (int(signature[1])-2) {
		t.Errorf("signature is in an incorrect length")
	}

	if int(signature[0]) != 48 {
		t.Errorf("signature has an incorrect first-byte")
	}

	if v, err = NewEcdsaSha256Verifier([]byte(publicKey)); err != nil {
		t.Errorf("should be able to create validator")
	}

	if !v.Verify([]byte("hello world!"), signature) {
		t.Errorf("signature should be verified")
	}
}
