package api

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

type UserHandler struct {
	store    domain.Store
	sessions *scs.SessionManager
}

// Register POST /api/users/register
func (h *UserHandler) Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		form := new(RegisterForm)
		if err := ctx.Bind(form); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error binding register-form: %v", err.Error()))
		}

		if _, err := h.store.GetUserByUsername(form.Username); err == nil {
			form.UsernameTaken = true
		}
		if _, err := h.store.GetUserByEmail(form.Email); err == nil {
			form.EmailTaken = true
		}

		if !form.Validate() {
			return ctx.JSON(http.StatusUnprocessableEntity, form.FormErrors)
		}

		return nil
	}
}
