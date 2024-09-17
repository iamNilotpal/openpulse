package users_handler

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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

	web.Respond(w, http.StatusOK, web.Success(userId))
}
