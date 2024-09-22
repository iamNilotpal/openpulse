package user_handler

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/user"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Handler struct {
	userRepo *user.Repository
}

func New(userRepo *user.Repository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) QueryById(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		return err
	}

	user, err := h.userRepo.QueryById(r.Context(), userId)
	if err != nil {
		return err
	}

	return web.Success(w, http.StatusOK, user)
}
