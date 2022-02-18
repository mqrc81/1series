package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type GenreStore struct {
	*sqlx.DB
}

func (s *GenreStore) GetGenres() (genres []domain.Genre, err error) {

	if err = s.Select(
		&genres,
		"SELECT g.* FROM genres g",
	); err != nil {
		err = fmt.Errorf("error getting genres: %w", err)
	}

	return genres, err
}
