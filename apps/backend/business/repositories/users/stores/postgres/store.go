package users_store

import (
	"context"
	"strings"

	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (User, error)
	Create(context context.Context, payload NewUser, permissions []UserPermissions) (int, error)
	QueryByEmail(context context.Context, email string) (User, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (p *postgresStore) Create(
	context context.Context, payload NewUser, permissions []UserPermissions,
) (int, error) {
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

	var args []any
	query = "INSERT INTO users_permissions (user_id, permission_id, enabled, updated_by) VALUES "

	params := database.MultipleQueryParams(
		permissions,
		func(index int, v UserPermissions, isLast bool) string {
			args = append(args, id, v.Permission.Id, true, id)
			return "(?, ?, ?, ?)"
		},
	)

	query += strings.Join(params, ", ")

	println()
	println("Insert Into User Permissions Query", query)
	println()

	_, err = p.db.ExecContext(context, query, args...)

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (p *postgresStore) QueryById(context context.Context, id int) (User, error) {
	query := `
		SELECT
		id, first_name, last_name, email, role_id, avatar_url, account_status, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user User
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
		return User{}, err
	}

	return user, nil
}

func (p *postgresStore) QueryByEmail(context context.Context, email string) (User, error) {
	query := `
		SELECT
		id, first_name, last_name, email, role_id, avatar_url, account_status, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	var user User
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
		return User{}, err
	}

	return user, nil
}
