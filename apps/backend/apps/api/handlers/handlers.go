package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/cors"
	v1 "github.com/iamNilotpal/openpulse/apps/api/handlers/v1"
	"github.com/iamNilotpal/openpulse/business/core/user"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	Cores    Cores
	DB       *sqlx.DB
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Config   *config.OpenpulseApiConfig
}

type Cores struct {
	UserCore *user.Core
}

func NewHandler(cfg HandlerConfig) http.Handler {
	app := web.NewApp(web.AppConfig{Shutdown: cfg.Shutdown}, cors.Options{
		MaxAge:           300,
		AllowCredentials: true,
		AllowedOrigins:   cfg.Config.Web.AllowedOrigins,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodHead,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
	})

	apiV1 := v1.New(app, cfg.Cores.UserCore, cfg.Log, cfg.Config)
	apiV1.SetupRoutes()

	return app
}
