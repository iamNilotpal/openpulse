package users_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type DBStore interface {
	Create(context.Context, CreateUserDBPayload) (int, error)
	QueryByUserId(context context.Context, id int) (DBUser, error)
}

type PostgresStore struct {
	DB *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *PostgresStore {
	return &PostgresStore{DB: db}
}

func (p *PostgresStore) Create(context context.Context, payload CreateUserDBPayload) (int, error) {
	query := `
		INSERT INTO
		users(first_name, last_name, email, password_hash, role_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`

	result, err := p.DB.ExecContext(
		context, query,
		payload.FirstName, payload.LastName, payload.Email, payload.PasswordHash, payload.RoleId,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *PostgresStore) QueryByUserId(context context.Context, id int) (DBUser, error) {
	query := `
		SELECT
		id, first_name, last_name, email, role_id, avatar_url, account_status, created_at, updated_at
		FROM users
		WHERE id = $1;
	`
	row := p.DB.QueryRowContext(context, query, id)

	var user DBUser
	err := row.Scan(&user)
	if err != nil {
		return DBUser{}, err
	}

	err = row.Err()
	if err != nil {
		return DBUser{}, err
	}

	return user, nil
}
