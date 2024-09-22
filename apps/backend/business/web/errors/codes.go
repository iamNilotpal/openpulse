package errors

type ErrorCode string

const (
	TokenExpiredErrorCode      ErrorCode = "TOKEN_EXPIRED"
	AccountSuspendedErrorCode  ErrorCode = "ACCOUNT_SUSPENDED"
	InvalidAuthHeaderErrorCode ErrorCode = "INVALID_TOKEN_HEADER"
	InvalidTokenSignature      ErrorCode = "INVALID_TOKEN_SIGNATURE"

	DuplicateValueErrorCode ErrorCode = "DUPLICATE_VALUE"
	NotFoundErrorCode       ErrorCode = "RESOURCE_NOT_FOUND"
	AlreadyExistsErrorCode  ErrorCode = "RESOURCE_ALREADY_EXISTS"

	InternalErrorCode           ErrorCode = "INTERNAL_ERROR"
	DatabaseErrorCode           ErrorCode = "DATABASE_ERROR"
	ServiceUnavailableErrorCode ErrorCode = "SERVICE_UNAVAILABLE"

	UnknownErrorCode              ErrorCode = "UNKNOWN_ERROR"
	InvalidInputErrorCode         ErrorCode = "INVALID_INPUT"
	MissingRequiredFieldErrorCode ErrorCode = "MISSING_REQUIRED_FIELD"
)

func ToString(code ErrorCode) string {
	return string(code)
}
