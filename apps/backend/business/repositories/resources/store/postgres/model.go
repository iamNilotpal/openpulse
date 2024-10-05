package resources_store

import (
	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    string
	CreatedBy   modified_by.ModifiedBy
	UpdatedBy   modified_by.ModifiedBy
	CreatedAt   string
	UpdatedAt   string
}

type NewResource struct {
	CreatorId   int
	Name        string
	Resource    string
	Description string
}

type ResourceAccessConfig struct {
	Id       int
	Resource string
}

type ResourceWithPermission struct {
	Resource   ResourceAccessConfig
	Permission permissions_store.PermissionAccessConfig
}
