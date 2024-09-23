package v1

import (
	"github.com/go-chi/chi/v5"
	roles_handler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/roles"
	users_handler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/users"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/middlewares"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const apiV1 = "/api/v1"

type cfg struct {
	app          *web.App
	auth         *auth.Auth
	log          *zap.SugaredLogger
	repositories repositories.Repositories
	config       *config.OpenpulseApiConfig
}

func New(
	app *web.App,
	auth *auth.Auth,
	log *zap.SugaredLogger,
	config *config.OpenpulseApiConfig,
	repositories repositories.Repositories,
) *cfg {
	return &cfg{app: app, auth: auth, log: log, config: config, repositories: repositories}
}

func (c *cfg) SetupRoutes() {
	errorMiddleware := middlewares.ErrorResponder(c.log)

	usersHandler := users_handler.New(c.repositories.User)
	rolesHandler := roles_handler.New(c.repositories.Roles)

	c.app.Route(apiV1, func(r chi.Router) {
		r.Route("/roles", func(r chi.Router) {
			r.Post("/", errorMiddleware(rolesHandler.Create))
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", errorMiddleware(usersHandler.QueryById))
		})
	})
}
