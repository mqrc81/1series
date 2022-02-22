// Package postgres is the PostgreSQL database access layer
package postgres

import (
	_ "database/sql"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dataSourceName string) (*Store, sessions.Store, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %w", err)
	}

	sessionsStore, err := postgres.NewStore(db.DB)

	return &Store{
		&UserStore{db},
		&GenreStore{db},
		&NetworkStore{db},
		&ReleaseStore{db},
	}, sessionsStore, nil
}

type Store struct {
	*UserStore
	*GenreStore
	*NetworkStore
	*ReleaseStore
}
