package domain

type Genre struct {
	GenreId int    `db:"genre_id" json:"id,omitempty"`
	Name    string `db:"name" json:"name,omitempty"`
}
