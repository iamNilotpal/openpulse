package web

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Data       any      `json:"data,omitempty"`
	Error      apiError `json:"error,omitempty"`
}

type apiError struct {
	Trace      string    `json:"-"`
	Message    string    `json:"message"`
	ErrorCode  errorType `json:"errorCode"`
	StatusCode int       `json:"statusCode"`
	Fields     any       `json:"fields,omitempty"`
}

func RespondSuccess(w http.ResponseWriter, code int, data any) error {
	response := apiResponse{Success: true, Data: data, StatusCode: code}
	return respond(w, code, response)
}

func RespondError(
	w http.ResponseWriter, code int, errorCode errorType, message string, fields any,
) error {
	response := apiResponse{
		Data:    nil,
		Success: false,
		Error: apiError{
			StatusCode: code,
			Fields:     fields,
			Message:    message,
			ErrorCode:  errorCode,
		},
	}

	return respond(w, code, response)
}

func respond(w http.ResponseWriter, statusCode int, data apiResponse) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
