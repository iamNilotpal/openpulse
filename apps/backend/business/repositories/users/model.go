package users

import (
	"database/sql"
	"time"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type User struct {
	ID            int
	FirstName     string
	LastName      string
	Email         string
	PasswordHash  []byte
	RoleID        int
	AvatarUrl     string
	AccountStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type NewUser struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
	AvatarUrl    string
	RoleId       int
}

type UpdateUser struct {
	FirstName string
	LastName  string
	Email     string
	AvatarUrl string
}

type Role struct {
	Id int
}

type Permission struct {
	Id int
}

type UserPermissions struct {
	Role
	Permission
}

func ToNewDBUser(p NewUser) users_store.NewUser {
	return users_store.NewUser{
		Email:        p.Email,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		PasswordHash: p.PasswordHash,
		AvatarUrl:    sql.NullString{String: p.AvatarUrl, Valid: p.AvatarUrl != ""},
		RoleID:       p.RoleId,
	}
}

func FromDBUser(p users_store.User) User {
	createdAt, _ := time.Parse("", p.CreatedAt)
	updatedAt, _ := time.Parse("", p.UpdatedAt)

	return User{
		ID:            p.Id,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Email:         p.Email,
		RoleID:        p.RoleID,
		AvatarUrl:     p.AvatarUrl,
		AccountStatus: p.AccountStatus,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func ToDBUserPermission(p UserPermissions) users_store.UserPermissions {
	return users_store.UserPermissions{
		Role:       users_store.Role{Id: p.Role.Id},
		Permission: users_store.Permission{Id: p.Permission.Id},
	}
}
