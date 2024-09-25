package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"go.uber.org/zap"
)

type Config struct {
	AuthConfig config.Auth
	Logger     *zap.SugaredLogger
	UserRepo   *users.PostgresRepository
}

type Auth struct {
	cfg      config.Auth
	parser   *jwt.Parser
	method   jwt.SigningMethod
	logger   *zap.SugaredLogger
	userRepo *users.PostgresRepository
}

func New(cfg Config) *Auth {
	return &Auth{
		logger:   cfg.Logger,
		userRepo: cfg.UserRepo,
		cfg:      cfg.AuthConfig,
		method:   jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
			jwt.WithAudience(cfg.AuthConfig.Issuer),
			jwt.WithIssuer(cfg.AuthConfig.Issuer),
			jwt.WithExpirationRequired(),
			jwt.WithIssuedAt(),
		),
	}
}

func (a *Auth) GenerateAccessToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	signedToken, err := token.SignedString(a.cfg.AccessTokenSecret)

	if err != nil {
		return "", NewAuthError(
			"Internal Server Error", errors.InternalServerError, http.StatusInternalServerError,
		)
	}

	return signedToken, nil
}

func (a *Auth) GenerateRefreshToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	signedToken, err := token.SignedString(a.cfg.RefreshTokenSecret)

	if err != nil {
		return "", NewAuthError(
			"Internal Server Error", errors.InternalServerError, http.StatusInternalServerError,
		)
	}

	return signedToken, nil
}

func (a *Auth) Authenticate(context context.Context, bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return Claims{}, NewAuthError(
			"Authorization header missing. Expected format 'Bearer <token>'",
			errors.InvalidAuthHeader,
			http.StatusUnauthorized,
		)
	}

	var claims Claims
	tokenStr := parts[1]

	if _, err := a.parser.ParseWithClaims(
		tokenStr, &claims, func(t *jwt.Token,
		) (interface{}, error) {
			if t.Method != a.method {
				return "", NewAuthError(
					"Invalid token signature", errors.InvalidTokenSignature, http.StatusUnauthorized,
				)
			}
			return a.cfg.AccessTokenSecret, nil
		}); err != nil {
		return Claims{}, NewAuthError(
			"Invalid token", errors.InvalidTokenSignature, http.StatusUnauthorized,
		)
	}

	return claims, nil
}
