package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type validator interface {
	Validate() error
}

func GetParam(r *http.Request, key string) string {
	value := chi.URLParam(r, key)
	return strings.TrimSpace(value)
}

func GetQuery(r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	return strings.TrimSpace(value)
}

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value. If the provided value is a struct then it is checked for
// validation tags. If the value implements a validate function, it is executed.
func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(val); err != nil {
		return err
	}

	if v, ok := val.(validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
