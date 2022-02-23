package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cyruzin/golang-tmdb"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	"go.uber.org/zap"
)

const (
	releasesPerRequest = 20
)

type ShowHandler struct {
	store  domain.Store
	tmdb   *tmdb.Client
	trakt  *trakt.Client
	mapper *DtoMapper
	log    *zap.SugaredLogger
}

// PopularShows GET /api/shows/popular
func (h *ShowHandler) PopularShows() echo.HandlerFunc {
	const pageParam QueryParam = "page"

	return func(ctx echo.Context) error {
		page, _ := strconv.Atoi(ctx.QueryParam(pageParam))
		if page < 1 {
			page = 1
		}

		traktShows, err := h.trakt.GetShowsWatchedWeekly(page, 20)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("trakt error fetching most-watched-shows [%v]: %v", page, err.Error()))
		}

		var shows []domain.Show
		for _, traktShow := range traktShows {
			tmdbShow, err := h.tmdb.GetTVDetails(traktShow.TmdbId(), nil)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError,
					fmt.Sprintf("tmdb error fetching tv-details [%v]: %v", traktShow.Ids(), err.Error()))
			}

			shows = append(shows, h.mapper.ShowFromTmdbShow(tmdbShow))
		}

		return ctx.JSON(http.StatusOK, shows)
	}
}

// Show GET /api/shows/{showId}
func (h *ShowHandler) Show() echo.HandlerFunc {
	const showIdParam UrlParam = "showId"

	return func(ctx echo.Context) error {
		showId, err := strconv.Atoi(ctx.Param(showIdParam))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid show-id [%v]", ctx.Param(showIdParam)))
		}

		tmdbShow, err := h.tmdb.GetTVDetails(showId, nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("tmdb error fetching tv-details [%d]: %v", showId, err.Error()))
		}

		return ctx.JSON(http.StatusOK, h.mapper.ShowFromTmdbShow(tmdbShow))
	}
}

// SearchShows GET /api/shows/search
func (h *ShowHandler) SearchShows() echo.HandlerFunc {
	const searchTermParam QueryParam = "searchTerm"

	return func(ctx echo.Context) error {
		searchTerm := ctx.Param(searchTermParam)
		if searchTerm == "" {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid search-term [%v]", searchTerm))
		}

		tmdbShows, err := h.tmdb.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError,
				fmt.Sprintf("tmdb error fetching search-tv-show [%v]: %v", searchTerm, err.Error()))
		}

		return ctx.JSON(http.StatusOK, h.mapper.ShowsFromTmdbShowsSearch(tmdbShows, 8))
	}
}

// UpcomingReleases GET /api/shows/releases
func (h *ShowHandler) UpcomingReleases() echo.HandlerFunc {
	const pageParam QueryParam = "page"

	return func(ctx echo.Context) error {
		pastReleases, err := h.store.GetPastReleasesCount()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		amount, offset := calculateRange(ctx.QueryParam(pageParam), pastReleases)

		releasesRef, err := h.store.GetReleases(amount, offset)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var releases []domain.Release
		for _, releaseRef := range releasesRef {
			tmdbRelease, err := h.tmdb.GetTVDetails(releaseRef.ShowId,
				map[string]string{"append_to_response": "translations"})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError,
					fmt.Sprintf("tmdb error fetching tv-details [%v]: %v", tmdbRelease.Name, err.Error()))
			}

			releases = append(releases,
				h.mapper.ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate))
		}

		return ctx.JSON(http.StatusOK, releases)
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
