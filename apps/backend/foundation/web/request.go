package web

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetParam(r *http.Request, key string) string {
	value := chi.URLParam(r, key)
	return strings.TrimSpace(value)
}
