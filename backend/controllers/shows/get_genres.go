package shows

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"
)

func (c *showsController) GetGenres(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	genres, err := c.genresRepository.FindAll()
	if err != nil {
		return errors.FromDatabase(err, "genres", nil)
	}

	// Output
	return ctx.JSON(http.StatusOK, genres)
}
