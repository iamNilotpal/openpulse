package roles_store

import (
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Role         int
	Name         string
	Description  string
	CreatedAt    string
	UpdatedAt    string
}

// type RoleResources struct {
// 	Resource   resources_store.Resource
// 	Permission permissions_store.Permission
// }

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

type RoleAccessControl struct {
	Role       RoleAccessConfig
	Resource   resources_store.ResourceAccessConfig
	Permission permissions_store.PermissionAccessConfig
}
