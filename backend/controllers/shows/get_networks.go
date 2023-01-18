package shows

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"
)

func (c *showsController) GetNetworks(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	networks, err := c.networksRepository.FindAll()
	if err != nil {
		return errors.FromDatabase(err, "genres", nil)
	}

	// Output
	return ctx.JSON(http.StatusOK, networks)
}
