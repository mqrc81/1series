package api

import (
	"net/http"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/trakt"
)

type ShowController struct {
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

func (h *ShowController) PopularShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowController) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowController) SearchShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}
