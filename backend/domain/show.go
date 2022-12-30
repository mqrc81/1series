package domain

type Show struct {
	Id            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Overview      string    `json:"overview,omitempty"`
	Year          int       `json:"year,omitempty"`
	Poster        string    `json:"poster,omitempty"`
	Rating        float32   `json:"rating,omitempty"`
	Homepage      string    `json:"homepage,omitempty"`
	SeasonsCount  int       `json:"seasonsCount"`
	EpisodesCount int       `json:"episodesCount"`
	Genres        []Genre   `json:"genres"`
	Networks      []Network `json:"networks"`
}

type TrackedShow struct {
	UserId int `db:"user_id"`
	ShowId int `db:"show_id"`
	Rating int `db:"rating"`
}
