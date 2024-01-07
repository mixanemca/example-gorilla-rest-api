// Package app for handlers
package app

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"time"

	pg "github.com/mixanemca/example-gorilla-rest-api/internal/storage/postgres"

	"github.com/gorilla/mux"
	_ "github.com/mixanemca/example-gorilla-rest-api/docs"
	v1 "github.com/mixanemca/example-gorilla-rest-api/internal/app/api/handler/v1"
	"github.com/mixanemca/example-gorilla-rest-api/internal/app/api/middleware"
	"github.com/mixanemca/example-gorilla-rest-api/internal/app/service"
	"github.com/mixanemca/example-gorilla-rest-api/internal/config"
	"github.com/mixanemca/example-gorilla-rest-api/internal/storage/sqlite"

	httpSwagger "github.com/swaggo/http-swagger"
)

type app struct {
	config           config.Config
	logger           *slog.Logger
	publicHTTPServer *http.Server
	service          *service.Service
}

func New(cfg config.Config, logger *slog.Logger) (*app, error) {
	logger.Debug("Create new API app")

	var userRepo v1.UserRepository
	switch cfg.Database.DBType { // set database type
	case config.DBTypePostgres:
		db, err := pg.NewConnection(cfg, logger)
		if err != nil {
			return nil, err
		}
		userRepo, err = v1.NewUserRepositoryPg(db)
		if err != nil {
			return nil, err
		}
	case config.DBTypeSQLite:
		db, err := sqlite.NewConnection(logger)
		if err != nil {
			return nil, err
		}
		userRepo, err = v1.NewUserRepositorySqlite(db)
		if err != nil {
			return nil, err
		}
	default:
		log.Fatal("field to sen any database type")
	}

	service := service.NewService(userRepo)

	return &app{
		config: cfg,
		logger: logger,
		publicHTTPServer: &http.Server{
			Addr: cfg.HTTP.Address,
		},
		service: service,
	}, nil
}

func (a *app) Run() {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware(a.logger))
	router.Use(middleware.PanicRecover(a.logger))

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("doc.json")))
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("", a.service.CreateUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/list", a.service.GetUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9a-f\\-]+}", a.service.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9a-f\\-]+}", a.service.UpdateUser).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id:[0-9a-f\\-]+}", a.service.DeleteUser).Methods(http.MethodDelete)

	http.Handle("/", router)

	// start HTTP server
	go func() {
		a.logger.Info("service run on " + a.config.HTTP.Address)
		if err := a.publicHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.publicHTTPServer.ErrorLog.Fatalf("error occurred while running http server: %s\n", err.Error()) // TODO: change when external logger was added
		}
	}()
}

func (a *app) Shutdown(ctx context.Context) {
	<-ctx.Done()
	a.logger.Info("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(a.config.ShutdownTimeout))
	defer cancel()

	if err := a.publicHTTPServer.Shutdown(shutdownCtx); err != nil {
		a.logger.Error("Stopping service error: %v", err)
	}
	a.logger.Info("HTTP server successfully stopped")
}
