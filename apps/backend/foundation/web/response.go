package web

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"statusCode"`
	Data       any      `json:"data,omitempty"`
	Error      APIError `json:"error,omitempty"`
}

type APIError struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	Fields    any    `json:"fields,omitempty"`
}

func Success(w http.ResponseWriter, code int, data any) error {
	response := apiResponse{Success: true, Data: data, StatusCode: code}
	return respond(w, code, response)
}

func Error(w http.ResponseWriter, code int, err APIError) error {
	response := apiResponse{Success: false, StatusCode: code, Error: err}
	return respond(w, response.StatusCode, response)
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
