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

func (c *usersController) ImportImdbWatchlist(ctx echo.Context) (err error) {
	// Input
	user, err := GetUserFromSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "no user is logged in")
	}
	// formFile, err := ctx.FormFile(imdbWatchlistImportFileName)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "invalid imdb watchlist export file: "+err.Error())
	// }
	// file, err := formFile.Open()
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "unable to open the imdb watchlist export file")
	// }
	// defer file.Close()

	var exportedImdbWatchlist []*exportedImdbWatchlistRow
	var textContent struct {
		Content string `json:"content"`
	}
	if err = ctx.Bind(&textContent); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "unable to parse content: "+err.Error())
	}
	if err = gocsv.UnmarshalString(textContent.Content, &exportedImdbWatchlist); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse imdb watchlist file: "+err.Error())
	}

	// Use-Case
	var failedImports []failedImdbWatchlistImports
	for _, row := range exportedImdbWatchlist {
		if !(row.TitleType == imdbTitleTypeTvSeries || row.TitleType == imdbTitleTypeTvMiniSeries) {
			continue
		}
		results, err := c.tmdbClient.GetFindByID(row.Const, map[string]string{"external_source": "imdb_id"})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to find tmdb show by imdb id: "+err.Error())
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
