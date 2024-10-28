package auth

import (
	"context"

	"github.com/iamNilotpal/openpulse/business/repositories/users"
)

type ctxKey int

const userKey ctxKey = 1
const accessControlKey ctxKey = 2

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

func SetUserAccessControl(ctx context.Context, accessControl UserAccessControlMap) context.Context {
	return context.WithValue(ctx, accessControlKey, accessControl)
}

func GetUserAccessControl(ctx context.Context) UserAccessControlMap {
	v, ok := ctx.Value(accessControlKey).(UserAccessControlMap)
	if !ok {
		return UserAccessControlMap{}
	}
	return v
}
