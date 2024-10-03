package roles_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, permission NewRole) (int, error)
	GetAll(context context.Context) ([]Role, error)
	QueryById(context context.Context, id int) (Role, error)
	QueryByName(context context.Context, name string) (Role, error)
	GetRolesWithPermissions(context context.Context) ([]RoleWithPermission, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, nr NewRole) (int, error) {
	query := `
		INSERT INTO ROLES (name, description, is_system_role, created_by)
		VALUES ($1, $2, $3, $4) RETURNING id;
	`

	result, err := s.db.ExecContext(
		context, query, nr.Name, nr.Description, nr.IsSystemRole, nr.CreatorId,
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

func (s *postgresStore) GetAll(context context.Context) ([]Role, error) {
	query := `
		SELECT
			r.id AS roleId,
			r.name AS roleName,
			r.description AS roleDescription,
			r.is_system_role AS isSystemRole,
			rcb.id AS roleAuthorId,
			rcb.email AS roleAuthorEmail,
			rcb.first_name AS roleAuthorFirstName,
			rcb.last_name AS roleAuthorLastName,
			rub.id AS roleAuthorId,
			rub.email AS roleAuthorEmail,
			rub.first_name AS roleUpdaterFirstName,
			rub.last_name AS roleUpdatedLastName,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt
		FROM
			roles r
			LEFT JOIN users rcb ON rcb.id = r.created_by
			LEFT JOIN users rub ON rub.id = r.updated_by;
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
			&role.CreatedBy.Id,
			&role.CreatedBy.Email,
			&role.CreatedBy.FirstName,
			&role.CreatedBy.LastName,
			&role.UpdatedBy.Id,
			&role.UpdatedBy.Email,
			&role.UpdatedBy.FirstName,
			&role.UpdatedBy.LastName,
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
			rcb.id AS roleAuthorId,
			rcb.email AS roleAuthorEmail,
			rcb.first_name AS roleAuthorFirstName,
			rcb.last_name AS roleAuthorLastName,
			rub.id AS roleAuthorId,
			rub.email AS roleAuthorEmail,
			rub.first_name AS roleUpdaterFirstName,
			rub.last_name AS roleUpdatedLastName,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt
		FROM
			roles r
			LEFT JOIN users rcb ON rcb.id = r.created_by
			LEFT JOIN users rub ON rub.id = r.updated_by
		WHERE
			r.id = $1;
	`

	if err := s.db.QueryRowContext(context, query, id).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.IsSystemRole,
		&role.CreatedBy.Id,
		&role.CreatedBy.Email,
		&role.CreatedBy.FirstName,
		&role.CreatedBy.LastName,
		&role.UpdatedBy.Id,
		&role.UpdatedBy.Email,
		&role.UpdatedBy.FirstName,
		&role.UpdatedBy.LastName,
		&role.CreatedAt,
		&role.UpdatedAt,
	); err != nil {
		return Role{}, err
	}

	return role, nil
}

func (s *postgresStore) QueryByName(context context.Context, name string) (Role, error) {
	var role Role
	query := `
		SELECT
			r.id AS roleId,
			r.name AS roleName,
			r.description AS roleDescription,
			r.is_system_role AS isSystemRole,
			rcb.id AS roleAuthorId,
			rcb.email AS roleAuthorEmail,
			rcb.first_name AS roleAuthorFirstName,
			rcb.last_name AS roleAuthorLastName,
			rub.id AS roleAuthorId,
			rub.email AS roleAuthorEmail,
			rub.first_name AS roleUpdaterFirstName,
			rub.last_name AS roleUpdatedLastName,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt
		FROM
			roles r
			LEFT JOIN users rcb ON rcb.id = r.created_by
			LEFT JOIN users rub ON rub.id = r.updated_by
		WHERE
			r.name = $1;
	`

	if err := s.db.QueryRowContext(context, query, name).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.IsSystemRole,
		&role.CreatedBy.Id,
		&role.CreatedBy.Email,
		&role.CreatedBy.FirstName,
		&role.CreatedBy.LastName,
		&role.UpdatedBy.Id,
		&role.UpdatedBy.Email,
		&role.UpdatedBy.FirstName,
		&role.UpdatedBy.LastName,
		&role.CreatedAt,
		&role.UpdatedAt,
	); err != nil {
		return Role{}, err
	}

	return role, nil
}

func (s *postgresStore) GetRolesWithPermissions(context context.Context) ([]RoleWithPermission, error) {
	query := `
		SELECT
			r.id AS roleID,
			r.name AS roleName,
			r.description AS roleDescription,
			r.is_system_role AS isSystemRole,
			r.created_at AS roleCreatedAt,
			r.updated_at AS roleUpdatedAt,
			p.id AS permissionId,
			p.name AS permissionName,
			p.description AS permissionDescription,
			p.action AS permissionAction,
			p.created_at AS permissionCreatedAt,
			p.updated_at AS permissionUpdatedAt,
			rp.created_at AS rolePermissionCreatedAt,
			rp.updated_at AS rolePermissionUpdatedAt
		FROM
			roles r
			JOIN roles_permissions rp ON rp.role_id = r.id
			JOIN permissions p ON p.id = rp.permission_id;
	`

	rows, err := s.db.QueryContext(context, query)
	if err != nil {
		return []RoleWithPermission{}, err
	}

	defer rows.Close()
	rolesWithPermissions := make([]RoleWithPermission, 0)

	for rows.Next() {
		var row RoleWithPermission

		if err := rows.Scan(
			&row.Role.Id,
			&row.Role.Name,
			&row.Role.Description,
			&row.Role.IsSystemRole,
			&row.Role.CreatedAt,
			&row.Role.UpdatedAt,
			&row.Permission.Id,
			&row.Permission.Name,
			&row.Permission.Description,
			&row.Permission.Action,
			&row.Permission.CreatedAt,
			&row.Permission.UpdatedAt,
			&row.CreatedAt,
			&row.UpdatedAt,
		); err != nil {
			return []RoleWithPermission{}, err
		}

		rolesWithPermissions = append(rolesWithPermissions, row)
	}

	if err := rows.Err(); err != nil {
		return []RoleWithPermission{}, err
	}

	return rolesWithPermissions, nil
}
