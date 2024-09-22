package middlewares

import (
	"net/http"
)

func Authorize(handler http.Handler) http.Handler {
	m := func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(m)
}
