package teams

import (
	"time"

	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
)

type Team struct {
	Id             int
	Name           string
	Description    string
	TotalMembers   int
	InvitationCode string
	Creator        Creator
	Organization   Organization
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Organization struct {
	Id      int
	Name    string
	LogoURL string
}

type Creator struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

type NewTeam struct {
	CreatorId      int
	CreatorRoleId  int
	OrgId          int
	Name           string
	Description    string
	InvitationCode string
	UserRBAC       []UserRBAC
}

type UserRBAC struct {
	RoleId       int
	UserId       int
	ResourceId   int
	PermissionId int
}

func ToNewDBTeam(t NewTeam) teams_store.NewTeam {
	tu := make([]teams_store.UserRBAC, len(t.UserRBAC))
	for i, t := range t.UserRBAC {
		tu[i] = teams_store.UserRBAC{
			RoleId:       t.RoleId,
			UserId:       t.UserId,
			ResourceId:   t.ResourceId,
			PermissionId: t.PermissionId,
		}
	}

	return teams_store.NewTeam{
		UserRBAC:       tu,
		Name:           t.Name,
		OrgId:          t.OrgId,
		CreatorId:      t.CreatorId,
		Description:    t.Description,
		CreatorRoleId:  t.CreatorRoleId,
		InvitationCode: t.InvitationCode,
	}
}

func FromDBTeam(t teams_store.Team) Team {
	return Team{
		Id:             t.Id,
		Name:           t.Name,
		Description:    t.Description,
		TotalMembers:   t.TotalMembers,
		InvitationCode: t.InvitationCode,
		Organization: Organization{
			Id:      t.Organization.Id,
			Name:    t.Organization.Name,
			LogoURL: t.Organization.LogoURL,
		},
		Creator: Creator{
			Id:        t.Creator.Id,
			Email:     t.Creator.Email,
			FirstName: t.Creator.FirstName,
			LastName:  t.Creator.LastName,
		},
	}
}
