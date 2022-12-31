package domain

type Genre struct {
	Id     int    `db:"id" json:"id,omitempty"`
	TmdbId int    `db:"tmdb_id" json:",omitempty"`
	Name   string `db:"name" json:"name,omitempty"`
}
