package error

import (
	"testing"

	"github.com/scalar-labs/dl/ledger/status_code"
)

func Test_NewClientError_StatusCodeAndMessage_ShouldBeSuccessful(t *testing.T) {
	var err LedgerError = LedgerError{
		statusCode: 400,
		message:    "invalid signature",
	}

	if err.StatusCode() != status_code.INVALID_SIGNATURE {
		t.Errorf("should be created with correct status code")
	}

	if err.Error() != "invalid signature" {
		t.Errorf("should be created with correct error message")
	}
}
