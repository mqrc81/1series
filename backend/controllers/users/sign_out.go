package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *usersController) SignUserOut(ctx echo.Context) (err error) {
	// Input
	// -

	// Use-Case
	if err = c.unauthenticateUser(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}
