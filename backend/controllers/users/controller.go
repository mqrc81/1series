package users

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/repositories"
)

const (
	imdbWatchlistImportFileName = "WATCHLIST.csv"
	imdbTitleTypeTvSeries       = "tvSeries"
	imdbTitleTypeTvMiniSeries   = "tvMiniSeries"
)

type usersController struct {
	usersRepository        repositories.UsersRepository
	trackedShowsRepository repositories.TrackedShowsRepository
	tmdbClient             *tmdb.Client
	emailClient            *email.Client
	validate               *validator.Validate
}

type Controller interface {
	SignUserUp(ctx echo.Context) error
	SignUserIn(ctx echo.Context) error
	SignUserOut(ctx echo.Context) error
	ImportImdbWatchlist(ctx echo.Context) error
}

func NewController(
	usersRepository repositories.UsersRepository,
	trackedShowsRepository repositories.TrackedShowsRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
	validate *validator.Validate,
) Controller {
	return &usersController{
		usersRepository,
		trackedShowsRepository,
		tmdbClient,
		emailClient,
		validate,
	}
}
