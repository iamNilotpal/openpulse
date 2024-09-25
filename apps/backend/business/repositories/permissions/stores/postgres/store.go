package permissions_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (DBPermission, error)
	Create(context context.Context, permission DBNewPermission) (int, error)
	QueryByUserId(context context.Context, userId int) ([]DBUserPermission, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, permission DBNewPermission) (int, error) {
	query := `
		INSERT INTO permissions (name, description, action, resource)
		VALUES ($1, $2, $3, $4);
	`

	result, err := s.db.ExecContext(
		context,
		query,
		permission.Name,
		permission.Description,
		permission.Action,
		permission.Resource,
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

func (s *postgresStore) QueryById(context context.Context, id int) (DBPermission, error) {
	var permission DBPermission
	query := `
		SELECT id, name, description, action, resource, created_at, updated_at
		WHERE id = $1;
	`

	if err := s.db.QueryRowContext(context, query, id).Scan(
		&permission.Id,
		&permission.Name,
		&permission.Description,
		&permission.Action,
		&permission.Resource,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	); err != nil {
		return DBPermission{}, err
	}

	return permission, nil
}

func (s *postgresStore) QueryByUserId(context context.Context, userId int) ([]DBUserPermission, error) {
	query := `
		SELECT
			p.id AS id,
			p.name AS name,
			p.action AS action,
			up.enabled AS enabled,
			p.resource AS resource,
			p.description AS description
			up.updated_by AS updated_by
			p.created_At AS created_at
			p.updated_at AS updated_at
		FROM
			users_permissions up
			JOIN permissions p ON p.id = up.permission_id
		WHERE
			up.user_id = $1;
	`

	rows, err := s.db.QueryContext(context, query, userId)
	if err != nil {
		return []DBUserPermission{}, err
	}

	defer rows.Close()
	permissions := make([]DBUserPermission, 0)

	for rows.Next() {
		var permission DBUserPermission

		if err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Action,
			&permission.Enabled,
			&permission.Resource,
			&permission.Description,
			&permission.UpdatedBy,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return []DBUserPermission{}, err
		}

		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return []DBUserPermission{}, err
	}

	return permissions, nil
}
