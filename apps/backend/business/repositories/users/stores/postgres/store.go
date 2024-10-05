package users_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	// Create(context context.Context, payload NewUser, permissions []AccessControl) (int, error)
	QueryById(context context.Context, id int) (User, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

// func (p *postgresStore) Create(
// 	context context.Context, payload NewUser, permissions []AccessControl,
// ) (int, error) {
// 	query := `
// 		INSERT INTO
// 		users(first_name, last_name, email, password_hash, role_id)
// 		VALUES ($1, $2, $3, $4, $5) RETURNING id;
// 	`

// 	result, err := p.db.ExecContext(
// 		context, query,
// 		payload.FirstName, payload.LastName, payload.Email, payload.PasswordHash, payload.RoleId,
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	var args []any
// 	query = "INSERT INTO users_permissions (user_id, permission_id, enabled, updated_by) VALUES "

// 	params := database.MultipleQueryParams(
// 		permissions,
// 		func(index int, isLast bool, v AccessControl) string {
// 			args = append(args, id, v.Permission.Id, true, id)
// 			return "(?, ?, ?, ?)"
// 		},
// 	)

// 	query += strings.Join(params, ", ")

// 	println()
// 	println("Insert Into User Permissions Query", query)
// 	println()

// 	_, err = p.db.ExecContext(context, query, args...)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(id), nil
// }

func (p *postgresStore) QueryById(context context.Context, id int) (User, error) {
	query := `
		SELECT
			us.id AS userId,
			us.email AS email,
			us.first_name AS firstName,
			us.last_name AS lastName,
			us.phone_number as phoneNumber,
			us.avatar_url as avatarUrl,
			us.account_status as accountStatus,
			us.created_at as createdAt,
			us.updated_at as updatedAt,
			ro.id AS roleId,
			ro.name AS roleName,
			ro.description AS roleDescription,
			ro.role AS role,
			ro.is_system_role AS isSystemRole,
			up.id AS userPreferenceId,
			up.appearance as appearance,
			up.created_at as userPreferenceCreatedAt,
			up.updated_at as userPreferenceUpdatedAt
		FROM
			users us
			JOIN roles ro ON ro.id = us.role_id
			JOIN users_preferences up ON up.user_id = us.id
		WHERE
			us.id = $1;
	`

	var user User
	if err := p.db.QueryRowContext(context, query, id).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.AvatarUrl,
		&user.AccountStatus,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role.Id,
		&user.Role.Name,
		&user.Role.Description,
		&user.Role.Role,
		&user.Role.IsSystemRole,
		&user.Preference.Id,
		&user.Preference.Appearance,
		&user.Preference.CreatedAt,
		&user.Preference.UpdatedAt,
	); err != nil {
		return User{}, err
	}

	query = `
		SELECT
			res.id as resourceId,
			res.resource as resource,
			res.display_name as resourceName,
			res.description as resourceDescription,
			pes.id AS permissionsId,
			pes.name AS permissionName,
			pes.description AS permissionDescription,
			pes.action AS action,
			uac.enabled as permissionEnabled
		FROM
			users_access_controls uac
			JOIN resources res ON res.id = uac.resource_id
			JOIN permissions pes ON pes.id = uac.permission_id
		WHERE
			uac.user_id = $1;
	`

	rows, err := p.db.QueryContext(context, query, id)
	if err != nil {
		return User{}, err
	}

	defer rows.Close()
	resources := make([]ResourcePermission, 0)

	for rows.Next() {
		var r ResourcePermission
		if err := rows.Scan(
			&r.Resource.Id,
			&r.Resource.Resource,
			&r.Resource.Name,
			&r.Resource.Description,
			&r.Permission.Id,
			&r.Permission.Name,
			&r.Permission.Description,
			&r.Permission.Action,
			&r.Permission.Enabled,
		); err != nil {
			return User{}, err
		}

		resources = append(resources, r)
	}

	if err = rows.Err(); err != nil {
		return User{}, err
	}

	user.Resources = resources
	return user, nil
}
