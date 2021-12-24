package model

import (
	"fmt"

	"github.com/scalar-labs/dl"
	"github.com/scalar-labs/dl/ledger/asset"
)

// ContractExecutionResult defines the result of a contract execution.
// It contains the result of the contract execution along with a list of asset proofs from Ledger and Auditor.
type ContractExecutionResult struct {
	Result        dl.JSONObject
	Proofs        []asset.Proof
	AuditorProofs []asset.Proof
}

// Equal checks if two contract execution results have the same values.
func (r ContractExecutionResult) Equal(another ContractExecutionResult) (equal bool) {
	var (
		myResult      dl.JSONObject = r.Result
		anotherResult dl.JSONObject = another.Result
	)

	if (myResult == nil && anotherResult != nil) || (myResult != nil && anotherResult == nil) {
		return false
	}

	if !myResult.Equal(anotherResult) {
		return false
	}

	if len(r.Proofs) != len(another.Proofs) {
		return false
	}

	var anotherProofsInMap map[string]asset.Proof = make(map[string]asset.Proof)

	for _, p := range another.Proofs {
		key := fmt.Sprintf("%s-%d", p.ID, p.Age)
		anotherProofsInMap[key] = p
	}

	for _, p := range r.Proofs {
		key := fmt.Sprintf("%s-%d", p.ID, p.Age)
		a, ok := anotherProofsInMap[key]

		if !ok || !p.Equal(a) {
			return false
		}
	}

	if len(r.AuditorProofs) != len(another.AuditorProofs) {
		return false
	}

	var anotherAuditorProofsInMap map[string]asset.Proof = make(map[string]asset.Proof)

	for _, p := range another.AuditorProofs {
		key := fmt.Sprintf("%s-%d", p.ID, p.Age)
		anotherAuditorProofsInMap[key] = p
	}

	for _, p := range r.AuditorProofs {
		key := fmt.Sprintf("%s-%d", p.ID, p.Age)
		a, ok := anotherAuditorProofsInMap[key]

		if !ok || !p.Equal(a) {
			return false
		}
	}

	return true
}
