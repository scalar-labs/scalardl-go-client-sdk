package crypto

// Verifier defines an interface to verify if the message and the signature are matched.
type Verifier interface {
	Verify(message []byte, signature []byte) (verified bool)
}
