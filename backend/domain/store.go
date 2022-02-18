package domain

type Store interface {
	UserStore
	GenreStore
	NetworkStore
}

type UserStore interface {
	GetUser(userId int) (User, error)
}

type GenreStore interface {
	GetGenres() ([]Genre, error)
	// AddGenre(genre Genre) error
}

type NetworkStore interface {
	GetNetworks() ([]Network, error)
	// AddNetwork(network Network) error
}
