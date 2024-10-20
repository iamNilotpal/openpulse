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

/* ========================================================= */

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
type ResourcePermsMap map[resources.AppResource]ResourcePermConfig
type RBACMap map[roles.AppRole]ResourcePermsMap

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

type AccessControlPolicy struct {
	Role       RoleConfig
	Resource   ResourceConfig
	Permission PermissionConfig
}
