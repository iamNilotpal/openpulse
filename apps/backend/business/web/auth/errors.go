package auth

import (
	stdErrors "errors"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type AuthError struct {
	Status  int
	Message string
	Code    errors.ErrorCode
}

func NewAuthError(msg string, code errors.ErrorCode, status int) *AuthError {
	return &AuthError{Message: msg, Code: code, Status: status}
}

func (a *AuthError) Error() string {
	return a.Message
}

func IsAuthError(err error) bool {
	var a *AuthError
	return stdErrors.As(err, &a)
}

func ExtractAuthError(err error) *AuthError {
	var a *AuthError
	if !stdErrors.As(err, &a) {
		return nil
	}

	return a
}
