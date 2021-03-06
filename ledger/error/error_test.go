package error

import (
	"testing"

	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
)

func TestNewLedgerError(t *testing.T) {
	var err LedgerError = NewLedgerError(400, "invalid signature")

	if err.StatusCode() != statuscode.InvalidSignature {
		t.Errorf("should be created with correct status code")
	}

	if err.Error() != "invalid signature" {
		t.Errorf("should be created with correct error message")
	}
}
