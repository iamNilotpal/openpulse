package teams_store

type Team struct {
	Id           int
	Name         string
	Description  string
	TotalMembers int
	AdminId      int
	CreatedAt    string
	UpdatedAt    string
}

type NewTeam struct {
	AdminId     int
	Name        string
	Description string
}
