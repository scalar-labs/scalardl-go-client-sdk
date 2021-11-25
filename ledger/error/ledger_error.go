package error

import "github.com/scalar-labs/dl/ledger/status_code"

type LedgerError struct {
	message    string
	statusCode status_code.StatusCode
}

// NewLedgerError creates the ledger error instance.
func NewLedgerError(statusCode status_code.StatusCode, message string) LedgerError {
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
func (e LedgerError) StatusCode() status_code.StatusCode {
	return e.statusCode
}
