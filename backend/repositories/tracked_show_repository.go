package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type trackedShowRepository struct {
	*sql.Database
}

func (r *trackedShowRepository) FindAll() (trackedShows []domain.TrackedShow, err error) {

	if err = r.Select(
		&trackedShows,
		`SELECT ws.* FROM tracked_shows ws`,
	); err != nil {
		err = fmt.Errorf("error finding tracked shows: %w", err)
	}

	return trackedShows, err
}

func (r *trackedShowRepository) FindAllByUser(user domain.User) (trackedShows []domain.TrackedShow, err error) {

	if err = r.Select(
		&trackedShows,
		`SELECT ws.* FROM tracked_shows ws WHERE ws.user_id = $1`,
		user.Id,
	); err != nil {
		err = fmt.Errorf("error finding tracked shows [%v]: %w", user.Id, err)
	}

	return trackedShows, err
}

func (r *trackedShowRepository) Save(trackedShow domain.TrackedShow) (err error) {

	if _, err = r.Exec(
		`INSERT INTO tracked_shows(user_id, show_id, rating) VALUES($1, $2, $3)`,
		trackedShow.UserId,
		trackedShow.ShowId,
		trackedShow.Rating,
	); err != nil {
		err = fmt.Errorf("error saving tracked show [%v]: %w", trackedShow, err)
	}

	return err
}
