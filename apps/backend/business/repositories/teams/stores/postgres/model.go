package teams_store

type DBTeam struct {
	Id           int
	Name         string
	Description  string
	TotalMembers int
	AdminId      int
	CreatedAt    string
	UpdatedAt    string
}

type DBNewTeam struct {
	AdminId     int
	Name        string
	Description string
}
