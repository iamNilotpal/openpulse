package permissions_store

type Permission struct {
	Id          int
	Name        string
	Description string
	Action      int
	CreatedAt   string
	UpdatedAt   string
}

type NewPermission struct {
	Name        string
	Description string
	Action      int
}

type PermissionAccessConfig struct {
	Id     int
	Action int
}
