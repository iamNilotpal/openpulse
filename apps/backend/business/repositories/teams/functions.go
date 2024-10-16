package teams

import teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"

func FromDBTeam(cmd teams_store.Team) Team {
	return Team{
		Id:             cmd.Id,
		Name:           cmd.Name,
		Description:    cmd.Description,
		TotalMembers:   cmd.TotalMembers,
		InvitationCode: cmd.InvitationCode,
		Organization: Organization{
			Id:      cmd.Organization.Id,
			Name:    cmd.Organization.Name,
			LogoURL: cmd.Organization.LogoURL,
		},
		Creator: Creator{
			Id:        cmd.Creator.Id,
			Email:     cmd.Creator.Email,
			FirstName: cmd.Creator.FirstName,
			LastName:  cmd.Creator.LastName,
		},
	}
}
