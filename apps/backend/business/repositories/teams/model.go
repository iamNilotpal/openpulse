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

func ToDBNewTeam(t NewTeam) teams_store.DBNewTeam {
	return teams_store.DBNewTeam{
		Name:        t.Name,
		AdminId:     t.AdminId,
		Description: t.Description,
	}
}

func ToTeam(t teams_store.DBTeam) Team {
	return Team{
		Id:           t.Id,
		Name:         t.Name,
		AdminId:      t.AdminId,
		Description:  t.Description,
		TotalMembers: t.TotalMembers,
	}
}
