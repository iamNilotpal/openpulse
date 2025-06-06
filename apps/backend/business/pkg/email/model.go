package email

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type SendOptions struct {
	To   string
	From string
}
