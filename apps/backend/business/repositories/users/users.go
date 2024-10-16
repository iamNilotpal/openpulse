package users

import (
	"context"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type Repository interface {
	QueryById(context context.Context, id int) (User, error)
	Create(context context.Context, payload NewUser) (int, error)
	CreateTeam(context context.Context, cmd NewTeam) (int, error)
	CreateOrganization(context context.Context, cmd NewOrganization) (int, error)
}

type postgresRepository struct {
	store users_store.Store
}

func NewPostgresRepository(store users_store.Store) *postgresRepository {
	return &postgresRepository{store: store}
}

func (r *postgresRepository) Create(context context.Context, payload NewUser) (int, error) {
	id, err := r.store.Create(context, ToNewDBUser(payload))
	return id, err
}

func (r *postgresRepository) CreateOrganization(
	context context.Context, cmd NewOrganization,
) (int, error) {
	return r.store.CreateOrganization(context, ToNewDBOrganization(cmd))
}

func (r *postgresRepository) CreateTeam(context context.Context, cmd NewTeam) (int, error) {
	return r.store.CreateTeam(context, ToNewDBTeam(cmd))
}

func (r *postgresRepository) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := r.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return FromDBUser(dbUser), nil
}
