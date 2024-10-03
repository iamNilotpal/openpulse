package users

import (
	"context"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

type Repository interface {
	QueryById(context context.Context, id int) (User, error)
	Create(context context.Context, payload NewUser, permissions []UserPermissions) (int, error)
	QueryByEmail(context context.Context, email string) (User, error)
	Query(context context.Context, query QueryFilter) ([]User, error)
}

type PostgresRepository struct {
	store users_store.Store
}

func NewPostgresRepository(store users_store.Store) *PostgresRepository {
	return &PostgresRepository{store: store}
}

func (c *PostgresRepository) Create(
	context context.Context, payload NewUser, permissions []UserPermissions,
) (int, error) {
	perms := make([]users_store.UserPermissions, 0, len(permissions))
	for i, p := range permissions {
		perms[i] = ToDBUserPermission(p)
	}

	id, err := c.store.Create(context, ToNewDBUser(payload), perms)
	return id, err
}

func (c *PostgresRepository) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := c.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return FromDBUser(dbUser), nil
}

func (c *PostgresRepository) QueryByEmail(context context.Context, email string) (User, error) {
	dbUser, err := c.store.QueryByEmail(context, email)
	if err != nil {
		return User{}, err
	}

	return FromDBUser(dbUser), nil
}

func (c *PostgresRepository) Query(context context.Context, query QueryFilter) ([]User, error) {
	if err := query.Validate(); err != nil {
		return []User{}, err
	}

	return []User{}, nil
}
