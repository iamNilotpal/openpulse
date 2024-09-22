package roles_store

type DBRole struct {
	Id          int
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type NewDBRole struct {
	Name        string
	Description string
}
