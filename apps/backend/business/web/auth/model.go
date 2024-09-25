package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
)

type PermissionsMap map[string][]Permissions

type Claims struct {
	RoleId    int
	SessionId int
	jwt.RegisteredClaims
}

type UserRole struct {
	Id int
}

type UserPermission struct {
	Id       int
	Enabled  bool
	Action   string
	Resource string
}

type Role struct {
	Id int
}

type Permission struct {
	Id       int
	Action   string
	Resource string
}

type Permissions struct {
	Role
	Permission
}
type UserPermissions struct {
	Role       UserRole
	Permission UserPermission
}

func ToRole(role roles.Role) Role {
	return Role{Id: role.Id}
}

func ToPermission(p permissions.Permission) Permission {
	return Permission{Id: p.Id, Action: p.Action, Resource: p.Resource}
}

func ToUserRole(role roles.Role) UserRole {
	return UserRole{Id: role.Id}
}

func ToUserPermission(p permissions.UserPermission) UserPermission {
	return UserPermission{Id: p.Id, Action: p.Action, Resource: p.Resource, Enabled: p.Enabled}
}

func ToPermissions(role Role, permission Permission) Permissions {
	return Permissions{
		Role: Role{Id: role.Id},
		Permission: Permission{
			Id:       permission.Id,
			Action:   permission.Action,
			Resource: permission.Resource,
		},
	}
}

func ToUserPermissions(role UserRole, permission UserPermission) UserPermissions {
	return UserPermissions{
		Role: UserRole{Id: role.Id},
		Permission: UserPermission{
			Id:       permission.Id,
			Action:   permission.Action,
			Enabled:  permission.Enabled,
			Resource: permission.Resource,
		},
	}
}
