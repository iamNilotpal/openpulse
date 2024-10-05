package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	v1 "github.com/iamNilotpal/openpulse/apps/api/handlers/v1"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	DB                     *sqlx.DB
	Auth                   *auth.Auth
	Cache                  *redis.Client
	Shutdown               chan os.Signal
	Log                    *zap.SugaredLogger
	RolesMap               auth.RoleConfigMap
	Repositories           repositories.Repositories
	Config                 *config.OpenpulseApiConfig
	ResourcePermissionsMap auth.ResourcePermissionsMap
	RolePermissionsMap     auth.RoleResourcesPermissionsMap
}

func NewHandler(cfg HandlerConfig) http.Handler {
	app := web.NewApp(web.AppConfig{
		Shutdown: cfg.Shutdown,
		Cors: cors.Options{
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
		},
	},
		middleware.Logger,
		middleware.RealIP,
		middleware.RequestID,
		middleware.Recoverer,
	)

	apiV1 := v1.New(
		app,
		cfg.Auth,
		cfg.Log,
		cfg.Config,
		cfg.Repositories,
		cfg.RolesMap,
		cfg.ResourcePermissionsMap,
		cfg.RolePermissionsMap,
	)

	apiV1.SetupRoutes()
	return app
}
