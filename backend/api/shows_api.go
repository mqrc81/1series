package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

type ShowHandler struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

// PopularShows GET /api/shows/popular
func (h *ShowHandler) PopularShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var shows []domain.Show

		page := 1
		if req.URL.Query().Has("page") {
			page, _ = strconv.Atoi(req.URL.Query().Get("page"))
		}

		traktShows, err := h.trakt.GetShowsWatchedWeekly(page, 20)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, traktShow := range traktShows {
			tmdbShow, err := h.tmdb.GetTVDetails(traktShow.TmdbId(), nil)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			shows = append(shows, domain.ShowFromDto(tmdbShow))
		}

		if err = h.respond(res, shows); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Show GET /api/shows/{show_id}
func (h *ShowHandler) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(req, "show_id"))

		tmdbShow, err := h.tmdb.GetTVDetails(id, nil)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		show := domain.ShowFromDto(tmdbShow)

		if err = h.respond(res, show); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// SearchShows GET /api/shows/search
func (h *ShowHandler) SearchShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO
	}
}

func (h *ShowHandler) respond(res http.ResponseWriter, body interface{}) error {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res.Header().Add("Content-Type", "application/json")
	if _, err = res.Write(bodyJSON); err != nil {
		return err
	}

	return nil
}
