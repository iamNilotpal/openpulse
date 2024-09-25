package roles_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, permission NewDBRole) (int, error)
	GetAll(context context.Context) ([]DBRole, error)
	QueryById(context context.Context, id int) (DBRole, error)
	QueryByName(context context.Context, name string) (DBRole, error)
	QueryRolesWithPermissions(context context.Context) ([]DBRolePermissions, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, nr NewDBRole) (int, error) {
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

func (s *postgresStore) GetAll(context context.Context) ([]DBRole, error) {
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

func (s *postgresStore) QueryById(context context.Context, id int) (DBRole, error) {
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

func (s *postgresStore) QueryByName(context context.Context, name string) (DBRole, error) {
	var role DBRole
	query := `
		SELECT id, name, description, created_at, updated_at
		FROM roles
		WHERE name = $1;
	`

	if err := s.db.QueryRowContext(context, query, name).Scan(
		&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt,
	); err != nil {
		return DBRole{}, err
	}

	return role, nil
}

func (s *postgresStore) QueryRolesWithPermissions(context context.Context) ([]DBRolePermissions, error) {
	query := `
		SELECT
			r.id AS role_id,
			r.name AS role_name,
			r.description AS role_description,
			r.created_at AS role_created_at,
			r.updated_at AS role_updated_at,
			p.id AS permission_id,
			p.name AS permission_name,
			p.description AS permission_description,
			p.action AS permission_action,
			p.resource AS permission_resource,
			p.created_at AS permission_created_at,
			p.updated_at AS permission_updated_at
		FROM
			roles_permissions AS rp
			JOIN roles r ON r.id = rp.role_id
			JOIN permissions p ON p.id = rp.permission_id;
	`

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return []DBRolePermissions{}, err
	}

	defer rows.Close()
	rolesWithPermissions := make([]DBRolePermissions, 0)

	for rows.Next() {
		var row DBRolePermissions

		if err := rows.Scan(
			&row.Role.Id,
			&row.Role.Name,
			&row.Role.Description,
			&row.Role.CreatedAt,
			&row.Role.UpdatedAt,
			&row.Permission.Id,
			&row.Permission.Name,
			&row.Permission.Description,
			&row.Permission.Action,
			&row.Permission.Resource,
			&row.Permission.CreatedAt,
			&row.Permission.UpdatedAt,
		); err != nil {
			return []DBRolePermissions{}, err
		}

		rolesWithPermissions = append(rolesWithPermissions, row)
	}

	if err := rows.Err(); err != nil {
		return []DBRolePermissions{}, err
	}

	return rolesWithPermissions, nil
}
