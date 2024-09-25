package permissions

import (
	"time"

	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type NewPermission struct {
	Name        string
	Description string
	Action      string
	Resource    string
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      string
	Resource    string
	CreateAt    time.Time
	UpdatedAt   time.Time
}

type UserPermission struct {
	Id          int
	Enabled     bool
	Name        string
	Description string
	Action      string
	Resource    string
	CreateAt    string
	UpdatedAt   string
	UpdatedBy   string
}

func ToDBNewPermission(p NewPermission) permissions_store.DBNewPermission {
	return permissions_store.DBNewPermission{
		Name:        p.Name,
		Action:      p.Action,
		Resource:    p.Resource,
		Description: p.Description,
	}
}

func ToPermission(p permissions_store.DBPermission) Permission {
	createdAt, _ := time.Parse("", p.CreatedAt)
	updatedAt, _ := time.Parse("", p.UpdatedAt)

	return Permission{
		Id:          p.Id,
		Name:        p.Name,
		Action:      p.Action,
		CreateAt:    createdAt,
		UpdatedAt:   updatedAt,
		Resource:    p.Resource,
		Description: p.Description,
	}
}

func ToUserPermission(p permissions_store.DBUserPermission) UserPermission {
	return UserPermission{
		Id:          p.Id,
		Name:        p.Name,
		Action:      p.Action,
		Enabled:     p.Enabled,
		Resource:    p.Resource,
		CreateAt:    p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		UpdatedBy:   p.UpdatedBy,
		Description: p.Description,
	}
}
