// Package api is the internal API interface used solely in Angular
package api

import (
	"io"
	"net/http"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/util"
)

func NewHandler(store domain.Store, sessionStore sessions.Store, tmdbClient *tmdb.Client,
	traktClient *trakt.Client) (*Handler, error) {

	h := &Handler{
		echo.New(),
		store,
	}

	h.Logger.SetOutput(io.Discard)
	h.HideBanner = true
	h.Validator = NewValidator(store)

	shows := &ShowHandler{store, tmdbClient, traktClient, &DtoMapper{}}
	users := &UserHandler{store}

	h.Use(
		middleware.RequestID(),
		middleware.Recover(),
		h.logRequest(),
		// middleware.CSRF(),
		session.Middleware(sessionStore),
		h.withUser(),
	)

	showsApi := h.Group("/api/shows")
	{
		showsApi.GET("/popular", shows.PopularShows())
		showsApi.GET("/:showsId", shows.Show())
		showsApi.GET("/search", shows.SearchShows())
		showsApi.GET("/releases", shows.UpcomingReleases())
	}

	usersApi := h.Group("/api/users")
	{
		usersApi.POST("/register", users.Register())
		// usersApi.POST("/login", users.Login())
		// usersApi.POST("/logout", users.Logout())
	}

	h.GET("/api/ping", h.Ping())

	return h, nil
}

func (h *Handler) Ping() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		util.LogWarn("Pong",
			"idk", 3)
		return ctx.String(http.StatusOK, "Pong!")
	}
}

func (h *Handler) withUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if mySession, err := session.Get("my-session", ctx); err == nil {
				if userId, ok := mySession.Values["userId"].(int); ok {
					if user, err := h.store.GetUser(userId); err == nil {
						ctx.Set("user", user)
					}
				}
			}
			return next(ctx)
		}
	}
}

func (h *Handler) logRequest() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				util.LogError("Http error occurred: request=[%v %v %v] error=[%v] latency=[%v]",
					v.Method, v.URI, v.Status, v.Error, v.Latency)
			} else if v.Latency > 3*time.Second {
				util.LogWarn("Latency surpassed 3 seconds: request=[%v %v %v] error=[%v] latency=[%v]",
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

type Handler struct {
	*echo.Echo
	store domain.Store
}

// QueryParam & UrlParam don't serve a real purpose other than
// clearer documentation of all params used in each endpoint
type QueryParam = string
type UrlParam = string
