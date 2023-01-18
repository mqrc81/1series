package controllers

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/1series/controllers/errors"
	"github.com/mqrc81/1series/controllers/users"
	"github.com/mqrc81/1series/domain"
	"github.com/mqrc81/1series/env"
	"github.com/mqrc81/1series/logger"
	"net/http"
	"time"
)

const (
	csrfTokenKey = "_csrf"
)

var (
	endpointsExpectingIncreasedLatency = []string{
		"/api/users/importImdbWatchlist",
	}
)

func (c *controller) withMiddleware() *controller {
	c.Use(
		middleware.RequestID(),
		middleware.Recover(),
		c.logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{env.Config().FrontendUrl}, AllowCredentials: true}),
		middleware.CSRFWithConfig(middleware.CSRFConfig{TokenLookup: "token:" + csrfTokenKey}),
		c.session(),
		c.withUser(),
	)
	return c
}

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
	key := []byte(env.Config().SessionKey)
	store := sessions.NewCookieStore(key)
	return session.Middleware(store)
}

func (c *controller) withUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			sess, err := session.Get(users.SessionKey, ctx)
			if err != nil {
				return err
			}

			userId, ok := sess.Values[users.SessionUserIdKey].(int)
			if !ok {
				cookie, err := ctx.Cookie(users.RememberLoginTokenCookieName)
				if err != nil {
					return next(ctx)
				}

				token, err := c.tokensRepository.Find(cookie.Value)
				if err != nil || token.IsExpired() {
					cookie.Expires = time.Now()
					ctx.SetCookie(cookie)
					return next(ctx)
				}

				userId = token.UserId

				newToken := domain.CreateToken(domain.RememberLogin, userId)
				if err = c.tokensRepository.SaveOrReplace(newToken); err != nil {
					cookie.Expires = time.Now()
					ctx.SetCookie(cookie)
					return next(ctx)
				}

				ctx.SetCookie(&http.Cookie{
					Name:     users.RememberLoginTokenCookieName,
					Value:    newToken.Id,
					Expires:  newToken.ExpiresAt,
					SameSite: http.SameSiteDefaultMode,
				})
			}

			user, err := c.usersRepository.Find(userId)
			if err != nil {
				delete(sess.Values, users.SessionUserIdKey)
				_ = sess.Save(ctx.Request(), ctx.Response())
				return next(ctx)
			}

			ctx.Set(users.ContextUserKey, user)
			return next(ctx)
		}
	}
}

func (c *controller) adminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if user, err := users.GetAuthenticatedUser(ctx); err != nil || !users.IsAdmin(user) {
				return errors.AdminOnly()
			}
			return next(ctx)
		}
	}
}
