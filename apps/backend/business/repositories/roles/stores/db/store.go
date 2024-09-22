package roles_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(context.Context) ([]DBRole, error)
	Create(context.Context, NewDBRole) (int, error)
	QueryById(context.Context, int) (DBRole, error)
}

type PostgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) Create(context context.Context, nr NewDBRole) (int, error) {
	query := `INSERT INTO ROLES (name, description) VALUES ($1, $2) RETURNING id;`

	result, err := s.db.ExecContext(context, query, nr.Name, nr.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *PostgresStore) GetAll(context context.Context) ([]DBRole, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM roles;`

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return []DBRole{}, err
	}

	defer rows.Close()
	roles := make([]DBRole, 0)

	for rows.Next() {
		var role DBRole

		if err := rows.Scan(
			&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt,
		); err != nil {
			return []DBRole{}, err
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return []DBRole{}, err
	}

	return roles, nil
}

func (s *PostgresStore) QueryById(context context.Context, id int) (DBRole, error) {
	var role DBRole
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM roles
		WHERE id = $1;
	`

	if err := s.db.QueryRowContext(context, query, id).Scan(
		&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt,
	); err != nil {
		return DBRole{}, err
	}

	return role, nil
}
