package jobs

import (
	"testing"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/trakt"
)

type TraktShow struct {
	Ids struct {
		Tmdb int    `json:"tmdb"`
		Imdb string `json:"imdb"`
		Tvdb int    `json:"tvdb"`
		Slug string `json:"slug"`
	} `json:"ids"`
}

type TraktIds struct {
	Tmdb int    `json:"tmdb"`
	Imdb string `json:"imdb"`
	Tvdb int    `json:"tvdb"`
	Slug string `json:"slug"`
}

func Test_hasRelevantIds(t *testing.T) {
	tests := []struct {
		name string
		args trakt.SeasonPremieresDto
		want bool
	}{
		{
			name: "#1",
			args: trakt.SeasonPremieresDto{
				Show: TraktShow{
					Ids: TraktIds{
						Tmdb: 123,
						Imdb: "abc",
						Tvdb: 456,
						Slug: "def",
					},
				},
			},
			want: true,
		},
		{
			name: "#2",
			args: trakt.SeasonPremieresDto{
				Show: TraktShow{
					Ids: TraktIds{
						Imdb: "abc",
						Tvdb: 456,
						Slug: "def",
					},
				},
			},
			want: false,
		},
		{
			name: "#3",
			args: trakt.SeasonPremieresDto{
				Show: TraktShow{
					Ids: TraktIds{
						Tmdb: 123,
						Tvdb: 456,
						Slug: "def",
					},
				},
			},
			want: false,
		},
		{
			name: "#4",
			args: trakt.SeasonPremieresDto{
				Show: TraktShow{
					Ids: TraktIds{
						Tmdb: 123,
						Imdb: "abc",
						Slug: "def",
					},
				},
			},
			want: false,
		},
		{
			name: "#5",
			args: trakt.SeasonPremieresDto{
				Show: TraktShow{
					Ids: TraktIds{
						Tmdb: 123,
						Imdb: "abc",
						Tvdb: 456,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRelevantIds(tt.args); got != tt.want {
				t.Errorf("hasRelevantIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TmdbGenres []struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TmdbNetworks []struct {
	Name          string `json:"name"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	OriginCountry string `json:"origin_country"`
}

type TmdbSeasons []struct {
	AirDate      string `json:"air_date"`
	EpisodeCount int    `json:"episode_count"`
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
}

type TmdbTranslations []struct {
	Iso3166_1   string `json:"iso_3166_1"`
	Iso639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name"`
	Data        struct {
		Name     string `json:"name"`
		Overview string `json:"overview"`
		Tagline  string `json:"tagline"`
		Homepage string `json:"homepage"`
	} `json:"data"`
}

func Test_hasRelevantInfo(t *testing.T) {
	tests := []struct {
		name string
		args *tmdb.TVDetails
		want bool
	}{
		{
			name: "#1",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{{}},
				Overview: "Valid",
				Seasons:  TmdbSeasons{{}},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: true,
		},
		{
			name: "#2",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{},
				Networks: TmdbNetworks{{}},
				Overview: "Valid",
				Seasons:  TmdbSeasons{{}},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: false,
		},
		{
			name: "#3",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{},
				Overview: "Valid",
				Seasons:  TmdbSeasons{{}},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: false,
		},
		{
			name: "#4",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{{}},
				Overview: "",
				Seasons:  TmdbSeasons{{}},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: false,
		},
		{
			name: "#5",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{{}},
				Overview: "Valid",
				Seasons:  TmdbSeasons{},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: false,
		},
		{
			name: "#6",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{{}},
				Overview: "Valid",
				Seasons:  TmdbSeasons{{}},
				Type:     "Porn",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           123,
					Translations: TmdbTranslations{{Iso639_1: "en"}},
				}},
			},
			want: false,
		},
		{
			name: "#7",
			args: &tmdb.TVDetails{
				Genres:   TmdbGenres{{}},
				Networks: TmdbNetworks{{}},
				Overview: "Valid",
				Seasons:  TmdbSeasons{{}},
				Type:     "Scripted",
				TVTranslationsAppend: &tmdb.TVTranslationsAppend{Translations: &tmdb.TVTranslations{
					ID:           432,
					Translations: TmdbTranslations{{Iso639_1: "fr"}},
				}},
			},
			want: false,
		},
		{
			name: "#8",
			args: &tmdb.TVDetails{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRelevantInfo(tt.args); got != tt.want {
				t.Errorf("hasRelevantInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
