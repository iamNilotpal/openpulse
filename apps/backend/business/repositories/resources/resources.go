package resources

import (
	"context"

	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
)

type Repository interface {
	Create(context context.Context, nr NewResource) (int, error)
	QueryById(context context.Context, id int) (Resource, error)
	QueryAllResourcesWithPermissions(context context.Context) ([]ResourceWithPermission, error)
}

type PostgresRepository struct {
	store resources_store.Store
}

func NewPostgresRepository(store resources_store.Store) *PostgresRepository {
	return &PostgresRepository{store: store}
}

func (r *PostgresRepository) Create(context context.Context, nr NewResource) (int, error) {
	return r.store.Create(context, ToNewDBResource(nr))
}

func (r *PostgresRepository) QueryById(context context.Context, id int) (Resource, error) {
	dbResource, err := r.store.QueryById(context, id)
	if err != nil {
		return Resource{}, err
	}

	return FromDBResource(dbResource), nil
}

func (r *PostgresRepository) QueryAllResourcesWithPermissions(context context.Context) (
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