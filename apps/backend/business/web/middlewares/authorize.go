package middlewares

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Options struct {
	Strict              bool
	RequiredRole        auth.RoleConfig
	RequiredPermissions []auth.PermissionConfig
}

func Authorize(options Options) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			userRole := auth.GetRole(r.Context())
			userResources := auth.GetResourcesMap(r.Context())

			if len(userResources) == 0 ||
				!auth.CheckRoleAccessControl(options.RequiredRole, userRole) {
				web.Error(
					w,
					http.StatusForbidden,
					web.CreateAPIError(
						http.StatusText(http.StatusForbidden),
						errors.FromErrorCode(errors.Forbidden),
						nil,
					),
				)
				return
			}

			userPermissions := make([]auth.UserPermissionConfig, 0)
			for _, v := range userResources {
				userPermissions = append(userPermissions, v...)
			}

			if !auth.CheckPermissionAccessControl(
				options.Strict, userPermissions, options.RequiredPermissions,
			) {
				web.Error(
					w,
					http.StatusForbidden,
					web.CreateAPIError(
						http.StatusText(http.StatusForbidden),
						errors.FromErrorCode(errors.Forbidden),
						nil,
					),
				)
				return
			}

			handler.ServeHTTP(w, r)
		}
		return http.HandlerFunc(m)
	}
	return a
}
