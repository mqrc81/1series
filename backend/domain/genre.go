package domain

type Genre struct {
	Id     int    `db:"id" json:"omitempty"`
	TmdbId int    `db:"tmdb_id" json:"id"`
	Name   string `db:"name" json:"name,omitempty"`
}
