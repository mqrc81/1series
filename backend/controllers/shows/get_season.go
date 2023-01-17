package shows

import (
	"github.com/mqrc81/zeries/controllers/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c *showsController) GetSeason(ctx echo.Context) error {
	// Input
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return errors.MissingParameter("show-id")
	}
	seasonNumber, err := strconv.Atoi(ctx.Param("seasonNumber"))
	if err != nil {
		return errors.MissingParameter("season-number")
	}

	// Use-Case
	tmdbSeason, err := c.tmdbClient.GetTVSeasonDetails(showId, seasonNumber, nil)
	if err != nil {
		return errors.FromTmdb(err, "season", errors.Params{"show-id": strconv.Itoa(showId), "season-number": strconv.Itoa(seasonNumber)})
	}

	// Output
	return ctx.JSON(http.StatusOK, SeasonFromTmdbSeason(tmdbSeason, showId))
}
