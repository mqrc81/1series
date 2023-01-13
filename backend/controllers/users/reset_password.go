package users

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type resetPasswordForm struct {
	Password string `json:"password" validate:"required,min=3"`
}

func (c *usersController) ResetPassword(ctx echo.Context) (err error) {
	// Input
	tokenParam := ctx.QueryParam("token")
	form := new(resetPasswordForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	token, err := c.tokensRepository.Find(tokenParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid token")
	}
	if token.IsExpired() {
		return echo.NewHTTPError(http.StatusBadRequest, "token expired")
	}

	user, err := c.usersRepository.Find(token.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "error hashing password: "+err.Error())
	}
	user.Password = string(hashedPassword)
	if err = c.usersRepository.Update(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = c.tokensRepository.Delete(token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}
