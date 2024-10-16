package resources

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
)

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    AppResource
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewResource struct {
	Name        string
	Description string
	Resource    AppResource
}

type ResourceAccessConfig struct {
	Id       int
	Resource AppResource
}

type ResourceWithPermission struct {
	Resource   ResourceAccessConfig
	Permission permissions.PermissionAccessConfig
}
