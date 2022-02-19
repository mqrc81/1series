package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/chi/v5"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

var (
	tmdbLanguageEnglishOptions    = map[string]string{"language": "en-US"}
	tmdbAppendTranslationsOptions = map[string]string{"append_to_response": "translations"}
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

		tmdbShows, err := h.tmdb.GetSearchTVShow(searchTerm, tmdbLanguageEnglishOptions)
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

		startDate, days := calculateDateAndDays(req.URL.Query().Get("startDate"))
		traktReleases, err := h.trakt.GetSeasonPremieres(startDate, days)

		// TODO: The relevant upcoming releases should be computed by a scheduler daily
		//  which then stores the tmdb-id, season-number & air-date.
		//  This would drastically improve performance & reduce the amount of external api calls.

		for _, traktRelease := range traktReleases {

			if hasRelevantIds(traktRelease) {
				tmdbRelease, err := h.tmdb.GetTVDetails(traktRelease.TmdbId(), tmdbAppendTranslationsOptions)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				if hasRelevantInfo(tmdbRelease) {
					releases = append(releases,
						h.mapper.ReleaseFromTmdbShow(tmdbRelease, traktRelease.SeasonNumber(), traktRelease.AirDate()))
				}
			}
		}

		if err = respond(res, releases); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func calculateDateAndDays(dateStr string) (time.Time, int) {
	startDate, _ := time.Parse("2006-01-02", dateStr)
	days := 5

	timeDiffInWeeks := startDate.Sub(time.Now()).Hours() / 24 / 7
	if timeDiffInWeeks > 0 {
		days += int(timeDiffInWeeks * 3)
	}
	return startDate, days
}

func hasRelevantIds(release trakt.SeasonPremieresDto) bool {
	ids := release.Show.Ids
	return ids.Tmdb != 0 && ids.Tvdb != 0 && ids.Imdb != "" && ids.Slug != ""
}

func hasRelevantInfo(show *tmdb.TVDetails) bool {
	return len(show.Genres) > 0 &&
		len(show.Networks) > 0 &&
		len(show.Overview) > 0 &&
		hasEnglishTranslation(show.Translations) &&
		hasRelevantType(show)
}

func hasEnglishTranslation(translations *tmdb.TVTranslations) bool {
	for _, translation := range translations.Translations {
		if translation.Iso639_1 == "en" {
			return true
		}
	}
	return false
}

func hasRelevantType(show *tmdb.TVDetails) bool {
	t := show.Type
	if t == "Scripted" || t == "Miniseries" || t == "Documentary" {
		return true
	}
	if t != "Reality" && t != "News" && t != "Talk Show" {
		log.Printf("Unknown type [%v] detected for show [%d, %v]\n", show.Type, show.ID, show.Name)
	}
	return false
}

func respond(res http.ResponseWriter, body interface{}) error {
	res.Header().Add("Content-Type", "application/json")

	// TODO: check if escaping unicode is actually necessary for angular
	e := json.NewEncoder(res)
	e.SetEscapeHTML(false)
	if err := e.Encode(body); err != nil {
		return fmt.Errorf("unable to encode [%v]: %w", body, err)
	}
	return nil
}
