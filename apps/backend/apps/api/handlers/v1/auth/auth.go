package auth_handlers

import (
	"database/sql"
	stdErrors "errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamNilotpal/openpulse/business/pkg/email"
	"github.com/iamNilotpal/openpulse/business/repositories/emails"
	"github.com/iamNilotpal/openpulse/business/repositories/organizations"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	"github.com/iamNilotpal/openpulse/business/repositories/sessions"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/errors"
	"github.com/iamNilotpal/openpulse/foundation/hash"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

type Config struct {
	Auth          *auth.Auth
	EmailService  *email.Email
	Config        *config.APIConfig
	HashService   hash.Hasher
	Users         users.Repository
	Emails        emails.Repository
	Sessions      sessions.Repository
	Organizations organizations.Repository
	RoleMap       auth.RoleConfigMap
	RBACMap       auth.RBACMap
}

type handler struct {
	auth          *auth.Auth
	emailService  *email.Email
	hashService   hash.Hasher
	config        *config.APIConfig
	users         users.Repository
	sessions      sessions.Repository
	organizations organizations.Repository
	emails        emails.Repository
	rolesMap      auth.RoleConfigMap
	rbacMap       auth.RBACMap
}

func New(cfg Config) *handler {
	return &handler{
		auth:          cfg.Auth,
		config:        cfg.Config,
		sessions:      cfg.Sessions,
		rolesMap:      cfg.RoleMap,
		users:         cfg.Users,
		emails:        cfg.Emails,
		organizations: cfg.Organizations,
		hashService:   cfg.HashService,
		emailService:  cfg.EmailService,
		rbacMap:       cfg.RBACMap,
	}
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var input SignUpInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	passwordHash, err := h.hashService.Hash([]byte(input.Password))
	if err != nil {
		return errors.NewRequestError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	userId, err := h.users.Create(
		r.Context(),
		users.NewUser{
			FirstName:    input.FirstName,
			LastName:     input.LastName,
			Email:        input.Email,
			PasswordHash: passwordHash,
			RoleId:       h.rolesMap[roles.RoleOrgAdmin].Id,
		},
	)
	if err != nil {
		if err := database.CheckPQError(
			err,
			func(err *pq.Error) error {
				if err.Column == "email" && err.Code == pgerrcode.UniqueViolation {
					return errors.NewRequestError(
						"User with same email already exists.",
						http.StatusConflict,
						errors.DuplicateValue,
					)
				}
				return nil
			},
		); err != nil {
			return err
		}
		return err
	}

	token, err := h.emailService.NewToken(
		email.Claims{
			Email: input.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    h.config.Auth.Issuer,
				Subject:   strconv.Itoa(userId),
				Audience:  jwt.ClaimStrings{h.config.Auth.Audience},
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Email.EmailExpTime)),
			},
		},
	)
	if err != nil {
		return web.Error(
			w,
			http.StatusInternalServerError,
			web.NewAPIError(
				"Error while sending verification mail.",
				errors.FromErrorCode(errors.FlowIncomplete),
				RegisterUserResponse{
					State: RegistrationState{
						EmailSentState:    verificationMailNotSent,
						RegistrationState: userRegistrationIncomplete,
					},
				},
			),
		)
	}

	if err = h.emailService.Send(email.SendOptions{}); err != nil {
		return web.Error(
			w,
			http.StatusInternalServerError,
			web.NewAPIError(
				"Error while sending verification mail.",
				errors.FromErrorCode(errors.FlowIncomplete),
				RegisterUserResponse{
					State: RegistrationState{
						EmailSentState:    verificationMailNotSent,
						RegistrationState: userRegistrationComplete,
					},
				},
			),
		)
	}

	if err = h.emails.SaveEmailVerificationDetails(
		r.Context(),
		emails.EmailVerificationDetails{
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
						EmailSentState:    verificationMailNotSent,
						RegistrationState: userRegistrationComplete,
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
			URL:    h.config.Web.ClientAPIHost + "verify/email" + "/" + token,
			State: RegistrationState{
				EmailSentState:    verificationMailSent,
				RegistrationState: userRegistrationComplete,
			},
		},
	)
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) error {
	var input SignInInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	user, err := h.users.QueryByEmail(r.Context(), input.Email)
	if err != nil {
		if stdErrors.Is(err, sql.ErrNoRows) {
			return errors.NewRequestError(
				"Invalid email or password.", http.StatusUnauthorized, errors.Unauthorized,
			)
		}
		return err
	}

	if !h.hashService.Compare([]byte(user.Password), []byte(input.Password)) {
		return errors.NewRequestError(
			"Invalid email or password.", http.StatusUnauthorized, errors.Unauthorized,
		)
	}

	if !user.IsEmailVerified {
		return errors.NewRequestError(
			"Please verify your email to sign in.", http.StatusUnauthorized, errors.UserNotVerified,
		)
	}

	isCompleted, err := h.organizations.CheckOnboardingStatus(r.Context(), user.Id)
	if err != nil {
		return errors.NewRequestError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	onboardingToken, err := h.auth.NewOnboardingToken(
		auth.OnBoardingClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   strconv.Itoa(user.Id),
				Issuer:    h.config.Onboarding.Issuer,
				Audience:  jwt.ClaimStrings{h.config.Onboarding.Audience},
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Onboarding.TokenExpTime)),
			},
		},
	)
	if err != nil {
		return errors.NewRequestError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	if !isCompleted && user.Team.Id == 0 {
		return web.Error(
			w,
			http.StatusBadRequest,
			web.NewAPIError(
				"Onboarding process not completed.",
				errors.FromErrorCode(errors.FlowIncomplete),
				IncompleteOnboardingResponse{
					UserId:           user.Id,
					AccessToken:      onboardingToken,
					OnboardingStatus: ONBOARDING_CREATE_ORGANIZATION,
				},
			),
		)
	}

	if user.Team.Id == 0 {
		return web.Error(
			w,
			http.StatusBadRequest,
			web.NewAPIError(
				"Onboarding process not completed.",
				errors.FromErrorCode(errors.FlowIncomplete),
				IncompleteOnboardingResponse{
					UserId:           user.Id,
					AccessToken:      onboardingToken,
					OnboardingStatus: ONBOARDING_CREATE_TEAM,
				},
			),
		)
	}

	aToken, err := h.auth.NewAccessToken(
		auth.AccessTokenClaims{
			RoleId: user.Role.Id,
			TeamId: user.Team.Id,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    h.config.Auth.Issuer,
				Subject:   strconv.Itoa(user.Id),
				Audience:  jwt.ClaimStrings{h.config.Auth.Audience},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Auth.AccessTokenExpTime)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	)
	if err != nil {
		return err
	}

	rToken, err := h.auth.NewRefreshToken(
		auth.RefreshTokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    h.config.Auth.Issuer,
				Subject:   strconv.Itoa(user.Id),
				Audience:  jwt.ClaimStrings{h.config.Auth.Audience},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Auth.AccessTokenExpTime)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	)
	if err != nil {
		return err
	}

	sessionId, err := h.sessions.Create(
		r.Context(),
		sessions.NewSession{
			Token:     rToken,
			UserId:    user.Id,
			IpAddress: r.RemoteAddr,
			UserAgent: r.UserAgent(),
		},
	)
	if err != nil {
		return err
	}

	return web.Success(
		w,
		http.StatusOK,
		"Logged in successfully.",
		SignInResponse{AccessToken: aToken, RefreshToken: rToken, UserId: user.Id, SessionId: sessionId},
	)
}

