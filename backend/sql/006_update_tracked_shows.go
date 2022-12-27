package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUpdateTrackedShowsTable, downUpdateTrackedShowsTable)
}

func upUpdateTrackedShowsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE ONLY watched_shows ALTER COLUMN rating SET DEFAULT 0;
		ALTER TABLE ONLY watched_shows DROP CONSTRAINT watched_shows_rating_check;
		ALTER TABLE ONLY watched_shows ADD CONSTRAINT tracked_shows_rating_check CHECK (rating BETWEEN 0 AND 10);
		UPDATE watched_shows SET rating = 0 WHERE rating IS NULL;
		ALTER TABLE ONLY watched_shows ALTER COLUMN rating SET NOT NULL;
		ALTER TABLE watched_shows RENAME TO tracked_shows;
		`)
	return err
}

func downUpdateTrackedShowsTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE ONLY tracked_shows ALTER COLUMN rating SET DEFAULT NULL;
		ALTER TABLE ONLY tracked_shows DROP CONSTRAINT tracked_shows_rating_check;
		ALTER TABLE ONLY tracked_shows ADD CONSTRAINT watched_shows_rating_check CHECK (rating IS NULL OR (rating >= 0 AND rating <= 100));
		ALTER TABLE tracked_shows RENAME TO watched_shows;
		`)
	return err
}
