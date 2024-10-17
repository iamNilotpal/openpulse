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
		Email:        cmd.Email,
		RoleId:       cmd.RoleId,
		FirstName:    cmd.FirstName,
		LastName:     cmd.LastName,
		PasswordHash: cmd.PasswordHash,
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

func FromDBResourceWithPermission(cmd users_store.ResourcePermission) ResourcePermission {
	return ResourcePermission{
		Resource:   FromDBResource(cmd.Resource),
		Permission: FromDBPermission(cmd.Permission),
	}
}

func FromDBUser(cmd users_store.User) User {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	resources := make([]ResourcePermission, 0, len(cmd.Resources))
	for i, r := range cmd.Resources {
		resources[i] = FromDBResourceWithPermission(r)
	}

	return User{
		Id:            cmd.Id,
		Email:         cmd.Email,
		LastName:      cmd.LastName,
		FirstName:     cmd.FirstName,
		AvatarUrl:     cmd.AvatarUrl,
		Password:      cmd.Password,
		AccountStatus: ParseStatusInt(cmd.AccountStatus),
		Resources:     resources,
		Team:          Team{Id: cmd.Team.Id, Name: cmd.Team.Name, LogoURL: cmd.Team.LogoURL},
		Role:          FromDBRole(cmd.Role),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
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
	tu := make([]teams_store.UserRBAC, len(cmd.UserRBAC))

	for i, t := range cmd.UserRBAC {
		tu[i] = teams_store.UserRBAC{
			RoleId:       t.RoleId,
			UserId:       t.UserId,
			ResourceId:   t.ResourceId,
			PermissionId: t.PermissionId,
		}
	}

	return users_store.NewTeam{
		UserRBAC:       tu,
		Name:           cmd.Name,
		OrgId:          cmd.OrgId,
		CreatorId:      cmd.CreatorId,
		Description:    cmd.Description,
		CreatorRoleId:  cmd.CreatorRoleId,
		InvitationCode: cmd.InvitationCode,
	}
}
