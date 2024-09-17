package users_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/core/users"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Handler struct {
	store users.Repository
}

func NewHandler(store users.Repository) *Handler {
	return &Handler{store: store}
}

func (h *Handler) QueryUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		if _, ok := err.(*strconv.NumError); ok {
			web.Respond(w, http.StatusBadRequest, web.Fail(web.ErrorResponse{Message: "Invalid id param"}))
			return
		}

		web.Respond(
			w, http.StatusInternalServerError,
			web.Fail(web.ErrorResponse{Message: http.StatusText(http.StatusInternalServerError)}),
		)
		return
	}

	user, err := h.store.QueryByUserId(r.Context(), userId)
	if err != nil {
		web.Respond(
			w, http.StatusNotFound,
			web.Fail(web.ErrorResponse{Message: fmt.Sprintf("User with id %d doesn't exists.", userId)}),
		)
		return
	}

	web.Respond(w, http.StatusOK, web.Success(user))
}
