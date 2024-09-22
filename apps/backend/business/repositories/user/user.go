package user

import (
	"context"
	"database/sql"
	"errors"

	user_store "github.com/iamNilotpal/openpulse/business/repositories/user/stores/db"
)

var (
	ErrNotFound   = errors.New("user not found")
	ErrUpdateUser = errors.New("update user data failed")
)

type Repository struct {
	store user_store.Store
}

func NewRepository(store user_store.Store) *Repository {
	return &Repository{store: store}
}

func (c *Repository) Create(context context.Context, payload NewUser) (int, error) {
	id, err := c.store.Create(context, ToNewDBUser(payload))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotFound
		}
		return 0, err
	}

	return id, nil
}

func (c *Repository) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := c.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return ToUser(dbUser), nil
}

func (c *Repository) QueryByEmail(context context.Context, email string) (User, error) {
	dbUser, err := c.store.QueryByEmail(context, email)
	if err != nil {
		return User{}, err
	}

	return ToUser(dbUser), nil
}

func (c *Repository) Query(context context.Context, query QueryFilter) ([]User, error) {
	if err := query.Validate(); err != nil {
		return []User{}, err
	}

	return []User{}, nil
}
