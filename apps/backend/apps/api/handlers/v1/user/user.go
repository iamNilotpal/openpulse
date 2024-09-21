package user_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repository/user"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Handler struct {
	userRepo *user.Repository
}

func New(userRepo *user.Repository) *Handler {
	return &Handler{userRepo: userRepo}
}

func (h *Handler) QueryById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		if _, ok := err.(*strconv.NumError); ok {
			web.Error(w, http.StatusBadRequest, web.InvalidInputErrorCode, "Invalid ID", nil)
			return
		}

		web.Error(w, http.StatusInternalServerError, web.InternalErrorCode, "Invalid ID", nil)
		return
	}

	user, err := h.userRepo.QueryById(r.Context(), userId)
	if err != nil {
		web.Error(
			w,
			http.StatusNotFound,
			web.NotFoundErrorCode,
			fmt.Sprintf("User with id %d doesn't exists.", userId),
			nil,
		)
		return
	}

	web.Success(w, http.StatusOK, user)
}
