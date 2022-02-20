package domain

import (
	"time"
)

type Store interface {
	UserStore
	GenreStore
	NetworkStore
	ReleaseStore
}

type UserStore interface {
	GetUser(userId int) (User, error)
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
	SaveRelease(release ReleaseRef, expiry time.Time) error
	ClearExpiredReleases(now time.Time) error
	SetPastReleasesCount(amount int) error
	GetPastReleasesCount() (int, error)
}
