package shows

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

const (
	tmdbImageUrl       = "https://image.tmdb.org/t/p/original"
	releasesPerRequest = 20
)

type useCase struct {
	userRepository    repositories.UserRepository
	releaseRepository repositories.ReleaseRepository
	traktClient       *trakt.Client
	tmdbClient        *tmdb.Client
}

type UseCase interface {
	GetShow(showId int) (domain.Show, error)
	GetPopularShows(page int) ([]domain.Show, error)
	GetUpcomingReleases(page int) ([]domain.Release, error)
	SearchShows(searchTerm string) ([]domain.Show, error)
}

func NewUseCase(
	userRepository repositories.UserRepository, releaseRepository repositories.ReleaseRepository,
	traktClient *trakt.Client, tmdbClient *tmdb.Client,
) UseCase {
	return &useCase{userRepository, releaseRepository, traktClient, tmdbClient}
}
