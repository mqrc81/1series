package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateReleasesTable, downCreateReleasesTable)
}

func upCreateReleasesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE releases
		(
			show_id            INT       NOT NULL,
			season_number      INT       NOT NULL,
			air_date           TIMESTAMP NOT NULL,
			anticipation_level INT       NOT NULL DEFAULT 0,
			PRIMARY KEY (show_id, season_number)
		);
		
		-- singleton table
		CREATE TABLE past_releases
		(
			past_releases_id INT PRIMARY KEY DEFAULT 69,
			amount           INT NOT NULL    DEFAULT 0,
			CHECK ( past_releases_id = 69 )
		);
		
		INSERT INTO past_releases DEFAULT VALUES;
		
		CREATE OR REPLACE RULE ignore_past_releases_delete AS ON DELETE TO past_releases DO NOTHING;
		CREATE OR REPLACE RULE ignore_past_releases_insert AS ON INSERT TO past_releases DO NOTHING;
		`)
	return err
}

func downCreateReleasesTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS releases;
		
		DROP TABLE IF EXISTS past_releases;
		`)
	return err
}
