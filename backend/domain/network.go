package domain

type Network struct {
	NetworkId int    `db:"network_id" json:"id,omitempty"`
	Name      string `db:"name" json:"name,omitempty"`
	Logo      string `db:"logo" json:"logo,omitempty"`
}
