package v1

import (
	"github.com/go-chi/chi/v5"
	auth_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/auth"
	invitations_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/invitations"
	onboarding_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/onboarding"
	roles_handler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/roles"
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
	App                         *web.App
	Auth                        *auth.Auth
	EmailService                *email.Email
	HashService                 hash.Hasher
	Log                         *zap.SugaredLogger
	Repositories                *repositories.Repositories
	APIConfig                   *config.APIConfig
	RolesMap                    auth.RoleConfigMap
	ResourcePermissionsMap      auth.ResourcePermissionsMap
	RoleResourcesPermissionsMap auth.RoleAccessControlMap
}

func SetupRoutes(cfg Config) {
	errorMiddleware := middlewares.ErrorResponder(cfg.Log)
	authHandler := auth_handlers.New(
		auth_handlers.Config{
			Auth:                        cfg.Auth,
			RolesMap:                    cfg.RolesMap,
			Config:                      cfg.APIConfig,
			HashService:                 cfg.HashService,
			EmailService:                cfg.EmailService,
			Users:                       cfg.Repositories.Users,
			Emails:                      cfg.Repositories.Emails,
			RoleResourcesPermissionsMap: cfg.RoleResourcesPermissionsMap,
		},
	)
	onboardingHandler := onboarding_handlers.New(
		onboarding_handlers.Config{
			Config: cfg.APIConfig,
			Users:  cfg.Repositories.Users,
		},
	)
	invitationHandler := invitations_handlers.New(invitations_handlers.Config{})
	rolesHandler := roles_handler.New(roles_handler.Config{Roles: cfg.Repositories.Roles})

	cfg.App.Route(apiV1, func(r chi.Router) {
		r.Post("/roles", errorMiddleware(rolesHandler.Create))

		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", errorMiddleware(authHandler.SignUp))
			r.Post("/signin", errorMiddleware(authHandler.SignIn))
		})

		r.Route("/onboard", func(r chi.Router) {
			r.Use(middlewares.Authenticate(cfg.Auth, cfg.Repositories.Users))

			r.Post("/organization", errorMiddleware(onboardingHandler.CreateOrganization))
			r.Post("/team", errorMiddleware(onboardingHandler.CreateTeam))
			r.Post("/invite", errorMiddleware(invitationHandler.InviteMembers))
		})
	})
}
