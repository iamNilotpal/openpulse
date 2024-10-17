package sessions

import (
	"context"

	sessions_store "github.com/iamNilotpal/openpulse/business/repositories/sessions/store/postgres"
)

type Repository interface {
	Create(context context.Context, cmd NewSession) (int, error)
}

type postgresRepository struct {
	store sessions_store.Store
}

func NewPostgresRepository(store sessions_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) Create(context context.Context, cmd NewSession) (int, error) {
	return r.store.Create(context, ToNewDBSession(cmd))
}
