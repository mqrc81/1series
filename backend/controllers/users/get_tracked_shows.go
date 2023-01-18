package users

import (
	"github.com/mqrc81/1series/controllers/errors"
	dtos "github.com/mqrc81/1series/controllers/shows"
	"github.com/mqrc81/1series/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

//goland:noinspection GoPreferNilSlice
func (c *usersController) GetTrackedShows(ctx echo.Context) (err error) {
	// Input
	// -

	// Use-Case
	user, err := GetAuthenticatedUser(ctx)
	if err != nil {
		return errors.Unauthorized()
	}

	trackedShows, err := c.trackedShowsRepository.FindAllByUser(user)
	if err != nil {
		return errors.FromDatabase(err, "tracked shows", nil)
	}

	shows := []domain.Show{}
	for _, trackedShow := range trackedShows {
		tmdbShow, err := c.tmdbClient.GetTVDetails(trackedShow.ShowId, nil)
		if err != nil {
			return errors.FromTmdb(err, "show", errors.Params{"id": trackedShow.ShowId})
		}

		shows = append(shows, dtos.ShowFromTmdbShow(tmdbShow))
	}

	// Output
	return ctx.JSON(http.StatusOK, shows)
}
