package web

type ErrorResponse struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

type responseSuccess struct {
	Data   any            `json:"data"`
	Status responseStatus `json:"status"`
}

type responseError struct {
	Error  ErrorResponse  `json:"error"`
	Status responseStatus `json:"status"`
}

// Fail converts an error to valid error response.
func Fail(err ErrorResponse) responseError {
	return responseError{Error: err, Status: statusError}
}

// Success converts any data to valid success response.
func Success(data any) responseSuccess {
	return responseSuccess{Data: data, Status: statusOk}
}
