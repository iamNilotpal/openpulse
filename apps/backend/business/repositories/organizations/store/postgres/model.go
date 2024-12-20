package organizations_store

type Organization struct {
	Id             int
	Name           string
	Description    string
	LogoURL        string
	TotalEmployees string
	Admin          OrgAdmin
	CreatedAt      string
	UpdatedAt      string
}

type OrgAdmin struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}
