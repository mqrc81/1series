// Package api is the internal API interface used solely in Angular
package api

import (
	"net/http"

	"github.com/cyruzin/golang-tmdb"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

func Init(store domain.Store, sessionStore sessions.Store,
	tmdbClient *tmdb.Client, traktClient *trakt.Client) (*Handler, error) {

	h := &Handler{
		gin.Default(),
		store,
	}

	shows := &ShowHandler{store, tmdbClient, traktClient, new(DtoMapper)}
	users := &UserHandler{store}

	h.Use(
		gin.Logger(),
		gin.Recovery(),
		sessions.Sessions("session", sessionStore),
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

func (h *Handler) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong!")
	}
}

func (h *Handler) withUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if userId, ok := session.Get("userId").(int); ok {
			if user, err := h.store.GetUser(userId); err == nil {
				ctx.Set("user", user)
			}
		}

		ctx.Next()
	}
}

func httpError400(ctx *gin.Context, err error) {
	_ = ctx.AbortWithError(http.StatusBadRequest, err)
}

func httpError500(ctx *gin.Context, err error) {
	_ = ctx.AbortWithError(http.StatusInternalServerError, err)
}

type Handler struct {
	*gin.Engine
	store domain.Store
}

// UrlQuery & UrlParam don't serve a real purpose other than
// clearer documentation of all params used in each endpoint
type UrlQuery = string
type UrlParam = string
