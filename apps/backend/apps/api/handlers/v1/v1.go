package v1

import (
	users_handler "github.com/iamNilotpal/openpulse/apps/api/handlers/v1/users"
	"github.com/iamNilotpal/openpulse/business/sys/config"
	"github.com/iamNilotpal/openpulse/foundation/web"
	"go.uber.org/zap"
)

const version = "/api/v1"
const usersRoute = version + "/users"

type V1Config struct {
	Log    *zap.SugaredLogger
	Config *config.OpenpulseApiConfig
}

func SetupRoutes(app *web.App, cfg V1Config) {
	usersHandler := users_handler.NewHandler()

	app.Mux.Get(usersRoute+"/{id}", usersHandler.QueryUserById)
}
