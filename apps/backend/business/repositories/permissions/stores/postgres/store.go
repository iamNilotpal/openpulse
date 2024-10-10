package permissions_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (Permission, error)
	Create(context context.Context, permission NewPermission) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, np NewPermission) (int, error) {
	query := `
		INSERT INTO permissions (name, description, action)
		VALUES ($1, $2, $3);
	`

	result, err := s.db.ExecContext(context, query, np.Name, np.Description, np.Action)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Permission, error) {
	var permission Permission
	query := `
		SELECT
			p.id AS permissionId,
			p.name AS permissionName,
			P.description AS permissionDescription,
			p.action AS permissionAction,
			p.created_at as permissionCreatedAt,
			p.updated_at as permissionUpdatedAt
		FROM
			permissions p
		WHERE
			id = $1;
	`

	if err := s.db.QueryRowContext(context, query, id).Scan(
		&permission.Id,
		&permission.Name,
		&permission.Description,
		&permission.Action,
		&permission.CreatedAt,
		&permission.UpdatedAt,
	); err != nil {
		return Permission{}, err
	}

	return permission, nil
}
