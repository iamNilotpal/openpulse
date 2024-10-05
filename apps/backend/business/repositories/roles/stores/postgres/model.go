package roles_store

import (
	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Role         string
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
	Role         string
	Name         string
	Description  string
}

type RoleAccessConfig struct {
	Id   int
	Role string
}

type RoleAccessControl struct {
	Role       RoleAccessConfig
	Resource   resources_store.ResourceAccessConfig
	Permission permissions_store.PermissionAccessConfig
}
