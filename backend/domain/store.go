package domain

type Store interface {
	UserStore
	GenreStore
	NetworkStore
	ReleaseStore
}

type UserStore interface {
	GetUser(userId int) (User, error)
	GetUserByUsername(username string) (User, error)
	GetUserByEmail(email string) (User, error)
}

type GenreStore interface {
	GetGenres() ([]Genre, error)
	AddGenre(genre Genre) error
}

type NetworkStore interface {
	GetNetworks() ([]Network, error)
	AddNetwork(network Network) error
}

type ReleaseStore interface {
	GetReleases(amount int, offset int) ([]ReleaseRef, error)
	SaveReleases(releases []ReleaseRef, pastReleasesCount int) error
	GetPastReleasesCount() (int, error)
}
