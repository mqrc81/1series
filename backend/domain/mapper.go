package domain

import (
	"time"

	"github.com/cyruzin/golang-tmdb"
)

const (
	tmdbImageUrl = "https://image.tmdb.org/t/p/original"
)

func ShowFromDto(show *tmdb.TVDetails) Show {

	var genres []Genre
	for _, genre := range show.Genres {
		genres = append(genres, Genre{
			Id:   int(genre.ID),
			Name: genre.Name,
		})
	}

	var networks []Network
	for _, network := range show.Networks {
		networks = append(networks, Network{
			Id:   int(network.ID),
			Name: network.Name,
			Logo: tmdbImageUrl + network.LogoPath,
		})
	}

	airDate, err := time.Parse("2006-01-02", show.FirstAirDate)
	if err != nil {
		airDate, _ = time.Parse("2006-01-02", show.Seasons[0].AirDate)
	}

	return Show{
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
