package users

import (
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"net/http"
	"time"
)

type exportedImdbWatchlistRow struct {
	Position    int       `csv:"-"`
	Const       string    `csv:"Const"`
	Created     time.Time `csv:"-"`
	Modified    time.Time `csv:"-"`
	Description string    `csv:"-"`
	Title       string    `csv:"-"`
	Url         string    `csv:"-"`
	TitleType   string    `csv:"Title Type"`
	ImdbRating  float32   `csv:"-"`
	Runtime     int       `csv:"-"`
	Year        int       `csv:"-"`
	Genres      []string  `csv:"-"`
	NumVotes    int       `csv:"-"`
	ReleaseDate time.Time `csv:"-"`
	Directors   time.Time `csv:"-"`
	YourRating  int       `csv:"Your Rating"`
	DateRated   time.Time `csv:"-"`
}

type failedImdbWatchlistImports struct {
	ImdbId string
	Title  string
	Reason string
}

//goland:noinspection GoPreferNilSlice
func (c *usersController) ImportImdbWatchlist(ctx echo.Context) (err error) {
	// Input
	user, err := GetUserFromSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "no user is logged in")
	}

	var exportedImdbWatchlist []*exportedImdbWatchlistRow
	reader := gocsv.DefaultCSVReader(ctx.Request().Body)
	if err = gocsv.UnmarshalCSV(reader, &exportedImdbWatchlist); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse imdb watchlist file: "+err.Error())
	}

	// Use-Case
	failedImports := []failedImdbWatchlistImports{}
	for _, row := range exportedImdbWatchlist {
		if !(row.TitleType == imdbTitleTypeTvSeries || row.TitleType == imdbTitleTypeTvMiniSeries) {
			continue
		}
		results, err := c.tmdbClient.GetFindByID(row.Const, map[string]string{"external_source": "imdb_id"})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to find tmdb show by imdb id")
		}
		if len(results.TvResults) == 1 {
			if err = c.trackedShowsRepository.Save(domain.TrackedShow{
				UserId: user.Id,
				ShowId: int(results.TvResults[0].ID),
				Rating: row.YourRating,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		} else if len(results.TvResults) < 1 {
			logger.Warning("No tmdb shows found for imdb id %v", row.Const)
			failedImports = append(failedImports, failedImdbWatchlistImports{
				ImdbId: row.Const,
				Title:  row.Title,
				Reason: "No series found in the TMDb-database for that IMDb-ID",
			})
		} else {
			logger.Warning("Multiple tmdb shows found for imdb id %v", row.Const)
			failedImports = append(failedImports, failedImdbWatchlistImports{
				ImdbId: row.Const,
				Title:  row.Title,
				Reason: "Multiple series found in the TMDb-database for that IMDb-ID",
			})
		}
	}

	// Output
	return ctx.JSON(http.StatusOK, failedImports)
}
