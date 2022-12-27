package users

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/usecases"
	"mime/multipart"
)

type UseCase interface {
	RegisterUser(form usecases.RegisterForm) (domain.User, error)
	LoginUser(form usecases.LoginForm) (domain.User, error)
	ImportImdbWatchlist(file multipart.File) error
}

type useCase struct {
	userRepository        repositories.UserRepository
	trackedShowRepository repositories.TrackedShowRepository
	tmdbClient            *tmdb.Client
	emailClient           *email.Client
}

func NewUseCase(
	userRepository repositories.UserRepository,
	trackedShowRepository repositories.TrackedShowRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
) UseCase {
	return &useCase{
		userRepository,
		trackedShowRepository,
		tmdbClient,
		emailClient,
	}
}
