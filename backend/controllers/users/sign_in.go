package users

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/controllers/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type signInForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *usersController) SignUserIn(ctx echo.Context) (err error) {
	// Input
	form := new(signInForm)
	if err = ctx.Bind(form); err != nil {
		return errors.Internal(err)
	}
	if err = c.validate.Struct(form); err != nil {
		return errors.FromValidation(err)
	}

	// Use-Case
	user, err := c.usersRepository.FindByEmail(form.Email)
	if err != nil {
		return errors.FromDatabase(err, "user", errors.Params{"email": form.Email})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return errors.InvalidBody("Invalid credentials.")
	}
	if err = c.authenticateUser(ctx, user); err != nil {
		return errors.Internal(err)
	}

	// Output
	return ctx.JSON(http.StatusOK, user)
}
