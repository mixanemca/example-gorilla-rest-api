package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	api "github.com/mixanemca/example-gorilla-rest-api/internal/app/api"
	"github.com/mixanemca/example-gorilla-rest-api/internal/config"
	"github.com/mixanemca/example-gorilla-rest-api/internal/logger"
)

var (
	version string = "unknown"
	build   string = "unknown"
)

func main() {
	cfg, err := config.New(version, build)
	if err != nil {
		log.Fatalf("error occurred while reading config: %s\n", err.Error()) // TODO: change when external logger was added
	}
	if cfg == nil {
		slog.Error("config is empty")
		return
	}

	// set logger
	log, err := logger.New(cfg.Logger.Level, cfg.Logger.Format)
	if err != nil {
		slog.Error("create logger", slog.String("error", err.Error()))
		return
	}
	log.Info(
		"starting example-gorilla-rest-api",
		slog.String("version", cfg.Version),
		slog.String("build", cfg.Build),
		slog.String("log level", cfg.Logger.Level),
		slog.String("log format", cfg.Logger.Format),
	)
	log.Debug("debug messages are enabled")

	// start HTTP server
	apiApp := api.NewApp(*cfg, log)
	apiApp.Run()

	// gracefully shutdown HTTP server
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	apiApp.Shutdown(ctx)
}
