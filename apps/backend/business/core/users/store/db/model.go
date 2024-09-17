package users_store

type CreateUserDBPayload struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash []byte
	AvatarUrl    string
	RoleId       int
}

type DBUser struct {
	Id            int    `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	Email         string `db:"email"`
	RoleId        int    `db:"role_id"`
	AvatarUrl     string `db:"avatar_url"`
	AccountStatus string `db:"account_status"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
