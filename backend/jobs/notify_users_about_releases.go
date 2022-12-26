package jobs

import (
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
)

func (job notifyUsersAboutReleasesJob) name() string {
	return "NOTIFY-USERS-ABOUT-RELEASES job"
}

func (job notifyUsersAboutReleasesJob) execute() error {
	logger.Info("Yet to implement " + job.name())

	return nil
}

type notifyUsersAboutReleasesJob struct {
	userRepository        repositories.UserRepository
	watchedShowRepository repositories.WatchedShowRepository
}
