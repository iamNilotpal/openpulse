package repositories

import (
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type Repositories struct {
	User  *users.Repository
	Roles *roles.Repository
}
