package roles

type AppRole string

var (
	RoleSuperAdmin       AppRole = "super_admin"
	RoleAdmin            AppRole = "admin"
	RoleTeamAdmin        AppRole = "team_admin"
	RoleTeamBillingAdmin AppRole = "team_billing_admin"
	RoleTeamLead         AppRole = "team_lead"
	RoleTeamResponder    AppRole = "team_responder"
	RoleTeamMember       AppRole = "team_member"
)

func FromAppRole(s AppRole) string {
	return string(s)
}

func ToAppRole(s string) AppRole {
	return AppRole(s)
}
