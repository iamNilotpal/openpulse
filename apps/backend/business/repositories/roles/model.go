package roles

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
)

type Role struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RolePermissions struct {
	Role       Role
	Permission permissions.Permission
}

type NewRole struct {
	Name        string
	Description string
}

type UpdateRole struct {
	Name        string
	Description string
}

func ToNewRole(name, description string) NewRole {
	return NewRole{Name: name, Description: description}
}

func ToNewDBRole(r NewRole) roles_store.NewDBRole {
	return roles_store.NewDBRole{Name: r.Name, Description: r.Description}
}

func ToRole(r roles_store.DBRole) Role {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Role{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func ToRoleWithPermissions(r roles_store.DBRolePermissions) RolePermissions {
	return RolePermissions{
		Role:       ToRole(r.Role),
		Permission: permissions.ToPermission(r.Permission),
	}
}
