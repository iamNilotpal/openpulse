package user_store

type NewDBUser struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
	AvatarUrl    string
	RoleID       int
}

type DBUser struct {
	Id            int
	FirstName     string
	LastName      string
	Email         string
	RoleID        int
	AvatarUrl     string
	AccountStatus string
	CreatedAt     string
	UpdatedAt     string
}
