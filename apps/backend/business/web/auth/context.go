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

func SetRole(ctx context.Context, role UserRoleConfig) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetRole(ctx context.Context) UserRoleConfig {
	v, ok := ctx.Value(roleKey).(UserRoleConfig)
	if !ok {
		return UserRoleConfig{}
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
