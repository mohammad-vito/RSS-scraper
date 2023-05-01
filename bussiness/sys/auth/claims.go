package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
)

type ctxKey int

const key ctxKey = 1

func SetClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, key, claims)
}

func GetClaims(ctx context.Context) (jwt.Claims, error) {
	v, ok := ctx.Value(key).(jwt.Claims)
	if !ok {
		return jwt.StandardClaims{}, errors.New("claim value missing from context")
	}
	return v, nil
}
