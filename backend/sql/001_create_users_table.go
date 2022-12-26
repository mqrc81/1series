package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE users
		(
			id                     SERIAL PRIMARY KEY,
			username               TEXT UNIQUE NOT NULL,
			password               TEXT        NOT NULL,
			email                  TEXT UNIQUE NOT NULL,
			notify_releases        BOOL        NOT NULL DEFAULT FALSE,
			notify_recommendations BOOL        NOT NULL DEFAULT FALSE
		);
		`)
	return err
}

func downCreateUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP TABLE IF EXISTS users;
		`)
	return err
}
