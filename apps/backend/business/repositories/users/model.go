package users

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type NewUser struct {
	RoleId       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
}

type User struct {
	ID            int
	Role          Role
	FirstName     string
	LastName      string
	Email         string
	AvatarUrl     string
	AccountStatus AccountStatus
	Resources     []ResourcePermission
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	Role         roles.AppRole
}

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    resources.AppResource
}

type Permission struct {
	Id          int
	Enabled     bool
	Name        string
	Description string
	Action      permissions.PermissionAction
}

type ResourcePermission struct {
	Resource   Resource
	Permission Permission
}

func ToNewDBUser(p NewUser) users_store.NewUser {
	return users_store.NewUser{
		Email:        p.Email,
		RoleId:       p.RoleId,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		PasswordHash: p.PasswordHash,
	}
}

func FromDBRole(r users_store.Role) Role {
	return Role{
		Id:           r.Id,
		Name:         r.Name,
		Description:  r.Description,
		IsSystemRole: r.IsSystemRole,
		Role:         roles.ParseRoleInt(r.Role),
	}
}

func FromDBResource(r users_store.Resource) Resource {
	return Resource{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Resource:    resources.ParseAppResourceInt(r.Resource),
	}
}

func FromDBPermission(p users_store.Permission) Permission {
	return Permission{
		Id:          p.Id,
		Name:        p.Name,
		Enabled:     p.Enabled,
		Description: p.Description,
		Action:      permissions.ParseActionInt(p.Action),
	}
}

func FromDBResourceWithPermission(r users_store.ResourcePermission) ResourcePermission {
	return ResourcePermission{
		Resource:   FromDBResource(r.Resource),
		Permission: FromDBPermission(r.Permission),
	}
}

func FromDBUser(u users_store.User) User {
	createdAt, _ := time.Parse(time.UnixDate, u.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, u.UpdatedAt)

	resources := make([]ResourcePermission, 0, len(u.Resources))
	for i, r := range u.Resources {
		resources[i] = FromDBResourceWithPermission(r)
	}

	return User{
		ID:            u.Id,
		Email:         u.Email,
		LastName:      u.LastName,
		FirstName:     u.FirstName,
		AvatarUrl:     u.AvatarUrl,
		AccountStatus: ParseStatusInt(u.AccountStatus),
		Resources:     resources,
		Role:          FromDBRole(u.Role),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
