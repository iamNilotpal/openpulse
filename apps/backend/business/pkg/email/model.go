package email

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

type SendOptions struct {
	To   string
	From string
}
