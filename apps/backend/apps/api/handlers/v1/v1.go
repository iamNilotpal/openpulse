package v1

import (
	"github.com/go-chi/chi/v5"
	auth_handlers "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/auth"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/middlewares"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const apiV1 = "/api/v1"

type cfg struct {
	app                         *web.App
	auth                        *auth.Auth
	rolesMap                    auth.RoleConfigMap
	log                         *zap.SugaredLogger
	repositories                *repositories.Repositories
	config                      *config.OpenpulseApiConfig
	resourcePermissionsMap      auth.ResourcePermissionsMap
	roleResourcesPermissionsMap auth.RoleResourcesPermissionsMap
}

func New(
	app *web.App,
	auth *auth.Auth,
	log *zap.SugaredLogger,
	config *config.OpenpulseApiConfig,
	repositories *repositories.Repositories,
	rolesMap auth.RoleConfigMap,
	resourcePermissionsMap auth.ResourcePermissionsMap,
	rolePermissionsMap auth.RoleResourcesPermissionsMap,
) *cfg {
	return &cfg{
		app:                         app,
		log:                         log,
		auth:                        auth,
		config:                      config,
		rolesMap:                    rolesMap,
		repositories:                repositories,
		roleResourcesPermissionsMap: rolePermissionsMap,
		resourcePermissionsMap:      resourcePermissionsMap,
	}
}

func (c *cfg) SetupRoutes() {
	errorMiddleware := middlewares.ErrorResponder(c.log)
	authHandler := auth_handlers.New(
		c.config.Auth,
		c.auth,
		c.rolesMap,
		c.repositories.Users,
		c.roleResourcesPermissionsMap,
	)

	/* Auth Routes - Register, Login */
	c.app.Route(apiV1, func(r chi.Router) {
		r.Post("/auth/register", errorMiddleware(authHandler.Register))
		r.Post("/auth/login", errorMiddleware(authHandler.Login))
	})
}
