package model

import (
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

	var (
		myProofs      []asset.Proof = r.Proofs
		anotherProofs []asset.Proof = another.Proofs
	)

	if len(myProofs) != len(anotherProofs) {
		return false
	}

	for i := range myProofs {
		if !myProofs[i].Equal(anotherProofs[i]) {
			return false
		}
	}

	var (
		myAuditorProofs      []asset.Proof = r.AuditorProofs
		anotherAuditorProofs []asset.Proof = another.AuditorProofs
	)

	if len(myAuditorProofs) != len(anotherAuditorProofs) {
		return false
	}

	for i := range myAuditorProofs {
		if !myAuditorProofs[i].Equal(anotherAuditorProofs[i]) {
			return false
		}
	}

	return true
}
