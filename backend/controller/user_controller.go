package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/usecase"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

type UserController interface {
	RegisterUser(ctx echo.Context) error
}

func newUserController(uc usecase.UserUseCase) UserController {
	return &userController{uc}
}

func (c *userController) RegisterUser(ctx echo.Context) error {
	form := new(usecase.RegisterForm)
	if err := ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	user, err := c.userUseCase.RegisterUser(*form, context.Background())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}
