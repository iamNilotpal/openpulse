package roles_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(ctx context.Context, permission NewRole) (int, error)
	QueryAll(ctx context.Context) ([]Role, error)
	QueryById(ctx context.Context, id int) (Role, error)
	QueryAccessControl(ctx context.Context) ([]AccessControl, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(ctx context.Context, nr NewRole) (int, error) {
	query := `
		INSERT INTO
			roles (name, description, is_system_role, role)
		VALUES
			($1, $2, $3, $4) RETURNING id;
	`

	var id int
	if err := s.db.QueryRowContext(
		ctx, query, nr.Name, nr.Description, nr.IsSystemRole, nr.Role,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *postgresStore) QueryAll(ctx context.Context) ([]Role, error) {
	query := `
		SELECT
			r.id AS roleId,
			r.name AS roleName,
			r.description AS roleDescription,
			r.is_system_role AS isSystemRole,
			r.role AS role,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt
		FROM
			roles r
		ORDER BY r.id;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return []Role{}, err
	}

	defer rows.Close()
	roles := make([]Role, 0)

	for rows.Next() {
		var role Role
		if err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.IsSystemRole,
			&role.Role,
			&role.CreatedAt,
			&role.UpdatedAt,
		); err != nil {
			return []Role{}, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return []Role{}, err
	}
	return roles, nil
}

func (s *postgresStore) QueryById(ctx context.Context, id int) (Role, error) {
	var role Role
	query := `
		SELECT
			r.id AS roleId,
			r.name AS roleName,
			r.description AS roleDescription,
			r.is_system_role AS isSystemRole,
			r.role AS role,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt
		FROM
			roles r
		WHERE
			r.id = $1;
	`

	if err := s.db.QueryRowContext(ctx, query, id).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.IsSystemRole,
		&role.Role,
		&role.CreatedAt,
		&role.UpdatedAt,
	); err != nil {
		return Role{}, err
	}

	return role, nil
}

func (s *postgresStore) QueryAccessControl(ctx context.Context) ([]AccessControl, error) {
	query := `
		SELECT
			ro.id AS roleId,
			ro.role AS role,
			res.id AS resourceId,
			res.resource AS resourceType,
			ps.id AS permissionId,
			ps.action AS permissionAction
		FROM
			roles ro
			JOIN roles_resources rr ON rr.role_id = ro.id
			JOIN resources res ON res.id = rr.resource_id
			JOIN resource_permissions rp ON rp.resource_id = res.id
			JOIN permissions ps ON ps.id = rp.permission_id
		ORDER BY ro.id, res.id, ps.id;
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return []AccessControl{}, err
	}

	defer rows.Close()
	accessControls := make([]AccessControl, 0)

	for rows.Next() {
		var row AccessControl
		if err := rows.Scan(
			&row.Role.Id,
			&row.Role.Role,
			&row.Resource.Id,
			&row.Resource.Resource,
			&row.Permission.Id,
			&row.Permission.Action,
		); err != nil {
			return []AccessControl{}, err
		}
		accessControls = append(accessControls, row)
	}

	if err := rows.Err(); err != nil {
		return []AccessControl{}, err
	}
	return accessControls, nil
}
