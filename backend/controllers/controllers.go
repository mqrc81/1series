package controllers

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/controllers/jobs"
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
	userRepository repositories.UserRepository
}

func NewController(
	database *sql.Database,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
	emailClient *email.Client,
	scheduler *gocron.Scheduler,
) (Controller, error) {

	userRepository := repositories.NewUserRepository(database)
	releaseRepository := repositories.NewReleaseRepository(database)
	genreRepository := repositories.NewGenreRepository(database)
	networkRepository := repositories.NewNetworkRepository(database)
	trackedShowRepository := repositories.NewTrackedShowRepository(database)

	validate := validator.New()

	baseController := newController(userRepository)
	userController := users.NewController(userRepository, trackedShowRepository, tmdbClient, emailClient, validate)
	showController := shows.NewController(userRepository, releaseRepository, genreRepository, networkRepository, traktClient, tmdbClient)
	jobController := jobs.NewController(scheduler)

	baseRouter := baseController.Group("/api")
	{
		baseRouter.GET("/ping", baseController.Ping)
	}

	showRouter := baseRouter.Group("/shows")
	{
		showRouter.GET("/:showId", showController.GetShow)
		showRouter.GET("/popular", showController.GetPopularShows)
		showRouter.GET("/releases", showController.GetUpcomingReleases)
		showRouter.GET("/search", showController.SearchShows)
		showRouter.GET("/genres", showController.GetGenres)
		showRouter.GET("/networks", showController.GetNetworks)
	}

	userRouter := baseRouter.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.LoginUser)
		userRouter.POST("/importImdbWatchlist", userController.ImportImdbWatchlist)
	}

	jobRouter := baseRouter.Group("/jobs")
	{
		jobRouter.GET("/runByTag", jobController.RunJobsByTag)
	}

	baseController.Use(
		middleware.RequestID(),
		middleware.Recover(),
		baseController.logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: corsAllowOrigins, AllowCredentials: true}),
		middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLookup: "token:_csrf"}),
		baseController.session(),
		baseController.withUser(),
	)

	return baseController, nil
}

func newController(userRepository repositories.UserRepository) controller {
	echoEngine := echo.New()
	echoEngine.Logger.SetOutput(io.Discard)
	echoEngine.HideBanner = true
	return controller{echoEngine, userRepository}
}

func (c *controller) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}
