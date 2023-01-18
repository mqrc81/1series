package controllers

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/1series/controllers/admin"
	"github.com/mqrc81/1series/controllers/shows"
	"github.com/mqrc81/1series/controllers/users"
	"github.com/mqrc81/1series/email"
	"github.com/mqrc81/1series/repositories"
	"github.com/mqrc81/1series/sql"
	"github.com/mqrc81/1series/trakt"
	"io"
	"net/http"
)

type Controller interface {
	Start(address string) error
}

type controller struct {
	*echo.Echo
	usersRepository  repositories.UsersRepository
	tokensRepository repositories.TokensRepository
}

func NewController(
	database *sql.Database,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
	emailClient *email.Client,
	scheduler *gocron.Scheduler,
) (Controller, error) {

	usersRepository := repositories.NewUsersRepository(database)
	releasesRepository := repositories.NewReleasesRepository(database)
	genresRepository := repositories.NewGenresRepository(database)
	networksRepository := repositories.NewNetworksRepository(database)
	trackedShowsRepository := repositories.NewTrackedShowsRepository(database)
	tokensRepository := repositories.NewTokensRepository(database)

	validate := validator.New()

	c := newController(usersRepository, tokensRepository).withMiddleware()

	adminController := admin.NewController(scheduler)
	usersController := users.NewController(usersRepository, trackedShowsRepository, tokensRepository, tmdbClient, emailClient, validate)
	showsController := shows.NewController(usersRepository, releasesRepository, genresRepository, networksRepository, traktClient, tmdbClient)

	baseRouter := c.Group("/api")
	{
		baseRouter.GET("/ping", c.Ping)
		baseRouter.GET("/init", c.Init)
	}

	adminRouter := baseRouter.Group("/admin", c.adminOnly())
	{
		adminRouter.GET("/triggerJobs", adminController.TriggerJobs)
	}

	showsRouter := baseRouter.Group("/shows")
	{
		showsRouter.GET("/:showId", showsController.GetShow)
		showsRouter.GET("/popular", showsController.GetPopularShows)
		showsRouter.GET("/releases", showsController.GetUpcomingReleases)
		showsRouter.GET("/search", showsController.SearchShows)
		showsRouter.GET("/genres", showsController.GetGenres)
		showsRouter.GET("/networks", showsController.GetNetworks)
	}

	usersRouter := baseRouter.Group("/users")
	{
		usersRouter.POST("/signUp", usersController.SignUserUp)
		usersRouter.POST("/signIn", usersController.SignUserIn)
		usersRouter.POST("/signOut", usersController.SignUserOut)
		usersRouter.POST("/forgotPassword", usersController.ForgotPassword)
		usersRouter.POST("/resetPassword", usersController.ResetPassword)
		usersRouter.POST("/importImdbWatchlist", usersController.ImportImdbWatchlist)
	}

	return c, nil
}

func newController(
	usersRepository repositories.UsersRepository,
	tokensRepository repositories.TokensRepository,
) *controller {
	e := echo.New()
	e.Logger.SetOutput(io.Discard)
	e.HideBanner = true
	return &controller{e, usersRepository, tokensRepository}
}

func (c *controller) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}

func (c *controller) Init(ctx echo.Context) error {
	if user, err := users.GetAuthenticatedUser(ctx); err != nil {
		return ctx.NoContent(http.StatusOK)
	} else {
		return ctx.JSON(http.StatusOK, user)
	}
}
