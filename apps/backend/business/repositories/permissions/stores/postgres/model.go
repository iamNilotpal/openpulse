package permissions_store

import (
	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      string
	CreatedBy   modified_by.ModifiedBy
	UpdatedBy   modified_by.ModifiedBy
	CreatedAt   string
	UpdatedAt   string
}

type NewPermission struct {
	CreatorId   int
	Name        string
	Description string
	Action      string
}

type PermissionWithResource struct {
	Permission Permission
	Resource   resources_store.Resource
	UpdatedBy  modified_by.ModifiedBy
	CreatedBy  modified_by.ModifiedBy
	CreatedAt  string
	UpdatedAt  string
}
