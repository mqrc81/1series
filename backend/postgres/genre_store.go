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

func (s *GenreStore) AddGenre(genre domain.Genre) (err error) {

	if _, err = s.Exec(
		"INSERT INTO genres(id, name) VALUES ($1, $2)",
		genre.Id,
		genre.Name,
	); err != nil {
		err = fmt.Errorf("error adding genre [%v, %v]: %w", genre.Id, genre.Name, err)
	}

	return err
}
