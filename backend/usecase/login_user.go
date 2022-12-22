package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (uc *userUseCase) LoginUser(form LoginForm) (user domain.User, err error) {

	user, err = uc.userRepository.FindByUsername(form.EmailOrUsername)
	if err != nil {
		user, err = uc.userRepository.FindByEmail(form.EmailOrUsername)
		if err != nil {
			return user, echo.NewHTTPError(http.StatusBadRequest, "invalid email or username")
		}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		return user, echo.NewHTTPError(http.StatusBadRequest, "invalid password")
	}

	return user, err
}
