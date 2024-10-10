package roles_handler

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type handler struct {
	roles roles.Repository
}

func New(rolesRepo roles.Repository) *handler {
	return &handler{roles: rolesRepo}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var role NewAppRole
	if err := web.Decode(r, &role); err != nil {
		return err
	}

	claims := auth.GetClaims(r.Context())

	id, err := h.roles.Create(
		r.Context(), roles.ConstructRole(role.Name, role.Description, claims.UserId, false),
	)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == pgerrcode.UniqueViolation {
				return errors.NewRequestError(
					"Role already exists", http.StatusConflict, errors.AlreadyExists,
				)
			}

			return errors.NewRequestError(
				"Unable to create role", http.StatusInternalServerError, errors.InternalServerError,
			)
		}

		return err
	}

	return web.Success(w, http.StatusCreated, "", map[string]int{"id": id})
}
