package auth

import (
	"context"
)

type ctxKey int

const roleKey ctxKey = 1
const claimKey ctxKey = 2
const permissionsKey ctxKey = 3

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

func SetUserRole(ctx context.Context, role UserRole) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetUserRole(ctx context.Context) UserRole {
	v, ok := ctx.Value(roleKey).(UserRole)
	if !ok {
		return UserRole{}
	}

	return v
}

func SetUserPermissions(ctx context.Context, permissions []UserAccessControl) context.Context {
	return context.WithValue(ctx, permissionsKey, permissions)
}

func GetUserPermissions(ctx context.Context) []UserAccessControl {
	v, ok := ctx.Value(permissionsKey).([]UserAccessControl)
	if !ok {
		return []UserAccessControl{}
	}

	return v
}
