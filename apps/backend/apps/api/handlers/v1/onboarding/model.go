package onboarding_handlers

import "github.com/iamNilotpal/openpulse/business/sys/validate"

type CreateOrganizationInput struct {
	Name         string `json:"orgName" validate:"required,min=1"`
	Description  string `json:"orgDescription" validate:"required,min=50,max=255"`
	LogoURL      string `json:"orgLogoURL"`
	Designation  string `json:"designation" validate:"required,min=1"`
	MembersCount string `json:"membersCount" validate:"required,min=1"`
}

func (oop CreateOrganizationInput) Validate() error {
	return validate.Check(oop)
}

type CreateOrganizationResponse struct {
	OrgId int `json:"orgId"`
}

type CreateTeamInput struct {
	OrgId           int    `json:"orgId"`
	UserId          int    `json:"userId"`
	TeamName        string `json:"teamName"`
	TeamLogoURL     string `json:"teamLogoURL"`
	TeamDescription string `json:"teamDescription"`
}

type CreateTeamResponse struct {
	TeamId int `json:"teamId"`
}
