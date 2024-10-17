package auth

import (
	"context"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type ctxKey int

const userKey ctxKey = 1
const resourcesKey ctxKey = 2

func SetUser(ctx context.Context, user users.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUser(ctx context.Context) users.User {
	v, ok := ctx.Value(userKey).(users.User)
	if !ok {
		return users.User{}
	}
	return v
}

func SetResourcesMap(ctx context.Context, resPermissions UserAccessControlMap) context.Context {
	return context.WithValue(ctx, resourcesKey, resPermissions)
}

func GetResourcesMap(ctx context.Context) UserAccessControlMap {
	v, ok := ctx.Value(resourcesKey).(UserAccessControlMap)
	if !ok {
		return UserAccessControlMap{}
	}

	return v
}
