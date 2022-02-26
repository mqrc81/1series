// Package postgres is the PostgreSQL database access layer
package postgres

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, *sql.DB, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		&UserStore{db},
		&GenreStore{db},
		&NetworkStore{db},
		&ReleaseStore{db},
	}, db.DB, nil
}

type Store struct {
	*UserStore
	*GenreStore
	*NetworkStore
	*ReleaseStore
}
