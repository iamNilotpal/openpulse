package main

import (
	"os"
	"runtime"

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

	return nil
}
