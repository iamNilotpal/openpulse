package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/iamNilotpal/openpulse/apps/api/handlers"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/sys/cache"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/foundation/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Load envs.
	godotenv.Load()

	log := logger.New("Openpulse Backend")
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorf("Startup", "error", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// GOMAXPROCS
	log.Infow("Startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// Initialize Config
	cfg := config.NewOpenpulseConfig()
	log.Infow("Config", "config", cfg)

	// Initialize Database
	db, err := database.Open(cfg.DB)
	if err != nil {
		log.Infow("DATABASE CONNECTION ERROR", "error", err)
		return err
	}

	err = database.StatusCheck(context.Background(), db)
	if err != nil {
		log.Infow("DATABASE Status Check Error", "error", err)
		return err
	}

	// Initialize Cache
	redis, err := cache.Open(cfg.Cache)
	if err != nil {
		log.Infow("CACHE DATABASE CONNECTION ERROR", "error", err)
		return err
	}

	if err = cache.StatusCheck(context.Background(), redis); err != nil {
		log.Infow("CACHE Status Check Error", "error", err)
		return err
	}

	// Initialize repositories
	usersStore := users_store.NewPostgresStore(db)
	usersRepository := users.NewPostgresRepository(usersStore)

	rolesStore := roles_store.NewPostgresStore(db)
	rolesRepository := roles.NewPostgresRepository(rolesStore)

	permissionsStore := permissions_store.NewPostgresStore(db)
	permissionsRepository := permissions.NewPostgresRepository(permissionsStore)

	repositories := repositories.Repositories{
		Users:       usersRepository,
		Roles:       rolesRepository,
		Permissions: permissionsRepository,
	}

	// Get roles with permissions
	rolesWithPermissions, err := rolesRepository.QueryRolesWithPermissions(context.Background())
	if err != nil {
		return err
	}

	// Build the Permissions Map
	permissionsMap := buildPermissionsMap(rolesWithPermissions)

	// Initialize authentication support
	auth := auth.New(auth.Config{AuthConfig: cfg.Auth, UserRepo: usersRepository, Logger: log})

	// Shutdown Signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Initialize API support
	mux := handlers.NewHandler(
		handlers.HandlerConfig{
			DB:             db,
			Log:            log,
			Config:         cfg,
			Auth:           auth,
			Cache:          redis,
			Shutdown:       shutdown,
			Repositories:   repositories,
			PermissionsMap: permissionsMap,
		},
	)

	api := http.Server{
		Handler:      mux,
		Addr:         cfg.Web.APIHost,
		ReadTimeout:  cfg.Web.ReadTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Infow("Server Listening", "address", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Graceful Shutdown of Server
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("Shutting Down Server", "signal", sig)
		defer log.Infow("Shutdown Complete", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func buildPermissionsMap(permissions []roles.RolePermissions) auth.PermissionsMap {
	permissionsMap := make(auth.PermissionsMap)

	for _, permission := range permissions {
		stored, ok := permissionsMap[permission.Role.Name]
		if !ok {
			stored = []auth.Permissions{
				auth.ToPermissions(
					auth.ToRole(permission.Role), auth.ToPermission(permission.Permission),
				),
			}
		}

		stored = append(
			stored,
			auth.ToPermissions(
				auth.ToRole(permission.Role), auth.ToPermission(permission.Permission),
			),
		)

		permissionsMap[permission.Permission.Name] = stored
	}

	return permissionsMap
}
