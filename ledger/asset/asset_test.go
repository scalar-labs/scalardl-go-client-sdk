package asset

import (
	"bytes"
	"testing"

	"github.com/scalar-labs/dl/v3"
	"github.com/scalar-labs/dl/v3/crypto"
	"github.com/scalar-labs/dl/v3/rpc"
)

func TestProofKey_Equal(t *testing.T) {
	var shouldBeTrue = ProofKey{
		ID:  "foo",
		Age: 1,
	}.Equal(ProofKey{
		ID:  "foo",
		Age: 1,
	})

	if !shouldBeTrue {
		t.Errorf("should be equal")
	}

	var shouldBeFalse = ProofKey{
		ID:  "foo",
		Age: 1,
	}.Equal(ProofKey{
		ID:  "bar",
		Age: 2,
	})

	if shouldBeFalse {
		t.Errorf("should not be equal")
	}
}

func TestProofKey_Compare(t *testing.T) {
	var shouldBe0 = ProofKey{
		ID:  "foo",
		Age: 1,
	}.Compare(ProofKey{
		ID:  "foo",
		Age: 1,
	})

	if shouldBe0 != 0 {
		t.Errorf("should be 0 when two ProofKeys are the same")
	}

	var shouldBePositive = ProofKey{
		ID:  "b",
		Age: 1,
	}.Compare(ProofKey{
		ID:  "a",
		Age: 1,
	})

	if shouldBePositive <= 0 {
		t.Errorf("should be positive when ID is bigger")
	}

	shouldBePositive = ProofKey{
		ID:  "a",
		Age: 1,
	}.Compare(ProofKey{
		ID:  "a",
		Age: 0,
	})

	if shouldBePositive <= 0 {
		t.Errorf("should be positive when Age is bigger")
	}

	var shouldBeNegtive = ProofKey{
		ID:  "a",
		Age: 1,
	}.Compare(ProofKey{
		ID:  "b",
		Age: 1,
	})

	if shouldBeNegtive >= 0 {
		t.Errorf("should be negtive when ID is smaller")
	}

	shouldBeNegtive = ProofKey{
		ID:  "a",
		Age: 0,
	}.Compare(ProofKey{
		ID:  "a",
		Age: 1,
	})

	if shouldBeNegtive >= 0 {
		t.Errorf("should be negtive when Age is smaller")
	}
}

func TestProof_FromGRPC(t *testing.T) {
	var proof = FromGRPC(&rpc.AssetProof{
		AssetId:   "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     `{"argument":"parameter"}`,
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0xAA, 0xDD},
	})

	if !proof.Key.Equal(ProofKey{ID: "foo", Age: 999}) {
		t.Errorf("ProofKey is not correctly parsed")
	}

	if proof.ID != "foo" {
		t.Errorf("ID is not correctly parsed")
	}

	if proof.Age != 999 {
		t.Errorf("Age is not correctly parsed")
	}

	if proof.Nonce != "a-nonce" {
		t.Errorf("Nonce is not correctly parsed")
	}

	if !proof.Input.Equal(dl.JSONObject{"argument": "parameter"}) {
		t.Errorf("Input is not correctly parsed")
	}

	if !bytes.Equal(proof.Hash, []byte{0x00, 0x01}) {
		t.Errorf("Hash is not correctly parsed")
	}

	if !bytes.Equal(proof.PrevHash, []byte{0xCA, 0xFE}) {
		t.Errorf("PrevHash is not correctly parsed")
	}

	if !bytes.Equal(proof.Signature, []byte{0xAA, 0xDD}) {
		t.Errorf("Signature is not correctly parsed")
	}
}

func TestProof_Equal(t *testing.T) {
	var shouldBeTrue = Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0xAA, 0xDD},
		Key:       ProofKey{ID: "foo", Age: 999},
	}.Equal(Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0xAA, 0xDD},
		Key:       ProofKey{ID: "foo", Age: 999},
	})

	if !shouldBeTrue {
		t.Errorf("should be true if two Proofs are the same")
	}

	var shouldBeFalse = Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x55, 0x66},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0xAA, 0xDD},
		Key:       ProofKey{ID: "foo", Age: 999},
	}.Equal(Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0xAA, 0xDD},
		Key:       ProofKey{ID: "foo", Age: 999},
	})

	if shouldBeFalse {
		t.Errorf("should be false if two Proofs have different hash")
	}
}

