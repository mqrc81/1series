package domain

type Network struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Logo string `json:"logo,omitempty"`
}
