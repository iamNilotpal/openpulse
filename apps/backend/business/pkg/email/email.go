package email

import (
	stdErrors "errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"go.uber.org/zap"
)

type Config struct {
	Config *config.Email
	Logger *zap.SugaredLogger
}

type Email struct {
	parser *jwt.Parser
	config *config.Email
	method jwt.SigningMethod
	logger *zap.SugaredLogger
}

func New(cfg Config) *Email {
	return &Email{
		config: cfg.Config,
		logger: cfg.Logger,
		method: jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
			jwt.WithAudience(cfg.Config.Issuer),
			jwt.WithIssuer(cfg.Config.Issuer),
			jwt.WithExpirationRequired(),
			jwt.WithIssuedAt(),
		),
	}
}

func (e *Email) NewToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(e.method, c)

	signedToken, err := token.SignedString([]byte(e.config.Secret))
	if err != nil {
		return "", errors.NewRequestError(
			"Internal Server Error.", http.StatusInternalServerError, errors.InternalServerError,
		)
	}

	return signedToken, nil
}

func (e *Email) Send(opts SendOptions) error {
	return nil
}

func (e *Email) VerifyToken(token string) (Claims, error) {
	var claims Claims
	_, err := e.parser.ParseWithClaims(
		token,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != e.method {
				return "", stdErrors.New("invalid token")
			}
			return []byte(e.config.Secret), nil
		},
	)

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, errors.NewRequestError(
				"Invitation expired.", http.StatusBadRequest, errors.TokenExpired,
			)
		}
		return Claims{}, errors.NewRequestError(
			"Invalid token.", http.StatusBadRequest, errors.InvalidTokenSignature,
		)
	}

	return claims, nil
}
