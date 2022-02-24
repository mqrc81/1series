package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

type UserHandler struct {
	store domain.Store
}

// Register POST /api/users/register
func (h *UserHandler) Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		form := new(RegisterForm)
		if err := ctx.Bind(form); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("error binding register-form: %v", err.Error()))
		}
		if err := ctx.Validate(form); err != nil {
			return echo.NewHTTPError(http.StatusPartialContent, err)
		}
		return nil
	}
}
