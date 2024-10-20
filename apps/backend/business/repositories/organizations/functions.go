package organizations

import (
	"time"

	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
)

func FromDBOrg(cmd organizations_store.Organization) Organization {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	return Organization{
		Id:             cmd.Id,
		Name:           cmd.Name,
		LogoURL:        cmd.LogoURL,
		Description:    cmd.Description,
		TotalEmployees: cmd.TotalEmployees,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Admin: OrgAdmin{
			Id:        cmd.Admin.Id,
			FirstName: cmd.Admin.FirstName,
			LastName:  cmd.Admin.LastName,
			Email:     cmd.Admin.Email,
		},
	}
}
