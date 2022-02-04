package asset

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
	"unsafe"

	"github.com/scalar-labs/scalardl-go-client-sdk/v3/crypto"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/jsonobject"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/rpc"
)

// Proof defines a proof stored in client-side to validate the server ledger states.
type Proof struct {
	ID        string
	Age       int32
	Nonce     string
	Input     jsonobject.JSONObject
	Hash      []byte
	PrevHash  []byte
	Signature []byte
	Key       ProofKey
}

// FromGRPC converts a rpc.AssetProof to a Proof.
func FromGRPC(p *rpc.AssetProof) Proof {
	if p == nil {
		return Proof{}
	}

	var input jsonobject.JSONObject
	input, _ = jsonobject.FromJSON(p.GetInput())

	return Proof{
		ID:        p.GetAssetId(),
		Age:       int32(p.GetAge()),
		Nonce:     p.GetNonce(),
		Input:     input,
		Hash:      p.GetHash(),
		PrevHash:  p.GetPrevHash(),
		Signature: p.GetSignature(),
		Key:       ProofKey{ID: p.GetAssetId(), Age: int32(p.GetAge())},
	}
}

// Equal checks if the asset proof has the same value with another one.
func (p Proof) Equal(another Proof) bool {
	return strings.Compare(p.ID, another.ID) == 0 &&
		p.Age == another.Age &&
		strings.Compare(p.Nonce, another.Nonce) == 0 &&
		p.Input.Equal(another.Input) &&
		bytes.Equal(p.Hash, another.Hash) &&
		bytes.Equal(p.PrevHash, another.PrevHash) &&
		bytes.Equal(p.Signature, another.Signature)
}

// ValueEqual checks if the asset proof has the same values except the signature with another one.
func (p Proof) ValueEqual(another Proof) bool {
	return strings.Compare(p.ID, another.ID) == 0 &&
		p.Age == another.Age &&
		strings.Compare(p.Nonce, another.Nonce) == 0 &&
		p.Input.Equal(another.Input) &&
		bytes.Equal(p.Hash, another.Hash) &&
		bytes.Equal(p.PrevHash, another.PrevHash)
}

// Serialize serialize the asset proof into a byte array.
func (p Proof) Serialize() (serialized []byte) {
	var ageBytes []byte = make([]byte, unsafe.Sizeof(p.Age))
	binary.BigEndian.PutUint32(ageBytes, (uint32)(p.Age))

	serialized = append(serialized, []byte(p.ID)...)
	serialized = append(serialized, ageBytes...)
	serialized = append(serialized, []byte(p.Nonce)...)
	serialized = append(serialized, []byte(p.Input.String())...)
	serialized = append(serialized, p.Hash...)
	serialized = append(serialized, p.PrevHash...)

	return
}

// VerifyWith checks if the signature of the asset proof can be verified by serialized value.
func (p Proof) VerifyWith(verifier crypto.Verifier) bool {
	return verifier.Verify(p.Serialize(), p.Signature)
}

// String returns the text formatted as same as the Java package: com.google.common.base.MoreObjects.toStringHelper
func (p Proof) String() (s string) {
	return fmt.Sprintf(`Proof{id=%s, age=%d, nonce=%s, input=%s, hash=%s, prev_hash=%s, signature=%s}`,
		p.ID,
		p.Age,
		p.Nonce,
		p.Input.String(),
		base64.StdEncoding.EncodeToString(p.Hash),
		base64.StdEncoding.EncodeToString(p.PrevHash),
		base64.StdEncoding.EncodeToString(p.Signature),
	)
}
