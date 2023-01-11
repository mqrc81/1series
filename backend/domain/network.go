package domain

import "time"

type Network struct {
	NetworkId int       `db:"network_id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name,omitempty"`
	Logo      string    `db:"logo" json:"logo,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
