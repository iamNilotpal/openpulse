package middlewares

import (
	"net/http"
	"strings"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func Authorize(
	requiredPermissions ...[]auth.AuthAccessControl,
) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			userPermissions := auth.GetUserPermissions(r.Context())

			if !checkPermissions(requiredPermissions, userPermissions) {
				web.Error(
					w,
					http.StatusForbidden,
					web.NewAPIError(
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

func checkPermissions(
	requiredPermissions [][]auth.AuthAccessControl, userPermissions []auth.UserAccessControl,
) bool {
	if len(requiredPermissions) == 0 {
		return true
	}

	for _, rr := range requiredPermissions {
		for i := 0; i < len(requiredPermissions); i++ {
			rPermission := rr[i]
			uPermission := userPermissions[i]

			if uPermission.Permission.Enabled &&
				uPermission.Role.Id == rPermission.Role.Id &&
				uPermission.Permission.Id == rPermission.Permission.Id &&
				strings.EqualFold(
					permissions.FromPermissionAction(uPermission.Permission.Action),
					permissions.FromPermissionAction(rPermission.Permission.Action),
				) {
				return true
			}
		}
	}

	return false
}
