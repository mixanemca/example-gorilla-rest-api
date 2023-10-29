// Package pg for postgress database
package pg

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixanemca/example-gorilla-rest-api/internal/config"
)

// NewConnection for create connection to database
func NewConnection(cfg config.Config) *pgxpool.Pool {
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
		log.Fatalf("filed to parse config for database %s", err.Error())
	}

	// we set the maximum number of connections that can be in waiting.
	poolConfig.MaxConns = cfg.Database.MaxConns

	conn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("filed to create database pool %s", err.Error())
	}

	// check connection
	_, err = conn.Exec(context.Background(), ";")
	if err != nil {
		log.Fatalf("filed to set database connection %s", err.Error())
	}

	return conn
}
