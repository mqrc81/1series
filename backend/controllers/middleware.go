package controllers

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/controllers/users"
	"github.com/mqrc81/zeries/logger"
	"net/http"
	"time"
)

var (
	endpointsExpectingIncreasedLatency = []string{
		"/api/users/importImdbWatchlist",
	}
)

func (c *controller) logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.Error("Http error occurred: request=[%v %v %v] error=[%v] latency=[%v]",
					v.Method, v.URI, v.Status, v.Error, v.Latency)
			} else if v.Latency > 5*time.Second && !endpointExpectsIncreasedLatency(ctx.Request()) {
				logger.Warning("Latency surpassed 5 seconds: request=[%v %v %v] error=[%v] latency=[%v]",
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

func endpointExpectsIncreasedLatency(req *http.Request) bool {
	for _, endpoint := range endpointsExpectingIncreasedLatency {
		if endpoint == req.URL.Path {
			return true
		}
	}
	return false
}

func (c *controller) session() echo.MiddlewareFunc {
	key := securecookie.GenerateRandomKey(32)
	store := sessions.NewCookieStore(key)
	return session.Middleware(store)
}

func (c *controller) withUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			currentSession, err := session.Get(users.SessionKey, ctx)
			if err != nil {
				return err
			}

			userId, ok := currentSession.Values[users.SessionUserIdKey].(int)
			if !ok {
				return next(ctx)
			}

			user, err := c.usersRepository.Find(userId)
			if err != nil {
				return next(ctx)
			}

			ctx.Set(users.SessionUserKey, user)
			return next(ctx)
		}
	}
}
