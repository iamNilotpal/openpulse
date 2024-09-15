package web

type ErrorResponse struct {
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

type responseSuccess[T any] struct {
	Data   T              `json:"data"`
	Status responseStatus `json:"status"`
}

type responseError[T ErrorResponse] struct {
	Error  T              `json:"error"`
	Status responseStatus `json:"status"`
}

// Fail converts an error to valid error response.
func Fail(err ErrorResponse) responseError[ErrorResponse] {
	return responseError[ErrorResponse]{Error: err, Status: statusError}
}

// Success converts any data to valid success response.
func Success[T any](data T) responseSuccess[T] {
	return responseSuccess[T]{Data: data, Status: statusOk}
}
