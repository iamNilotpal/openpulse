package users_store

import "database/sql"

type NewDBUser struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
	AvatarUrl    sql.NullString
	RoleID       int
}

type DBUser struct {
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
