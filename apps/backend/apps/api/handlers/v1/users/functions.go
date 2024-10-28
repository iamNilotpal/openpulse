package users_handler

import (
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

func toRole(cmd users.Role) Role {
	return Role{
		Id:           cmd.Id,
		Name:         cmd.Name,
		Description:  cmd.Description,
		IsSystemRole: cmd.IsSystemRole,
		Role:         string(cmd.Role),
	}
}

func toTeam(cmd users.Team) *Team {
	if cmd.Id == 0 || cmd.Name == "" {
		return nil
	}
	return &Team{Id: cmd.Id, Name: cmd.Name, LogoURL: cmd.LogoURL}
}

func toResource(cmd users.Resource) Resource {
	return Resource{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Description: cmd.Description,
		Resource:    string(cmd.Resource),
	}
}

func toPermission(cmd users.Permission) Permission {
	return Permission{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Enabled:     cmd.Enabled,
		Description: cmd.Description,
		Action:      string(cmd.Action),
	}
}

func toAccessControl(cmd users.AccessControl) AccessControl {
	return AccessControl{
		Resource:   toResource(cmd.Resource),
		Permission: toPermission(cmd.Permission),
	}
}

func toOAuthAccount(cmd users.OAuthAccount) OAuthAccount {
	return OAuthAccount{
		Id:         cmd.Id,
		Scope:      cmd.Scope,
		Metadata:   cmd.Metadata,
		Provider:   cmd.Provider,
		ExternalId: cmd.ExternalId,
		CreatedAt:  cmd.CreatedAt.String(),
		UpdatedAt:  cmd.CreatedAt.String(),
	}
}

func toUser(cmd users.User) User {
	resources := make([]AccessControl, 0, len(cmd.AccessControl))
	for i, r := range cmd.AccessControl {
		resources[i] = toAccessControl(r)
	}

	oauthAccounts := make([]OAuthAccount, 0)
	for i, ac := range cmd.OAuthAccounts {
		oauthAccounts[i] = toOAuthAccount(ac)
	}

	return User{
		Id:              cmd.Id,
		Email:           cmd.Email,
		LastName:        cmd.LastName,
		FirstName:       cmd.FirstName,
		Phone:           cmd.Phone,
		AvatarUrl:       cmd.AvatarUrl,
		Designation:     cmd.Designation,
		IsEmailVerified: cmd.IsEmailVerified,
		AccountStatus:   string(cmd.AccountStatus),
		AccessControl:   resources,
		OAuthAccounts:   oauthAccounts,
		Team:            toTeam(cmd.Team),
		Role:            toRole(cmd.Role),
		CreatedAt:       cmd.CreatedAt.String(),
		UpdatedAt:       cmd.UpdatedAt.String(),
	}
}
