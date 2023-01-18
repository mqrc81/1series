package users

import (
	"fmt"
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/1series/domain"
	"golang.org/x/crypto/bcrypt"
)

type signUpForm struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=16"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

func (c *usersController) SignUserUp(ctx echo.Context) (err error) {
	// Input
	form := new(signUpForm)
	if err = ctx.Bind(form); err != nil {
		return errors.Internal(err)
	}
	if err = c.validate.Struct(form); err != nil {
		return errors.FromValidation(err)
	}

	// Use-Case
	if _, err = c.usersRepository.FindByUsername(form.Username); err == nil {
		return errors.InvalidParam("Username is already taken.")
	}
	if _, err = c.usersRepository.FindByEmail(form.Email); err == nil {
		return errors.InvalidParam("Email is already taken.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Internal(fmt.Errorf("error hashing password: %w", err))
	}

	if err = c.usersRepository.Save(domain.User{
		Username: form.Username,
		Email:    form.Email,
		Password: string(hashedPassword),
		NotificationOptions: domain.NotificationOptions{
			Releases:        true,
			Recommendations: true,
		},
	}); err != nil {
		return errors.FromDatabase(err, "user", nil)
	}

	user, err := c.usersRepository.FindByUsername(form.Username)
	if err != nil {
		return errors.Internal(err)
	}
	if err = c.authenticateUser(ctx, user); err != nil {
		return errors.Internal(err)
	}

	// Output
	return ctx.JSON(http.StatusOK, user)
}
