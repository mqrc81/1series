package shows

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *showsController) SearchShows(ctx echo.Context) error {
	// Input
	searchTerm := ctx.QueryParam("searchTerm")
	if searchTerm == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid search-term")
	}

	// Use-Case
	tmdbShows, err := c.tmdbClient.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error searching tmdb shows: "+err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, ShowsFromTmdbShowsSearch(tmdbShows, showSearchesPerRequest))
}
