package auth

import (
	"context"
)

type ctxKey int

const roleKey ctxKey = 1
const claimKey ctxKey = 2
const resourcesKey ctxKey = 3

func SetClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

func GetClaims(ctx context.Context) Claims {
	v, ok := ctx.Value(claimKey).(Claims)
	if !ok {
		return Claims{}
	}
	return v
}

func SetUserRole(ctx context.Context, role UserRoleConfig) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetUserRole(ctx context.Context) UserRoleConfig {
	v, ok := ctx.Value(roleKey).(UserRoleConfig)
	if !ok {
		return UserRoleConfig{}
	}

	return v
}

func SetUserResources(ctx context.Context, resources UserResourcePermissionsMap) context.Context {
	return context.WithValue(ctx, resourcesKey, resources)
}

func GetUserResources(ctx context.Context) UserResourcePermissionsMap {
	v, ok := ctx.Value(resourcesKey).(UserResourcePermissionsMap)
	if !ok {
		return UserResourcePermissionsMap{}
	}

	return v
}
