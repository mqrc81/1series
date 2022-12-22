package jobs

import (
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
)

func (job notifyUsersJob) name() string {
	return "NOTIFY-USERS job"
}

func (job notifyUsersJob) execute() error {
	logger.Info("Yet to implement " + job.name())

	// TODO

	return nil
}

type notifyUsersJob struct {
	userRepository repositories.UserRepository
}
