package permissions

import (
	"time"

	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

func NewDBPermission(cmd NewPermission) permissions_store.NewPermission {
	return permissions_store.NewPermission{
		Name:        cmd.Name,
		Description: cmd.Description,
		Action:      ParseAction(cmd.Action),
	}
}

func FromDBPermission(cmd permissions_store.Permission) Permission {
	createdAt, _ := time.Parse("", cmd.CreatedAt)
	updatedAt, _ := time.Parse("", cmd.UpdatedAt)

	return Permission{
		Id:          cmd.Id,
		Name:        cmd.Name,
		CreateAt:    createdAt,
		UpdatedAt:   updatedAt,
		Description: cmd.Description,
		Action:      ParseActionInt(cmd.Action),
	}
}

func FromDBPermissionAccessDetails(cmd permissions_store.PermissionAccessConfig) PermissionAccessConfig {
	return PermissionAccessConfig{Id: cmd.Id, Action: ParseActionInt(cmd.Action)}
}
