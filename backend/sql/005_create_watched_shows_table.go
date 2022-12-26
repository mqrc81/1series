package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateWatchedShowsTable, downCreateWatchedShowsTable)
}

func upCreateWatchedShowsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE watched_shows
		(
			user_id INT REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
			show_id INT NOT NULL,
			rating  INT,
			PRIMARY KEY (user_id, show_id),
			CHECK ( rating IS NULL OR (rating >= 0 AND rating <= 100))
		);
		`)
	return err
}

func downCreateWatchedShowsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS watched_shows;
		`)
	return err
}