func TestProof_ValueEqual(t *testing.T) {
	var shouldBeTrue = Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0x11, 0x22},
		Key:       ProofKey{ID: "foo", Age: 999},
	}.ValueEqual(Proof{
		ID:        "foo",
		Age:       999,
		Nonce:     "a-nonce",
		Input:     dl.JSONObject{"argument": "parameter"},
		Hash:      []byte{0x00, 0x01},
		PrevHash:  []byte{0xCA, 0xFE},
		Signature: []byte{0x33, 0x44}, // different
		Key:       ProofKey{ID: "foo", Age: 999},
	})

	if !shouldBeTrue {
		t.Errorf("should be true if two Proofs are the same in values")
	}

	var shouldBeFalse = Proof{
		ID:       "foo",
		Age:      999,
		Nonce:    "a-nonce",
		Input:    dl.JSONObject{"argument": "parameter"},
		Hash:     []byte{0x55, 0x66},
		PrevHash: []byte{0xCA, 0xFE},
	}.Equal(Proof{
		ID:       "foo",
		Age:      999,
		Nonce:    "a-nonce",
		Input:    dl.JSONObject{"argument": "parameter"},
		Hash:     []byte{0x00, 0x01},
		PrevHash: []byte{0xCA, 0xFE},
	})

	if shouldBeFalse {
		t.Errorf("should be false if two Proofs have different hash")
	}
}

func TestProof_Serialize(t *testing.T) {
	var (
		proof = Proof{
			ID:       "foo",
			Age:      999,
			Nonce:    "a-nonce",
			Input:    dl.JSONObject{"argument": "parameter"},
			Hash:     []byte{0x00, 0x01},
			PrevHash: []byte{0xCA, 0xFE},
		}

		serialized = []byte{
			102, 111, 111, 0, 0, 3, 231, 97, 45, 110, 111, 110,
			99, 101, 123, 34, 97, 114, 103, 117, 109, 101, 110,
			116, 34, 58, 34, 112, 97, 114, 97, 109, 101, 116, 101, 114, 34, 125, 0, 1, 202, 254,
		}
	)

	if !bytes.Equal(proof.Serialize(), serialized) {
		t.Errorf("should serialize correctly")
	}

	if bytes.Equal(proof.Serialize(), []byte{0, 1, 2, 4}) {
		t.Errorf("should not serialize to a wrong byte array")
	}
}

func TestProof_String(t *testing.T) {
	var (
		proof = Proof{
			ID:        "foo",
			Age:       999,
			Nonce:     "a-nonce",
			Input:     dl.JSONObject{"argument": "parameter"},
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0xCA, 0xFE},
			Signature: []byte{0x11, 0x22},
		}

		stringtified = `Proof{id=foo, age=999, nonce=a-nonce, input={"argument":"parameter"}, hash=AAE=, prev_hash=yv4=, signature=ESI=}`
	)

	if proof.String() != stringtified {
		t.Errorf("should generate correct string")
	}
}

func TestProof_VerifyWith(t *testing.T) {
	var proof = Proof{
		ID:       "foo",
		Age:      999,
		Nonce:    "a-nonce",
		Input:    dl.JSONObject{"argument": "parameter"},
		Hash:     []byte{0x00, 0x01},
		PrevHash: []byte{0xCA, 0xFE},
		Signature: []byte{
			48, 69, 2, 33, 0, 182,
			176, 5, 207, 234, 232,
			157, 63, 91, 91, 115, 57,
			234, 195, 27, 85, 192, 219,
			130, 177, 18, 163, 11, 192,
			68, 38, 184, 126, 183, 90,
			58, 170, 2, 32, 14, 174, 29,
			253, 223, 130, 145, 244, 75,
			117, 134, 139, 38, 237, 34,
			179, 151, 64, 160, 157, 232,
			184, 106, 187, 138, 115, 39,
			77, 104, 101, 75, 193,
		},
	}

	var certificate = `-----BEGIN CERTIFICATE-----
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

	verifier, _ := crypto.NewEcdsaSha256Verifier([]byte(certificate))

	if !proof.VerifyWith(verifier) {
		t.Errorf("shoule be verified")
	}
}
