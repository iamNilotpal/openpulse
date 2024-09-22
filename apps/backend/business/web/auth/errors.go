package auth

import (
	stdErrors "errors"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type AuthError struct {
	Message string
	Code    errors.ErrorCode
}

func NewAuthError(msg string, code errors.ErrorCode) *AuthError {
	return &AuthError{Message: msg, Code: code}
}

func (a *AuthError) Error() string {
	return a.Message
}

func IsAuthError(err error) bool {
	var a *AuthError
	return stdErrors.Is(err, a)
}

func ExtractAuthError(err error) *AuthError {
	var a *AuthError
	if !stdErrors.As(err, &a) {
		return nil
	}

	return a
}
