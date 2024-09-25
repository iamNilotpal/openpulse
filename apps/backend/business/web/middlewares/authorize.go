package middlewares

import (
	"net/http"
	"strings"

	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

func Authorize(
	requiredPermissions ...[]auth.Permissions,
) func(http.Handler) http.Handler {
	a := func(handler http.Handler) http.Handler {
		m := func(w http.ResponseWriter, r *http.Request) {
			userPermissions := auth.GetUserPermissions(r.Context())

			if !checkPermissions(requiredPermissions, userPermissions) {
				web.Error(w, http.StatusForbidden, web.APIError{
					Message:   http.StatusText(http.StatusForbidden),
					ErrorCode: errors.CodeToString(errors.Forbidden),
				})
				return
			}

			handler.ServeHTTP(w, r)
		}

		return http.HandlerFunc(m)
	}

	return a
}

func checkPermissions(
	requiredPermissions [][]auth.Permissions, userPermissions []auth.UserPermissions,
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
				strings.EqualFold(uPermission.Permission.Action, rPermission.Permission.Action) &&
				strings.EqualFold(uPermission.Permission.Resource, rPermission.Permission.Resource) {
				return true
			}
		}
	}

	return false
}
