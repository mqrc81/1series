package usecase

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (uc *userUseCase) LoginUser(form LoginForm, reqCtx context.Context) (domain.User, error) {

	user, err := uc.userRepository.FindByUsername(form.Identifier)
	if err != nil {

		user, err = uc.userRepository.FindByEmail(form.Identifier)
		if err != nil {
			form.IdentifierNotFound = true
		}
	}

	if formErrors := form.Validate(); formErrors != nil {
		return domain.User{}, echo.NewHTTPError(http.StatusBadRequest, formErrors)
	}

	uc.sessionManager.Put(reqCtx, "userId", user.Id)

	return user, nil
}
