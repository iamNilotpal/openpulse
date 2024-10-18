package middlewares

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func VerifiedUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := auth.GetUser(r.Context())

		if !user.IsEmailVerified {
			web.Error(
				w,
				http.StatusForbidden,
				web.NewAPIError("Please verify your email.", errors.FromErrorCode(errors.Forbidden), nil),
			)
			return
		}

		h.ServeHTTP(w, r)
	})
}
