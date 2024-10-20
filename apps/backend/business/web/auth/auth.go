package auth

import (
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
	AuthConfig    *config.Auth
	UserRepo      users.Repository
	OnboardingCfg *config.Onboarding
	Logger        *zap.SugaredLogger
}

type Auth struct {
	parser        *jwt.Parser
	method        jwt.SigningMethod
	logger        *zap.SugaredLogger
	authCfg       *config.Auth
	onboardingCfg *config.Onboarding
	userRepo      users.Repository
}

func New(cfg Config) *Auth {
	return &Auth{
		logger:        cfg.Logger,
		userRepo:      cfg.UserRepo,
		authCfg:       cfg.AuthConfig,
		onboardingCfg: cfg.OnboardingCfg,
		method:        jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
			jwt.WithAudience(cfg.AuthConfig.Issuer),
			jwt.WithIssuer(cfg.AuthConfig.Issuer),
			jwt.WithExpirationRequired(),
			jwt.WithIssuedAt(),
		),
	}
}

func generateToken(token *jwt.Token, secret string) (string, error) {
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", NewAuthError(
			"Internal Server Error", errors.InternalServerError, http.StatusInternalServerError,
		)
	}
	return signedToken, nil
}

func (a *Auth) NewAccessToken(claims AccessTokenClaims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	return generateToken(token, a.authCfg.AccessTokenSecret)
}

func (a *Auth) NewRefreshToken(claims RefreshTokenClaims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	return generateToken(token, a.authCfg.RefreshTokenSecret)
}

func (a *Auth) NewOnboardingToken(claims OnBoardingClaims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	return generateToken(token, a.onboardingCfg.Secret)
}

func (a *Auth) Authenticate(bearerToken string) (AccessTokenClaims, error) {
	parts := strings.Split(strings.TrimSpace(bearerToken), " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return AccessTokenClaims{}, NewAuthError(
			"Authorization header missing. Expected format 'Bearer <token>'",
			errors.InvalidAuthHeader,
			http.StatusUnauthorized,
		)
	}

	var claims AccessTokenClaims
	tokenStr := parts[1]

	_, err := a.parser.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != a.method {
				return "", stdErrors.New("invalid token")
			}
			return a.authCfg.AccessTokenSecret, nil
		},
	)

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return AccessTokenClaims{}, NewAuthError(
				"Session expired.", errors.TokenExpired, http.StatusUnauthorized,
			)
		}
		return AccessTokenClaims{}, NewAuthError(
			"Invalid token", errors.InvalidTokenSignature, http.StatusUnauthorized,
		)
	}

	return claims, nil
}

func (a *Auth) AuthenticateOnboard(bearerToken string) (OnBoardingClaims, error) {
	parts := strings.Split(strings.TrimSpace(bearerToken), " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return OnBoardingClaims{}, NewAuthError(
			"Authorization header missing. Expected format 'Bearer <token>'",
			errors.InvalidAuthHeader,
			http.StatusUnauthorized,
		)
	}

	var claims OnBoardingClaims
	tokenStr := parts[1]

	_, err := a.parser.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != a.method {
				return "", stdErrors.New("invalid token")
			}
			return a.onboardingCfg.Secret, nil
		},
	)

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return OnBoardingClaims{}, NewAuthError(
				"Session expired.", errors.TokenExpired, http.StatusUnauthorized,
			)
		}
		return OnBoardingClaims{}, NewAuthError(
			"Invalid token", errors.InvalidTokenSignature, http.StatusUnauthorized,
		)
	}

	return claims, nil
}

func CheckRoleAccessControl(requiredRoles []RoleConfig, userRole UserRoleConfig) bool {
	for _, rr := range requiredRoles {
		if rr.Id == userRole.Id &&
			strings.EqualFold(string(rr.Role), string(userRole.Role)) {
			return true
		}
	}
	return false
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
