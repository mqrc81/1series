package job

import (
	. "github.com/mqrc81/zeries/util"
)

func (e notifyUsersJob) Execute() error {
	LogInfo("Running notify-users job")

	// TODO
	LogInfo("Completed notify-users job with %d users notified", 0)
	return nil
}

type notifyUsersJob struct {
	// TODO
}
