package errors

import (
	"errors"
)

// RequestError is used to pass an error during the request through the
// application with web specific context.
type RequestError struct {
	Status int
	Msg    string
	Code   ErrorCode
}

// NewRequestError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRequestError(msg string, status int, code ErrorCode) error {
	return &RequestError{Msg: msg, Status: status, Code: code}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (re *RequestError) Error() string {
	return re.Msg
}

// IsRequestError checks if an error of type RequestError exists.
func IsRequestError(err error) bool {
	var re *RequestError
	return errors.Is(err, re)
}

// GetRequestError returns a copy of the RequestError pointer.
func GetRequestError(err error) *RequestError {
	var re *RequestError
	if !errors.As(err, &re) {
		return nil
	}
	return re
}
