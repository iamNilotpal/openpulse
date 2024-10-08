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
	FirstName     string
	LastName      string
	Email         string
	Phone         string
	AvatarUrl     string
	AccountStatus int
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
