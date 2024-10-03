package permissions

import (
	"time"

	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type NewPermission struct {
	CreatorId   int
	Name        string
	Description string
	Action      PermissionAction
}

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      PermissionAction
	CreatedBy   modified_by.ModifiedBy
	UpdatedBy   modified_by.ModifiedBy
	CreateAt    time.Time
	UpdatedAt   time.Time
}

func NewDBPermission(p NewPermission) permissions_store.NewPermission {
	return permissions_store.NewPermission{
		Name:        p.Name,
		CreatorId:   p.CreatorId,
		Description: p.Description,
		Action:      FromPermissionAction(p.Action),
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
		Action:      ToPermissionAction(p.Action),
		CreatedBy: modified_by.New(
			p.CreatedBy.Id, p.CreatedBy.Email, p.CreatedBy.FirstName, p.CreatedBy.LastName,
		),
		UpdatedBy: modified_by.New(
			p.UpdatedBy.Id, p.UpdatedBy.Email, p.UpdatedBy.FirstName, p.UpdatedBy.LastName,
		),
	}
}
