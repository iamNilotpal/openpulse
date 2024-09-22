package roles

import (
	"time"

	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/db"
)

type Role struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NewRole struct {
	Name        string
	Description string
}

type UpdateRole struct {
	Name        string
	Description string
}

func ToNewRole(name, description string) NewRole {
	return NewRole{Name: name, Description: description}
}

func ToNewDBRole(r NewRole) roles_store.NewDBRole {
	return roles_store.NewDBRole{Name: r.Name, Description: r.Description}
}

func ToRole(r roles_store.DBRole) Role {
	createdAt, _ := time.Parse(time.UnixDate, r.CreatedAt)
	updatedAt, _ := time.Parse(time.UnixDate, r.UpdatedAt)

	return Role{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
