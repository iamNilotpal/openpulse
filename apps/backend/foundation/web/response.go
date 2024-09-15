package web

import (
	"encoding/json"
	"net/http"
)

type responseStatus string

const (
	statusOk    responseStatus = "OK"
	statusError responseStatus = "ERROR"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(w http.ResponseWriter, statusCode int, data any) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
