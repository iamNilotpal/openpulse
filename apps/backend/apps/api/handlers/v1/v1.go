package v1

import (
	"github.com/go-chi/chi/v5"
	userHandler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/user"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const version = "/api/v1"

type cfg struct {
	app          *web.App
	log          *zap.SugaredLogger
	repositories config.Repositories
	config       *config.OpenpulseApiConfig
}

func New(
	app *web.App,
	log *zap.SugaredLogger,
	config *config.OpenpulseApiConfig,
	repositories config.Repositories,
) *cfg {
	return &cfg{app: app, log: log, config: config, repositories: repositories}
}

func (c *cfg) SetupRoutes() {
	usersHandler := userHandler.New(c.repositories.User)

	c.app.Route(version, func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", usersHandler.QueryById)
		})
	})
}
