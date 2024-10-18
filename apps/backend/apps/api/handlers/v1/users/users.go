package users_handler

import (
	"database/sql"
	stdErrors "errors"
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Config struct {
	Users users.Repository
}

type handler struct {
	users users.Repository
}

func New(cfg Config) *handler {
	return &handler{users: cfg.Users}
}

func (h *handler) QueryById(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		return errors.NewRequestError("Invalid user id.", http.StatusUnprocessableEntity, errors.BadRequest)
	}

	user, err := h.users.QueryById(r.Context(), userId)
	if err != nil {
		if stdErrors.Is(err, sql.ErrNoRows) {
			return errors.NewRequestError("User not found.", http.StatusNotFound, errors.NotFound)
		}
		return err
	}

	return web.Success(w, http.StatusOK, "OK", map[string]any{"user": FromAppUser(user)})
}
