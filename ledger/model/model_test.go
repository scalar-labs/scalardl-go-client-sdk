package model

import (
	"testing"

	"github.com/scalar-labs/scalardl-go-client-sdk/v3/json"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/asset"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
)

func TestLedgerValidationResult_Equal(t *testing.T) {
	shouldBeTrue := LedgerValidationResult{
		Code: statuscode.OK,
		Proof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
	}.Equal(LedgerValidationResult{
		Code: statuscode.OK,
		Proof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
	})

	if !shouldBeTrue {
		t.Errorf("two identitical LedgerValidationResult should be equal")
	}

	shouldBeFalse := LedgerValidationResult{
		Code: statuscode.OK,
		Proof: asset.Proof{
			ID:        "bar",
			Age:       1,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "bar",
			Age:       1,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
	}.Equal(LedgerValidationResult{
		Code: statuscode.OK,
		Proof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
	})

	if shouldBeFalse {
		t.Errorf("two differenct LedgerValidationResult should not be equal")
	}
}

func TestContractExecutionResult_Equal(t *testing.T) {
	shouldBeTrue := ContractExecutionResult{
		Result: json.Object{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	}.Equal(ContractExecutionResult{
		Result: json.Object{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	})

	if !shouldBeTrue {
		t.Errorf("two identitical ContractExecutionResult should be equal")
	}

	shouldBeFalse := ContractExecutionResult{
		Result: json.Object{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	}.Equal(ContractExecutionResult{
		Result: json.Object{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "bar",
			Age:       0,
			Input:     json.Object{},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "bar", Age: 0},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     json.Object{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	})

	if shouldBeFalse {
		t.Errorf("two differenct ContractExecutionResult should not be equal")
	}
}
