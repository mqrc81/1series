package shows

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"net/http"
	"strconv"
)

type PopularShowsDto struct {
	nextPage int
	shows    []domain.Show
}

//goland:noinspection GoPreferNilSlice
func (c *showsController) GetPopularShows(ctx echo.Context) error {
	// Input
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	// Use-Case
	traktShows, err := c.traktClient.GetShowsWatchedWeekly(page, popularShowsPerRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error fetching trakt shows watched weekly: "+err.Error())
	}

	shows := []domain.Show{}
	for _, traktShow := range traktShows {
		if traktShow.TmdbId() == 0 {
			logger.Warning("tmdb show for trakt show %v not found", traktShow.Show.Ids.Trakt)
			continue
		}
		tmdbShow, err := c.tmdbClient.GetTVDetails(traktShow.TmdbId(), nil)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error fetching tmdb show: "+err.Error())
		}

		shows = append(shows, ShowFromTmdbShow(tmdbShow))
	}
	if page >= 25 {
		page = -1
	}

	// Output
	return ctx.JSON(http.StatusOK, popularShowsDto{
		NextPage: page + 1,
		Shows:    shows,
	})
}
