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
	Repositories          *repositories.Repositories
	RoleMap               auth.RoleMappings
	ResourceMap           auth.ResourceMappings
	PermissionMap         auth.PermissionMappings
	ResourcePermissionMap auth.ResourceToPermissionsMap
	AccessControlMap      auth.RoleNameToAccessControlMap
}

func SetupRoutes(cfg Config) {
	/* =========== BUILD RespondErrorFunc =========== */
	RespondErrorFunc := middlewares.ErrorResponder(cfg.Log)

	/* =========== BUILD AUTH HANDLERS =========== */
	authHandler := auth_handlers.New(
		auth_handlers.Config{
			Auth:             cfg.Auth,
			RoleMap:          cfg.RoleMap,
			Config:           cfg.APIConfig,
			HashService:      cfg.HashService,
			EmailService:     cfg.EmailService,
			AccessControlMap: cfg.AccessControlMap,
			Users:            cfg.Repositories.Users,
			Emails:           cfg.Repositories.Emails,
			Sessions:         cfg.Repositories.Sessions,
			Organizations:    cfg.Repositories.Organizations,
		},
	)

	/* =========== BUILD ONBOARDING HANDLERS =========== */
	onboardingHandlers := onboarding_handlers.New(
		onboarding_handlers.Config{
			RoleMap:          cfg.RoleMap,
			Config:           cfg.APIConfig,
			AccessControlMap: cfg.AccessControlMap,
			Users:            cfg.Repositories.Users,
			Organizations:    cfg.Repositories.Organizations,
		},
	)

	/* =========== BUILD ROLES HANDLERS =========== */
	rolesHandler := roles_handler.New(
		roles_handler.Config{
			Roles: cfg.Repositories.Roles,
		},
	)

	/* =========== BUILD PERMISSIONS HANDLERS =========== */
	permissionsHandler := permissions_handlers.New(
		permissions_handlers.Config{
			Permissions: cfg.Repositories.Permissions,
		},
	)

	/* =========== BUILD RESOURCE HANDLERS =========== */
	resourcesHandler := resources_handlers.New(
		resources_handlers.Config{
			Resources: cfg.Repositories.Resources,
		},
	)

	/* =========== DEFINE ROUTES =========== */
	cfg.App.Route(apiV1, func(r chi.Router) {
		/* =========== Roles Routes =========== */
		r.Route("/roles", func(r chi.Router) {
			r.Post("/", RespondErrorFunc(rolesHandler.Create))
		})

		/* =========== Resources Routes =========== */
		r.Route("/resources", func(r chi.Router) {
			r.Post("/", RespondErrorFunc(resourcesHandler.Create))
		})

		/* =========== Permissions Routes =========== */
		r.Route("/permissions", func(r chi.Router) {
			r.Post("/", RespondErrorFunc(permissionsHandler.Create))
		})

		/* =========== Auth Routes =========== */
		r.Post("/oauth/{provider}", RespondErrorFunc(authHandler.OauthSignup))
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", RespondErrorFunc(authHandler.SignUp))
			r.Post("/signin", RespondErrorFunc(authHandler.SignIn))
			r.Post("/verify-email", RespondErrorFunc(authHandler.VerifyEmail))
		})

		/* =========== Onboarding Routes =========== */
		r.Route("/onboard", func(r chi.Router) {
			/* =========== ONBOARDING MIDDLEWARES =========== */
			r.Use(middlewares.AuthenticateOnboard(cfg.Auth, cfg.Repositories.Users))
			r.Use(middlewares.VerifiedUser)

			r.Post("/organization/create", RespondErrorFunc(onboardingHandlers.CreateOrganization))
			r.Post("/team/create", RespondErrorFunc(onboardingHandlers.CreateTeam))
		})
	})
}
