package controller

import (
	"io"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/usecase"
	. "github.com/mqrc81/zeries/util"
)

type controller struct {
	*echo.Echo
	showController ShowController
	userController UserController
}

type Controller interface {
	Start(port string) error
	Ping(ctx echo.Context) error
}

func NewController(
	database *sqlx.DB, sessionManager *scs.SessionManager, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (Controller, error) {

	userRepository := repository.NewUserRepository(database)
	releaseRepository := repository.NewReleaseRepository(database)

	showUseCase := usecase.NewShowUseCase(userRepository, releaseRepository, traktClient, tmdbClient)
	userUseCase := usecase.NewUserUseCase(userRepository, sessionManager)

	echoEngine := NewEcho()

	router := echoEngine.Group("/api")
	controller := &controller{
		echoEngine,
		NewShowController(showUseCase, router.Group("/shows")),
		NewUserController(userUseCase, router.Group("/users")),
	}

	router.GET("/ping", func(ctx echo.Context) error {
		return controller.Ping(ctx)
	})

	controller.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.CORS(),
		controller.logRequest(),
		// middleware.CSRF(),
		controller.withUser(sessionManager, userRepository),
	)

	return controller, nil
}

func NewEcho() *echo.Echo {
	echoEngine := echo.New()
	echoEngine.Logger.SetOutput(io.Discard)
	echoEngine.HideBanner = true
	return echoEngine
}

func (c *controller) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Pong!")
}

func (c *controller) logRequest() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				LogError("Http error occurred: request=[%v %v %v] error=[%v] latency=[%v]",
					v.Method, v.URI, v.Status, v.Error, v.Latency)
			} else if v.Latency > 5*time.Second {
				LogWarning("Latency surpassed 5 seconds: request=[%v %v %v] error=[%v] latency=[%v]",
					v.Method, v.URI, v.Status, v.Error, v.Latency)
			}
			return nil
		},
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogError:   true,
		LogLatency: true,
	})
}

func (c *controller) withUser(
	sessionManager *scs.SessionManager, userRepository repository.UserRepository,
) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			userId := sessionManager.GetInt(ctx.Request().Context(), "userId")
			if userId <= 0 {
				return next(ctx)
			}

			user, err := userRepository.Find(userId)
			if err != nil {
				return next(ctx)
			}

			ctx.Set("user", user)
			return next(ctx)
		}
	}
}
