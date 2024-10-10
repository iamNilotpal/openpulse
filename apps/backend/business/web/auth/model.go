package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type Claims struct {
	TeamId       int
	UserId       int
	RoleId       int
	SessionToken string
	jwt.RegisteredClaims
}

type UserAccessControlMap map[resources.AppResource][]UserPermissionConfig

type UserRoleConfig struct {
	Id   int
	Role roles.AppRole
}

type UserPermissionConfig struct {
	Id      int
	Enabled bool
	Action  permissions.PermissionAction
}

type UserResourceConfig struct {
	Id       int
	Resource resources.AppResource
}

type UserAccessControlPolicy struct {
	Role       UserRoleConfig
	Resource   UserResourceConfig
	Permission UserPermissionConfig
}

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

type RoleConfigMap map[roles.AppRole]RoleConfig
type ResourcePermissionsMap map[resources.AppResource][]PermissionConfig
type RoleAccessControlMap map[roles.AppRole]ResourcePermissionsMap

type RoleConfig struct {
	Id   int
	Role roles.AppRole
}

type PermissionConfig struct {
	Id     int
	Action permissions.PermissionAction
}

type ResourceConfig struct {
	Id       int
	Resource resources.AppResource
}

type AccessControlPolicy struct {
	Role       RoleConfig
	Resource   ResourceConfig
	Permission PermissionConfig
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
