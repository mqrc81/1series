package jobs

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

type jobController struct {
	scheduler *gocron.Scheduler
}

type Controller interface {
	RunJobsByTag(ctx echo.Context) error
}

func NewController(
	scheduler *gocron.Scheduler,
) Controller {
	return &jobController{scheduler}
}
