package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mqrc81/zeries/domain"
)

type UserHandler struct {
	store domain.Store
}

// Register POST /api/users/register
func (h *UserHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO
	}
}
