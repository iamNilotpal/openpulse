package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
)

type AuthedRolesMap map[string]AuthRole
type AuthedPermissionsMap map[string][]AuthAccessControl

type Claims struct {
	RoleId    int
	UserId    int
	SessionId int
	jwt.RegisteredClaims
}

type UserRole struct {
	Id int
}

type UserPermission struct {
	Id      int
	Enabled bool
	Action  permissions.PermissionAction
}

type UserResource struct {
	Id       int
	Resource resources.ResourceType
}

type UserAccessControl struct {
	Role       UserRole
	Resource   UserResource
	Permission UserPermission
}

type AuthRole struct {
	Id int
}

type AuthPermission struct {
	Id     int
	Action permissions.PermissionAction
}

type AuthResource struct {
	Id       int
	Resource resources.ResourceType
}

type AuthAccessControl struct {
	Role       AuthRole
	Resource   AuthResource
	Permission AuthPermission
}

func ToAuthedRole(role roles.Role) AuthRole {
	return AuthRole{Id: role.Id}
}

func ToAuthedPermission(p permissions.Permission) AuthPermission {
	return AuthPermission{Id: p.Id, Action: p.Action}
}

func ToAuthedUserRole(role roles.Role) UserRole {
	return UserRole{Id: role.Id}
}

func ToAuthedPermissions(role AuthRole, permission AuthPermission) AuthAccessControl {
	return AuthAccessControl{
		Role: AuthRole{Id: role.Id},
		Permission: AuthPermission{
			Id:     permission.Id,
			Action: permission.Action,
		},
	}
}

func ToAuthedUserPermissions(role UserRole, permission UserPermission) UserAccessControl {
	return UserAccessControl{
		Role: UserRole{Id: role.Id},
		Permission: UserPermission{
			Id:      permission.Id,
			Action:  permission.Action,
			Enabled: permission.Enabled,
		},
	}
}
