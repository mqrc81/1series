package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/session"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/usecases/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/usecases"
)

const (
	imdbWatchlistExportFileName = "WATCHLIST.csv"
)

type userController struct {
	userUseCase users.UseCase
	validate    *validator.Validate
}

type UserController interface {
	RegisterUser(ctx echo.Context) error
	LoginUser(ctx echo.Context) error
	ImportImdbWatchlist(ctx echo.Context) error
}

func newUserController(userUseCase users.UseCase, validate *validator.Validate) UserController {
	return &userController{userUseCase, validate}
}

func (c *userController) RegisterUser(ctx echo.Context) (err error) {
	// Input
	form := new(usecases.RegisterForm)
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
	form := new(usecases.LoginForm)
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

func (c *userController) ImportImdbWatchlist(ctx echo.Context) (err error) {
	// Input
	formFile, err := ctx.FormFile(imdbWatchlistExportFileName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid imdb watchlist export file")
	}
	file, err := formFile.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "unable to open the imdb watchlist export file")
	}
	defer file.Close()

	// Use-Case
	if err = c.userUseCase.ImportImdbWatchlist(file); err != nil {
		return err
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}

func (c *userController) addUserToSession(ctx echo.Context, user domain.User) error {
	currentSession, err := session.Get(sessionKey, ctx)
	if err == nil {
		currentSession.Values[sessionUserIdKey] = user.Id
		err = currentSession.Save(ctx.Request(), ctx.Response())
	}
	return err
}
