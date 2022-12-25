package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type UserRepository interface {
	Find(userId int) (domain.User, error)
	FindByUsername(username string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Save(user domain.User) (int, error)
}

type GenreRepository interface {
	FindAll() ([]domain.Genre, error)
	Save(genre domain.Genre) error
	ReplaceAll(genres []domain.Genre) error
}

type NetworkRepository interface {
	FindAll() ([]domain.Network, error)
	Save(network domain.Network) error
}

type ReleaseRepository interface {
	FindAllInRange(amount int, offset int) ([]domain.ReleaseRef, error)
	ReplaceAll(releases []domain.ReleaseRef, pastReleasesCount int) error
	CountPastReleases() (int, error)
}

type WatchedShowRepository interface {
	FindAllByUser(user domain.User) ([]domain.WatchedShow, error)
}

func NewUserRepository(database *sqlx.DB) UserRepository {
	return &userRepository{database}
}

func NewReleaseRepository(database *sqlx.DB) ReleaseRepository {
	return &releaseRepository{database}
}

func NewGenreRepository(database *sqlx.DB) GenreRepository {
	return &genreRepository{database}
}

func NewNetworkRepository(database *sqlx.DB) NetworkRepository {
	return &networkRepository{database}
}

func NewWatchedShowRepository(database *sqlx.DB) WatchedShowRepository {
	return &watchedShowRepository{database}
}

func newId(res sql.Result) int {
	id, _ := res.LastInsertId()
	return int(id)
}
