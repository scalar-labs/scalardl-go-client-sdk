package error

import "github.com/scalar-labs/dl/ledger/status_code"

// ClientError is used when ClientService has errors.
// It implements the Error interface.
type ClientError struct {
	message    string
	statusCode status_code.StatusCode
}

// Error just returns the error message.
func (e ClientError) Error() string {
	return e.message
}

// StatusCode returns the status code that represents the type of errors.
func (e ClientError) StatusCode() status_code.StatusCode {
	return e.statusCode
}
