package api

import (
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/util"
)

const (
	tmdbImageUrl = "https://image.tmdb.org/t/p/original"
)

func showFromTmdbShow(dto *tmdb.TVDetails) (show domain.Show) {

	var genres []domain.Genre
	for _, genre := range dto.Genres {
		genres = append(genres, domain.Genre{
			Id:   int(genre.ID),
			Name: genre.Name,
		})
	}

	var networks []domain.Network
	for _, network := range dto.Networks {
		networks = append(networks, domain.Network{
			Id:   int(network.ID),
			Name: network.Name,
			Logo: tmdbImageUrl + network.LogoPath,
		})
	}

	airDate, err := time.Parse("2006-01-02", dto.FirstAirDate)
	if err != nil {
		airDate, _ = time.Parse("2006-01-02", dto.Seasons[0].AirDate)
	}

	return domain.Show{
		Id:            int(dto.ID),
		Name:          dto.Name,
		Overview:      dto.Overview,
		Year:          airDate.Year(),
		Poster:        tmdbImageUrl + dto.PosterPath,
		Rating:        dto.VoteAverage,
		Genres:        genres,
		Networks:      networks,
		Homepage:      dto.Homepage,
		SeasonsCount:  dto.NumberOfSeasons,
		EpisodesCount: dto.NumberOfEpisodes,
	}
}

func showsFromTmdbShowsSearch(dto *tmdb.SearchTVShows, maxResults int) (shows []domain.Show) {
	for _, result := range dto.Results[:maxResults] {
		shows = append(shows, domain.Show{
			Id:     int(result.ID),
			Name:   result.Name,
			Poster: tmdbImageUrl + result.PosterPath,
			Rating: result.VoteAverage,
		})
	}
	return shows
}

func releaseFromTmdbShow(dto *tmdb.TVDetails, seasonNumber int, airDate time.Time) domain.Release {
	return domain.Release{
		Show:    showFromTmdbShow(dto),
		Season:  seasonFromTmdbShow(dto, seasonNumber),
		AirDate: airDate,
	}
}

func seasonFromTmdbShow(dto *tmdb.TVDetails, seasonNumber int) domain.Season {
	if seasonNumber > len(dto.Seasons) {
		util.LogError("Tmdb show [%d, %v] has no season [%v]", dto.ID, dto.Name, seasonNumber)
		return domain.Season{}
	}
	season := dto.Seasons[seasonNumber-1]
	return domain.Season{
		ShowId:        int(dto.ID),
		Number:        season.SeasonNumber,
		Name:          season.Name,
		Overview:      season.Overview,
		Poster:        tmdbImageUrl + season.PosterPath,
		EpisodesCount: season.EpisodeCount,
	}
}
