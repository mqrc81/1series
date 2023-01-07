package shows

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c *showsController) GetSeason(ctx echo.Context) error {
	// Input
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid show-id")
	}
	seasonNumber, err := strconv.Atoi(ctx.Param("seasonNumber"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid season-number")
	}

	// Use-Case
	tmdbSeason, err := c.tmdbClient.GetTVSeasonDetails(showId, seasonNumber, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, SeasonFromTmdbSeason(tmdbSeason, showId))
}
