package permissions_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Permissions permissions.Repository
}

type handler struct {
	permissions permissions.Repository
}

func New(cfg Config) *handler {
	return &handler{permissions: cfg.Permissions}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	var input NewPermissionInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	action, err := permissions.ParseActionString(input.Action)
	if err != nil {
		return err
	}

	id, err := h.permissions.Create(r.Context(), permissions.NewPermission{
		Action:      action,
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		if err := database.CheckPQError(
			err,
			func(err *pq.Error) error {
				if err.Column == "action" && err.Code == pgerrcode.CheckViolation {
					return errors.NewRequestError("Invalid action.", http.StatusConflict, errors.DuplicateValue)
				}
				return nil
			},
		); err != nil {
			return err
		}
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Permission created.",
		NewPermissionResponse{Id: id, Name: input.Name, Action: string(action)},
	)
}
