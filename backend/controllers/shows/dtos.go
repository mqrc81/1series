package shows

import (
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
)

const (
	tmdbVoteAverageInflationMultiplier = 1.03
)

type (
	popularShowsDto struct {
		NextPage int           `json:"nextPage,omitempty"`
		Shows    []domain.Show `json:"shows"`
	}

	upcomingReleasesDto struct {
		PreviousPage int              `json:"previousPage,omitempty"`
		NextPage     int              `json:"nextPage,omitempty"`
		Releases     []domain.Release `json:"releases"`
	}
)

//goland:noinspection GoPreferNilSlice
func ShowFromTmdbShow(dto *tmdb.TVDetails) (show domain.Show) {

	genres := []domain.Genre{}
	for _, genre := range dto.Genres {
		genres = append(genres, domain.Genre{
			GenreId: int(genre.ID),
			Name:    genre.Name,
		})
	}

	networks := []domain.Network{}
	for _, network := range dto.Networks {
		networks = append(networks, domain.Network{
			NetworkId: int(network.ID),
			Name:      network.Name,
			Logo:      tmdbImageUrlFromImagePath(network.LogoPath),
		})
	}

	airDate, err := time.Parse("2006-01-02", dto.FirstAirDate)
	if err != nil {
		airDate, _ = time.Parse("2006-01-02", dto.Seasons[0].AirDate)
	}

	return domain.Show{
		Id:            int(dto.ID),
		Name:          dto.Name,
		OriginalName:  dto.OriginalName,
		Overview:      dto.Overview,
		Year:          airDate.Year(),
		Poster:        tmdbImageUrlFromImagePath(dto.PosterPath),
		Backdrop:      tmdbImageUrlFromImagePath(dto.BackdropPath),
		Rating:        ratingFromTmdbVoteAverage(dto.VoteAverage),
		Genres:        genres,
		Networks:      networks,
		Homepage:      dto.Homepage,
		SeasonsCount:  dto.NumberOfSeasons,
		EpisodesCount: dto.NumberOfEpisodes,
	}
}

func ReleaseFromTmdbShow(dto *tmdb.TVDetails, seasonNumber int, airDate time.Time, anticipationLevel domain.AnticipationLevel) domain.Release {
	return domain.Release{
		Show:              ShowFromTmdbShow(dto),
		Season:            SeasonFromTmdbShow(dto, seasonNumber),
		AirDate:           airDate,
		AnticipationLevel: anticipationLevel,
	}
}

func SeasonFromTmdbShow(dto *tmdb.TVDetails, seasonNumber int) domain.Season {
	if seasonNumber > len(dto.Seasons) {
		logger.Error("Tmdb show [%d, %v] has no season [%v]", dto.ID, dto.Name, seasonNumber)
		return domain.Season{}
	}
	season := dto.Seasons[seasonNumber-1]
	return domain.Season{
		ShowId:        int(dto.ID),
		Number:        season.SeasonNumber,
		Name:          season.Name,
		Overview:      season.Overview,
		Poster:        tmdbImageUrlFromImagePath(season.PosterPath),
		EpisodesCount: season.EpisodeCount,
	}
}

func SeasonFromTmdbSeason(dto *tmdb.TVSeasonDetails, showId int) domain.Season {
	return domain.Season{
		ShowId:        showId,
		Number:        dto.SeasonNumber,
		Name:          dto.Name,
		Overview:      dto.Overview,
		Poster:        tmdbImageUrlFromImagePath(dto.PosterPath),
		EpisodesCount: len(dto.Episodes),
	}
}

//goland:noinspection GoNameStartsWithPackageName,GoPreferNilSlice
func ShowsFromTmdbShowsSearch(dto *tmdb.SearchTVShows, maxResults int) []domain.Show {
	shows := []domain.Show{}
	if maxResults > len(dto.Results) {
		maxResults = len(dto.Results)
	}
	for _, result := range dto.Results[:maxResults] {
		shows = append(shows, domain.Show{
			Id:           int(result.ID),
			Name:         result.Name,
			OriginalName: result.OriginalName,
			Poster:       tmdbImageUrlFromImagePath(result.PosterPath),
			Backdrop:     tmdbImageUrlFromImagePath(result.BackdropPath),
			Rating:       ratingFromTmdbVoteAverage(result.VoteAverage),
		})
	}
	return shows
}

func tmdbImageUrlFromImagePath(imagePath string) string {
	if len(imagePath) > 0 {
		return tmdbImageBaseUrl + imagePath
	}
	return ""
}

func ratingFromTmdbVoteAverage(voteAverage float32) float32 {
	// ratings are
	if rating := voteAverage * tmdbVoteAverageInflationMultiplier; rating > 10 {
		return 10
	} else {
		return rating
	}
}
