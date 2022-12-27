package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"golang.org/x/crypto/bcrypt"
)

type registerForm struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=16"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

func (c *userController) RegisterUser(ctx echo.Context) (err error) {
	// Input
	form := new(registerForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	if _, err = c.userRepository.FindByUsername(form.Username); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "username is already taken")
	}
	if _, err = c.userRepository.FindByEmail(form.Email); err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "email is already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "error hashing password: "+err.Error())
	}

	userId, err := c.userRepository.Save(domain.User{
		Username: form.Username,
		Email:    form.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	user, err := c.userRepository.Find(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	// Output
	if err = AddUserToSession(ctx, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}
