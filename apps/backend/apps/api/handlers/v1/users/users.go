package users_handler

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type handler struct {
	usersRepo users.Repository
}

func New(usersRepo users.Repository) *handler {
	return &handler{usersRepo: usersRepo}
}

func (h *handler) QueryById(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		return err
	}

	user, err := h.usersRepo.QueryById(r.Context(), userId)
	if err != nil {
		return err
	}

	return web.Success(w, http.StatusOK, "", user)
}
