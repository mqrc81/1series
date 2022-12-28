package repositories

import (
	sqlx "database/sql"
	"time"

	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type UsersRepository interface {
	Find(userId int) (domain.User, error)
	FindAll() ([]domain.User, error)
	FindByUsername(username string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Save(user domain.User) (int, error)
}

type GenresRepository interface {
	FindAll() ([]domain.Genre, error)
	Save(genre domain.Genre) error
	ReplaceAll(genres []domain.Genre) error
}

type NetworksRepository interface {
	FindAll() ([]domain.Network, error)
	Save(network domain.Network) error
}

type ReleasesRepository interface {
	FindAllInRange(amount int, offset int) ([]domain.ReleaseRef, error)
	FindAllAiringBetween(startDate time.Time, endDate time.Time) ([]domain.ReleaseRef, error)
	ReplaceAll(releases []domain.ReleaseRef, pastReleasesCount int) error
	CountPastReleases() (int, error)
}

type TrackedShowsRepository interface {
	FindAll() ([]domain.TrackedShow, error)
	FindAllByUser(user domain.User) ([]domain.TrackedShow, error)
	Save(trackedShow domain.TrackedShow) error
}

func NewUsersRepository(database *sql.Database) UsersRepository {
	return &usersRepository{database}
}

func NewReleasesRepository(database *sql.Database) ReleasesRepository {
	return &releasesRepository{database}
}

func NewGenresRepository(database *sql.Database) GenresRepository {
	return &genresRepository{database}
}

func NewNetworksRepository(database *sql.Database) NetworksRepository {
	return &networksRepository{database}
}

func NewTrackedShowsRepository(database *sql.Database) TrackedShowsRepository {
	return &trackedShowsRepository{database}
}

func newId(res sqlx.Result) int {
	id, _ := res.LastInsertId()
	return int(id)
}
