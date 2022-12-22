package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/usecases/jobs"
	"net/http"
)

type jobController struct {
	jobUseCase jobs.UseCase
}

type JobController interface {
	RunJobsByTag(ctx echo.Context) error
}

func newJobController(jobUseCase jobs.UseCase) JobController {
	return &jobController{jobUseCase}
}

func (c *jobController) RunJobsByTag(ctx echo.Context) (err error) {
	// Input
	tag := ctx.QueryParam("tag")
	// currentSession, _ := session.Get(sessionKey, ctx)
	// if currentSession.Values[sessionUserKey] != "marc" {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Only the big boss is allowed to run jobs manually you peasant")
	// }

	// Use-Case
	amountOfJobs, err := c.jobUseCase.RunJobsByTag(tag)
	if err != nil {
		return err
	}

	// Output
	return ctx.JSON(http.StatusOK, fmt.Sprintf("Running %d jobs", amountOfJobs))
}
