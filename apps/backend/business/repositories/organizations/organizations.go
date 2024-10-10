package organizations

import (
	"context"

	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
)

type Repository interface {
	Create(context context.Context, org NewOrganization) (int, error)
}

type postgresRepository struct {
	store organizations_store.Store
}

func NewPostgresRepository(store organizations_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) Create(context context.Context, org NewOrganization) (int, error) {
	return r.store.Create(context, ToNewDBOrg(org))
}
