package teams_store

import (
	"context"
	"fmt"
	"strings"

	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Create(context context.Context, team NewTeam) (int, error)
	QueryById(context context.Context, id int) (Team, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) Create(context context.Context, team NewTeam) (int, error) {
	var id int
	if err := database.WithTx(
		context,
		s.db,
		nil,
		func(tx *sqlx.Tx) error {
			query := `
				INSERT INTO
					roles (name, description, total_members, invitation_code, creator_id, org_id)
				VALUES
					($1, $2, $3, $4, $5, $6) RETURNING id;
			`

			if err := tx.QueryRowContext(
				context,
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

			var args []any
			query = `
				INSERT INTO
					team_users (team_id, user_id, role_id, resource_id, permission_id)
				VALUES
			`

			params := database.BuildQueryParams(
				team.UserRBAC,
				func(index int, isLast bool, v UserRBAC) string {
					args = append(args, id, v.UserId, v.RoleId, v.ResourceId, v.PermissionId)
					return "(?, ?, ?, ?, ?)"
				},
			)

			query += strings.Join(params, ", ")
			fmt.Printf("\nQUERY : %s", query)
			if _, err := tx.ExecContext(context, query, args...); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *postgresStore) AddTeamMember(
	context context.Context, teamId int, userRBAC []UserRBAC,
) error {
	var args []any
	query := `
		INSERT INTO
			team_users (team_id, user_id, role_id, resource_id, permission_id)
		VALUES
	`

	params := database.BuildQueryParams(
		userRBAC,
		func(index int, isLast bool, v UserRBAC) string {
			args = append(args, teamId, v.UserId, v.RoleId, v.ResourceId, v.PermissionId)
			return "(?, ?, ?, ?, ?)"
		},
	)

	query += strings.Join(params, ", ")
	fmt.Printf("\nQUERY : %s\n", query)
	if _, err := s.db.ExecContext(context, query, args...); err != nil {
		return err
	}

	return nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Team, error) {
	var team Team
	// query := `
	// 	SELECT id, name, description, total_members, admin_id, created_at, updated_at
	// 	FROM teams
	// 	WHERE id = $1;
	// `
	// if err := s.db.QueryRowContext(context, query, id).Scan(
	// 	&team.Id, &team.Name, &team.Description, &team.AdminId, &team.CreatedAt, &team.UpdatedAt,
	// ); err != nil {
	// 	return Team{}, err
	// }

	return team, nil
}
