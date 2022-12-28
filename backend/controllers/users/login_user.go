package users

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type loginForm struct {
	EmailOrUsername string `json:"emailOrUsername" validate:"required,email|alphanum"`
	Password        string `json:"password" validate:"required"`
}

func (c *usersController) LoginUser(ctx echo.Context) (err error) {
	// Input
	form := new(loginForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	user, err := c.usersRepository.FindByUsername(form.EmailOrUsername)
	if err != nil {
		user, err = c.usersRepository.FindByEmail(form.EmailOrUsername)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid email or username")
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid password")
	}

	// Output
	if err = AddUserToSession(ctx, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}
