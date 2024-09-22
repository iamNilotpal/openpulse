package user

import (
	"database/sql"
	"fmt"
	"net/mail"
	"time"

	user_store "github.com/iamNilotpal/openpulse/business/repositories/user/stores/db"
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

func ToNewDBUser(p NewUser) user_store.NewDBUser {
	return user_store.NewDBUser{
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email.Address,
		PasswordHash: p.PasswordHash,
		AvatarUrl:    sql.NullString{String: p.AvatarUrl, Valid: p.AvatarUrl != ""},
		RoleID:       p.RoleId,
	}
}

func ToUser(p user_store.DBUser) User {
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
