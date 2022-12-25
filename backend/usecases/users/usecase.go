package users

import (
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/usecases"
)

type UseCase interface {
	RegisterUser(form usecases.RegisterForm) (domain.User, error)
	LoginUser(form usecases.LoginForm) (domain.User, error)
}

type useCase struct {
	userRepository repositories.UserRepository
}

func NewUseCase(
	userRepository repositories.UserRepository,
) UseCase {
	return &useCase{
		userRepository,
	}
}