// Package postgres is the PostgreSQL database access layer
package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dataSourceName string) (*Store, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return &Store{
		&UserStore{db},
	}, nil
}

type Store struct {
	*UserStore
}
