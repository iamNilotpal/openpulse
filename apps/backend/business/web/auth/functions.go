package auth

import (
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

func NewUserRoleConfig(role users.Role) UserRoleConfig {
	return UserRoleConfig{Id: role.Id, Role: role.Role}
}

func NewUserPermissionConfig(p users.Permission) UserPermissionConfig {
	return UserPermissionConfig{Id: p.Id, Action: p.Action, Enabled: p.Enabled}
}

func NewUserResourceConfig(r users.Resource) UserResourceConfig {
	return UserResourceConfig{Id: r.Id, Resource: r.Resource}
}

func NewUserAccessControlPolicy(
	ur UserRoleConfig, up UserPermissionConfig, res UserResourceConfig,
) UserAccessControlPolicy {
	return UserAccessControlPolicy{
		Role:       UserRoleConfig{Id: ur.Id, Role: ur.Role},
		Resource:   UserResourceConfig{Id: res.Id, Resource: res.Resource},
		Permission: UserPermissionConfig{Id: up.Id, Action: up.Action, Enabled: up.Enabled},
	}
}

func NewRoleConfig(role roles.RoleAccessConfig) RoleConfig {
	return RoleConfig{Id: role.Id, Role: role.Role}
}

func NewPermissionConfig(p permissions.PermissionAccessConfig) PermissionConfig {
	return PermissionConfig{Id: p.Id, Action: p.Action}
}

func NewResourceConfig(r resources.ResourceAccessConfig) ResourceConfig {
	return ResourceConfig{Id: r.Id, Resource: r.Resource}
}

func NewAccessControlPolicy(
	role RoleConfig, resource ResourceConfig, permission PermissionConfig,
) AccessControlPolicy {
	return AccessControlPolicy{
		Role:       RoleConfig{Id: role.Id, Role: role.Role},
		Resource:   ResourceConfig{Id: resource.Id, Resource: resource.Resource},
		Permission: PermissionConfig{Id: permission.Id, Action: permission.Action},
	}
}
