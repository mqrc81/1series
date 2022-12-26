package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateGenresTable, downCreateGenresTable)
}

func upCreateGenresTable(tx *sql.Tx) error {
	_, err := tx.Exec(`		
		CREATE TABLE genres
		(
			id      SERIAL PRIMARY KEY,
			tmdb_id INT         NOT NULL,
			name    TEXT UNIQUE NOT NULL
		);
		`)
	return err
}

func downCreateGenresTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS genres;
		`)
	return err
}
