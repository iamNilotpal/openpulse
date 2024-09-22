package auth

import (
	"context"
	"errors"
)

type ctxKey int

const userKey ctxKey = 1
const claimKey ctxKey = 2

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

func SetUserID(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, userKey, userId)
}

func GetUserId(ctx context.Context) (int, error) {
	v, ok := ctx.Value(userKey).(int)
	if !ok {
		return 0, errors.New("not found")
	}
	return v, nil
}
