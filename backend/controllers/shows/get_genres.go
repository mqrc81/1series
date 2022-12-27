package shows

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *showController) GetGenres(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	genres, err := c.genreRepository.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, genres)
}
