package permissions_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryAll(ctx context.Context) ([]Permission, error)
	QueryById(ctx context.Context, id int) (Permission, error)
	Create(ctx context.Context, permission NewPermission) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(ctx context.Context, np NewPermission) (int, error) {
	var id int
	query := `
		INSERT INTO
			permissions (name, description, action)
		VALUES
			($1, $2, $3) RETURNING id;
	`

	err := s.db.QueryRowContext(ctx, query, np.Name, np.Description, np.Action).Scan(&id)
	return id, err
}

func (s *postgresStore) QueryById(ctx context.Context, id int) (Permission, error) {
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

	if err := s.db.QueryRowContext(ctx, query, id).Scan(
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

func (s *postgresStore) QueryAll(ctx context.Context) ([]Permission, error) {
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
		ORDER BY p.id;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return []Permission{}, err
	}

	defer rows.Close()
	permissions := make([]Permission, 0)

	for rows.Next() {
		var permission Permission
		if err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Description,
			&permission.Action,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return []Permission{}, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return []Permission{}, err
	}

	return permissions, nil
}
