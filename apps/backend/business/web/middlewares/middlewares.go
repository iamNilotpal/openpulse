package middlewares

import "net/http"

type middleware func(handler handler) http.HandlerFunc

type handler func(w http.ResponseWriter, r *http.Request) error
