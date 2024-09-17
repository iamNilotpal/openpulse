package users

import (
	"context"

	users_store "github.com/iamNilotpal/openpulse/business/core/users/store/db"
)

type Repository interface {
	Create(context.Context, CreateUserPayload) (int, error)
	QueryByUserId(context context.Context, id int) (AppUser, error)
}

type Core struct {
	Store users_store.DBStore
}

func NewCore(store users_store.DBStore) *Core {
	return &Core{Store: store}
}

func (c *Core) Create(context context.Context, payload CreateUserPayload) (int, error) {
	return c.Store.Create(context, ToCreateDBUser(payload))
}

func (c *Core) QueryByUserId(context context.Context, id int) (AppUser, error) {
	dbUser, err := c.Store.QueryByUserId(context, id)
	if err != nil {
		return AppUser{}, err
	}

	return ToAppUser(dbUser), nil
}
