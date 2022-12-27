package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-co-op/gocron"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
	"time"
)

const (
	PopulateDatabaseTag = "DATABASE"
	NotifyUsersTag      = "NOTIFICATION"
)

type job interface {
	execute() error
	name() string
}

func RegisterUpdateGenresJob(
	scheduler *gocron.Scheduler,
	genreRepository repositories.GenreRepository,
	tmdbClient *tmdb.Client,
) error {
	return registerJob(
		updateGenresJob{genreRepository, tmdbClient},
		scheduler.Every(1).Monday().At("00:00").Tag(PopulateDatabaseTag).Do,
	)
}

func RegisterUpdateReleasesJob(
	scheduler *gocron.Scheduler,
	releaseRepository repositories.ReleaseRepository,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
) error {
	return registerJob(
		updateReleasesJob{releaseRepository, tmdbClient, traktClient},
		scheduler.Every(1).Day().At("00:05").Tag(PopulateDatabaseTag).Do,
	)
}

func RegisterNotifyUsersAboutReleasesJob(
	scheduler *gocron.Scheduler,
	userRepository repositories.UserRepository,
	releaseRepository repositories.ReleaseRepository,
	trackedShowRepository repositories.TrackedShowRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
) error {
	return registerJob(
		notifyUsersAboutReleasesJob{userRepository, releaseRepository, trackedShowRepository, tmdbClient, emailClient},
		scheduler.Every(1).Monday().At("00:10").Tag(NotifyUsersTag).Do,
	)
}

func RegisterNotifyUsersAboutRecommendationsJob(
	scheduler *gocron.Scheduler,
	userRepository repositories.UserRepository,
	trackedShowRepository repositories.TrackedShowRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
) error {
	return registerJob(
		notifyUsersAboutRecommendationsJob{userRepository, trackedShowRepository, tmdbClient, emailClient},
		scheduler.Every(5).Weekday(time.Friday).At("00:15").Tag(NotifyUsersTag).Do,
	)
}

type scheduleJobFunc = func(jobFun interface{}, params ...interface{}) (*gocron.Job, error)

func registerJob(job job, scheduleJobFunc scheduleJobFunc) error {
	executeJobFunc := func() {
		if err := job.execute(); err != nil {
			logger.Error("error executing %v: %v", job.name(), err.Error())
		}
	}
	_, err := scheduleJobFunc(executeJobFunc)
	return err
}
