package sql

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddTokensUniqueConstraint, downAddTokensUniqueConstraint)
}

//goland:noinspection SqlResolve
func upAddTokensUniqueConstraint(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE tokens RENAME COLUMN token_id TO id;
		ALTER TABLE tokens ADD CONSTRAINT tokens_user_id_purpose_key UNIQUE(user_id, purpose);
		`)
	return err
}

func downAddTokensUniqueConstraint(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE tokens RENAME COLUMN id TO token_id;
		ALTER TABLE tokens DROP CONSTRAINT tokens_user_id_purpose_key;
		`)
	return err
}
