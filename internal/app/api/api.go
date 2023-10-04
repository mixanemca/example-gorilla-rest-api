// Package api for handlers
package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"
	"github.com/mixanemca/example-gorilla-rest-api/internal/app/service"
	"github.com/mixanemca/example-gorilla-rest-api/internal/config"
)

type app struct {
	config           config.Config
	logger           *slog.Logger
	publicHTTPServer *http.Server
	service          *service.Service
}

func NewApp(cfg config.Config, logger *slog.Logger) *app {
	logger.Debug("Create new API app")

	userRepo := v1.NewUserRepository()
	service := service.NewService(userRepo)

	return &app{
		config: cfg,
		logger: logger,
		publicHTTPServer: &http.Server{
			Addr: cfg.HTTP.Address,
		},
		service: service,
	}
}

func (a *app) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/user", v1.NewUserRepository().CreateUser).Methods(http.MethodPost)
	http.Handle("/", router)

	// start HTTP server
	go func() {
		a.logger.Info("service run on " + a.config.HTTP.Address)
		if err := a.publicHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.publicHTTPServer.ErrorLog.Fatalf("error occurred while running http server: %s\n", err.Error()) // TODO: change when external logger was added
		}
	}()
	time.Sleep(time.Minute * 2)
}
