package domain

type Genre struct {
	Id     int    `db:"id"`
	TmdbId int    `db:"tmdb_id"`
	Name   string `db:"name"`
}
