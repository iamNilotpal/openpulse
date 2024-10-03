package teams

import (
	"time"

	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
)

type Team struct {
	Id           int
	Name         string
	Description  string
	TotalMembers int
	AdminId      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type NewTeam struct {
	AdminId     int
	Name        string
	Description string
}

func ToNewDBTeam(t NewTeam) teams_store.NewTeam {
	return teams_store.NewTeam{
		Name:        t.Name,
		AdminId:     t.AdminId,
		Description: t.Description,
	}
}

func FromDBTeam(t teams_store.Team) Team {
	return Team{
		Id:           t.Id,
		Name:         t.Name,
		AdminId:      t.AdminId,
		Description:  t.Description,
		TotalMembers: t.TotalMembers,
	}
}
