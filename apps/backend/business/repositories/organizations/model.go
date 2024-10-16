package organizations

import (
	"time"
)

type NewOrganization struct {
	AdminId        int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
}

type Organization struct {
	Id             int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
	Admin          OrgAdmin
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type OrgAdmin struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
