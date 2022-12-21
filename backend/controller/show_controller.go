package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/usecase"
)

type showController struct {
	showUseCase usecase.ShowUseCase
}

type ShowController interface {
	GetShow(ctx echo.Context) error
	GetPopularShows(ctx echo.Context) error
	SearchShows(ctx echo.Context) error
	GetUpcomingReleases(ctx echo.Context) error
}

func newShowController(uc usecase.ShowUseCase) ShowController {
	return &showController{uc}
}

func (c *showController) GetShow(ctx echo.Context) error {
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid show-id")
	}

	show, err := c.showUseCase.GetShow(showId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, show)
}

func (c *showController) GetPopularShows(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	output, err := c.showUseCase.GetPopularShows(page)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, output)
}

func (c *showController) SearchShows(ctx echo.Context) error {
	searchTerm := ctx.Param("searchTerm")
	if searchTerm == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid search-term")
	}

	shows, err := c.showUseCase.SearchShows(searchTerm)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, shows)
}

func (c *showController) GetUpcomingReleases(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	releases, err := c.showUseCase.GetUpcomingReleases(page)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, releases)
}
