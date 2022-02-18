package api

import (
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
)

const (
	tmdbImageUrl = "https://image.tmdb.org/t/p/original"
)

func showFromDto(show *tmdb.TVDetails) domain.Show {

	var genres []domain.Genre
	for _, genre := range show.Genres {
		genres = append(genres, domain.Genre{
			Id:   int(genre.ID),
			Name: genre.Name,
		})
	}

	var networks []domain.Network
	for _, network := range show.Networks {
		networks = append(networks, domain.Network{
			Id:   int(network.ID),
			Name: network.Name,
			Logo: tmdbImageUrl + network.LogoPath,
		})
	}

	airDate, err := time.Parse("2006-01-02", show.FirstAirDate)
	if err != nil {
		airDate, _ = time.Parse("2006-01-02", show.Seasons[0].AirDate)
	}

	return domain.Show{
		Id:            int(show.ID),
		Name:          show.Name,
		Description:   show.Overview,
		Year:          airDate.Year(),
		Poster:        tmdbImageUrl + show.PosterPath,
		Rating:        show.VoteAverage,
		Genres:        genres,
		Networks:      networks,
		Homepage:      show.Homepage,
		SeasonsCount:  show.NumberOfSeasons,
		EpisodesCount: show.NumberOfEpisodes,
	}
}
