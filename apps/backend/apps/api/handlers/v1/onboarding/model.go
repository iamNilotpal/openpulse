package onboarding_handlers

import "github.com/iamNilotpal/openpulse/foundation/validate"

type CreateOrganizationInput struct {
	Name         string `json:"name" validate:"required,min=1"`
	Description  string `json:"description" validate:"required,min=50,max=255"`
	LogoURL      string `json:"logoURL" validate:"url"`
	Designation  string `json:"designation" validate:"required,min=1"`
	MembersCount string `json:"membersCount" validate:"required,min=1"`
}

func (v CreateOrganizationInput) Validate() error {
	return validate.Check(v)
}

type CreateOrganizationResponse struct {
	OrgId int `json:"orgId"`
}

type CreateTeamInput struct {
	OrgId           int    `json:"orgId" validate:"required,number"`
	TeamName        string `json:"teamName" validate:"required,min=1"`
	TeamLogoURL     string `json:"teamLogoURL" validate:"http_url"`
	TeamDescription string `json:"teamDescription" validate:"required,min=1,max=200"`
}

func (v CreateTeamInput) Validate() error {
	return validate.Check(v)
}

type CreateTeamResponse struct {
	TeamId int `json:"teamId"`
}
