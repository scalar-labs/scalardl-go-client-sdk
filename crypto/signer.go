package crypto

// Signer defines an interface to sign messages.
type Signer interface {
	Sign(message []byte) (signature []byte, err error)
}
