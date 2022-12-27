package shows

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c *showController) GetShow(ctx echo.Context) error {
	// Input
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid show-id")
	}

	// Use-Case
	tmdbShow, err := c.tmdbClient.GetTVDetails(showId, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, ShowFromTmdbShow(tmdbShow))
}
