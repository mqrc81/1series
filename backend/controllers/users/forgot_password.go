package users

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"net/http"
)

type forgotPasswordForm struct {
	Email string `json:"email" validate:"required,email"`
}

func (c *usersController) ForgotPassword(ctx echo.Context) (err error) {
	// Input
	form := new(forgotPasswordForm)
	if err = ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}
	if err = c.validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// Use-Case
	user, err := c.usersRepository.FindByEmail(form.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "unknown email")
	}

	token := domain.CreateToken(domain.ResetPassword, user.Id)
	if err = c.tokensRepository.SaveOrReplace(token); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = c.emailClient.Send(email.PasswordResetEmail{
		Recipient: user,
		Token:     token,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error sending email: "+err.Error())
	}

	// Output
	return ctx.NoContent(http.StatusOK)
}
