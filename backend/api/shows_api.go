package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

type ShowHandler struct {
	store  domain.Store
	tmdb   *tmdb.Client
	trakt  *trakt.Client
	mapper *DtoMapper
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

			shows = append(shows, h.mapper.ShowFromTmdbShow(tmdbShow))
		}

		if err = respond(res, shows); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Show GET /api/shows/{showId}
func (h *ShowHandler) Show() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var show domain.Show

		id, _ := strconv.Atoi(chi.URLParam(req, "showId"))

		tmdbShow, err := h.tmdb.GetTVDetails(id, nil)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		show = h.mapper.ShowFromTmdbShow(tmdbShow)

		if err = respond(res, show); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// SearchShows GET /api/shows/search
func (h *ShowHandler) SearchShows() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var shows []domain.Show

		searchTerm := req.URL.Query().Get("searchTerm")
		if searchTerm == "" {
			http.Error(res, "empty search-term", http.StatusBadRequest)
			return
		}

		tmdbShows, err := h.tmdb.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		shows = h.mapper.ShowsFromTmdbShowsSearch(tmdbShows, 8)

		if err = respond(res, shows); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// UpcomingReleases GET /api/shows/upcoming
func (h *ShowHandler) UpcomingReleases() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var releases []domain.Release

		startDate, _ := time.Parse("2006-01-02", req.URL.Query().Get("startDate"))

		traktReleases, err := h.trakt.GetSeasonPremieres(startDate, 5)

		for _, traktRelease := range traktReleases {
			if traktRelease.IsRelevant() {
				tmdbRelease, err := h.tmdb.GetTVDetails(traktRelease.TmdbId(), nil)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}
				releases = append(releases,
					h.mapper.ReleaseFromTmdbShow(tmdbRelease, traktRelease.SeasonNumber(), traktRelease.AirDate()))
			}
		}

		if err = respond(res, releases); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func respond(res http.ResponseWriter, body interface{}) error {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("unable to parse %v: %w", body, err)
	}

	res.Header().Add("Content-Type", "application/json")
	if _, err = res.Write(bodyJson); err != nil {
		return fmt.Errorf("unable to respond with %v: %w", bodyJson, err)
	}

	return nil
}
