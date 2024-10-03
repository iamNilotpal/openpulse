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
	app            *web.App
	auth           *auth.Auth
	rolesMap       auth.AuthedRolesMap
	log            *zap.SugaredLogger
	permissionsMap auth.AuthedPermissionsMap
	repositories   repositories.Repositories
	config         *config.OpenpulseApiConfig
}

func New(
	app *web.App,
	auth *auth.Auth,
	log *zap.SugaredLogger,
	rolesMap auth.AuthedRolesMap,
	config *config.OpenpulseApiConfig,
	permissionsMap auth.AuthedPermissionsMap,
	repositories repositories.Repositories,

) *cfg {
	return &cfg{
		app:            app,
		log:            log,
		auth:           auth,
		config:         config,
		rolesMap:       rolesMap,
		repositories:   repositories,
		permissionsMap: permissionsMap,
	}
}

func (c *cfg) SetupRoutes() {
	errorMiddleware := middlewares.ErrorResponder(c.log)
	authHandler := auth_handlers.New(
		c.config.Auth,
		c.auth,
		c.rolesMap,
		c.repositories.Users,
		c.permissionsMap,
	)

	/* Auth Routes - Register, Login */
	c.app.Route(apiV1, func(r chi.Router) {
		r.Post("/auth/register", errorMiddleware(authHandler.Register))
		r.Post("/auth/login", errorMiddleware(authHandler.Login))
	})
}
