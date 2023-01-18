package shows

import (
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (c *showsController) GetShow(ctx echo.Context) error {
	// Input
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return errors.MissingParameter("show-id")
	}

	// Use-Case
	tmdbShow, err := c.tmdbClient.GetTVDetails(showId, nil)
	if err != nil {
		return errors.FromTmdb(err, "show", errors.Params{"id": showId})
	}

	// Output
	return ctx.JSON(http.StatusOK, ShowFromTmdbShow(tmdbShow))
}
