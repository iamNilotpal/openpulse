package teams

import (
	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
)

type Repository interface {
	// QueryById(context context.Context, id int) (Team, error)
}

type postgresRepository struct {
	s teams_store.Store
}

func NewRepository(store teams_store.Store) *postgresRepository {
	return &postgresRepository{s: store}
}
