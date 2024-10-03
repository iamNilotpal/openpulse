package roles_store

import (
	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	CreatedBy    modified_by.ModifiedBy
	UpdatedBy    modified_by.ModifiedBy
	CreatedAt    string
	UpdatedAt    string
}

type NewRole struct {
	CreatorId    int
	IsSystemRole bool
	Name         string
	Description  string
}

type RoleWithPermission struct {
	Role       Role
	Permission permissions_store.Permission
	CreatedAt  string
	UpdatedAt  string
}
