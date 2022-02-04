package error

import (
	"testing"

	"github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"
)

func TestNewClientError(t *testing.T) {
	var err ClientError = ClientError{
		statusCode: 401,
		message:    "error message",
	}

	if err.StatusCode() != statuscode.UnloadableKey {
		t.Errorf("should be created with correct status code")
	}

	if err.Error() != "error message" {
		t.Errorf("should be created with correct error message")
	}
}
