package trakt

import (
	"fmt"
	"time"
)

// ShowsWatchedDto represents the relevant fields of Trakt's DTO.
// We only need the TMDb-ID, since we fetch all the details from TMDb.
type ShowsWatchedDto struct {
	Show struct {
		Ids struct {
			Tmdb int    `json:"tmdb"`
			Slug string `json:"slug"`
		} `json:"ids"`
	} `json:"show"`
}

func (dto ShowsWatchedDto) TmdbId() int {
	return dto.Show.Ids.Tmdb
}

func (dto ShowsWatchedDto) SlugId() string {
	return dto.Show.Ids.Slug
}

func (dto ShowsWatchedDto) Ids() string {
	return fmt.Sprintf("%d, %v", dto.TmdbId(), dto.SlugId())
}

// SeasonPremieresDto represents the relevant fields of Trakt's DTO.
// Apart from the TMDb-ID we also need several other field to filter out irrelevant shows.
type SeasonPremieresDto struct {
	FirstAired string `json:"first_aired"`
	Episode    struct {
		Season int `json:"season"`
	} `json:"episode"`
	Show struct {
		Ids struct {
			Tmdb int    `json:"tmdb"`
			Imdb string `json:"imdb"`
			Tvdb int    `json:"tvdb"`
			Slug string `json:"slug"`
		} `json:"ids"`
	} `json:"show"`
}

func (dto SeasonPremieresDto) TmdbId() int {
	return dto.Show.Ids.Tmdb
}

func (dto SeasonPremieresDto) SlugId() string {
	return dto.Show.Ids.Slug
}

func (dto SeasonPremieresDto) SeasonNumber() int {
	return dto.Episode.Season
}

func (dto SeasonPremieresDto) AirDate() time.Time {
	airDate, _ := time.Parse("2006-01-02T15:04:05.000Z", dto.FirstAired)
	return airDate
}

func (dto SeasonPremieresDto) Ids() string {
	return fmt.Sprintf("%d, %v", dto.TmdbId(), dto.SlugId())
}

// ShowsAnticipatedDto represents the relevant fields of Trakt's DTO.
// We only need the TMDb-ID, since we fetch all the details from TMDb.
type ShowsAnticipatedDto struct {
	Show struct {
		Ids struct {
			Tmdb int `json:"tmdb"`
		} `json:"ids"`
	} `json:"show"`
}

func (dto ShowsAnticipatedDto) TmdbId() int {
	return dto.Show.Ids.Tmdb
}
