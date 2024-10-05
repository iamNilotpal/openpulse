package permissions_store

import (
	modified_by "github.com/iamNilotpal/openpulse/business/data/modified-by"
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

type PermissionAccessConfig struct {
	Id     int
	Action string
}
