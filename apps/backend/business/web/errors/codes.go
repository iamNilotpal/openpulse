package errors

type ErrorCode string

const (
	DuplicateValue ErrorCode = "DUPLICATE_VALUE"
	NotFound       ErrorCode = "RESOURCE_NOT_FOUND"
	AlreadyExists  ErrorCode = "RESOURCE_ALREADY_EXISTS"

	DatabaseError       ErrorCode = "DATABASE_ERROR"
	InternalServerError ErrorCode = "INTERNAL_ERROR"
	ServiceUnavailable  ErrorCode = "SERVICE_UNAVAILABLE"

	InvalidInput         ErrorCode = "INVALID_INPUT"
	MissingRequiredField ErrorCode = "MISSING_REQUIRED_FIELD"

	Forbidden             ErrorCode = "FORBIDDEN"
	Unauthorized          ErrorCode = "UNAUTHORIZED"
	TokenExpired          ErrorCode = "TOKEN_EXPIRED"
	AccountSuspended      ErrorCode = "ACCOUNT_SUSPENDED"
	InvalidAuthHeader     ErrorCode = "INVALID_TOKEN_HEADER"
	InvalidTokenSignature ErrorCode = "INVALID_TOKEN_SIGNATURE"
)

func FromErrorCode(code ErrorCode) string {
	return string(code)
}

func ToErrorCode(code string) ErrorCode {
	return ErrorCode(code)
}
