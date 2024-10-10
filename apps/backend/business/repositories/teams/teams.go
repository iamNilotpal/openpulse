package teams

import (
	"context"

	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
)

type Repository interface {
	Create(context context.Context, team NewTeam) (int, error)
	QueryById(context context.Context, id int) (Team, error)
}

type postgresRepository struct {
	s teams_store.Store
}

func NewRepository(store teams_store.Store) *postgresRepository {
	return &postgresRepository{s: store}
}

func (r *postgresRepository) Create(context context.Context, team NewTeam) (int, error) {
	id, err := r.s.Create(context, ToNewDBTeam(team))
	return id, err
}

func (r *postgresRepository) QueryById(context context.Context, id int) (Team, error) {
	team, err := r.s.QueryById(context, id)
	if err != nil {
		return Team{}, err
	}

	return FromDBTeam(team), err
}
