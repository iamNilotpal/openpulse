package roles

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

func ToNewDBRole(cmd NewRole) roles_store.NewRole {
	return roles_store.NewRole{
		Name:         cmd.Name,
		Description:  cmd.Description,
		IsSystemRole: cmd.IsSystemRole,
		Role:         ParseRole(cmd.Role),
	}
}

func FromDBRole(cmd roles_store.Role) Role {
	createdAt, _ := time.Parse(time.UnixDate, cmd.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, cmd.UpdatedAt)

	return Role{
		Id:           cmd.Id,
		Name:         cmd.Name,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Description:  cmd.Description,
		IsSystemRole: cmd.IsSystemRole,
		Role:         ParseRoleInt(cmd.Role),
	}
}

func FromDBRoleAccessConfig(cmd roles_store.RoleAccessConfig) RoleAccessConfig {
	return RoleAccessConfig{Id: cmd.Id, Role: ParseRoleInt(cmd.Role)}
}

func FromDBRoleAccessControl(cmd roles_store.AccessControl) RoleAccessControl {
	return RoleAccessControl{
		Role:       FromDBRoleAccessConfig(cmd.Role),
		Resource:   resources.FromDBResourceAccessDetails(cmd.Resource),
		Permission: permissions.FromDBPermissionAccessDetails(cmd.Permission),
	}
}
