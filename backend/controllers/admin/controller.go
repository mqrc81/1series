package admin

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
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
