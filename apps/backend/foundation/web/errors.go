package web

type errorCode string

const (
	InternalErrorCode           errorCode = "INTERNAL_ERROR"
	DatabaseErrorCode           errorCode = "DATABASE_ERROR"
	ServiceUnavailableErrorCode errorCode = "SERVICE_UNAVAILABLE"

	UnknownErrorCode              errorCode = "UNKNOWN_ERROR"
	InvalidInputErrorCode         errorCode = "INVALID_INPUT"
	MissingRequiredFieldErrorCode errorCode = "MISSING_REQUIRED_FIELD"

	DuplicateValueErrorCode errorCode = "DUPLICATE_VALUE"
	NotFoundErrorCode       errorCode = "RESOURCE_NOT_FOUND"
	AlreadyExistsErrorCode  errorCode = "RESOURCE_ALREADY_EXISTS"
)

func ToErrorCode(code string) errorCode {
	return errorCode(code)
}
