package user

import (
	"context"
	"database/sql"
	"errors"

	user_store "github.com/iamNilotpal/openpulse/business/core/user/store/db"
)

var (
	ErrNotFound   = errors.New("user not found")
	ErrUpdateUser = errors.New("update user data failed")
)

type Core struct {
	store user_store.Store
}

func NewCore(store user_store.Store) *Core {
	return &Core{store: store}
}

func (c *Core) Create(context context.Context, payload NewUser) (int, error) {
	id, err := c.store.Create(context, ToNewDBUser(payload))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotFound
		}
		return 0, err
	}

	return id, nil
}

func (c *Core) QueryById(context context.Context, id int) (User, error) {
	dbUser, err := c.store.QueryById(context, id)
	if err != nil {
		return User{}, err
	}

	return ToUser(dbUser), nil
}

func (c *Core) Query(context context.Context, query QueryFilter) ([]User, error) {
	if err := query.Validate(); err != nil {
		return []User{}, err
	}

	return []User{}, nil
}
