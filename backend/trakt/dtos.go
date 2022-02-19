package trakt

import (
	"time"
)

// ShowsWatchedDto represents the relevant fields of Trakt's DTO.
// We only need the TMDb-ID, since we fetch all the details from TMDb.
type ShowsWatchedDto struct {
	Show struct {
		Ids struct {
			Tmdb int `json:"tmdb"`
		} `json:"ids"`
	} `json:"show"`
}

func (dto ShowsWatchedDto) TmdbId() int {
	return dto.Show.Ids.Tmdb
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
		} `json:"ids"`
		Genres                []string `json:"genres"`
		Network               string   `json:"network"`
		AvailableTranslations []string `json:"available_translations"`
		Overview              string   `json:"overview"`
		Rating                float32  `json:"rating"`
	} `json:"show"`
}

func (dto SeasonPremieresDto) TmdbId() int {
	return dto.Show.Ids.Tmdb
}

func (dto SeasonPremieresDto) SeasonNumber() int {
	return dto.Episode.Season
}

func (dto SeasonPremieresDto) AirDate() time.Time {
	airDate, _ := time.Parse("2006-01-02T15:04:05.000Z", dto.FirstAired)
	return airDate
}

func (dto SeasonPremieresDto) IsRelevant() bool {
	return dto.hasRelevantIds() &&
		dto.Show.Overview != "" &&
		dto.Show.Rating != 0 &&
		dto.hasRelevantAvailableTranslation() &&
		len(dto.Show.Genres) > 0 &&
		dto.Show.Network != ""
}

func (dto SeasonPremieresDto) hasRelevantIds() bool {
	ids := dto.Show.Ids
	return ids.Tmdb != 0 && ids.Tvdb != 0 && ids.Imdb != ""
}

func (dto SeasonPremieresDto) hasRelevantAvailableTranslation() bool {
	for _, availableTranslation := range dto.Show.AvailableTranslations {
		if availableTranslation == "en" {
			return true
		}
	}
	return false
}
