package usecase

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
)

const (
	tmdbImageUrl       = "https://image.tmdb.org/t/p/original"
	releasesPerRequest = 20
)

type showUseCase struct {
	userRepository    repository.UserRepository
	releaseRepository repository.ReleaseRepository
	traktClient       *trakt.Client
	tmdbClient        *tmdb.Client
}

type userUseCase struct {
	userRepository repository.UserRepository
}

type ShowUseCase interface {
	GetShow(showId int) (domain.Show, error)
	GetPopularShows(page int) ([]domain.Show, error)
	GetUpcomingReleases(page int) ([]domain.Release, error)
	SearchShows(searchTerm string) ([]domain.Show, error)
}

type UserUseCase interface {
	RegisterUser(form RegisterForm) (domain.User, error)
	LoginUser(form LoginForm) (domain.User, error)
}

func NewShowUseCase(
	userRepository repository.UserRepository, releaseRepository repository.ReleaseRepository,
	traktClient *trakt.Client, tmdbClient *tmdb.Client,
) ShowUseCase {
	return &showUseCase{userRepository, releaseRepository, traktClient, tmdbClient}
}

func NewUserUseCase(
	userRepository repository.UserRepository,
) UserUseCase {
	return &userUseCase{userRepository}
}
