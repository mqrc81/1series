package controllers

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/controllers/admin"
	"github.com/mqrc81/zeries/controllers/shows"
	"github.com/mqrc81/zeries/controllers/users"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/sql"
	"github.com/mqrc81/zeries/trakt"
	"io"
	"net/http"
)

var (
	corsAllowOrigins = []string{"http://127.0.0.1:4000", "https://next.up.railway.app", "https://dev-next.up.railway.app"}
)

type Controller interface {
	Start(address string) error
}

type controller struct {
	*echo.Echo
	usersRepository repositories.UsersRepository
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

	validate := validator.New()

	c := newController(usersRepository)
	adminController := admin.NewController(scheduler)
	usersController := users.NewController(usersRepository, trackedShowsRepository, tmdbClient, emailClient, validate)
	showsController := shows.NewController(usersRepository, releasesRepository, genresRepository, networksRepository, traktClient, tmdbClient)

	baseRouter := c.Group("/api")
	{
		baseRouter.GET("/ping", c.Ping)
	}

	adminRouter := baseRouter.Group("/admin")
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
		usersRouter.POST("/register", usersController.RegisterUser)
		usersRouter.POST("/login", usersController.LoginUser)
		usersRouter.POST("/importImdbWatchlist", usersController.ImportImdbWatchlist)
	}

	c.Use(
		middleware.RequestID(),
		middleware.Recover(),
		c.logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: corsAllowOrigins, AllowCredentials: true}),
		middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLookup: "token:_csrf"}),
		c.session(),
		c.withUser(),
	)

	return c, nil
}

func newController(usersRepository repositories.UsersRepository) controller {
	echoEngine := echo.New()
	echoEngine.Logger.SetOutput(io.Discard)
	echoEngine.HideBanner = true
	return controller{echoEngine, usersRepository}
}

func (c *controller) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}
