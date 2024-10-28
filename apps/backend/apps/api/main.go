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
	"github.com/iamNilotpal/openpulse/business/pkg/email"
	"github.com/iamNilotpal/openpulse/business/repositories"
	"github.com/iamNilotpal/openpulse/business/repositories/emails"
	emails_store "github.com/iamNilotpal/openpulse/business/repositories/emails/store/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/organizations"
	organizations_store "github.com/iamNilotpal/openpulse/business/repositories/organizations/store/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/permissions"
	permissions_store "github.com/iamNilotpal/openpulse/business/repositories/permissions/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/resources"
	resources_store "github.com/iamNilotpal/openpulse/business/repositories/resources/store/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/roles"
	roles_store "github.com/iamNilotpal/openpulse/business/repositories/roles/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/sessions"
	sessions_store "github.com/iamNilotpal/openpulse/business/repositories/sessions/store/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/teams"
	teams_store "github.com/iamNilotpal/openpulse/business/repositories/teams/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/repositories/users"
	users_store "github.com/iamNilotpal/openpulse/business/repositories/users/stores/postgres"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/business/sys/database"
	"github.com/iamNilotpal/openpulse/business/web/auth"
	"github.com/iamNilotpal/openpulse/foundation/hash"
	"github.com/iamNilotpal/openpulse/foundation/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	/* ========= Load Env =========  */
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
	/* ========= GOMAXPROCS =========  */
	log.Infow("Startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	/* ========= INITIALIZE CONFIG =========  */
	cfg := config.NewAPIConfig()
	if err := config.Validate(*cfg); err != nil {
		return err
	}

	log.Infow("Config", "config", cfg)

	/* ========= INITIALIZE DATABASE SUPPORT =========  */
	db, err := database.Open(cfg.DB)
	if err != nil {
		log.Infow("DATABASE CONNECTION ERROR", "error", err)
		return err
	}

	/* ========= DATABASE STATUS CHECK =========  */
	err = database.StatusCheck(context.Background(), db)
	if err != nil {
		log.Infow("DATABASE Status Check Error", "error", err)
		return err
	}

	// Initialize Cache
	// redis, err := cache.Open(cfg.Cache)
	// if err != nil {
	// 	log.Infow("CACHE DATABASE CONNECTION ERROR", "error", err)
	// 	return err
	// }

	// if err = cache.StatusCheck(context.Background(), redis); err != nil {
	// 	log.Infow("CACHE Status Check Error", "error", err)
	// 	return err
	// }

	/* ========= INITIALIZE REPOSITORIES =========  */
	usersStore := users_store.NewPostgresStore(db)
	usersRepo := users.NewPostgresRepository(usersStore)

	emailsStore := emails_store.NewPostgresStore(db)
	emailsRepo := emails.NewPostgresRepository(emailsStore)

	teamsStore := teams_store.NewPostgresStore(db)
	teamsRepo := teams.NewRepository(teamsStore)

	orgsStore := organizations_store.NewPostgresStore(db)
	orgsRepo := organizations.NewPostgresRepository(orgsStore)

	rolesStore := roles_store.NewPostgresStore(db)
	rolesRepo := roles.NewPostgresRepository(rolesStore)

	resourceStore := resources_store.NewPostgresStore(db)
	resourceRepo := resources.NewPostgresRepository(resourceStore)

	permissionsStore := permissions_store.NewPostgresStore(db)
	permissionsRepo := permissions.NewPostgresRepository(permissionsStore)

	sessionsStore := sessions_store.NewPostgresRepository(db)
	sessionsRepo := sessions.NewPostgresRepository(sessionsStore)

	/* ========= BUILD REPOSITORIES STRUCT =========  */
	repositories := repositories.Repositories{
		Organizations: orgsRepo,
		Teams:         teamsRepo,
		Users:         usersRepo,
		Emails:        emailsRepo,
		Roles:         rolesRepo,
		Resources:     resourceRepo,
		Permissions:   permissionsRepo,
		Sessions:      sessionsRepo,
	}

	/* ========= GET ACCESS CONTROL =========  */
	accessControls, err := rolesRepo.QueryAccessControl(context.Background())
	if err != nil {
		return err
	}

	/* ========= GET ROLES =========  */
	appRoles, err := rolesRepo.QueryAll(context.Background())
	if err != nil {
		return err
	}

	/* ========= GET RESOURCES =========  */
	appResources, err := resourceRepo.QueryAll(context.Background())
	if err != nil {
		return err
	}

	/* ========= GET PERMISSIONS =========  */
	appPermissions, err := permissionsRepo.QueryAll(context.Background())
	if err != nil {
		return err
	}

	/* ========= BUILD RBAC AND RESOURCE PERMISSIONS MAP =========  */
	resourcePermsMap, accessControlMap := auth.BuildAccessControlMaps(accessControls)

	/* ========= BUILD ROLES, RESOURCES AND PERMISSIONS MAP =========  */
	roleMapping, resourceMapping, permissionMapping := auth.BuildAuthorizationMaps(
		appRoles, appResources, appPermissions,
	)

	/* ========= AUTHENTICATION SUPPORT =========  */
	auth := auth.New(auth.Config{Logger: log, AuthConfig: cfg.Auth, UserRepo: usersRepo})

	/* ========= INITIALIZE EMAIL SERVICE =========  */
	emailService := email.New(email.Config{Config: cfg.Email, Logger: log})

	/* ========= INITIALIZE HASH SERVICE =========  */
	bcryptHasher := hash.NewBcryptHasher()

	/* ========= SHUTDOWN SIGNAL =========  */
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	/* ========= INITIALIZE API SUPPORT =========  */
	mux := handlers.NewHandler(
		handlers.HandlerConfig{
			DB:                     db,
			Log:                    log,
			APIConfig:              cfg,
			Auth:                   auth,
			Shutdown:               shutdown,
			EmailService:           emailService,
			HashService:            bcryptHasher,
			Repositories:           &repositories,
			RoleMap:                roleMapping,
			ResourceMap:            resourceMapping,
			AccessControlMap:       accessControlMap,
			PermissionMap:          permissionMapping,
			ResourcePermissionsMap: resourcePermsMap,
		},
	)

	serverErrors := make(chan error, 1)
	api := http.Server{
		Handler:      mux,
		Addr:         cfg.Web.APIHost,
		ReadTimeout:  cfg.Web.ReadTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	/* ========= START THE API SERVER =========  */
	go func() {
		log.Infow("Server Listening", "address", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	/* ========= GRACEFUL SHUTDOWN =========  */
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
