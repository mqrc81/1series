package domain

import "time"

type Genre struct {
	GenreId   int       `db:"genre_id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
