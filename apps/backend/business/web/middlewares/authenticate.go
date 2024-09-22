package middlewares

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func Authenticate(authenticator *auth.Auth) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			claims, err := authenticator.Authenticate(r.Context(), token)
			if err != nil {
				if auth.IsAuthError(err) {
					authErr := auth.ExtractAuthError(err)

					web.Error(w, http.StatusUnauthorized, web.APIError{
						Message:   authErr.Error(),
						ErrorCode: errors.ToString(authErr.Code),
					})
					return
				}

				web.Error(w, http.StatusInternalServerError, web.APIError{
					ErrorCode: errors.ToString(errors.InternalErrorCode),
					Message:   http.StatusText(http.StatusInternalServerError),
				})
				return
			}

			auth.SetClaims(r.Context(), claims)
			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(m)
	}

	return a
}
