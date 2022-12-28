package shows

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

const (
	tmdbImageBaseUrl           = "https://image.tmdb.org/t/p/original"
	upcomingReleasesPerRequest = 20
	popularShowsPerRequest     = 20
	showSearchesPerRequest     = 8
)

type showsController struct {
	usersRepository    repositories.UsersRepository
	releasesRepository repositories.ReleasesRepository
	genresRepository   repositories.GenresRepository
	networksRepository repositories.NetworksRepository
	traktClient        *trakt.Client
	tmdbClient         *tmdb.Client
}

type Controller interface {
	GetShow(ctx echo.Context) error
	GetPopularShows(ctx echo.Context) error
	GetUpcomingReleases(ctx echo.Context) error
	SearchShows(ctx echo.Context) error
	GetGenres(ctx echo.Context) error
	GetNetworks(ctx echo.Context) error
}

func NewController(
	usersRepository repositories.UsersRepository,
	releasesRepository repositories.ReleasesRepository,
	genresRepository repositories.GenresRepository,
	networksRepository repositories.NetworksRepository,
	traktClient *trakt.Client, tmdbClient *tmdb.Client,
) Controller {
	return &showsController{
		usersRepository,
		releasesRepository,
		genresRepository,
		networksRepository,
		traktClient,
		tmdbClient,
	}
}
