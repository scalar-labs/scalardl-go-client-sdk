package model

import (
	"testing"

	"github.com/scalar-labs/dl"
	"github.com/scalar-labs/dl/ledger/asset"
	"github.com/scalar-labs/dl/ledger/statuscode"
)

func Test_LedgerValidationResult_Equal(t *testing.T) {
	shouldBeTrue := LedgerValidationResult{
		Code: statuscode.OK,
		Proof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
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
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
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
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "bar",
			Age:       1,
			Input:     dl.JSONObject{"argument": "parameter"},
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
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		},
		AuditorProof: asset.Proof{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
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

func Test_ContractExecutionResult_Equal(t *testing.T) {
	shouldBeTrue := ContractExecutionResult{
		Result: dl.JSONObject{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	}.Equal(ContractExecutionResult{
		Result: dl.JSONObject{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
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
		Result: dl.JSONObject{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0xCC, 0xDD},
			Key:       asset.ProofKey{ID: "foo", Age: 999},
		}},
	}.Equal(ContractExecutionResult{
		Result: dl.JSONObject{"argument": "parameter"},
		Proofs: []asset.Proof{{
			ID:        "bar",
			Age:       0,
			Input:     dl.JSONObject{},
			Nonce:     "i-am-nonce",
			Hash:      []byte{0x00, 0x01},
			PrevHash:  []byte{0x11, 0x22},
			Signature: []byte{0x00, 0x11},
			Key:       asset.ProofKey{ID: "bar", Age: 0},
		}},
		AuditorProofs: []asset.Proof{{
			ID:        "foo",
			Age:       999,
			Input:     dl.JSONObject{"argument": "parameter"},
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
