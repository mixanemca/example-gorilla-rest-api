// Package pg for postgress database
package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mixanemca/example-gorilla-rest-api/internal/config"
	_ "github.com/mixanemca/example-gorilla-rest-api/internal/migrations"
	"github.com/pressly/goose/v3"
)

// NewConnection for create connection to database
func NewConnection(cfg config.Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Database.User),
		url.QueryEscape(cfg.Database.Password),
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Timeout)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Error("filed to parse config for database", err)
		return nil, err
	}

	// we set the maximum number of connections that can be in waiting.
	poolConfig.MaxConns = cfg.Database.MaxConns

	conn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		logger.Error("filed to create database pool", err)
		return nil, err
	}

	// check connection
	_, err = conn.Exec(context.Background(), ";")
	if err != nil {
		logger.Error("filed to set database connection", err)
		return nil, err
	}

	// make sql conn and init migrations
	mdb, _ := sql.Open("postgres", poolConfig.ConnString())
	err = mdb.Ping()
	if err != nil {
		logger.Error("filed to set database connection for migrations", err)
		return nil, err
	}
	err = goose.Up(mdb, "/var")
	if err != nil {
		logger.Error("filed to init migrations", err)
		return nil, err
	}
	logger.Info("connected to Postgres DB succsessfuly")

	return conn, nil
}
