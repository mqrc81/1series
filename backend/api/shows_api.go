package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

const (
	releasesPerRequest = 20
)

type ShowHandler struct {
	store    domain.Store
	sessions *scs.SessionManager
	tmdb     *tmdb.Client
	trakt    *trakt.Client
}

// Popular GET /api/shows/popular
func (h *ShowHandler) Popular() echo.HandlerFunc {
	const pageParam QueryParam = "page"

	return func(ctx echo.Context) error {
		page, _ := strconv.Atoi(ctx.QueryParam(pageParam))
		if page < 1 {
			page = 1
		}

		traktShows, err := h.trakt.GetShowsWatchedWeekly(page, 20)
		if err != nil {
			return NewHttpError(Trakt, err, page)
		}

		var shows []domain.Show
		for _, traktShow := range traktShows {
			tmdbShow, err := h.tmdb.GetTVDetails(traktShow.TmdbId(), nil)
			if err != nil {
				return NewHttpError(Tmdb, err, traktShow.Ids())
			}

			shows = append(shows, showFromTmdbShow(tmdbShow))
		}

		return ctx.JSON(http.StatusOK, shows)
	}
}

// Show GET /api/shows/:showId
func (h *ShowHandler) Show() echo.HandlerFunc {
	const showIdParam UrlParam = "showId"

	return func(ctx echo.Context) error {
		showId, err := strconv.Atoi(ctx.Param(showIdParam))
		if err != nil {
			return NewHttpError(Parameter, fmt.Errorf("invalid show-id"), ctx.Param(showIdParam))
		}

		tmdbShow, err := h.tmdb.GetTVDetails(showId, nil)
		if err != nil {
			return NewHttpError(Tmdb, err, showId)
		}

		return ctx.JSON(http.StatusOK, showFromTmdbShow(tmdbShow))
	}
}

// Search GET /api/shows/search
func (h *ShowHandler) Search() echo.HandlerFunc {
	const searchTermParam QueryParam = "searchTerm"

	return func(ctx echo.Context) error {
		searchTerm := ctx.Param(searchTermParam)
		if searchTerm == "" {
			return NewHttpError(Parameter, fmt.Errorf("invalid search-term"), searchTerm)
		}

		tmdbShows, err := h.tmdb.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
		if err != nil {
			return NewHttpError(Tmdb, err, searchTerm)
		}

		return ctx.JSON(http.StatusOK, showsFromTmdbShowsSearch(tmdbShows, 8))
	}
}

// Releases GET /api/shows/releases
func (h *ShowHandler) Releases() echo.HandlerFunc {
	const pageParam QueryParam = "page"

	return func(ctx echo.Context) error {
		pastReleases, err := h.store.GetPastReleasesCount()
		if err != nil {
			return NewHttpError(Database, err)
		}

		amount, offset := calculateRange(ctx.QueryParam(pageParam), pastReleases)

		releasesRef, err := h.store.GetReleases(amount, offset)
		if err != nil {
			return NewHttpError(Database, err)
		}

		var releases []domain.Release
		for _, releaseRef := range releasesRef {
			tmdbRelease, err := h.tmdb.GetTVDetails(releaseRef.ShowId,
				map[string]string{"append_to_response": "translations"})
			if err != nil {
				return NewHttpError(Tmdb, err, tmdbRelease.Name)
			}

			releases = append(releases, releaseFromTmdbShow(
				tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate, releaseRef.AnticipationLevel))
		}

		return ctx.JSON(http.StatusOK, releases)
	}
}

func calculateRange(pageQueryParam string, pastReleases int) (int, int) {
	page, _ := strconv.Atoi(pageQueryParam)

	if page < 0 {
		return calculateRangeForPastReleases(pastReleases, page)
	}
	return calculateRangeForUpcomingReleases(pastReleases, page)
}

func calculateRangeForUpcomingReleases(pastReleases int, page int) (int, int) {
	// For pages 0+ return 20 releases
	return releasesPerRequest, pastReleases + releasesPerRequest*(page)
}

func calculateRangeForPastReleases(pastReleases int, page int) (int, int) {
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
