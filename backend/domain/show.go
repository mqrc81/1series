package domain

type Show struct {
	Id            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Overview      string    `json:"overview,omitempty"`
	Year          int       `json:"year,omitempty"`
	Poster        string    `json:"poster,omitempty"`
	Rating        float32   `json:"rating,omitempty"`
	Homepage      string    `json:"homepage,omitempty"`
	SeasonsCount  int       `json:"seasonsCount,omitempty"`
	EpisodesCount int       `json:"episodesCount,omitempty"`
	Genres        []Genre   `json:"genres,omitempty"`
	Networks      []Network `json:"networks,omitempty"`
}
