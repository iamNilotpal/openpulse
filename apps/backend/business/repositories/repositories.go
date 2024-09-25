package repositories

import (
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/teams"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type Repositories struct {
	Teams       teams.Repository
	Users       users.Repository
	Roles       roles.Repository
	Permissions permissions.Repository
}
