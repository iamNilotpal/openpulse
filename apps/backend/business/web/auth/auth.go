package auth

import (
	"context"
	stdErrors "errors"
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"go.uber.org/zap"
)

type Config struct {
	AuthConfig config.Auth
	UserRepo   users.Repository
	Logger     *zap.SugaredLogger
}

type Auth struct {
	cfg      config.Auth
	parser   *jwt.Parser
	userRepo users.Repository
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
	parts := strings.Split(strings.TrimSpace(bearerToken), " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return Claims{}, NewAuthError(
			"Authorization header missing. Expected format 'Bearer <token>'",
			errors.InvalidAuthHeader,
			http.StatusUnauthorized,
		)
	}

	var claims Claims
	tokenStr := parts[1]

	_, err := a.parser.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != a.method {
				return "", NewAuthError(
					"Invalid token signature", errors.InvalidTokenSignature, http.StatusUnauthorized,
				)
			}
			return a.cfg.AccessTokenSecret, nil
		},
	)

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, NewAuthError(
				"Session expired.", errors.TokenExpired, http.StatusUnauthorized,
			)
		}

		return Claims{}, NewAuthError(
			"Invalid token", errors.InvalidTokenSignature, http.StatusUnauthorized,
		)
	}

	return claims, nil
}

func CheckRoleAccessControl(requiredRole RoleConfig, userRole UserRoleConfig) bool {
	if requiredRole.Id != userRole.Id || requiredRole.Role != userRole.Role {
		return false
	}
	return true
}

func CheckResourceAccessControl(
	requiredResource ResourceConfig, userResource UserResourceConfig,
) bool {
	if requiredResource.Id != userResource.Id || requiredResource.Resource != userResource.Resource {
		return false
	}
	return true
}

func CheckPermissionAccessControl(
	strict bool,
	userPermissions []UserPermissionConfig,
	requiredPermissions []PermissionConfig,
) bool {
	if len(requiredPermissions) == 0 {
		return true
	}

	if len(userPermissions) == 0 || (strict && len(requiredPermissions) > len(userPermissions)) {
		return false
	}

	for _, rp := range requiredPermissions {
		index := slices.IndexFunc(
			userPermissions,
			func(v UserPermissionConfig) bool {
				return v.Id == rp.Id && v.Action == rp.Action
			},
		)

		if index == -1 {
			return false
		}

		if up := userPermissions[index]; !up.Enabled {
			return false
		}
	}

	return true
}
