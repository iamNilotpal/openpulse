package users

import (
	"time"

	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
)

type NewUser struct {
	RoleId       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
}

type User struct {
	Role          Role
	Id            int
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	Password      string
	AvatarUrl     string
	Designation   string
	IsVerified    bool
	AccountStatus AccountStatus
	Team          Team
	Resources     []ResourcePermission
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Team struct {
	Id      int
	Name    string
	LogoURL string
}

type Role struct {
	Id           int
	IsSystemRole bool
	Name         string
	Description  string
	Role         roles.AppRole
}

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    resources.AppResource
}

type Permission struct {
	Id          int
	Enabled     bool
	Name        string
	Description string
	Action      permissions.PermissionAction
}

type ResourcePermission struct {
	Resource   Resource
	Permission Permission
}

type NewOrganization struct {
	AdminId        int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
	Designation    string
}

type NewTeam struct {
	CreatorId      int
	CreatorRoleId  int
	OrgId          int
	Name           string
	Description    string
	InvitationCode string
	UserRBAC       []UserRBAC
}

type UserRBAC struct {
	RoleId       int
	UserId       int
	ResourceId   int
	PermissionId int
}
