package auth_handlers

import (
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/foundation/web"
)

type handler struct {
	auth           *auth.Auth
	cfg            config.Auth
	usersRepo      users.Repository
	permissionsMap auth.PermissionsMap
}

func New(
	cfg config.Auth,
	auth *auth.Auth,
	usersRepo users.Repository,
	permissionsMap auth.PermissionsMap,
) *handler {
	return &handler{cfg: cfg, auth: auth, usersRepo: usersRepo, permissionsMap: permissionsMap}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload RegisterUserPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	userId, err := h.usersRepo.Create(r.Context(), users.NewUser{
		RoleId:       1,
		AvatarUrl:    "",
		PasswordHash: []byte{},
		LastName:     payload.LastName,
		FirstName:    payload.FirstName,
		Email:        mail.Address{Address: payload.Email},
	})

	if err != nil {
		return err
	}

	accessToken, err := h.auth.GenerateAccessToken(
		auth.Claims{
			RoleId: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    h.cfg.Issuer,
				Subject:   strconv.Itoa(userId),
				Audience:  jwt.ClaimStrings{h.cfg.Issuer},
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.cfg.AccessTokenExpTime)),
			},
		},
	)
	if err != nil {
		return err
	}

	refreshToken, err := h.auth.GenerateRefreshToken(
		auth.Claims{
			RoleId: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    h.cfg.Issuer,
				Subject:   strconv.Itoa(userId),
				Audience:  jwt.ClaimStrings{h.cfg.Issuer},
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.cfg.RefreshTokenExpTime)),
			},
		},
	)
	if err != nil {
		return err
	}

	return web.Success(w, http.StatusCreated, RegisterUserResponse{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserEmail:    payload.Email,
	})
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) error {
	return nil
}
