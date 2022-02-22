package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cyruzin/golang-tmdb"
	"github.com/gin-gonic/gin"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

const (
	releasesPerRequest = 20
)

type ShowHandler struct {
	store  domain.Store
	tmdb   *tmdb.Client
	trakt  *trakt.Client
	mapper *DtoMapper
}

// PopularShows GET /api/shows/popular
func (h *ShowHandler) PopularShows() gin.HandlerFunc {
	const pageQuery UrlQuery = "page"

	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.Query(pageQuery))
		if page < 1 {
			page = 1
		}

		traktShows, err := h.trakt.GetShowsWatchedWeekly(page, 20)
		if err != nil {
			httpError500(ctx, fmt.Errorf("trakt error most-watched-shows [%v]: %w", page, err))
			return
		}

		var shows []domain.Show
		for _, traktShow := range traktShows {
			tmdbShow, err := h.tmdb.GetTVDetails(traktShow.TmdbId(), nil)
			if err != nil {
				httpError500(ctx, fmt.Errorf("tmdb error tv-details [%v]: %w", traktShow.Ids(), err))
				return
			}

			shows = append(shows, h.mapper.ShowFromTmdbShow(tmdbShow))
		}

		ctx.JSON(http.StatusOK, shows)
	}
}

// Show GET /api/shows/{showId}
func (h *ShowHandler) Show() gin.HandlerFunc {
	const showIdParam UrlParam = "showId"

	return func(ctx *gin.Context) {
		showId, err := strconv.Atoi(ctx.Param(showIdParam))
		if err != nil {
			httpError400(ctx, fmt.Errorf("invalid show-id [%v]", ctx.Param(showIdParam)))
			return
		}

		tmdbShow, err := h.tmdb.GetTVDetails(showId, nil)
		if err != nil {
			httpError500(ctx, fmt.Errorf("tmdb error tv-details [%d]: %w", showId, err))
			return
		}

		ctx.JSON(http.StatusOK, h.mapper.ShowFromTmdbShow(tmdbShow))
	}
}

// SearchShows GET /api/shows/search
func (h *ShowHandler) SearchShows() gin.HandlerFunc {
	const searchTermQuery UrlQuery = "searchTerm"

	return func(ctx *gin.Context) {
		searchTerm := ctx.Param(searchTermQuery)
		if searchTerm == "" {
			httpError400(ctx, fmt.Errorf("invalid search-term [%v]", searchTerm))
			return
		}

		tmdbShows, err := h.tmdb.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
		if err != nil {
			httpError500(ctx, fmt.Errorf("tmdb error search-tv-show [%v]: %w", searchTerm, err))
			return
		}

		ctx.JSON(http.StatusOK, h.mapper.ShowsFromTmdbShowsSearch(tmdbShows, 8))
	}
}

// UpcomingReleases GET /api/shows/releases
func (h *ShowHandler) UpcomingReleases() gin.HandlerFunc {
	const pageQuery UrlQuery = "page"

	return func(ctx *gin.Context) {
		pastReleases, err := h.store.GetPastReleasesCount()
		if err != nil {
			httpError500(ctx, err)
			return
		}

		amount, offset := calculateRange(ctx.Query(pageQuery), pastReleases)

		releasesRef, err := h.store.GetReleases(amount, offset)
		if err != nil {
			httpError500(ctx, err)
			return
		}

		var releases []domain.Release
		for _, releaseRef := range releasesRef {
			tmdbRelease, err := h.tmdb.GetTVDetails(releaseRef.ShowId,
				map[string]string{"append_to_response": "translations"})
			if err != nil {
				httpError500(ctx, fmt.Errorf("tmdb error tv-details [%v]: %w", tmdbRelease.Name, err))
				return
			}

			releases = append(releases,
				h.mapper.ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate))
		}

		ctx.JSON(http.StatusOK, releases)
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
