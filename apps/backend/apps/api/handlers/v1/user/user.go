package user_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/core/user"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Handler struct {
	userCore *user.Core
}

func New(userCore *user.Core) *Handler {
	return &Handler{userCore: userCore}
}

func (h *Handler) QueryById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(web.GetParam(r, "id"))
	if err != nil {
		if _, ok := err.(*strconv.NumError); ok {
			web.RespondError(w, http.StatusBadRequest, web.ValidationErrorType, "Invalid ID", nil)
			return
		}

		web.RespondError(w, http.StatusInternalServerError, web.InternalErrorType, "Invalid ID", nil)
		return
	}

	user, err := h.userCore.QueryById(r.Context(), userId)
	if err != nil {
		web.RespondError(
			w,
			http.StatusNotFound,
			web.NotFoundErrorType,
			fmt.Sprintf("User with id %d doesn't exists.", userId),
			nil,
		)
		return
	}

	web.RespondSuccess(w, http.StatusOK, user)
}
