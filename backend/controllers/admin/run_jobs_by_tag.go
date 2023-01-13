package admin

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/controllers/users"
	"net/http"
	"time"
)

func (c *adminController) TriggerJobs(ctx echo.Context) (err error) {
	// Input
	tag := ctx.QueryParam("tag")
	if user, err := users.GetAuthenticatedUser(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !isAdmin(user) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only the big boss is allowed to run jobs manually you peasant")
	}

	// Use-Case
	jobs, err := c.scheduler.FindJobsByTag(tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = c.scheduler.RunByTagWithDelay(tag, time.Second); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Output
	return ctx.JSON(http.StatusOK, fmt.Sprintf("Running %d jobs", len(jobs)))
}
