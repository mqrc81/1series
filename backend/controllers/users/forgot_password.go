package users

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	resetPasswordExpiration  = 1 * time.Hour
	resetPasswordTokenLength = 32
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

	token := generateToken(resetPasswordTokenLength)
	if err = c.tokensRepository.Save(domain.Token{
		TokenId:   token,
		UserId:    user.Id,
		Purpose:   domain.ResetPassword,
		ExpiresAt: time.Now().Add(resetPasswordExpiration),
	}); err != nil {
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

func generateToken(length int) string {
	b := make([]byte, length/2)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
