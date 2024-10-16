package teams

import (
	"time"
)

type Team struct {
	Id             int
	Name           string
	Description    string
	TotalMembers   int
	InvitationCode string
	Creator        Creator
	Organization   Organization
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Organization struct {
	Id      int
	Name    string
	LogoURL string
}

type Creator struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
