package roles

import (
	"net/http"
	"slices"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type AppRole string

const (
	RoleOrgAdmin         AppRole = "org_admin"
	RoleTeamAdmin        AppRole = "team_admin"
	RoleTeamBillingAdmin AppRole = "team_billing_admin"
	RoleTeamLead         AppRole = "team_lead"
	RoleTeamResponder    AppRole = "team_responder"
	RoleTeamMember       AppRole = "team_member"
)

const (
	RoleOrgAdminInt int = iota + 1
	RoleTeamAdminInt
	RoleTeamBillingAdminInt
	RoleTeamLeadInt
	RoleTeamResponderInt
	RoleTeamMemberInt
)

var roles = []AppRole{
	RoleOrgAdmin,
	RoleTeamLead,
	RoleTeamAdmin,
	RoleTeamMember,
	RoleTeamResponder,
	RoleTeamBillingAdmin,
}

var roleMapping = map[AppRole]int{
	RoleOrgAdmin:         RoleOrgAdminInt,
	RoleTeamAdmin:        RoleTeamAdminInt,
	RoleTeamBillingAdmin: RoleTeamBillingAdminInt,
	RoleTeamLead:         RoleTeamLeadInt,
	RoleTeamResponder:    RoleTeamResponderInt,
	RoleTeamMember:       RoleTeamMemberInt,
}

var roleMappingReverse = map[int]AppRole{
	RoleOrgAdminInt:         RoleOrgAdmin,
	RoleTeamAdminInt:        RoleTeamAdmin,
	RoleTeamBillingAdminInt: RoleTeamBillingAdmin,
	RoleTeamLeadInt:         RoleTeamLead,
	RoleTeamResponderInt:    RoleTeamResponder,
	RoleTeamMemberInt:       RoleTeamMember,
}

func ParseRoleString(s string) (AppRole, error) {
	if contains := slices.Contains(roles, AppRole(s)); contains {
		return AppRole(s), nil
	}
	return "", errors.NewRequestError("Invalid role type.", http.StatusBadRequest, errors.BadRequest)
}

func ParseRole(s AppRole) int {
	return roleMapping[s]
}

func ParseRoleInt(v int) AppRole {
	return roleMappingReverse[v]
}
