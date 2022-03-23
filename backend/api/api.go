// Package api is the internal API interface used solely in Angular
package api

import (
	"io"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	. "github.com/mqrc81/zeries/util"
	echoscs "github.com/spazzymoto/echo-scs-session"
)

func NewHandler(store domain.Store, sessionManager *scs.SessionManager, tmdbClient *tmdb.Client,
	traktClient *trakt.Client) (*Handler, error) {

	h := &Handler{
		echo.New(),
		store,
		sessionManager,
	}

	h.Logger.SetOutput(io.Discard)
	h.HideBanner = true

	shows := &ShowHandler{store, sessionManager, tmdbClient, traktClient}
	users := &UserHandler{store, sessionManager}

	h.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.CORS(),
		h.logRequest(),
		// middleware.CSRF(),
		echoscs.LoadAndSave(sessionManager),
		h.withUser(),
	)

	showsApi := h.Group("/api/shows")
	{
		showsApi.GET("/:showId", shows.Show())
		showsApi.GET("/popular", shows.Popular())
		showsApi.GET("/releases", shows.Releases())
		showsApi.GET("/search", shows.Search())
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
		return ctx.String(http.StatusOK, "Pong!")
	}
}

func (h *Handler) withUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if userId := h.sessions.GetInt(ctx.Request().Context(), "userId"); userId > 0 {
				if user, err := h.store.GetUser(userId); err == nil {
					ctx.Set("user", user)
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

type Handler struct {
	*echo.Echo
	store    domain.Store
	sessions *scs.SessionManager
}

// QueryParam & UrlParam don't serve a real purpose other than
// clearer documentation of all params used in each endpoint
type QueryParam = string
type UrlParam = string
