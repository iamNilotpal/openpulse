package roles_handler

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Roles roles.Repository
}

type handler struct {
	roles roles.Repository
}

func New(cfg Config) *handler {
	return &handler{roles: cfg.Roles}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var input NewRoleInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	appRole, err := roles.ParseRoleString(input.Role)
	if err != nil {
		return err
	}

	id, err := h.roles.Create(
		r.Context(),
		roles.NewRole{
			IsSystemRole: true,
			Role:         appRole,
			Name:         input.Name,
			Description:  input.Description,
		},
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

	return web.Success(
		w,
		http.StatusCreated,
		"Role created successfully.",
		NewRoleResponse{Id: id, Role: string(appRole), Name: input.Name},
	)
}