package middlewares

import (
	"net/http"
)

func Authenticate(handler http.Handler) http.Handler {
	m := func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(m)
}
