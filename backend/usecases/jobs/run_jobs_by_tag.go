package jobs

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (uc *useCase) RunJobsByTag(tag string) (int, error) {

	jobs, err := uc.scheduler.FindJobsByTag(tag)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = uc.scheduler.RunByTagWithDelay(tag, time.Second); err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return len(jobs), nil
}
