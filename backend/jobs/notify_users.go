package jobs

import (
	"log"
)

func (e NotifyUsersJobExecutor) Execute() error {
	e.logStart()
	// TODO
	return e.logEnd()
}

type NotifyUsersJobExecutor struct {
	// TODO
	actions int
}

func (e NotifyUsersJobExecutor) logStart() {
	log.Println("Running notify-users job")
}

func (e NotifyUsersJobExecutor) logEnd() error {
	log.Printf("Completed notify-users job with %d users notified\n", e.actions)
	return nil
}
