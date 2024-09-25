package roles_store

import permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"

type DBRole struct {
	Id          int
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type NewDBRole struct {
	Name        string
	Description string
}

type DBRolePermissions struct {
	Role       DBRole
	Permission permissions_store.DBPermission
}
