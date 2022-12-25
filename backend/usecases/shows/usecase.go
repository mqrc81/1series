package shows

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

const (
	tmdbImageBaseUrl           = "https://image.tmdb.org/t/p/original"
	upcomingReleasesPerRequest = 20
	popularShowsPerRequest     = 20
	showSearchesPerRequest     = 8
)

type useCase struct {
	userRepository    repositories.UserRepository
	releaseRepository repositories.ReleaseRepository
	genreRepository   repositories.GenreRepository
	networkRepository repositories.NetworkRepository
	traktClient       *trakt.Client
	tmdbClient        *tmdb.Client
}

type UseCase interface {
	GetShow(showId int) (domain.Show, error)
	GetPopularShows(page int) ([]domain.Show, error)
	GetUpcomingReleases(page int) ([]domain.Release, bool, error)
	SearchShows(searchTerm string) ([]domain.Show, error)
	GetGenres() ([]domain.Genre, error)
	GetNetworks() ([]domain.Network, error)
}

func NewUseCase(
	userRepository repositories.UserRepository, releaseRepository repositories.ReleaseRepository,
	genreRepository repositories.GenreRepository, networkRepository repositories.NetworkRepository,
	traktClient *trakt.Client, tmdbClient *tmdb.Client,
) UseCase {
	return &useCase{
		userRepository,
		releaseRepository,
		genreRepository,
		networkRepository,
		traktClient,
		tmdbClient,
	}
}
