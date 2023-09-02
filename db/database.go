package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mixanemca/example-gorilla-rest-api/config"
)

var Conn *pgx.Conn

func InitDatabase() error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.Name)
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return err
	}
	Conn = conn

	return nil
}
