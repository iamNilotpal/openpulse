package auth_handlers

import (
	"net/http"
	"strconv"

	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/email"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Auth                        *auth.Auth
	EmailService                *email.Email
	AuthCfg                     *config.Auth
	UsersRepo                   users.Repository
	RolesMap                    auth.RoleConfigMap
	RoleResourcesPermissionsMap auth.RoleResourcesPermissionsMap
}

type handler struct {
	auth                        *auth.Auth
	authCfg                     *config.Auth
	emailService                *email.Email
	usersRepo                   users.Repository
	rolesMap                    auth.RoleConfigMap
	RoleResourcesPermissionsMap auth.RoleResourcesPermissionsMap
}

func New(cfg Config) *handler {
	return &handler{
		auth:                        cfg.Auth,
		authCfg:                     cfg.AuthCfg,
		rolesMap:                    cfg.RolesMap,
		usersRepo:                   cfg.UsersRepo,
		emailService:                cfg.EmailService,
		RoleResourcesPermissionsMap: cfg.RoleResourcesPermissionsMap,
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload RegisterUserPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	adminRole := h.rolesMap[roles.RoleTeamAdminString]
	userId, err := h.usersRepo.Create(
		r.Context(),
		users.NewUser{
			PasswordHash: []byte{},
			RoleId:       adminRole.Id,
			Email:        payload.Email,
			LastName:     payload.LastName,
			FirstName:    payload.FirstName,
		},
	)

	if err != nil {
		if ok := database.CheckPQError(
			err,
			func(err *pq.Error) bool {
				return err.Column == "email" && err.Code == pgerrcode.UniqueViolation
			},
		); ok {
			return errors.NewRequestError(
				"User with same email already exists.",
				http.StatusBadRequest,
				errors.DuplicateValue,
			)
		}

		return err
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Account registered successfully.",
		RegisterUserResponse{UserId: userId},
	)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handler) VerifyEmail(w http.ResponseWriter, r *http.Request) error {
	var payload VerifyEmailPayload
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	claims, err := h.emailService.VerifyToken(payload.Token)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return err
	}

	return web.Success(w, http.StatusOK, "", userId)
}
