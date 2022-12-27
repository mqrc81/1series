package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
)

func (job notifyUsersAboutRecommendationsJob) name() string {
	return "NOTIFY-USERS-ABOUT-RECOMMENDATIONS job"
}

func (job notifyUsersAboutRecommendationsJob) execute() error {
	logger.Warning("%v not implemented yet", job.name())

	// TODO

	return nil
}

type notifyUsersAboutRecommendationsJob struct {
	userRepository        repositories.UserRepository
	trackedShowRepository repositories.TrackedShowRepository
	tmdbClient            *tmdb.Client
	emailClient           *email.Client
}
