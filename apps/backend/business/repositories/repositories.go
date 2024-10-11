package repositories

import (
	"github.com/iamNilotpal/openpulse/business/repositories/emails"
	"github.com/iamNilotpal/openpulse/business/repositories/organizations"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/teams"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type Repositories struct {
	Users         users.Repository
	Emails        emails.Repository
	Organizations organizations.Repository
	Teams         teams.Repository
	Roles         roles.Repository
	Resources     resources.Repository
	Permissions   permissions.Repository
}
