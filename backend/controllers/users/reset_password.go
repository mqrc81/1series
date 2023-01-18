package users

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/controllers/errors"
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
		return errors.Internal(err)
	}
	if err = c.validate.Struct(form); err != nil {
		return errors.FromValidation(err)
	}

	// Use-Case
	token, err := c.tokensRepository.Find(tokenParam)
	if err != nil {
		return errors.FromDatabase(err, "token", nil)
	}
	if token.IsExpired() {
		return errors.InvalidParam("The link has expired. Please request a new password-reset.")
	}

	user, err := c.usersRepository.Find(token.UserId)
	if err != nil {
		return errors.FromDatabase(err, "user", nil)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Internal(fmt.Errorf("error hashing password: %w", err))
	}
	user.Password = string(hashedPassword)
	if err = c.usersRepository.Update(user); err != nil {
		return errors.FromDatabase(err, "user", nil)
	}

	if err = c.tokensRepository.Delete(token); err != nil {
		return errors.FromDatabase(err, "token", nil)
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}
