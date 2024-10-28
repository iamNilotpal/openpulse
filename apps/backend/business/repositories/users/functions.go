package users

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

func ToNewDBUser(cmd NewUser) users_store.NewUser {
	return users_store.NewUser{
		Email:     cmd.Email,
		RoleId:    cmd.RoleId,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
	}
}

func FromDBRole(cmd users_store.Role) Role {
	return Role{
		Id:           cmd.Id,
		Name:         cmd.Name,
		Description:  cmd.Description,
		IsSystemRole: cmd.IsSystemRole,
		Role:         roles.ParseRoleInt(cmd.Role),
	}
}

func FromDBResource(cmd users_store.Resource) Resource {
	return Resource{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Description: cmd.Description,
		Resource:    resources.ParseAppResourceInt(cmd.Resource),
	}
}

func FromDBPermission(cmd users_store.Permission) Permission {
	return Permission{
		Id:          cmd.Id,
		Name:        cmd.Name,
		Enabled:     cmd.Enabled,
		Description: cmd.Description,
		Action:      permissions.ParseActionInt(cmd.Action),
	}
}

func FromDBResourceWithPermission(cmd users_store.ResourcePermission) AccessControl {
	return AccessControl{
		Resource:   FromDBResource(cmd.Resource),
		Permission: FromDBPermission(cmd.Permission),
	}
}

func FromDBOAuthAccount(cmd users_store.OAuthAccount) OAuthAccount {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	return OAuthAccount{
		Id:         cmd.Id,
		Scope:      cmd.Scope,
		Provider:   cmd.Provider,
		ExternalId: cmd.ExternalId,
		Metadata:   cmd.Metadata.String,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

func FromDBUser(cmd users_store.User) User {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	resources := make([]AccessControl, 0, len(cmd.Resources))
	for i, r := range cmd.Resources {
		resources[i] = FromDBResourceWithPermission(r)
	}

	oauthAccounts := make([]OAuthAccount, 0)
	for i, ac := range cmd.OAuthAccounts {
		oauthAccounts[i] = FromDBOAuthAccount(ac)
	}

	return User{
		Id:              cmd.Id,
		Email:           cmd.Email,
		LastName:        cmd.LastName,
		FirstName:       cmd.FirstName,
		OAuthAccounts:   oauthAccounts,
		Phone:           cmd.Phone.String,
		IsEmailVerified: cmd.IsEmailVerified,
		AvatarUrl:       cmd.AvatarUrl.String,
		CountryCode:     cmd.CountryCode.String,
		Designation:     cmd.Designation.String,
		AccountStatus:   ParseStatusInt(cmd.AccountStatus),
		AccessControl:   resources,
		Role:            FromDBRole(cmd.Role),
		Team: Team{
			Name:    cmd.Team.Name.String,
			Id:      int(cmd.Team.Id.Int64),
			LogoURL: cmd.Team.LogoURL.String,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToNewDBOrganization(cmd NewOrganization) users_store.NewOrganization {
	return users_store.NewOrganization{
		AdminId:        cmd.AdminId,
		Name:           cmd.Name,
		Description:    cmd.Description,
		LogoURL:        cmd.LogoURL,
		TotalEmployees: cmd.TotalEmployees,
		Designation:    cmd.Designation,
	}
}

func ToNewDBTeam(cmd NewTeam) users_store.NewTeam {
	tu := make([]teams_store.UserAccessControl, len(cmd.UserAccessControl))

	for i, t := range cmd.UserAccessControl {
		tu[i] = teams_store.UserAccessControl{
			RoleId:       t.RoleId,
			UserId:       t.UserId,
			ResourceId:   t.ResourceId,
			PermissionId: t.PermissionId,
		}
	}

	return users_store.NewTeam{
		UserAccessControl: tu,
		Name:              cmd.Name,
		OrgId:             cmd.OrgId,
		CreatorId:         cmd.CreatorId,
		Description:       cmd.Description,
		CreatorRoleId:     cmd.CreatorRoleId,
		InvitationCode:    cmd.InvitationCode,
	}
}

func ToNewDBOauthUser(cmd NewOAuthUser) users_store.NewOAuthUser {
	return users_store.NewOAuthUser{
		RoleId:    cmd.RoleId,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Phone:     cmd.Phone,
		AvatarURL: cmd.AvatarURL,
	}
}

func ToNewDBOauthAccount(cmd NewOAuthAccount) users_store.NewOAuthAccount {
	return users_store.NewOAuthAccount{
		Scope:      cmd.Scope,
		Metadata:   cmd.Metadata,
		Provider:   cmd.Provider,
		ExternalId: cmd.ExternalId,
		User:       ToNewDBOauthUser(cmd.User),
	}
}
