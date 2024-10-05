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
	permissions, err := rolesRepository.GetRolesResourcesPermissions(context.Background())
	if err != nil {
		return err
	}

	// Build the Permissions Map
	rolesMap, resPermissionsMap, rolePermissionsMap := buildRBACMaps(permissions)

	// Initialize authentication support
	auth := auth.New(auth.Config{Logger: log, AuthConfig: cfg.Auth, UserRepo: usersRepository})

	// Shutdown Signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Initialize API support
	mux := handlers.NewHandler(
		handlers.HandlerConfig{
			DB:                     db,
			Log:                    log,
			Config:                 cfg,
			Auth:                   auth,
			Cache:                  redis,
			Shutdown:               shutdown,
			RolesMap:               rolesMap,
			Repositories:           repositories,
			ResourcePermissionsMap: resPermissionsMap,
			RolePermissionsMap:     rolePermissionsMap,
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

func buildRBACMaps(roleAccessControls []roles.RoleAccessControl) (
	auth.RoleConfigMap, auth.ResourcePermissionsMap, auth.RoleResourcesPermissionsMap,
) {
	rolesMap := make(auth.RoleConfigMap)
	resourcesPermissionsMap := make(auth.ResourcePermissionsMap)
	roleResourcesPermissionsMap := make(auth.RoleResourcesPermissionsMap)

	for _, rac := range roleAccessControls {
		// 1. Assign to roles map.
		if _, ok := rolesMap[rac.Role.Role]; !ok {
			rolesMap[rac.Role.Role] = auth.NewRoleConfig(rac.Role)
		}

		// 2. Check if any value exists for current role
		resPermsMap, ok := roleResourcesPermissionsMap[rac.Role.Role]

		// 3. If not then assign new value with resource and permissions
		if !ok {
			resPermsMap := make(auth.ResourcePermissionsMap)
			permissions := []auth.PermissionConfig{auth.NewPermissionConfig(rac.Permission)}

			resPermsMap[rac.Resource.Resource] = permissions
			roleResourcesPermissionsMap[rac.Role.Role] = resPermsMap
			resourcesPermissionsMap = resPermsMap
			continue
		}

		// 4. Check any value exists for current resource
		permissions, ok := resPermsMap[rac.Resource.Resource]

		// 5. If not then assign new value with permissions
		if !ok {
			permissions := []auth.PermissionConfig{auth.NewPermissionConfig(rac.Permission)}

			resPermsMap[rac.Resource.Resource] = permissions
			roleResourcesPermissionsMap[rac.Role.Role] = resPermsMap
			resourcesPermissionsMap = resPermsMap
			continue
		}

		// 6. If yes, append new permission to the existing resource.
		permissions = append(permissions, auth.NewPermissionConfig(rac.Permission))
		resPermsMap[rac.Resource.Resource] = permissions
		roleResourcesPermissionsMap[rac.Role.Role] = resPermsMap
		resourcesPermissionsMap = resPermsMap
	}

	return rolesMap, resourcesPermissionsMap, roleResourcesPermissionsMap
}
