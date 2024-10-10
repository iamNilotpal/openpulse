package onboarding_handlers

type OnboardingOrganizationPayload struct {
	UserId         int    `json:"userId"`
	OrgName        string `json:"orgName"`
	OrgDescription string `json:"orgDescription"`
	OrgLogoURL     string `json:"orgLogoURL"`
	Designation    string `json:"designation"`
	MembersCount   string `json:"membersCount"`
}

type OnboardingOrganizationResponse struct {
	OrgId int `json:"orgId"`
}

type OnboardingTeamPayload struct {
	OrgId           int    `json:"orgId"`
	UserId          int    `json:"userId"`
	TeamName        string `json:"teamName"`
	TeamLogoURL     string `json:"teamLogoURL"`
	TeamDescription string `json:"teamDescription"`
}

type OnboardingTeamResponse struct {
	TeamId int `json:"teamId"`
}
