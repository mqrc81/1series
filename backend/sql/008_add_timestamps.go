package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddTimestamps, downAddTimestamps)
}

func upAddTimestamps(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
		ALTER TABLE releases ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
		ALTER TABLE tracked_shows ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
		ALTER TABLE genres ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
		ALTER TABLE networks ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
		`)
	return err
}

func downAddTimestamps(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users DROP COLUMN created_at;
		ALTER TABLE releases DROP COLUMN created_at;
		ALTER TABLE tracked_shows DROP COLUMN created_at;
		ALTER TABLE genres DROP COLUMN created_at;
		ALTER TABLE networks DROP COLUMN created_at;
		`)
	return err
}
