package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/cors"
	v1 "github.com/iamNilotpal/openpulse/apps/api/handlers/v1"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	DB       *sqlx.DB
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Config   *config.OpenpulseApiConfig
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

	v1.SetupRoutes(app, v1.V1Config{Log: cfg.Log, Config: cfg.Config})

	return app.Mux
}
