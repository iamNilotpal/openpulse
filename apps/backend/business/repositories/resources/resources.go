package resources

import (
	"context"

	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Repository interface {
	Create(context context.Context, nr NewResource) (int, error)
	QueryAll(ctx context.Context) ([]Resource, error)
	QueryById(context context.Context, id int) (Resource, error)
	QueryAllResourcesWithPermissions(context context.Context) ([]ResourceWithPermission, error)
}

type postgresRepository struct {
	store resources_store.Store
}

func NewPostgresRepository(store resources_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) Create(context context.Context, nr NewResource) (int, error) {
	return r.store.Create(context, ToNewDBResource(nr))
}

func (r *postgresRepository) QueryById(context context.Context, id int) (Resource, error) {
	dbResource, err := r.store.QueryById(context, id)
	if err != nil {
		return Resource{}, err
	}

	return FromDBResource(dbResource), nil
}

func (r *postgresRepository) QueryAllResourcesWithPermissions(context context.Context) (
	[]ResourceWithPermission, error,
) {
	dbResourcesWithPermissions, err := r.store.QueryAllResourcesWithPermissions(context)
	if err != nil {
		return []ResourceWithPermission{}, err
	}

	rps := make([]ResourceWithPermission, 0, len(dbResourcesWithPermissions))
	for i, rp := range dbResourcesWithPermissions {
		rps[i] = FromDBResourceWithPermission(rp)
	}

	return rps, nil
}

func (r *postgresRepository) QueryAll(ctx context.Context) ([]Resource, error) {
	dbResources, err := r.store.QueryAll(ctx)
	if err != nil {
		return []Resource{}, err
	}

	resources := make([]Resource, len(dbResources))
	for i, res := range dbResources {
		resources[i] = FromDBResource(res)
	}

	return resources, nil
}
