package v1

import (
	"github.com/go-chi/chi/v5"
	userHandler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/user"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/business/web/middlewares"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const version = "/api/v1"

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
	usersHandler := userHandler.New(c.repositories.User)
	errorResponder := middlewares.ErrorResponder(c.log)

	c.app.Route(version, func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", errorResponder(usersHandler.QueryById))
		})
	})
}
