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
			claims, err := authenticator.Authenticate(token)

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

			resources := user.AccessControl
			userResourcesMap := make(auth.UserAccessControlMap)

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

			r = r.WithContext(auth.SetUser(r.Context(), user))
			r = r.WithContext(auth.SetUserAccessControl(r.Context(), userResourcesMap))

			handler.ServeHTTP(w, r)
		}
		return http.HandlerFunc(m)
	}
	return a
}

func AuthenticateOnboard(authenticator *auth.Auth, usersRepo users.Repository,
) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			claims, err := authenticator.AuthenticateOnboard(token)

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

			r = r.WithContext(auth.SetUser(r.Context(), user))
			handler.ServeHTTP(w, r)
		}
		return http.HandlerFunc(m)
	}
	return a
}
