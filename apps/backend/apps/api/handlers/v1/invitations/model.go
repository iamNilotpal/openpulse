package invitations_handlers

type InviteMemberPayload struct {
	Emails []string `json:"emails"`
	TeamId int      `json:"teamId"`
}
