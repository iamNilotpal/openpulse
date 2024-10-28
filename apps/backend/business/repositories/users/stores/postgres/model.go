package users_store

import (
	"database/sql"

	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
)

type NewUser struct {
	RoleId    int
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	Id              int
	FirstName       string
	LastName        string
	Email           string
	Phone           sql.NullString
	CountryCode     sql.NullString
	AvatarUrl       sql.NullString
	AccountStatus   int
	Designation     sql.NullString
	IsEmailVerified bool
	Role            Role
	Team            Team
	OAuthAccounts   []OAuthAccount
	Preference      Preference
	Resources       []ResourcePermission
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       sql.NullInt64
}

type Team struct {
	Id      sql.NullInt64
	Name    sql.NullString
	LogoURL sql.NullString
}

type Preference struct {
	Id         sql.NullInt64
	Appearance sql.NullString
	CreatedAt  sql.NullString
	UpdatedAt  sql.NullString
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

type OAuthAccount struct {
	Id         int
	Provider   string
	ExternalId string
	Scope      string
	Metadata   sql.NullString
	CreatedAt  string
	UpdatedAt  string
}

type NewOAuthUser struct {
	RoleId    int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	AvatarURL string
}

type NewOAuthAccount struct {
	Provider   string
	ExternalId string
	Scope      string
	Metadata   string
	User       NewOAuthUser
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
	CreatorId         int
	CreatorRoleId     int
	OrgId             int
	Name              string
	Description       string
	InvitationCode    string
	UserAccessControl []teams_store.UserAccessControl
}
