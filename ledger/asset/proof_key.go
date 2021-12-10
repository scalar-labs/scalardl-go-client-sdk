package asset

import "strings"

// ProofKey provides
type ProofKey struct {
	ID  string
	Age int32
}

// Equal tests if this key has the same values with another one.
func (k ProofKey) Equal(another ProofKey) bool {
	return k.ID == another.ID && k.Age == another.Age
}

// Compare compares two asset keys.
func (k ProofKey) Compare(another ProofKey) int {
	if k.Equal(another) {
		return 0
	}

	if strings.Compare(k.ID, another.ID) != 0 {
		return strings.Compare(k.ID, another.ID)
	}

	return (int)(k.Age - another.Age)
}
