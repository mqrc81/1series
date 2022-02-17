package api

import (
	"net/http"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

type ShowHandler struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

func (h *ShowHandler) PopularShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowHandler) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowHandler) SearchShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}
