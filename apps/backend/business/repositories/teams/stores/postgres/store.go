package teams_store

import (
	"context"
	"fmt"
	"strings"

	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	QueryById(context context.Context, id int) (Team, error)
}

type postgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(db *sqlx.DB) *postgresStore {
	return &postgresStore{db: db}
}

func (s *postgresStore) AddTeamMember(
	context context.Context, teamId int, userRBAC []UserAccessControl,
) error {
	var args []any
	query := `
		INSERT INTO
			team_users (team_id, user_id, role_id, resource_id, permission_id)
		VALUES
	`

	params := database.BuildQueryParams(
		userRBAC,
		func(index int, isLast bool, v UserAccessControl) string {
			args = append(args, teamId, v.UserId, v.RoleId, v.ResourceId, v.PermissionId)
			return "(?, ?, ?, ?, ?)"
		},
	)

	query += strings.Join(params, ", ")
	fmt.Printf("\nBefore Rebind QUERY : %s\n", query)

	query = s.db.Rebind(query)
	fmt.Printf("\nAfter Rebind QUERY : %s\n", query)

	if _, err := s.db.ExecContext(context, query, args...); err != nil {
		return err
	}

	return nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Team, error) {
	var team Team
	return team, nil
}
