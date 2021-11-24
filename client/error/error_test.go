package error

import (
	"testing"

	"github.com/scalar-labs/dl/ledger/status_code"
)

func Test_NewClientError_StatusCodeAndMessage_ShouldBeSuccessful(t *testing.T) {
	var err ClientError = ClientError{
		statusCode: 401,
		message:    "error message",
	}

	if err.StatusCode() != status_code.UNLOADABLE_KEY {
		t.Errorf("should be created with correct status code")
	}

	if err.Error() != "error message" {
		t.Errorf("should be created with correct error message")
	}
}
