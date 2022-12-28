package shows

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *showsController) GetNetworks(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	networks, err := c.networksRepository.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, networks)
}
