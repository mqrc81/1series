// Package api is the internal API interface used solely in Angular
package api

import (
	"net/http"

	"github.com/cyruzin/golang-tmdb"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

func Init(store domain.Store, sessionStore sessions.Store,
	tmdbClient *tmdb.Client, traktClient *trakt.Client) (*Handler, error) {

	h := &Handler{
		echo.New(),
		store,
	}

	shows := &ShowHandler{store, tmdbClient, traktClient, new(DtoMapper)}
	users := &UserHandler{store}

	h.Use(
		middleware.RequestID(),
		middleware.Logger(),
		middleware.Recover(),
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

	h.GET("/ping", h.Ping())

	return h, nil
}

func (h *Handler) Ping() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Pong!")
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

type Handler struct {
	*echo.Echo
	store domain.Store
}

// QueryParam & UrlParam don't serve a real purpose other than
// clearer documentation of all params used in each endpoint
type QueryParam = string
type UrlParam = string
