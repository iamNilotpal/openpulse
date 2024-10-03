package permissions

import (
	"context"

	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
)

type Repository interface {
	Create(context context.Context, permission NewPermission) (int, error)
	QueryById(context context.Context, id int) (Permission, error)
}

type PostgresRepository struct {
	s permissions_store.Store
}

func NewPostgresRepository(store permissions_store.Store) *PostgresRepository {
	return &PostgresRepository{s: store}
}

func (r *PostgresRepository) Create(context context.Context, permission NewPermission) (int, error) {
	id, err := r.s.Create(context, NewDBPermission(permission))
	return id, err
}

func (r *PostgresRepository) QueryById(context context.Context, id int) (Permission, error) {
	permission, err := r.s.QueryById(context, id)
	if err != nil {
		return Permission{}, nil
	}

	return FromDBPermission(permission), nil
}
