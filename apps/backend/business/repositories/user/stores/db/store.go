package user_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context.Context, NewDBUser) (int, error)
	QueryById(context context.Context, id int) (DBUser, error)
	QueryByEmail(context context.Context, email string) (DBUser, error)
}

type PostgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (p *PostgresStore) Create(context context.Context, payload NewDBUser) (int, error) {
	query := `
		INSERT INTO
		users(first_name, last_name, email, password_hash, role_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`

	result, err := p.db.ExecContext(
		context, query,
		payload.FirstName, payload.LastName, payload.Email, payload.PasswordHash, payload.RoleID,
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

func (p *PostgresStore) QueryById(context context.Context, id int) (DBUser, error) {
	query := `
		SELECT
		id, first_name, last_name, email, role_id, avatar_url, account_status, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user DBUser
	if err := p.db.QueryRowContext(context, query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
		&user.AvatarUrl,
		&user.AccountStatus,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return DBUser{}, err
	}

	return user, nil
}

func (p *PostgresStore) QueryByEmail(context context.Context, email string) (DBUser, error) {
	query := `
		SELECT
		id, first_name, last_name, email, role_id, avatar_url, account_status, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user DBUser
	if err := p.db.QueryRowContext(context, query, email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
		&user.AvatarUrl,
		&user.AccountStatus,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return DBUser{}, err
	}

	return user, nil
}
