package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type watchedShowRepository struct {
	*sqlx.DB
}

func (r *watchedShowRepository) FindAllByUser(user domain.User) (watchedShows []domain.WatchedShow, err error) {

	if err = r.Get(
		&watchedShows,
		`SELECT ws.* FROM watched_shows ws WHERE ws.user_id = $1`,
		user.Id,
	); err != nil {
		err = fmt.Errorf("error finding watched shows [%v]: %w", user.Id, err)
	}

	return watchedShows, err
}
