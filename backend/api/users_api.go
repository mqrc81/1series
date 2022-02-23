package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"go.uber.org/zap"
)

type UserHandler struct {
	store domain.Store
	log   *zap.SugaredLogger
}

// Register POST /api/users/register
func (h *UserHandler) Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// TODO
		return nil
	}
}
