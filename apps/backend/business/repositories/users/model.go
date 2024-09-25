package users

import (
	"database/sql"
	"fmt"
	"net/mail"
	"time"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type User struct {
	ID            int
	FirstName     string
	LastName      string
	Email         mail.Address
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
	Email        mail.Address
	PasswordHash []byte
	AvatarUrl    string
	RoleId       int
}

type UpdateUser struct {
	FirstName string
	LastName  string
	Email     mail.Address
	AvatarUrl string
}

func ToNewDBUser(p NewUser) users_store.DBNewUser {
	return users_store.DBNewUser{
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email.Address,
		PasswordHash: p.PasswordHash,
		AvatarUrl:    sql.NullString{String: p.AvatarUrl, Valid: p.AvatarUrl != ""},
		RoleID:       p.RoleId,
	}
}

func ToUser(p users_store.DBUser) User {
	createdAt, _ := time.Parse("", p.CreatedAt)
	updatedAt, _ := time.Parse("", p.UpdatedAt)

	return User{
		ID:            p.Id,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Email:         mail.Address{Name: fmt.Sprintf("%s %s", p.FirstName, p.LastName), Address: p.Email},
		RoleID:        p.RoleID,
		AvatarUrl:     p.AvatarUrl,
		AccountStatus: p.AccountStatus,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
