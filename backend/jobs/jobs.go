package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/go-co-op/gocron"
	"github.com/mqrc81/1series/email"
	"github.com/mqrc81/1series/logger"
	"github.com/mqrc81/1series/repositories"
	"github.com/mqrc81/1series/trakt"
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
	genresRepository repositories.GenresRepository,
	tmdbClient *tmdb.Client,
) error {
	return registerJob(
		updateGenresJob{genresRepository, tmdbClient},
		scheduler.Every(1).Monday().At("00:00").Tag(PopulateDatabaseTag).Do,
	)
}

func RegisterUpdateReleasesJob(
	scheduler *gocron.Scheduler,
	releasesRepository repositories.ReleasesRepository,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
) error {
	return registerJob(
		updateReleasesJob{releasesRepository, tmdbClient, traktClient},
		scheduler.Every(1).Day().At("00:05").Tag(PopulateDatabaseTag).Do,
	)
}

func RegisterNotifyUsersAboutReleasesJob(
	scheduler *gocron.Scheduler,
	usersRepository repositories.UsersRepository,
	releasesRepository repositories.ReleasesRepository,
	trackedShowsRepository repositories.TrackedShowsRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
) error {
	return registerJob(
		notifyUsersAboutReleasesJob{usersRepository, releasesRepository, trackedShowsRepository, tmdbClient, emailClient},
		scheduler.Every(1).Monday().At("00:10").Tag(NotifyUsersTag).Do,
	)
}

func RegisterNotifyUsersAboutRecommendationsJob(
	scheduler *gocron.Scheduler,
	usersRepository repositories.UsersRepository,
	trackedShowsRepository repositories.TrackedShowsRepository,
	tmdbClient *tmdb.Client,
	emailClient *email.Client,
) error {
	return registerJob(
		notifyUsersAboutRecommendationsJob{usersRepository, trackedShowsRepository, tmdbClient, emailClient},
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
