package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateTokensTable, downCreateTokensTable)
}

func upCreateTokensTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE tokens (
			token_id TEXT PRIMARY KEY CHECK ( token_id != '' ),
			user_id INT REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
			purpose INT NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT now()
		);
		`)
	return err
}

func downCreateTokensTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE tokens;
		`)
	return err
}
