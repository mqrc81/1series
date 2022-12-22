package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type genreRepository struct {
	*sqlx.DB
}

func (r *genreRepository) FindAll() (genres []domain.Genre, err error) {

	if err = r.Select(
		&genres,
		`SELECT g.* FROM genres g`,
	); err != nil {
		err = fmt.Errorf("error finding genres: %w", err)
	}

	return genres, err
}

func (r *genreRepository) Save(genre domain.Genre) (err error) {

	if _, err = r.Exec(
		`INSERT INTO genres(id, name) VALUES ($1, $2)`,
		genre.Id,
		genre.Name,
	); err != nil {
		err = fmt.Errorf("error adding genre [%v, %v]: %w", genre.Id, genre.Name, err)
	}

	return err
}
