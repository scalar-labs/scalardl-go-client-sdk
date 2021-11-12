package crypto

import (
	"encoding/pem"
	"testing"

	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
)

func Test_NewEcSha256Signer_WithIncorrectKey_ShouldGetError(t *testing.T) {
	if _, err := NewEcSha256Signer([]byte("not a key")); err == nil {
		t.Errorf("should get an error")
	}
}

func Test_NewEcSha256Signer_WithCorrectKey_ShouldGetInstance(t *testing.T) {
	var key string = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`

	if _, err := NewEcSha256Signer([]byte(key)); err != nil {
		t.Errorf("should get an correct signer instance")
	}
}

func Test_EcSha256Signer_WithCorrectKey_ShouldSignCorrectly(t *testing.T) {
	var privateKey string = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICcJGMEw3dyXUGFu/5a36HqY0ynZi9gLUfKgYWMYgr/IoAoGCCqGSM49
AwEHoUQDQgAEBGuhqumyh7BVNqcNKAQQipDGooUpURve2dO66pQCgjtSfu7lJV20
XYWdrgo0Y3eXEhvK0lsURO9N0nrPiQWT4A==
-----END EC PRIVATE KEY-----
`
	var publicKey string = `-----BEGIN CERTIFICATE-----
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
	var s EcSha256Signer
	var err error
	var signed []byte

	if s, err = NewEcSha256Signer([]byte(privateKey)); err != nil {
		t.Errorf("should get an correct signer instance")
	}

	if signed, err = s.Sign([]byte("hello world!")); err != nil {
		t.Errorf("should be able to sign")
	}

	if len(signed) < 2 || len(signed) < (int(signed[1])-2) {
		t.Errorf("signature is in an incorrect length")
	}

	if int(signed[0]) != 48 {
		t.Errorf("signature has an incorrect first-byte")
	}

	var block *pem.Block
	var certificate *x509.Certificate

	block, _ = pem.Decode([]byte(publicKey))
	if certificate, err = x509.ParseCertificate(block.Bytes); err != nil {
		t.Errorf("certificate should be parsed")
	}

	var hashed = sha256.Sum256([]byte("hello world!"))

	if !ecdsa.VerifyASN1(certificate.PublicKey.(*ecdsa.PublicKey), hashed[:], signed) {
		t.Errorf("signature should be verified")
	}
}
