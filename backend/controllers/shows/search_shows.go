package shows

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/1series/controllers/errors"
	"net/http"
)

func (c *showsController) SearchShows(ctx echo.Context) error {
	// Input
	searchTerm := ctx.QueryParam("searchTerm")
	if searchTerm == "" {
		return errors.MissingParameter("search-term")
	}

	// Use-Case
	tmdbShows, err := c.tmdbClient.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
	if err != nil {
		return errors.Internal(fmt.Errorf("error searching tmdb shows: %w", err))
	}

	// Output
	return ctx.JSON(http.StatusOK, ShowsFromTmdbShowsSearch(tmdbShows, showSearchesPerRequest))
}
