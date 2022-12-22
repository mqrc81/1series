package controllers

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/logger"
	"time"
)

const (
	sessionKey       = "session"
	sessionUserIdKey = "userId"
	sessionUserKey   = "user"
)

func (c *controller) logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				logger.Error("Http error occurred: request=[%v %v %v] error=[%v] latency=[%v]",
					v.Method, v.URI, v.Status, v.Error, v.Latency)
			} else if v.Latency > 5*time.Second {
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

func (c *controller) session() echo.MiddlewareFunc {
	key := securecookie.GenerateRandomKey(32)
	store := sessions.NewCookieStore(key)
	return session.Middleware(store)
}

func (c *controller) withUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			currentSession, err := session.Get(sessionKey, ctx)
			if err != nil {
				return err
			}

			userId, ok := currentSession.Values[sessionUserIdKey].(int)
			if !ok {
				return next(ctx)
			}

			user, err := c.userRepository.Find(userId)
			if err != nil {
				return next(ctx)
			}

			ctx.Set(sessionUserKey, user)
			return next(ctx)
		}
	}
}
