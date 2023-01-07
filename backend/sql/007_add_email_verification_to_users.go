package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddEmailVerificationToUsers, downAddEmailVerificationToUsers)
}

func upAddEmailVerificationToUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users ADD COLUMN email_verified BOOL NOT NULL DEFAULT FALSE;
		ALTER TABLE users ALTER COLUMN password DROP NOT NULL;
		`)
	return err
}

func downAddEmailVerificationToUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE users DROP COLUMN email_verified;
		ALTER TABLE users ALTER COLUMN password SET NOT NULL;
		`)
	return err
}
