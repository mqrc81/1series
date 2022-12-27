package repositories

import (
	sqlx "database/sql"
	"time"

	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type UserRepository interface {
	Find(userId int) (domain.User, error)
	FindAll() ([]domain.User, error)
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
	FindAllAiringBetween(startDate time.Time, endDate time.Time) ([]domain.ReleaseRef, error)
	ReplaceAll(releases []domain.ReleaseRef, pastReleasesCount int) error
	CountPastReleases() (int, error)
}

type WatchedShowRepository interface {
	FindAll() ([]domain.WatchedShow, error)
	FindAllByUser(user domain.User) ([]domain.WatchedShow, error)
}

func NewUserRepository(database *sql.Database) UserRepository {
	return &userRepository{database}
}

func NewReleaseRepository(database *sql.Database) ReleaseRepository {
	return &releaseRepository{database}
}

func NewGenreRepository(database *sql.Database) GenreRepository {
	return &genreRepository{database}
}

func NewNetworkRepository(database *sql.Database) NetworkRepository {
	return &networkRepository{database}
}

func NewWatchedShowRepository(database *sql.Database) WatchedShowRepository {
	return &watchedShowRepository{database}
}

func newId(res sqlx.Result) int {
	id, _ := res.LastInsertId()
	return int(id)
}
