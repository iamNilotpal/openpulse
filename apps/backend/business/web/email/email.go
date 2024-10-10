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
	Cfg    config.Email
	Logger *zap.SugaredLogger
}

type Email struct {
	parser *jwt.Parser
	cfg    config.Email
	method jwt.SigningMethod
	logger *zap.SugaredLogger
}

func New(cfg Config) *Email {
	return &Email{
		logger: cfg.Logger,
		cfg:    cfg.Cfg,
		method: jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
			jwt.WithAudience(cfg.Cfg.Issuer),
			jwt.WithIssuer(cfg.Cfg.Issuer),
			jwt.WithExpirationRequired(),
			jwt.WithIssuedAt(),
		),
	}
}

func (e *Email) GenerateVerificationToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(e.method, c)
	signedToken, err := token.SignedString(e.cfg.Secret)

	if err != nil {
		return "", errors.NewRequestError(
			"Internal Server Error", http.StatusInternalServerError, errors.InternalServerError,
		)
	}

	return signedToken, nil
}

func (e *Email) SendVerificationMail() error { return nil }

func (e *Email) VerifyToken(token string) (Claims, error) {
	var claims Claims

	_, err := e.parser.ParseWithClaims(
		token,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != e.method {
				return "", errors.NewRequestError(
					"Invalid token signature", http.StatusUnauthorized, errors.InvalidTokenSignature,
				)
			}
			return e.cfg.Secret, nil
		},
	)

	if err != nil {
		if stdErrors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, errors.NewRequestError(
				"Session expired.", http.StatusBadRequest, errors.TokenExpired,
			)
		}

		return Claims{}, errors.NewRequestError(
			"Invalid token", http.StatusBadRequest, errors.InvalidTokenSignature,
		)
	}

	return claims, nil
}
