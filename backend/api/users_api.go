package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/util"
)

type UserHandler struct {
	store domain.Store
	log   util.Logger
}

// Register POST /api/users/register
func (h *UserHandler) Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// TODO
		return nil
	}
}
