// Package postgres is the PostgreSQL database access layer
package postgres

import (
	_ "database/sql"
	"fmt"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dataSourceName string) (*Store, sessions.Store, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %w", err)
	}

	sessionStore, err := pgstore.NewPGStoreFromPool(db.DB)
	if err != nil {
		return nil, nil, fmt.Errorf("error initializing session store: %w", err)
	}

	return &Store{
		&UserStore{db},
		&GenreStore{db},
		&NetworkStore{db},
		&ReleaseStore{db},
	}, sessionStore, nil
}

type Store struct {
	*UserStore
	*GenreStore
	*NetworkStore
	*ReleaseStore
}
