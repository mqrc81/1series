package users

import (
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	tvSeriesTitleType     = "tvSeries"
	tvMiniSeriesTitleType = "tvMiniSeries"
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

func (uc *useCase) ImportImdbWatchlist(file multipart.File) (err error) {
	var exportedImdbWatchlist []*exportedImdbWatchlistRow
	if err = gocsv.Unmarshal(file, &exportedImdbWatchlist); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse imdb watchlist file: "+err.Error())
	}

	for _, row := range exportedImdbWatchlist {
		if row.TitleType != tvSeriesTitleType && row.TitleType != tvMiniSeriesTitleType {
			continue
		}
		results, err := uc.tmdbClient.GetFindByID(row.Const, map[string]string{"external_source": "imdb_id"})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to find tmdb show by imdb id: "+err.Error())
		}
		if len(results.TvResults) > 1 {
			logger.Warning("Multiple tmdb shows found for imdb id %v", row.Const)
		} else if len(results.TvResults) < 1 {
			// TODO return failed imports
		} else {
			if err = uc.trackedShowRepository.Save(domain.TrackedShow{
				UserId: 0,
				ShowId: int(results.TvResults[0].ID),
				Rating: row.YourRating,
			}); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
	}

	return err
}
