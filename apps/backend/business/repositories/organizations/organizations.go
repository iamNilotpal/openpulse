package organizations

import (
	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
)

type Repository interface {
}

type postgresRepository struct {
	store organizations_store.Store
}

func NewPostgresRepository(store organizations_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}
