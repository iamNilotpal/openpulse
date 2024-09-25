package middlewares

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func Authenticate(
	authenticator *auth.Auth, roles roles.Repository, permissions permissions.Repository,
) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			claims, err := authenticator.Authenticate(r.Context(), token)
			if err != nil {
				if auth.IsAuthError(err) {
					authErr := auth.ExtractAuthError(err)

					web.Error(w, authErr.Status, web.APIError{
						Message:   authErr.Error(),
						ErrorCode: errors.CodeToString(authErr.Code),
					})
					return
				}

				web.Error(
					w,
					http.StatusInternalServerError,
					web.APIError{
						ErrorCode: errors.CodeToString(errors.InternalServerError),
						Message:   http.StatusText(http.StatusInternalServerError),
					})
				return
			}

			role, err := roles.QueryById(r.Context(), claims.RoleId)
			if err != nil {
				web.Error(
					w,
					http.StatusUnauthorized,
					web.APIError{
						ErrorCode: errors.CodeToString(errors.Unauthorized),
						Message:   http.StatusText(http.StatusUnauthorized),
					},
				)
				return
			}

			userId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				web.Error(
					w,
					http.StatusUnauthorized,
					web.APIError{
						ErrorCode: errors.CodeToString(errors.Unauthorized),
						Message:   http.StatusText(http.StatusUnauthorized),
					},
				)
				return
			}

			permissions, err := permissions.QueryByUserId(r.Context(), userId)
			if err != nil {
				web.Error(
					w,
					http.StatusUnauthorized,
					web.APIError{
						ErrorCode: errors.CodeToString(errors.Unauthorized),
						Message:   http.StatusText(http.StatusUnauthorized),
					},
				)
				return
			}

			authPermissions := make([]auth.UserPermissions, len(permissions))
			for i, p := range permissions {
				authPermissions[i] = auth.ToUserPermissions(
					auth.ToUserRole(role),
					auth.ToUserPermission(p),
				)
			}

			r = r.WithContext(auth.SetClaims(r.Context(), claims))
			r = r.WithContext(auth.SetUserRole(r.Context(), auth.ToUserRole(role)))
			r = r.WithContext(auth.SetUserPermissions(r.Context(), authPermissions))

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(m)
	}

	return a
}
