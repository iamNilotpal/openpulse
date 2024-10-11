package errors

type ErrorCode string

const (
	FlowIncomplete ErrorCode = "INCOMPLETE_FLOW"

	DuplicateValue ErrorCode = "DUPLICATE_VALUE"
	NotFound       ErrorCode = "RESOURCE_NOT_FOUND"
	AlreadyExists  ErrorCode = "RESOURCE_ALREADY_EXISTS"

	BadRequest          ErrorCode = "BAD_REQUEST"
	InternalServerError ErrorCode = "INTERNAL_ERROR"
	ServiceUnavailable  ErrorCode = "SERVICE_UNAVAILABLE"

	InvalidInput          ErrorCode = "INVALID_INPUT"
	UnknownField          ErrorCode = "UNKNOWN_FIELD"
	MissingRequiredField  ErrorCode = "MISSING_REQUIRED_FIELD"
	MissingRequiredFields ErrorCode = "MISSING_REQUIRED_FIELDS"

	Forbidden             ErrorCode = "FORBIDDEN"
	Unauthorized          ErrorCode = "UNAUTHORIZED"
	TokenExpired          ErrorCode = "TOKEN_EXPIRED"
	InvalidAuthHeader     ErrorCode = "INVALID_AUTH_HEADER"
	InvalidTokenSignature ErrorCode = "INVALID_TOKEN_SIGNATURE"
)

func FromErrorCode(code ErrorCode) string {
	return string(code)
}

func ToErrorCode(code string) ErrorCode {
	return ErrorCode(code)
}
