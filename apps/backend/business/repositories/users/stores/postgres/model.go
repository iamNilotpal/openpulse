package users_store

import teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"

type NewUser struct {
	RoleId       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
}

type User struct {
	Id            int
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	Password      string
	AvatarUrl     string
	AccountStatus int
	Designation   string
	IsVerified    bool
	Role          Role
	Team          Team
	Preference    Preference
	Resources     []ResourcePermission
	CreatedAt     string
	UpdatedAt     string
}

type Team struct {
	Id      int
	Name    string
	LogoURL string
}

type Preference struct {
	Id         int
	Appearance string
	CreatedAt  string
	UpdatedAt  string
}

type Role struct {
	Id           int
	IsSystemRole bool
	Role         int
	Name         string
	Description  string
}

type Resource struct {
	Id          int
	Resource    int
	Name        string
	Description string
}

type Permission struct {
	Id          int
	Action      int
	Enabled     bool
	Name        string
	Description string
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
	UserRBAC       []teams_store.UserRBAC
}
