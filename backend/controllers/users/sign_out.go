package users

import (
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *usersController) SignUserOut(ctx echo.Context) (err error) {
	// Input
	// -

	// Use-Case
	if err = c.unauthenticateUser(ctx); err != nil {
		return errors.Internal(err)
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}
