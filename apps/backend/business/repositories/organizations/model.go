package organizations

import (
	"time"

	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
)

type NewOrganization struct {
	AdminId        int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
}

type Organization struct {
	Id             int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
	Admin          OrgAdmin
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrgAdmin struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

func ToNewDBOrg(org NewOrganization) organizations_store.NewOrganization {
	return organizations_store.NewOrganization{
		Name:           org.Name,
		LogoURL:        org.LogoURL,
		AdminId:        org.AdminId,
		Description:    org.Description,
		TotalEmployees: org.TotalEmployees,
	}
}
