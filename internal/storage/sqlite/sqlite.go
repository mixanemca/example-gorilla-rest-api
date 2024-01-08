// Package sqlite for sqlite database
package sqlite

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(15) NOT NULL
);

`

// NewConnection for create connection to database
func NewConnection(logger *slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:") // memory placed
	if err != nil {
		logger.Error("filed to create sqlite connection", err)
		return nil, err
	}
	if err := createUsersTable(db); err != nil {
		logger.Error("failed to create users table", err)
		return nil, err
	}
	logger.Info("connected to SQLite successfully")

	return db, nil
}

// createUsersTable for create table in database
func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(createTableQuery)
	return err
}