package controller

import (
	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/repository"
	. "github.com/mqrc81/zeries/util"
	"time"
)

func logRequest() echo.MiddlewareFunc {
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

func withUser(
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
