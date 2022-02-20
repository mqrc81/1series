package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

const (
	releasesPerRequest = 20
)

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

		pastReleases, err := h.store.GetPastReleasesCount()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		amount, offset := calculateRange(req.URL.Query().Get("page"), pastReleases)

		releasesRef, err := h.store.GetReleases(amount, offset)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, releaseRef := range releasesRef {
			tmdbRelease, err := h.tmdb.GetTVDetails(releaseRef.ShowId,
				map[string]string{"append_to_response": "translations"})
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			releases = append(releases,
				h.mapper.ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate))
		}

		if err = respond(res, releases); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func calculateRange(pageQueryParam string, pastReleases int) (int, int) {
	page, _ := strconv.Atoi(pageQueryParam)

	if page == 0 {
		// For first request, return 40 releases
		return releasesPerRequest * 2, pastReleases
	} else if page > 0 {
		// For pages 1+ return 20 releases
		return releasesPerRequest, pastReleases + releasesPerRequest*(page+1)
	} else {
		// For negative pages return 20 releases or max releases left for last page
		offset := pastReleases + releasesPerRequest*page
		amount := releasesPerRequest
		if offset <= 0 {
			// The last possible page for past releases has been reached
			amount += offset
			offset = 0
		}
		return amount, offset
	}
}

func respond(res http.ResponseWriter, body interface{}) error {
	res.Header().Add("Content-Type", "application/json")

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error parsing [%v] to json: %w", body, err)
	}

	if _, err = res.Write(bodyJson); err != nil {
		return fmt.Errorf("error responding with [%v]: %w", bodyJson, err)
	}

	return nil
}
