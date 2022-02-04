package error

import "github.com/scalar-labs/scalardl-go-client-sdk/v3/ledger/statuscode"

// LedgerError represents the errors fromo Ledger.
// It implements the Error interface.
type LedgerError struct {
	message    string
	statusCode statuscode.StatusCode
}

// NewLedgerError creates the ledger error instance.
func NewLedgerError(statusCode statuscode.StatusCode, message string) LedgerError {
	return LedgerError{
		message:    message,
		statusCode: statusCode,
	}
}

// Error just returns the error message.
func (e LedgerError) Error() string {
	return e.message
}

// StatusCode returns the status code that represents the type of errors.
func (e LedgerError) StatusCode() statuscode.StatusCode {
	return e.statusCode
}
