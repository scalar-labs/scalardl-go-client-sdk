package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// EcdsaSha256Signer is the signer implementation of ECDSA-SHA256.
type EcdsaSha256Signer struct {
	key *ecdsa.PrivateKey
}

// NewEcdsaSha256Signer parses given key and creates EcdsaSha256Signer.
func NewEcdsaSha256Signer(key []byte) (s EcdsaSha256Signer, err error) {
	var block *pem.Block
	block, _ = pem.Decode(key)

	if block == nil || block.Type != "EC PRIVATE KEY" {
		err = errors.New("not a EC private key")
		return
	}

	s.key, err = x509.ParseECPrivateKey(block.Bytes)

	return
}

// Sign signs given message.
func (s EcdsaSha256Signer) Sign(message []byte) (signed []byte, err error) {
	hash := sha256.Sum256(message)

	return ecdsa.SignASN1(rand.Reader, s.key, hash[:])
}

// EcdsaSha256Verifier is the verifier implementation of ECDSA-SHA256.
type EcdsaSha256Verifier struct {
	certificate *x509.Certificate
}

// NewEcdsaSha256Signer parses given certificate and creates EcdsaSha256Verifier.
func NewEcdsaSha256Verifier(certificate []byte) (v EcdsaSha256Verifier, err error) {
	var block *pem.Block
	block, _ = pem.Decode(certificate)

	var parsed *x509.Certificate
	if parsed, err = x509.ParseCertificate(block.Bytes); err != nil {
		err = errors.New("failed to parse certificate")
		return
	}

	v = EcdsaSha256Verifier{
		certificate: parsed,
	}

	return
}

// Verify verifies if given message and given signature are matched.
func (v EcdsaSha256Verifier) Verify(message []byte, signature []byte) bool {
	var hashed = sha256.Sum256(message)
	return ecdsa.VerifyASN1(v.certificate.PublicKey.(*ecdsa.PublicKey), hashed[:], signature)
}
