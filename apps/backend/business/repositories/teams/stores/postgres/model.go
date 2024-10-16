package teams_store

type Team struct {
	Id             int
	Name           string
	Description    string
	TotalMembers   int
	InvitationCode string
	Creator        Creator
	Organization   Organization
	CreatedAt      string
	UpdatedAt      string
}

type UserRBAC struct {
	RoleId       int
	UserId       int
	ResourceId   int
	PermissionId int
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
