package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Signer interface {
	Sign(message []byte) (signed []byte, err error)
}

type EcdsaSha256Signer struct {
	key *ecdsa.PrivateKey
}

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

func (s EcdsaSha256Signer) Sign(message []byte) (signed []byte, err error) {
	hash := sha256.Sum256(message)

	return ecdsa.SignASN1(rand.Reader, s.key, hash[:])
}
