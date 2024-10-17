package auth

import (
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
)

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

/* ========================================================= */

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
