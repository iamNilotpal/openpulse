package invitations_handlers

import (
	"net/http"

	"github.com/iamNilotpal/openpulse/foundation/web"
)

type Config struct{}

type handler struct{}

func New(cfg Config) *handler {
	return &handler{}
}

func (h *handler) InviteMembers(w http.ResponseWriter, r *http.Request) error {
	var payload InviteMemberPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	return web.Success(w, http.StatusOK, "Invitations sent.", nil)
}
