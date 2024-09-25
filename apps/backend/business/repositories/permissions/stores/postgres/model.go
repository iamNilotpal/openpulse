package permissions_store

type DBNewPermission struct {
	Name        string
	Description string
	Action      string
	Resource    string
}

type DBPermission struct {
	Id          int
	Name        string
	Description string
	Action      string
	Resource    string
	CreatedAt   string
	UpdatedAt   string
}

type DBUserPermission struct {
	Id          int
	Enabled     bool
	Name        string
	Description string
	Action      string
	Resource    string
	CreatedAt   string
	UpdatedAt   string
	UpdatedBy   string
}
