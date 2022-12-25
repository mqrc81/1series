package domain

type Season struct {
	ShowId        int    `json:"showId,omitempty"`
	Number        int    `json:"number"`
	Name          string `json:"name,omitempty"`
	Overview      string `json:"overview,omitempty"`
	Poster        string `json:"poster,omitempty"`
	EpisodesCount int    `json:"episodesCount"`
}
