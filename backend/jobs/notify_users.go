package jobs

import (
	"github.com/mqrc81/zeries/util"
)

func (e NotifyUsersJobExecutor) Execute() error {
	e.logStart()
	// TODO
	return e.logEnd(0)
}

type NotifyUsersJobExecutor struct {
	// TODO
}

func (e NotifyUsersJobExecutor) logStart() {
	util.LogInfo("Running notify-users job")
}

func (e NotifyUsersJobExecutor) logEnd(actions int) error {
	util.LogInfo("Completed notify-users job with %d users notified", actions)
	return nil
}