func (h *handler) OauthSignup(w http.ResponseWriter, r *http.Request) error {
	provider := web.GetParam(r, "provider")
	if provider == "" ||
		!strings.EqualFold(provider, providerGoogle) ||
		!strings.EqualFold(provider, providerGoogle) {
		return errors.NewRequestError("Unsupported provider.", http.StatusBadRequest, errors.BadRequest)
	}

	var input NewOAuthAccountInput
	if err := web.Decode(r, &input); err != nil {
		return err
	}

	_, err := h.users.CreateUsingOAuth(
		r.Context(),
		users.NewOAuthAccount{
			Provider:   provider,
			Scope:      input.Scope,
			Metadata:   input.Metadata,
			ExternalId: input.ExternalId,
			User: users.NewOAuthUser{
				Email:     input.User.Email,
				FirstName: input.User.FirstName,
				LastName:  input.User.LastName,
				Phone:     input.User.Phone,
				AvatarURL: input.User.AvatarURL,
				RoleId:    h.rolesMap[roles.RoleOrgAdmin].Id,
			},
		},
	)

	if err == nil {

	}

	if err != nil {
		if err := database.CheckPQError(
			err,
			func(err *pq.Error) error {
				return err
			},
		); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (h *handler) VerifyEmail(w http.ResponseWriter, r *http.Request) error {
	token := web.GetQuery(r, "token")
	if token == "" {
		return errors.NewRequestError("Missing verification token.", http.StatusBadRequest, errors.BadRequest)
	}

	claims, err := h.emailService.VerifyToken(token)
	if err != nil {
		return err
	}

	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return errors.NewRequestError("Invalid token.", http.StatusBadRequest, errors.InvalidTokenSignature)
	}

	if err = h.emails.ValidateVerificationDetails(
		r.Context(), token, userId, int(claims.ExpiresAt.UnixNano()),
	); err != nil {
		if stdErrors.Is(err, emails.ErrVerificationDataNotFound) ||
			stdErrors.Is(err, emails.ErrVerificationLimitExceed) {
			return errors.NewRequestError("Invitation expired.", http.StatusNotFound, errors.NotFound)
		}
		return err
	}

	onboardingToken, err := h.auth.NewOnboardingToken(
		auth.OnBoardingClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Subject:   claims.Subject,
				Issuer:    h.config.Onboarding.Issuer,
				Audience:  jwt.ClaimStrings{h.config.Onboarding.Audience},
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.config.Onboarding.TokenExpTime)),
			},
		},
	)
	if err != nil {
		return errors.NewRequestError(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			errors.InternalServerError,
		)
	}

	return web.Success(
		w,
		http.StatusOK,
		"Email verified.",
		IncompleteOnboardingResponse{
			UserId:           userId,
			AccessToken:      onboardingToken,
			OnboardingStatus: ONBOARDING_CREATE_ORGANIZATION,
		},
	)
}
