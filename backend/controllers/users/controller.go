package users

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/repositories"
)

type userController struct {
	userRepository        repositories.UserRepository
	trackedShowRepository repositories.TrackedShowRepository
	tmdbClient            *tmdb.Client
	emailClient           *email.Client
	validate              *validator.Validate
}

type Controller interface {
	RegisterUser(ctx echo.Context) error
	LoginUser(ctx echo.Context) error
	ImportImdbWatchlist(ctx echo.Context) error
}

func NewController(
	userRepository repositories.UserRepository,
	trackedShowRepository repositories.TrackedShowRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
	validate *validator.Validate,
) Controller {
	return &userController{
		userRepository,
		trackedShowRepository,
		tmdbClient,
		emailClient,
		validate,
	}
}
