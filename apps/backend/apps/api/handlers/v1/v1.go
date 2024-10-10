package v1

import (
	"github.com/go-chi/chi/v5"
	auth_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/auth"
	invitations_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/invitations"
	onboarding_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/onboarding"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/middlewares"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const apiV1 = "/api/v1"

type Config struct {
	App                         *web.App
	Auth                        *auth.Auth
	Log                         *zap.SugaredLogger
	Repositories                *repositories.Repositories
	APIConfig                   *config.OpenpulseAPIConfig
	RolesMap                    auth.RoleConfigMap
	ResourcePermissionsMap      auth.ResourcePermissionsMap
	RoleResourcesPermissionsMap auth.RoleResourcesPermissionsMap
}

func SetupRoutes(cfg Config) {
	errorMiddleware := middlewares.ErrorResponder(cfg.Log)

	authHandler := auth_handlers.New(auth_handlers.Config{Auth: cfg.Auth})
	onboardingHandler := onboarding_handlers.New(onboarding_handlers.Config{})
	invitationHandler := invitations_handlers.New(invitations_handlers.Config{})

	cfg.App.Route(apiV1, func(r chi.Router) {
		r.Post("/auth/register", errorMiddleware(authHandler.Register))
		r.Post("/auth/login", errorMiddleware(authHandler.Login))

	})

	cfg.App.Route(apiV1, func(r chi.Router) {
		r.Post("/onboard/organization", errorMiddleware(onboardingHandler.SaveOrganizationDetails))
		r.Post("/onboard/team", errorMiddleware(onboardingHandler.SaveTeamDetails))
		r.Post("/onboard/invite", errorMiddleware(invitationHandler.InviteMembers))
	})
}
