package middlewares

import "net/http"

type handler func(w http.ResponseWriter, r *http.Request) error

type middleware func(handler handler) http.HandlerFunc
