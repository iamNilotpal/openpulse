package roles

import (
	"net/http"
	"slices"

	"github.com/iamNilotpal/openpulse/business/web/errors"
)

type AppRole string

const (
	RoleOrgAdminString         AppRole = "org_admin"
	RoleTeamAdminString        AppRole = "team_admin"
	RoleTeamBillingAdminString AppRole = "team_billing_admin"
	RoleTeamLeadString         AppRole = "team_lead"
	RoleTeamResponderString    AppRole = "team_responder"
	RoleTeamMemberString       AppRole = "team_member"
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
	RoleOrgAdminString,
	RoleTeamLeadString,
	RoleTeamAdminString,
	RoleTeamMemberString,
	RoleTeamResponderString,
	RoleTeamBillingAdminString,
}

var roleMapping = map[AppRole]int{
	RoleOrgAdminString:         RoleOrgAdminInt,
	RoleTeamAdminString:        RoleTeamAdminInt,
	RoleTeamBillingAdminString: RoleTeamBillingAdminInt,
	RoleTeamLeadString:         RoleTeamLeadInt,
	RoleTeamResponderString:    RoleTeamResponderInt,
	RoleTeamMemberString:       RoleTeamMemberInt,
}

var roleMappingReverse = map[int]AppRole{
	RoleOrgAdminInt:         RoleOrgAdminString,
	RoleTeamAdminInt:        RoleTeamAdminString,
	RoleTeamBillingAdminInt: RoleTeamBillingAdminString,
	RoleTeamLeadInt:         RoleTeamLeadString,
	RoleTeamResponderInt:    RoleTeamResponderString,
	RoleTeamMemberInt:       RoleTeamMemberString,
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
