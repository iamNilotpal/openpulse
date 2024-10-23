package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
)

type AccessTokenClaims struct {
	TeamId int
	RoleId int
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}

type OnBoardingClaims struct {
	jwt.RegisteredClaims
}

/* ======================== User Access Controls  ======================== */
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

/* ======================== App Access Controls  ======================== */
type ResourceToPermissionsMap map[resources.AppResource]ResourcePermConfig
type RoleNameToAccessControlMap map[roles.AppRole]ResourceToPermissionsMap

type RoleIDMap map[int]RoleConfig
type RoleNameMap map[roles.AppRole]RoleConfig

type ResourceTypeIdMap map[int]ResourceConfig
type ResourceTypeMap map[resources.AppResource]ResourceConfig

type PermissionActionIdMap map[int]PermissionConfig
type PermissionActionMap map[permissions.PermissionAction]PermissionConfig

type RoleMappings struct {
	ByID   RoleIDMap
	ByName RoleNameMap
}

type PermissionMappings struct {
	ByAction PermissionActionMap
	ByID     PermissionActionIdMap
}

type ResourceMappings struct {
	ByID   ResourceTypeIdMap
	ByName ResourceTypeMap
}

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

type ResourcePermConfig struct {
	Resource    ResourceConfig
	Permissions []PermissionConfig
}
