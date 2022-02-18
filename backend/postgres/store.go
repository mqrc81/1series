// Package postgres is the PostgreSQL database access layer
package postgres

import (
	"fmt"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dataSourceName string) (*Store, *postgresstore.PostgresStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening database: %w", err)
	}

	sessionsStore := postgresstore.New(db.DB)

	return &Store{
		&UserStore{db},
		&GenreStore{db},
		&NetworkStore{db},
	}, sessionsStore, nil
}

type Store struct {
	*UserStore
	*GenreStore
	*NetworkStore
}
