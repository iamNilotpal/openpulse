package roles

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
)

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	Role         AppRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RoleAccessConfig struct {
	Id   int
	Role AppRole
}

type RoleAccessControl struct {
	Role       RoleAccessConfig
	Resource   resources.ResourceAccessConfig
	Permission permissions.PermissionAccessConfig
}

type NewRole struct {
	IsSystemRole bool
	Role         AppRole
	Name         string
	Description  string
}

type UpdateRole struct {
	Name        string
	Description string
}
