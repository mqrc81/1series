package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/session"
	"github.com/mqrc81/zeries/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/usecase"
)

type userController struct {
	userUseCase usecase.UserUseCase
	validate    *validator.Validate
}

type UserController interface {
	RegisterUser(ctx echo.Context) error
	LoginUser(ctx echo.Context) error
}

func newUserController(userUseCase usecase.UserUseCase, validate *validator.Validate) UserController {
	return &userController{userUseCase, validate}
}

func (c *userController) RegisterUser(ctx echo.Context) (err error) {
	// Input
	form := new(usecase.RegisterForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	user, err := c.userUseCase.RegisterUser(*form)
	if err != nil {
		return err
	}

	// Output
	if err = c.addUserToSession(ctx, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *userController) LoginUser(ctx echo.Context) (err error) {
	// Input
	form := new(usecase.LoginForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	user, err := c.userUseCase.LoginUser(*form)
	if err != nil {
		return err
	}

	// Output
	if err = c.addUserToSession(ctx, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *userController) addUserToSession(ctx echo.Context, user domain.User) error {
	currentSession, err := session.Get(sessionKey, ctx)
	if err == nil {
		currentSession.Values[sessionUserIdKey] = user.Id
		err = currentSession.Save(ctx.Request(), ctx.Response())
	}
	return err
}
