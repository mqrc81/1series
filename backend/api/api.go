// Package api
// Internal API interface
package api

import (
	"net/http"

	"github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mqrc81/zeries/trakt"
)

func Init(tmdbClient *tmdb.Client, traktClient *trakt.Client) *Controller {

	h := &Controller{
		Mux: chi.NewMux(),
	}

	shows := ShowController{tmdbClient, traktClient}
	// users := UserHandler{store, sessions}

	h.Use(middleware.Logger)
	// h.Use(sessions.LoadAndSave)
	// h.Use(h.withUser)

	h.Route("/api/shows", func(r chi.Router) {
		r.Get("/popular", shows.PopularShows())
		r.Get("/{show_id}", shows.Show())
		r.Get("/search", shows.SearchShows())
	})

	h.Get("/api/check", h.Check())

	return h
}

func (h *Controller) Check() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		if _, err := res.Write([]byte("Hello " + name)); err != nil {
			return
		}
	}
}

type Controller struct {
	*chi.Mux
}
