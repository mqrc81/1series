package users

import (
	"github.com/mqrc81/zeries/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uc *useCase) RegisterUser(form usecases.RegisterForm) (user domain.User, err error) {

	if _, err = uc.userRepository.FindByUsername(form.Username); err == nil {
		return user, echo.NewHTTPError(http.StatusBadRequest, "username is already taken")
	}
	if _, err = uc.userRepository.FindByEmail(form.Email); err == nil {
		return user, echo.NewHTTPError(http.StatusBadRequest, "email is already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, echo.NewHTTPError(http.StatusConflict, "error hashing password: "+err.Error())
	}

	userId, err := uc.userRepository.Save(domain.User{
		Username: form.Username,
		Email:    form.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return user, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	user, err = uc.userRepository.Find(userId)
	if err != nil {
		return user, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return user, nil
}
