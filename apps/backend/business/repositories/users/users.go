package users

import (
	"context"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type Repository interface {
	Create(context context.Context, payload NewUser) (int, error)
	QueryById(context context.Context, id int) (User, error)
}

type PostgresRepository struct {
	store users_store.Store
}

func NewPostgresRepository(store users_store.Store) *PostgresRepository {
	return &PostgresRepository{store: store}
}

func (c *PostgresRepository) Create(context context.Context, payload NewUser) (int, error) {
	id, err := c.store.Create(context, ToNewDBUser(payload))
	return id, err
}

func (c *PostgresRepository) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := c.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return FromDBUser(dbUser), nil
}
