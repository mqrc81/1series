package usecase

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uc *userUseCase) RegisterUser(form RegisterForm, reqCtx context.Context) (domain.User, error) {

	if _, err := uc.userRepository.FindByUsername(form.Username); err == nil {
		form.UsernameTaken = true
	}
	if _, err := uc.userRepository.FindByEmail(form.Email); err == nil {
		form.EmailTaken = true
	}

	if formErrors := form.Validate(); formErrors != nil {
		return domain.User{}, echo.NewHTTPError(http.StatusBadRequest, formErrors)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, echo.NewHTTPError(http.StatusConflict, "error hashing password: "+err.Error())
	}

	userId, err := uc.userRepository.Save(domain.User{
		Username: form.Username,
		Email:    form.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return domain.User{}, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	user, err := uc.userRepository.Find(userId)
	if err != nil {
		return domain.User{}, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	uc.sessionManager.Put(reqCtx, "userId", userId)

	return user, nil
}
