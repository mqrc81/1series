package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateNetworksTable, downCreateNetworksTable)
}

func upCreateNetworksTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE networks
		(
			id      SERIAL PRIMARY KEY,
			tmdb_id INT         NOT NULL,
			name    TEXT UNIQUE NOT NULL,
			logo    TEXT UNIQUE NOT NULL
		);
		`)
	return err
}

func downCreateNetworksTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS networks;
		`)
	return err
}
