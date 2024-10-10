package roles

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	Role         AppRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RoleAccessConfig struct {
	Id   int
	Role AppRole
}

type RoleAccessControl struct {
	Role       RoleAccessConfig
	Resource   resources.ResourceAccessConfig
	Permission permissions.PermissionAccessConfig
}

type NewRole struct {
	IsSystemRole bool
	Role         AppRole
	Name         string
	Description  string
}

type UpdateRole struct {
	Name        string
	Description string
}

func ToNewDBRole(r NewRole) roles_store.NewRole {
	return roles_store.NewRole{
		Name:         r.Name,
		Role:         ParseRole(r.Role),
		Description:  r.Description,
		IsSystemRole: r.IsSystemRole,
	}
}

func FromDBRole(r roles_store.Role) Role {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Role{
		Id:           r.Id,
		Name:         r.Name,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Description:  r.Description,
		IsSystemRole: r.IsSystemRole,
		Role:         ParseRoleInt(r.Role),
	}
}

func FromDBRoleAccessConfig(r roles_store.RoleAccessConfig) RoleAccessConfig {
	return RoleAccessConfig{Id: r.Id, Role: ParseRoleInt(r.Role)}
}

func FromDBRoleAccessControl(r roles_store.RoleAccessControl) RoleAccessControl {
	return RoleAccessControl{
		Role:       FromDBRoleAccessConfig(r.Role),
		Resource:   resources.FromDBResourceAccessDetails(r.Resource),
		Permission: permissions.FromDBPermissionAccessDetails(r.Permission),
	}
}
