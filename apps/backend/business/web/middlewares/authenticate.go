package middlewares

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func Authenticate(
	authenticator *auth.Auth, usersRepo users.Repository,
) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			claims, err := authenticator.Authenticate(r.Context(), token)

			if err != nil {
				if auth.IsAuthError(err) {
					authErr := auth.ExtractAuthError(err)
					web.Error(
						w,
						authErr.Status,
						web.NewAPIError(authErr.Error(), errors.FromErrorCode(authErr.Code), nil),
					)
					return
				}

				web.Error(
					w,
					http.StatusInternalServerError,
					web.NewAPIError(
						http.StatusText(http.StatusInternalServerError),
						errors.FromErrorCode(errors.InternalServerError),
						nil,
					),
				)
				return
			}

			userId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				web.Error(
					w,
					http.StatusUnauthorized,
					web.NewAPIError(
						http.StatusText(http.StatusUnauthorized),
						errors.FromErrorCode(errors.Unauthorized),
						nil,
					),
				)
				return
			}

			user, err := usersRepo.QueryById(r.Context(), userId)
			if err != nil {
				web.Error(
					w,
					http.StatusUnauthorized,
					web.NewAPIError(
						http.StatusText(http.StatusUnauthorized),
						errors.FromErrorCode(errors.Unauthorized),
						nil,
					),
				)
				return
			}

			role := user.Role
			resources := user.Resources
			userResourcesMap := make(auth.UserResourcePermissionsMap)

			for _, resource := range resources {
				val, ok := userResourcesMap[resource.Resource.Resource]
				if !ok {
					permissions := []auth.UserPermissionConfig{
						auth.NewUserPermissionConfig(resource.Permission),
					}

					userResourcesMap[resource.Resource.Resource] = permissions
					continue
				}

				permissions := append(val, auth.NewUserPermissionConfig(resource.Permission))
				userResourcesMap[resource.Resource.Resource] = permissions
			}

			r = r.WithContext(auth.SetClaims(r.Context(), claims))
			r = r.WithContext(auth.SetUserRole(r.Context(), auth.NewUserRoleConfig(role)))
			r = r.WithContext(auth.SetUserResources(r.Context(), userResourcesMap))

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(m)
	}

	return a
}
