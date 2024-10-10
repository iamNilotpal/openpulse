package users_store

type NewUser struct {
	RoleId       int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
}

type User struct {
	Id            int
	Role          Role
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	AvatarUrl     string
	AccountStatus string
	Preference    Preference
	Resources     []ResourcePermission
	CreatedAt     string
	UpdatedAt     string
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
	Name         string
	Description  string
	Role         string
}

type Resource struct {
	Id          int
	Name        string
	Description string
	Resource    string
}

type Permission struct {
	Id          int
	Enabled     bool
	Name        string
	Description string
	Action      string
}

type ResourcePermission struct {
	Resource   Resource
	Permission Permission
}
