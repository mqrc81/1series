package controller

import (
	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/usecase"
	session "github.com/spazzymoto/echo-scs-session"
	"io"
	"net/http"
)

type Controller interface {
	Start(address string) error
}

func NewController(
	database *sqlx.DB, sessionManager *scs.SessionManager, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (Controller, error) {

	userRepository := repository.NewUserRepository(database)
	releaseRepository := repository.NewReleaseRepository(database)

	showUseCase := usecase.NewShowUseCase(userRepository, releaseRepository, traktClient, tmdbClient)
	userUseCase := usecase.NewUserUseCase(userRepository, sessionManager)

	echoEngine := newEcho()

	baseRouter := echoEngine.Group("/api")
	echoEngine.GET("/api/ping", ping)

	showController := newShowController(showUseCase)
	showRouter := baseRouter.Group("/shows")
	showRouter.GET("/:showId", showController.GetShow)
	showRouter.GET("/popular", showController.GetPopularShows)
	showRouter.GET("/releases", showController.GetUpcomingReleases)
	showRouter.GET("/search", showController.SearchShows)

	userController := newUserController(userUseCase)
	userRouter := baseRouter.Group("/users")
	userRouter.POST("/register", userController.RegisterUser)

	echoEngine.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.CORS(),
		logRequest(),
		// middleware.CSRF(),
		session.LoadAndSave(sessionManager),
		withUser(sessionManager, userRepository),
	)

	return echoEngine, nil
}

func newEcho() *echo.Echo {
	echoEngine := echo.New()
	echoEngine.Logger.SetOutput(io.Discard)
	echoEngine.HideBanner = true
	return echoEngine
}

func ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}
