package domain

type Season struct {
	ShowId        int    `json:"showId"`
	Number        int    `json:"number"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	Poster        string `json:"poster"`
	EpisodesCount int    `json:"episodesCount"`
}
