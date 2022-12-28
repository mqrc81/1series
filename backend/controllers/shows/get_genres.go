package shows

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *showsController) GetGenres(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	genres, err := c.genresRepository.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, genres)
}
