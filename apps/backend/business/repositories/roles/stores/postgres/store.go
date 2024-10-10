package roles_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, permission NewRole) (int, error)
	GetAll(context context.Context) ([]Role, error)
	QueryById(context context.Context, id int) (Role, error)
	GetRolesAccessControl(context context.Context) ([]RoleAccessControl, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, nr NewRole) (int, error) {
	query := `
		INSERT INTO ROLES (name, description, is_system_role, role)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`

	result, err := s.db.ExecContext(context, query, nr.Name, nr.Description, nr.IsSystemRole, nr.Role)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *postgresStore) GetAll(context context.Context) ([]Role, error) {
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
	`

	rows, err := s.db.QueryContext(context, query)
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

func (s *postgresStore) QueryById(context context.Context, id int) (Role, error) {
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

	if err := s.db.QueryRowContext(context, query, id).Scan(
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

func (s *postgresStore) GetRolesAccessControl(context context.Context) ([]RoleAccessControl, error) {
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

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return []RoleAccessControl{}, err
	}

	defer rows.Close()
	accessControls := make([]RoleAccessControl, 0)

	for rows.Next() {
		var row RoleAccessControl

		if err := rows.Scan(
			&row.Role.Id,
			&row.Role.Role,
			&row.Resource.Id,
			&row.Resource.Resource,
			&row.Permission.Id,
			&row.Permission.Action,
		); err != nil {
			return []RoleAccessControl{}, err
		}

		accessControls = append(accessControls, row)
	}

	if err := rows.Err(); err != nil {
		return []RoleAccessControl{}, err
	}

	return accessControls, nil
}
