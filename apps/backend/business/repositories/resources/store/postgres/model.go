package resources_store

import (
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    int
	CreatedAt   string
	UpdatedAt   string
}

type NewResource struct {
	Resource    int
	Name        string
	Description string
}

type ResourceAccessConfig struct {
	Id       int
	Resource int
}

type ResourceWithPermission struct {
	Resource   ResourceAccessConfig
	Permission permissions_store.PermissionAccessConfig
}
