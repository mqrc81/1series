package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type Database struct {
	*sqlx.DB
}

func (db *Database) Migrate() error {
	return goose.Up(db.DB.DB, "./sql")
}
