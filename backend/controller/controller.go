package controller

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/usecase"
	"io"
	"net/http"
)

var (
	corsAllowOrigins = []string{"*localhost:*", "*127.0.0.1:*", "*.up.railway.app*"}
)

type Controller interface {
	Start(address string) error
}

type controller struct {
	*echo.Echo
	userRepository repository.UserRepository
}

func NewController(
	database *sqlx.DB, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (Controller, error) {

	validate := validator.New()

	userRepository := repository.NewUserRepository(database)
	releaseRepository := repository.NewReleaseRepository(database)

	showUseCase := usecase.NewShowUseCase(userRepository, releaseRepository, traktClient, tmdbClient)
	userUseCase := usecase.NewUserUseCase(userRepository)

	controller := newController(userRepository)

	baseRouter := controller.Group("/api")
	{
		baseRouter.GET("/ping", controller.ping)

	}

	showController := newShowController(showUseCase)
	showRouter := baseRouter.Group("/shows")
	{
		showRouter.GET("/:showId", showController.GetShow)
		showRouter.GET("/popular", showController.GetPopularShows)
		showRouter.GET("/releases", showController.GetUpcomingReleases)
		showRouter.GET("/search", showController.SearchShows)
	}

	userController := newUserController(userUseCase, validate)
	userRouter := baseRouter.Group("/users")
	{
		userRouter.POST("/register", userController.RegisterUser)
		userRouter.POST("/login", userController.LoginUser)
	}

	controller.Use(
		middleware.RequestID(),
		middleware.Recover(),
		controller.logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: corsAllowOrigins}),
		middleware.CSRF(),
		controller.session(),
		controller.withUser(),
	)

	return controller, nil
}

func newController(userRepository repository.UserRepository) controller {
	echoEngine := echo.New()
	echoEngine.Logger.SetOutput(io.Discard)
	echoEngine.HideBanner = true
	return controller{echoEngine, userRepository}
}

func (c *controller) ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}
