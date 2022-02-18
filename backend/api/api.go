// Package api
// Internal API interface
package api

import (
	"context"
	"net/http"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/postgres"
	"github.com/mqrc81/zeries/trakt"
)

func Init(store postgres.Store, tmdbClient *tmdb.Client, traktClient *trakt.Client) (*Handler, error) {

	h := &Handler{
		Mux:      chi.NewMux(),
		store:    store,
		sessions: newSessionsManager(store),
	}

	registerMiddleware(h)

	shows := ShowHandler{store, tmdbClient, traktClient}
	h.Route("/api/shows", func(r chi.Router) {
		r.Get("/popular", shows.PopularShows())
		r.Get("/{show_id}", shows.Show())
		r.Get("/search", shows.SearchShows())
	})

	users := UserHandler{store, h.sessions}
	h.Route("/api/users", func(r chi.Router) {
		r.Post("/register", users.Register())
		// r.Post("/login", users.Login())
		// r.Post("/logout", users.Logout())
	})

	h.Get("/api/check", h.HealthCheck())

	return h, nil
}

func registerMiddleware(h *Handler) {
	h.Use(middleware.RequestID)
	h.Use(middleware.Logger)
	h.Use(middleware.Recoverer)
	h.Use(middleware.Timeout(time.Minute))
	h.Use(h.sessions.LoadAndSave)
	h.Use(h.withUser)
}

func (h *Handler) HealthCheck() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write([]byte("Hello World"))
	}
}

func newSessionsManager(store postgres.Store) *scs.SessionManager {
	sessions := scs.New()
	sessions.Store = postgresstore.New(store.DB.DB)
	return sessions
}

func (h *Handler) withUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		userId := h.sessions.GetInt(req.Context(), "user_id")
		if userId == 0 {
			next.ServeHTTP(res, req)
			return
		}

		user, err := h.store.GetUser(userId)
		if err != nil {
			next.ServeHTTP(res, req)
			return
		}

		ctx := context.WithValue(req.Context(), "user", user)

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

type Handler struct {
	*chi.Mux

	store    domain.Store
	sessions *scs.SessionManager
}
