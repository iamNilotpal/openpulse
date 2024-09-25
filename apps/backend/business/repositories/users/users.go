package users

import (
	"context"
	"database/sql"
	"errors"

	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
)

var (
	ErrNotFound   = errors.New("user not found")
	ErrUpdateUser = errors.New("update user data failed")
)

type Repository interface {
	QueryById(context context.Context, id int) (User, error)
	Create(context context.Context, payload NewUser) (int, error)
	QueryByEmail(context context.Context, email string) (User, error)
	Query(context context.Context, query QueryFilter) ([]User, error)
}

type PostgresRepository struct {
	store users_store.Store
}

func NewPostgresRepository(store users_store.Store) *PostgresRepository {
	return &PostgresRepository{store: store}
}

func (c *PostgresRepository) Create(context context.Context, payload NewUser) (int, error) {
	id, err := c.store.Create(context, ToNewDBUser(payload))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotFound
		}
		return 0, err
	}

	return id, nil
}

func (c *PostgresRepository) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := c.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return ToUser(dbUser), nil
}

func (c *PostgresRepository) QueryByEmail(context context.Context, email string) (User, error) {
	dbUser, err := c.store.QueryByEmail(context, email)
	if err != nil {
		return User{}, err
	}

	return ToUser(dbUser), nil
}

func (c *PostgresRepository) Query(context context.Context, query QueryFilter) ([]User, error) {
	if err := query.Validate(); err != nil {
		return []User{}, err
	}

	return []User{}, nil
}
