package auth_handlers

import (
	stdErrors "errors"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/repositories/emails"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/email"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/hash"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Auth                        *auth.Auth
	EmailService                *email.Email
	HashService                 hash.Hasher
	UsersRepo                   users.Repository
	EmailsRepo                  emails.Repository
	RolesMap                    auth.RoleConfigMap
	RoleResourcesPermissionsMap auth.RoleAccessControlMap
	Config                      *config.OpenpulseAPIConfig
}

type handler struct {
	auth                        *auth.Auth
	emailService                *email.Email
	hashService                 hash.Hasher
	usersRepo                   users.Repository
	emailsRepo                  emails.Repository
	rolesMap                    auth.RoleConfigMap
	RoleResourcesPermissionsMap auth.RoleAccessControlMap
	config                      *config.OpenpulseAPIConfig
}

func New(cfg Config) *handler {
	return &handler{
		auth:                        cfg.Auth,
		config:                      cfg.Config,
		rolesMap:                    cfg.RolesMap,
		usersRepo:                   cfg.UsersRepo,
		emailsRepo:                  cfg.EmailsRepo,
		hashService:                 cfg.HashService,
		emailService:                cfg.EmailService,
		RoleResourcesPermissionsMap: cfg.RoleResourcesPermissionsMap,
	}
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var input SignUpInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	orgAdmin := h.rolesMap[roles.RoleOrgAdmin]
	passwordHash, err := h.hashService.Hash([]byte(input.Password))

	if err != nil {
		return errors.NewRequestError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	userId, err := h.usersRepo.Create(
		r.Context(),
		users.NewUser{
			RoleId:       orgAdmin.Id,
			PasswordHash: passwordHash,
			Email:        input.Email,
			LastName:     input.LastName,
			FirstName:    input.FirstName,
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

	token, err := h.emailService.GenerateVerificationToken(email.Claims{
		Email: input.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    h.config.Auth.Issuer,
			Subject:   strconv.Itoa(userId),
			Audience:  jwt.ClaimStrings{h.config.Auth.Audience},
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Email.TokenExpTime)),
		},
	})
	if err != nil {
		return web.Error(
			w,
			http.StatusInternalServerError,
			web.NewAPIError(
				"Error while sending verification mail.",
				errors.FromErrorCode(errors.FlowIncomplete),
				RegisterUserResponse{
					State: RegistrationState{
						EmailSentState:    AUTH_STATE_VERIFICATION_MAIL_NOT_SENT,
						RegistrationState: AUTH_STATE_USER_REGISTRATION_INCOMPLETE,
					},
				},
			),
		)
	}

	if err = h.emailService.SendVerificationMail(); err != nil {
		return web.Error(
			w,
			http.StatusInternalServerError,
			web.NewAPIError(
				"Error while sending verification mail.",
				errors.FromErrorCode(errors.FlowIncomplete),
				RegisterUserResponse{
					State: RegistrationState{
						EmailSentState:    AUTH_STATE_VERIFICATION_MAIL_NOT_SENT,
						RegistrationState: AUTH_STATE_USER_REGISTRATION_COMPLETE,
					},
				},
			),
		)
	}

	if err = h.emailsRepo.SaveEmailVerificationDetails(
		r.Context(),
		emails.EmailVerificationDetails{
			MaxAttempts:       5,
			VerificationToken: token,
			UserId:            userId,
			Email:             input.Email,
			ExpiresAt:         time.Now().Add(time.Minute * 30),
		},
	); err != nil {
		return web.Error(
			w,
			http.StatusInternalServerError,
			web.NewAPIError(
				http.StatusText(http.StatusInternalServerError),
				errors.FromErrorCode(errors.FlowIncomplete),
				RegisterUserResponse{
					State: RegistrationState{
						EmailSentState:    AUTH_STATE_VERIFICATION_MAIL_NOT_SENT,
						RegistrationState: AUTH_STATE_USER_REGISTRATION_COMPLETE,
					},
				},
			),
		)
	}

	return web.Success(
		w,
		http.StatusCreated,
		"Account registered successfully.",
		RegisterUserResponse{
			UserId: userId,
			State: RegistrationState{
				EmailSentState:    AUTH_STATE_VERIFICATION_MAIL_SENT,
				RegistrationState: AUTH_STATE_USER_REGISTRATION_COMPLETE,
			},
		},
	)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) error {
	var payload SignInInput
	if err := web.Decode(r, &payload); err != nil {
		return err
	}

	return nil
}

func (h *handler) VerifyEmail(w http.ResponseWriter, r *http.Request) error {
	token := web.GetParam(r, "invitationToken")
	if token == "" {
		return errors.NewRequestError(
			"Invalid invitation token.", http.StatusBadRequest, errors.BadRequest,
		)
	}

	claims, err := h.emailService.VerifyToken(token)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return errors.NewRequestError(
			"Invalid token signature.", http.StatusBadRequest, errors.InvalidTokenSignature,
		)
	}

	if err = h.emailsRepo.ValidateVerificationDetails(
		r.Context(), token, userId, int(claims.ExpiresAt.UnixNano()),
	); err != nil {
		if stdErrors.Is(err, emails.ErrVerificationDataNotFound) {
			return errors.NewRequestError("Invitation expired.", http.StatusNotFound, errors.NotFound)
		}
		if stdErrors.Is(err, emails.ErrVerificationLimitExceed) {
			return errors.NewRequestError("Invitation expired.", http.StatusNotFound, errors.NotFound)
		}
		return err
	}

	return web.Success(w, http.StatusOK, "Email verified successfully.", userId)
}
