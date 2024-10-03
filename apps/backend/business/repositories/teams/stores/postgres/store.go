package teams_store

import (
	"context"

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
	query := `
		INSERT INTO roles (name, description, admin_id, total_members)
		VALUES ($1, $2, $3, $4);
	`

	result, err := s.db.ExecContext(context, query, team.Name, team.Description, team.AdminId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *postgresStore) QueryById(context context.Context, id int) (Team, error) {
	var team Team
	query := `
		SELECT id, name, description, total_members, admin_id, created_at, updated_at
		FROM teams
		WHERE id = $1;
	`
	if err := s.db.QueryRowContext(context, query, id).Scan(
		&team.Id, &team.Name, &team.Description, &team.AdminId, &team.CreatedAt, &team.UpdatedAt,
	); err != nil {
		return Team{}, err
	}

	return team, nil
}
