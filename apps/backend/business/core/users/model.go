package users

import (
	"net/mail"
	"time"

	store "github.com/iamNilotpal/openpulse/business/core/users/store/db"
)

type AppUser struct {
	Id            int
	FirstName     string
	LastName      string
	Email         string
	RoleId        int
	AvatarUrl     string
	AccountStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateUserPayload struct {
	FirstName    string
	LastName     string
	Email        mail.Address
	PasswordHash []byte
	AvatarUrl    string
	RoleId       int
}

func ToCreateDBUser(p CreateUserPayload) store.CreateUserDBPayload {
	return store.CreateUserDBPayload{
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email.Address,
		PasswordHash: p.PasswordHash,
		AvatarUrl:    p.AvatarUrl,
		RoleId:       p.RoleId,
	}
}

func ToAppUser(p store.DBUser) AppUser {
	createdAt, _ := time.Parse("", p.CreatedAt)
	updatedAt, _ := time.Parse("", p.UpdatedAt)

	return AppUser{
		Id:            p.Id,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Email:         p.Email,
		RoleId:        p.RoleId,
		AvatarUrl:     p.AvatarUrl,
		AccountStatus: p.AccountStatus,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
