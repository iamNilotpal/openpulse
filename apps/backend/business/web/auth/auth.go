package auth

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/user"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"go.uber.org/zap"
)

type Claims struct {
	jwt.RegisteredClaims
}

type Config struct {
	AuthConfig config.Auth
	UserRepo   *user.Repository
	Logger     *zap.SugaredLogger
}

type Auth struct {
	cfg      config.Auth
	parser   *jwt.Parser
	userRepo *user.Repository
	method   jwt.SigningMethod
	logger   *zap.SugaredLogger
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

func (a *Auth) GenerateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)

	signedToken, err := token.SignedString(a.cfg.Secret)
	if err != nil {
		return "", NewAuthError("Internal server error", errors.InternalErrorCode)
	}

	return signedToken, nil
}

func (a *Auth) Authenticate(context context.Context, bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return Claims{}, NewAuthError(
			"Authorization header missing. Expected format 'Bearer <token>'.", errors.InvalidAuthHeaderErrorCode,
		)
	}

	var claims Claims
	tokenStr := parts[1]

	if _, err := a.parser.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != a.method {
			return "", NewAuthError("Invalid token signature", errors.InvalidTokenSignature)
		}
		return a.cfg.Secret, nil
	}); err != nil {
		return Claims{}, NewAuthError("Invalid token", errors.InvalidTokenSignature)
	}

	return claims, nil
}
