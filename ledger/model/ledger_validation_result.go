package model

import (
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/asset"
	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
)

// LedgerValidationResult defines the specified status code and the asset proof from Ledger and Auditor.
type LedgerValidationResult struct {
	Code         statuscode.StatusCode
	Proof        asset.Proof
	AuditorProof asset.Proof
}

// Equal checks if two ledger validation result have the same values.
func (r LedgerValidationResult) Equal(another LedgerValidationResult) bool {
	return r.Code == another.Code &&
		r.Proof.Equal(another.Proof) &&
		r.AuditorProof.Equal(another.AuditorProof)
}
