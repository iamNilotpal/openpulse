package users_handler

import (
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

func FromAppRole(cmd users.Role) Role {
	return Role{
		Id:           cmd.Id,
		Name:         cmd.Name,
		Description:  cmd.Description,
		IsSystemRole: cmd.IsSystemRole,
		Role:         string(cmd.Role),
	}
}

func FromAppResource(cmd users.Resource) Resource {
	return Resource{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Description: cmd.Description,
		Resource:    string(cmd.Resource),
	}
}

func FromAppPermission(cmd users.Permission) Permission {
	return Permission{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Enabled:     cmd.Enabled,
		Description: cmd.Description,
		Action:      string(cmd.Action),
	}
}

func FromAppResourceWithPermission(cmd users.ResourcePermission) ResourcePermission {
	return ResourcePermission{
		Resource:   FromAppResource(cmd.Resource),
		Permission: FromAppPermission(cmd.Permission),
	}
}

func FromAppOAuthAccount(cmd users.OAuthAccount) OAuthAccount {
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

func FromAppUser(cmd users.User) User {
	resources := make([]ResourcePermission, 0, len(cmd.Resources))
	for i, r := range cmd.Resources {
		resources[i] = FromAppResourceWithPermission(r)
	}

	oauthAccounts := make([]OAuthAccount, 0)
	for i, ac := range cmd.OAuthAccounts {
		oauthAccounts[i] = FromAppOAuthAccount(ac)
	}

	return User{
		Id:                  cmd.Id,
		Email:               cmd.Email,
		LastName:            cmd.LastName,
		FirstName:           cmd.FirstName,
		AvatarUrl:           cmd.AvatarUrl,
		Phone:               cmd.Phone,
		Designation:         cmd.Designation,
		IsEmailVerified:     cmd.IsEmailVerified,
		AccountStatus:       string(cmd.AccountStatus),
		ResourcePermissions: resources,
		Role:                FromAppRole(cmd.Role),
		OAuthAccounts:       oauthAccounts,
		Team: Team{
			Id:      cmd.Team.Id,
			Name:    cmd.Team.Name,
			LogoURL: cmd.Team.LogoURL,
		},
		CreatedAt: cmd.CreatedAt.String(),
		UpdatedAt: cmd.UpdatedAt.String(),
	}
}
