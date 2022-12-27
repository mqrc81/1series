package shows

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *showController) GetNetworks(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	networks, err := c.networkRepository.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, networks)
}
