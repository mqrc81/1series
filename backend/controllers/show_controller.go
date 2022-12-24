package controllers

import (
	"github.com/mqrc81/zeries/usecases/shows"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type showController struct {
	showUseCase shows.UseCase
}

type ShowController interface {
	GetShow(ctx echo.Context) error
	GetPopularShows(ctx echo.Context) error
	GetUpcomingReleases(ctx echo.Context) error
	SearchShows(ctx echo.Context) error
	GetGenres(ctx echo.Context) error
	GetNetworks(ctx echo.Context) error
}

func newShowController(uc shows.UseCase) ShowController {
	return &showController{uc}
}

func (c *showController) GetShow(ctx echo.Context) error {
	// Input
	showId, err := strconv.Atoi(ctx.Param("showId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid show-id")
	}

	// Use-Case
	show, err := c.showUseCase.GetShow(showId)
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, show)
}

func (c *showController) GetPopularShows(ctx echo.Context) error {
	// Input
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	// Use-Case
	shows, err := c.showUseCase.GetPopularShows(page)
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, popularShowsDto{
		NextPage: page + 1,
		Shows:    shows,
	})
}

func (c *showController) GetUpcomingReleases(ctx echo.Context) error {
	// Input
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	// Use-Case
	releases, err := c.showUseCase.GetUpcomingReleases(page)
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, upcomingReleasesDto{
		PreviousPage: page - 1,
		NextPage:     page + 1,
		Releases:     releases,
	})
}

func (c *showController) SearchShows(ctx echo.Context) error {
	// Input
	searchTerm := ctx.QueryParam("searchTerm")
	if searchTerm == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid search-term")
	}

	// Use-Case
	shows, err := c.showUseCase.SearchShows(searchTerm)
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, shows)
}

func (c *showController) GetGenres(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	genres, err := c.showUseCase.GetGenres()
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, genres)
}

func (c *showController) GetNetworks(ctx echo.Context) error {
	// Input
	// -

	// Use-Case
	networks, err := c.showUseCase.GetNetworks()
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, networks)
}
