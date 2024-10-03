package users_store

import "database/sql"

type NewUser struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
	AvatarUrl    sql.NullString
	RoleID       int
}

type User struct {
	Id            int
	FirstName     string
	LastName      string
	Email         string
	RoleID        int
	AvatarUrl     string
	AccountStatus string
	CreatedAt     string
	UpdatedAt     string
}

type Role struct {
	Id int
}

type Permission struct {
	Id       int
	Action   string
	Resource string
}

type UserPermissions struct {
	Role       Role
	Permission Permission
}
