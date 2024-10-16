package users_store

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (User, error)
	Create(context context.Context, cmd NewUser) (int, error)
	CreateTeam(context context.Context, team NewTeam) (int, error)
	CreateOrganization(context context.Context, cmd NewOrganization) (int, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, cmd NewUser) (int, error) {
	query := `
		INSERT INTO
			users(first_name, last_name, email, password_hash, role_id)
		VALUES
			($1, $2, $3, $4, $5) RETURNING id;
	`

	var id int
	if err := s.db.QueryRowContext(
		context,
		query,
		cmd.FirstName,
		cmd.LastName,
		cmd.Email,
		cmd.PasswordHash,
		cmd.RoleId,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *postgresStore) CreateOrganization(ctx context.Context, cmd NewOrganization) (int, error) {
	var id int
	err := database.WithTx(
		ctx,
		s.db,
		&sql.TxOptions{},
		func(tx *sqlx.Tx) error {
			query := `
				INSERT INTO
					organizations (
						name,
						description,
						logo_url,
						total_employees,
						admin_id
					)
				VALUES
					($1, $2, $3, $4, $5) RETURNING id;
			`

			var id int
			if err := tx.QueryRowContext(
				ctx, query, cmd.Name, cmd.Description, cmd.LogoURL, cmd.TotalEmployees, cmd.AdminId,
			).Scan(&id); err != nil {
				return err
			}

			query = `
				UPDATE users
				SET designation = $1
				WHERE id = $2;
			`

			_, err := tx.ExecContext(ctx, query, cmd.Designation, cmd.AdminId)
			return err
		},
	)

	return id, err
}

func (s *postgresStore) CreateTeam(ctx context.Context, team NewTeam) (int, error) {
	var id int
	err := database.WithTx(
		ctx,
		s.db,
		&sql.TxOptions{},
		func(tx *sqlx.Tx) error {
			query := `
				INSERT INTO
					roles (name, description, total_members, invitation_code, creator_id, org_id)
				VALUES
					($1, $2, $3, $4, $5, $6) RETURNING id;
			`

			if err := tx.QueryRowContext(ctx,
				query,
				team.Name,
				team.Description,
				1,
				team.InvitationCode,
				team.CreatorId,
				team.OrgId,
			).Scan(&id); err != nil {
				return err
			}

			query = `
				UPDATE users
					SET users.team_id = $1
				WHERE
					users.id = $2;
			`
			if _, err := tx.ExecContext(ctx, query, id, team.CreatorId); err != nil {
				return err
			}

			var args []any
			query = `
				INSERT INTO
					team_users (team_id, user_id, role_id, resource_id, permission_id)
				VALUES
			`

			params := database.BuildQueryParams(
				team.UserRBAC,
				func(index int, isLast bool, v teams_store.UserRBAC) string {
					args = append(args, id, v.UserId, v.RoleId, v.ResourceId, v.PermissionId)
					return "(?, ?, ?, ?, ?)"
				},
			)

			fmt.Printf("\nARGS : %+v\n", args)
			fmt.Printf("\nPARAMS : %+v\n", params)

			query += strings.Join(params, ", ")
			fmt.Printf("\nBefore Rebind QUERY : %s\n", query)

			query = tx.Rebind(query)
			fmt.Printf("\nAfter Rebind QUERY : %s\n", query)

			_, err := tx.ExecContext(ctx, query, args...)
			return err
		},
	)
	return id, err
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
