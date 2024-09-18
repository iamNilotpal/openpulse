package web

type errorType string

const (
	UnknownErrorType    errorType = "unknown_error"
	NotFoundErrorType   errorType = "notfound_error"
	ServiceErrorType    errorType = "service_error"
	InternalErrorType   errorType = "internal_error"
	ValidationErrorType errorType = "validation_error"
)
