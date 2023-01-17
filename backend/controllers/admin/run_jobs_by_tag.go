package admin

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/controllers/errors"
	"github.com/mqrc81/zeries/controllers/users"
	"net/http"
	"time"
)

func (c *adminController) TriggerJobs(ctx echo.Context) (err error) {
	// Input
	tag := ctx.QueryParam("tag")
	if tag == "" {
		return errors.MissingParameter("tag")
	}

	// Use-Case
	if user, err := users.GetAuthenticatedUser(ctx); err != nil || !isAdmin(user) {
		return errors.AdminOnly()
	}

	jobs, err := c.scheduler.FindJobsByTag(tag)
	if err != nil {
		return errors.NotFound("jobs", errors.Params{"tag": tag})
	}
	if err = c.scheduler.RunByTagWithDelay(tag, time.Second); err != nil {
		return errors.Internal(err)
	}

	// Output
	return ctx.String(http.StatusOK, fmt.Sprintf("Running %d jobs.", len(jobs)))
}
