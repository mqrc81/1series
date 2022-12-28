package admin

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/env"
)

type adminController struct {
	scheduler *gocron.Scheduler
}

type Controller interface {
	TriggerJobs(ctx echo.Context) error
}

func NewController(
	scheduler *gocron.Scheduler,
) Controller {
	return &adminController{scheduler}
}

func isAdmin(user domain.User) bool {
	for _, admin := range env.Config.Admins {
		if user.Username == admin {
			return true
		}
	}
	return false
}
