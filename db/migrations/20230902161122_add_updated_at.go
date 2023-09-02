package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddUpdatedAt, downAddUpdatedAt)
}

func upAddUpdatedAt(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downAddUpdatedAt(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
