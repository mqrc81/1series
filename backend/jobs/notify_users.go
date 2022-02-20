package jobs

import (
	"log"
)

func (e NotifyUsersJobExecutor) Execute() error {
	e.logStart()
	// TODO
	return e.logEnd(0)
}

type NotifyUsersJobExecutor struct {
	// TODO
}

func (NotifyUsersJobExecutor) logStart() {
	log.Println("Running notify-users job")
}

func (NotifyUsersJobExecutor) logEnd(actions int) error {
	log.Printf("Completed notify-users job with %d users notified\n", actions)
	return nil
}
