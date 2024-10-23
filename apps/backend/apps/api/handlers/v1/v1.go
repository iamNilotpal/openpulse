package v1

import (
	"github.com/go-chi/chi/v5"
	auth_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/auth"
	onboarding_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/onboarding"
	permissions_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/rbac/permissions"
	resources_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/rbac/resources"
	roles_handler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/rbac/roles"
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
	App                   *web.App
	Auth                  *auth.Auth
	HashService           hash.Hasher
	EmailService          *email.Email
	APIConfig             *config.APIConfig
	Log                   *zap.SugaredLogger
	RoleMap               auth.RoleMappings
	ResourceMap           auth.ResourceMappings
	PermissionMap         auth.PermissionMappings
	ResourcePermissionMap auth.ResourceToPermissionsMap
	AccessControlMap      auth.RoleNameToAccessControlMap
	Repositories          *repositories.Repositories
}

func SetupRoutes(cfg Config) {
	errorMiddleware := middlewares.ErrorResponder(cfg.Log)
	authHandler := auth_handlers.New(
		auth_handlers.Config{
			Auth:          cfg.Auth,
			RoleMap:       cfg.RoleMap,
			Config:        cfg.APIConfig,
			HashService:   cfg.HashService,
			EmailService:  cfg.EmailService,
			RBACMap:       cfg.AccessControlMap,
			Users:         cfg.Repositories.Users,
			Emails:        cfg.Repositories.Emails,
			Sessions:      cfg.Repositories.Sessions,
			Organizations: cfg.Repositories.Organizations,
		},
	)
	onboardingHandlers := onboarding_handlers.New(
		onboarding_handlers.Config{
			RoleMap:       cfg.RoleMap,
			Config:        cfg.APIConfig,
			Users:         cfg.Repositories.Users,
			RBACMap:       cfg.AccessControlMap,
			Organizations: cfg.Repositories.Organizations,
		},
	)
	rolesHandler := roles_handler.New(roles_handler.Config{Roles: cfg.Repositories.Roles})
	permissionsHandler := permissions_handlers.New(
		permissions_handlers.Config{Permissions: cfg.Repositories.Permissions},
	)
	resourcesHandler := resources_handlers.New(
		resources_handlers.Config{Resources: cfg.Repositories.Resources},
	)

	cfg.App.Route(apiV1, func(r chi.Router) {
		/* ====================== Roles Routes ======================  */
		r.Route("/roles", func(r chi.Router) {
			r.Post("/", errorMiddleware(rolesHandler.Create))
		})

		/* ====================== Resources Routes ======================  */
		r.Route("/resources", func(r chi.Router) {
			r.Post("/", errorMiddleware(resourcesHandler.Create))
		})

		/* ====================== Permissions Routes ======================  */
		r.Route("/permissions", func(r chi.Router) {
			r.Post("/", errorMiddleware(permissionsHandler.Create))
		})

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
