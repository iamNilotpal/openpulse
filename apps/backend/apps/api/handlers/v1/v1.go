package v1

import (
	"github.com/go-chi/chi/v5"
	auth_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/auth"
	onboarding_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/onboarding"
	"github.com/iamNilotpal/openpulse/business/pkg/email"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/middlewares"
	"github.com/iamNilotpal/openpulse/foundation/hash"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const apiV1 = "/api/v1"

type Config struct {
	App          *web.App
	Auth         *auth.Auth
	EmailService *email.Email
	HashService  hash.Hasher
	Log          *zap.SugaredLogger
	APIConfig    *config.APIConfig
	RolesMap     auth.RoleConfigMap
	ResPermsMap  auth.ResourcePermsMap
	RBACMaps     auth.RBACMap
	Repositories *repositories.Repositories
}

func SetupRoutes(cfg Config) {
	errorMiddleware := middlewares.ErrorResponder(cfg.Log)
	authHandler := auth_handlers.New(
		auth_handlers.Config{
			Auth:          cfg.Auth,
			RoleMap:       cfg.RolesMap,
			RBACMap:       cfg.RBACMaps,
			Config:        cfg.APIConfig,
			HashService:   cfg.HashService,
			EmailService:  cfg.EmailService,
			Users:         cfg.Repositories.Users,
			Emails:        cfg.Repositories.Emails,
			Sessions:      cfg.Repositories.Sessions,
			Organizations: cfg.Repositories.Organizations,
		},
	)
	onboardingHandlers := onboarding_handlers.New(
		onboarding_handlers.Config{
			Config:        cfg.APIConfig,
			Users:         cfg.Repositories.Users,
			Organizations: cfg.Repositories.Organizations,
			RoleMap:       cfg.RolesMap,
			RBACMap:       cfg.RBACMaps,
		},
	)

	cfg.App.Route(apiV1, func(r chi.Router) {
		/* ====================== Auth Routes ======================  */
		r.Post("/oauth/{provider}", errorMiddleware(authHandler.OauthSignup))
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", errorMiddleware(authHandler.SignUp))
			r.Post("/signin", errorMiddleware(authHandler.SignIn))
			r.Post("/verify-email", errorMiddleware(authHandler.VerifyEmail))
		})

		/* ====================== Onboarding Routes ======================  */
		r.Route("/onboard", func(r chi.Router) {
			r.Use(middlewares.AuthenticateOnboard(cfg.Auth, cfg.Repositories.Users))
			r.Use(middlewares.VerifiedUser)
			r.Post("/organization/create", errorMiddleware(onboardingHandlers.CreateOrganization))
			r.Post("/team/create", errorMiddleware(onboardingHandlers.CreateTeam))
		})
	})
}
