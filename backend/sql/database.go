package sql

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

const (
	migrationsDirectory = "./sql"
)

type Database struct {
	*sqlx.DB
}

func (db *Database) Migrate() error {
	return goose.Up(db.DB.DB, migrationsDirectory)
}

func (db *Database) Rollback() error {
	return goose.Down(db.DB.DB, migrationsDirectory)
}
