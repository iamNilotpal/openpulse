package roles_store

import (
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Role struct {
	Id           int
	Name         string
	Description  string
	Role         int
	IsSystemRole bool
	CreatedAt    string
	UpdatedAt    string
}

type NewRole struct {
	IsSystemRole bool
	Role         int
	Name         string
	Description  string
}

type RoleAccessConfig struct {
	Id   int
	Role int
}

type AccessControl struct {
	Role       RoleAccessConfig
	Resource   resources_store.ResourceAccessConfig
	Permission permissions_store.PermissionAccessConfig
}
