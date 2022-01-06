package error

import "github.com/scalar-labs/dl/v3/ledger/statuscode"

// ClientError is used when ClientService has errors.
// It implements the Error interface.
type ClientError struct {
	message    string
	statusCode statuscode.StatusCode
}

// NewClientError creates the client error instance.
func NewClientError(statusCode statuscode.StatusCode, message string) ClientError {
	return ClientError{
		message:    message,
		statusCode: statusCode,
	}
}

// Error just returns the error message.
func (e ClientError) Error() string {
	return e.message
}

// StatusCode returns the status code that represents the type of errors.
func (e ClientError) StatusCode() statuscode.StatusCode {
	return e.statusCode
}
