package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	v1 "github.com/iamNilotpal/openpulse/apps/api/handlers/v1"
	"github.com/iamNilotpal/openpulse/business/pkg/email"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/foundation/hash"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	DB                          *sqlx.DB
	Auth                        *auth.Auth
	HashService                 hash.Hasher
	EmailService                *email.Email
	Cache                       *redis.Client
	Shutdown                    chan os.Signal
	APIConfig                   *config.APIConfig
	Log                         *zap.SugaredLogger
	RoleMap                     auth.RoleMappings
	ResourceMap                 auth.ResourceMappings
	PermissionMap               auth.PermissionMappings
	Repositories                *repositories.Repositories
	ResourcePermissionsMap      auth.ResourceToPermissionsMap
	RoleResourcesPermissionsMap auth.RoleNameToAccessControlMap
}

func NewHandler(cfg HandlerConfig) http.Handler {
	app := web.NewApp(web.AppConfig{
		Shutdown: cfg.Shutdown,
		Cors: cors.Options{
			MaxAge:           300,
			AllowCredentials: true,
			AllowedOrigins:   cfg.APIConfig.Web.AllowedOrigins,
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

	v1.SetupRoutes(
		v1.Config{
			App:                   app,
			Log:                   cfg.Log,
			Auth:                  cfg.Auth,
			RoleMap:               cfg.RoleMap,
			APIConfig:             cfg.APIConfig,
			ResourceMap:           cfg.ResourceMap,
			HashService:           cfg.HashService,
			Repositories:          cfg.Repositories,
			EmailService:          cfg.EmailService,
			PermissionMap:         cfg.PermissionMap,
			ResourcePermissionMap: cfg.ResourcePermissionsMap,
			AccessControlMap:      cfg.RoleResourcesPermissionsMap,
		},
	)

	return app
}
