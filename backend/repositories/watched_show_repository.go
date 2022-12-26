package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type watchedShowRepository struct {
	*sql.Database
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
