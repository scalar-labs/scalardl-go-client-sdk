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

type EcSha256Signer struct {
	key *ecdsa.PrivateKey
}

func NewEcSha256Signer(raw []byte) (s EcSha256Signer, err error) {
	var block *pem.Block
	block, _ = pem.Decode(raw)

	if block == nil || block.Type != "EC PRIVATE KEY" {
		err = errors.New("not a EC private key")
		return
	}

	s.key, err = x509.ParseECPrivateKey(block.Bytes)

	return
}

func (s EcSha256Signer) Sign(message []byte) (signed []byte, err error) {
	hash := sha256.Sum256(message)

	return ecdsa.SignASN1(rand.Reader, s.key, hash[:])
}
