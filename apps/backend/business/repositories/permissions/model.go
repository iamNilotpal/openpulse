package permissions

import (
	"time"

	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type NewPermission struct {
	Name        string
	Description string
	Action      PermissionAction
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      PermissionAction
	CreateAt    time.Time
	UpdatedAt   time.Time
}

type PermissionAccessConfig struct {
	Id     int
	Action PermissionAction
}

func NewDBPermission(p NewPermission) permissions_store.NewPermission {
	return permissions_store.NewPermission{
		Name:        p.Name,
		Description: p.Description,
		Action:      ParseAction(p.Action),
	}
}

func FromDBPermission(p permissions_store.Permission) Permission {
	createdAt, _ := time.Parse("", p.CreatedAt)
	updatedAt, _ := time.Parse("", p.UpdatedAt)

	return Permission{
		Id:          p.Id,
		Name:        p.Name,
		CreateAt:    createdAt,
		UpdatedAt:   updatedAt,
		Description: p.Description,
		Action:      ParseActionInt(p.Action),
	}
}

func FromDBPermissionAccessDetails(r permissions_store.PermissionAccessConfig) PermissionAccessConfig {
	return PermissionAccessConfig{Id: r.Id, Action: ParseActionInt(r.Action)}
}
