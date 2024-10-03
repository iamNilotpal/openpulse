package auth_handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type handler struct {
	auth           *auth.Auth
	cfg            config.Auth
	rolesMap       auth.AuthedRolesMap
	usersRepo      users.Repository
	permissionsMap auth.AuthedPermissionsMap
}

func New(
	cfg config.Auth,
	auth *auth.Auth,
	rolesMap auth.AuthedRolesMap,
	usersRepo users.Repository,
	permissionsMap auth.AuthedPermissionsMap,
) *handler {
	return &handler{
		cfg:            cfg,
		auth:           auth,
		rolesMap:       rolesMap,
		usersRepo:      usersRepo,
		permissionsMap: permissionsMap,
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload RegisterUserPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	admin := h.rolesMap["admin"]
	adminPerms := h.permissionsMap["admin"]
	permissions := make([]users.UserPermissions, 0, len(adminPerms))

	for i, p := range adminPerms {
		permissions[i] = users.UserPermissions{
			Role:       users.Role{Id: p.Role.Id},
			Permission: users.Permission{Id: p.Permission.Id},
		}
	}

	userId, err := h.usersRepo.Create(
		r.Context(),
		users.NewUser{
			AvatarUrl:    "",
			RoleId:       admin.Id,
			LastName:     payload.LastName,
			FirstName:    payload.FirstName,
			Email:        payload.Email,
			PasswordHash: []byte{},
		},
		permissions,
	)

	if err != nil {
		if ok := database.CheckPQError(err, func(err *pq.Error) bool {
			return err.Code == pgerrcode.UniqueViolation
		}); ok {
			return errors.NewRequestError(
				"User with same email already exists",
				http.StatusBadRequest,
				errors.DuplicateValue,
			)
		}

		return err
	}

	accessToken, err := h.auth.GenerateAccessToken(
		auth.Claims{
			UserId: userId,
			RoleId: admin.Id,
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
			UserId: userId,
			RoleId: admin.Id,
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

	return web.Success(
		w,
		http.StatusCreated,
		RegisterUserResponse{
			UserId:       userId,
			RoleId:       admin.Id,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) error {
	return nil
}
