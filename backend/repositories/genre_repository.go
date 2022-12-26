package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type genreRepository struct {
	*sql.Database
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
		`INSERT INTO genres(tmdb_id, name) VALUES ($1, $2)`,
		genre.TmdbId,
		genre.Name,
	); err != nil {
		err = fmt.Errorf("error adding genre [%v, %v]: %w", genre.TmdbId, genre.Name, err)
	}

	return err
}

func (r *genreRepository) ReplaceAll(genres []domain.Genre) (err error) {

	txn, err := r.Beginx()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	defer func() {
		if err == nil {
			err = txn.Commit()
		} else {
			_ = txn.Rollback()
		}
	}()

	if err = r.deleteAllGenresInTransaction(txn); err != nil {
		return err
	}

	for _, genre := range genres {
		if err = r.saveGenresInTransaction(txn, genre); err != nil {
			return err
		}
	}

	return err
}

func (r *genreRepository) deleteAllGenresInTransaction(txn *sqlx.Tx) (err error) {
	//goland:noinspection SqlWithoutWhere
	if _, err = txn.Exec(`DELETE FROM genres`); err != nil {
		err = fmt.Errorf("error deleting genres: %w", err)
	}
	return err
}

func (r *genreRepository) saveGenresInTransaction(txn *sqlx.Tx, genre domain.Genre) (err error) {
	if _, err = txn.Exec(`INSERT INTO genres(tmdb_id, name) VALUES($1, $2)`,
		genre.TmdbId,
		genre.Name,
	); err != nil {
		err = fmt.Errorf("error saving genre: %w", err)
	}
	return err
}
