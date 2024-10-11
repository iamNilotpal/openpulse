package users_store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, payload NewUser) (int, error)
	QueryById(context context.Context, id int) (User, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (p *postgresStore) Create(context context.Context, payload NewUser) (int, error) {
	query := `
		INSERT INTO
			users(first_name, last_name, email, password_hash, role_id)
		VALUES
			($1, $2, $3, $4, $5) RETURNING id;
	`

	var id int
	if err := p.db.QueryRowContext(
		context,
		query,
		payload.FirstName,
		payload.LastName,
		payload.Email,
		payload.PasswordHash,
		payload.RoleId,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

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
			t.id as teamId,
			t.name as teamName,
			t.logo_url as teamLogo,
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
			JOIN teams t ON t.id = us.team_id
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
		&user.Team.Id,
		&user.Team.Name,
		&user.Team.LogoURL,
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
			tu.enabled as permissionEnabled
		FROM
			team_users tu
			JOIN resources res ON res.id = tu.resource_id
			JOIN permissions pes ON pes.id = tu.permission_id
		WHERE
			tu.team_id = $1 AND tu.user_id = $2;
	`

	rows, err := p.db.QueryContext(context, query, user.Team.Id, user.Role.Id)
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
