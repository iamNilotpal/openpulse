package organizations

import (
	"context"

	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
)

type Repository interface {
	CheckOnboardingStatus(ctx context.Context, userId int) (bool, error)
	QueryById(ctx context.Context, id int) (Organization, error)
	QueryByCreatorId(ctx context.Context, creatorId int) (Organization, error)
}

type postgresRepository struct {
	store organizations_store.Store
}

func NewPostgresRepository(store organizations_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) CheckOnboardingStatus(ctx context.Context, userId int) (bool, error) {
	return r.store.CheckOnboardingStatus(ctx, userId)
}

func (r *postgresRepository) QueryByCreatorId(
	ctx context.Context, creatorId int,
) (Organization, error) {
	org, err := r.store.QueryByCreatorId(ctx, creatorId)
	if err != nil {
		return Organization{}, err
	}

	return FromDBOrg(org), nil
}

func (r *postgresRepository) QueryById(ctx context.Context, id int) (Organization, error) {
	org, err := r.store.QueryById(ctx, id)
	if err != nil {
		return Organization{}, err
	}

	return FromDBOrg(org), nil
}
