// Package web contains a small web framework extension.
package web

import (
	"net/http"
	"os"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Middleware func(http.Handler) http.Handler

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers.
type App struct {
	*chi.Mux
	shutdown chan os.Signal
}

type AppConfig struct {
	Cors     cors.Options
	Shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(cfg AppConfig, middlewares ...Middleware) *App {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cfg.Cors))
	for _, m := range middlewares {
		mux.Use(m)
	}

	return &App{
		Mux:      mux,
		shutdown: cfg.Shutdown,
	}
}

// SignalShutdown is used to gracefully shutdown the app when an integrity issue is
// identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
