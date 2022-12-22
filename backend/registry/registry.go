package registry

import (
	"github.com/go-co-op/gocron"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/controllers"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

func NewDatabase(
	dataSourceName string,
) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dataSourceName)
}

func NewTmdbClient(
	tmdbKey string,
) (*tmdb.Client, error) {
	return tmdb.Init(tmdbKey)
}

func NewTraktClient(
	traktKey string,
) (*trakt.Client, error) {
	return trakt.Init(traktKey)
}

func NewController(
	database *sqlx.DB, tmdbClient *tmdb.Client, traktClient *trakt.Client, scheduler *gocron.Scheduler,
) (controllers.Controller, error) {
	return controllers.NewController(database, tmdbClient, traktClient, scheduler)
}

func NewScheduler(
	database *sqlx.DB, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (*gocron.Scheduler, error) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SetMaxConcurrentJobs(1, gocron.WaitMode)

	updateGenresJob := jobs.NewUpdateGenresJob(repositories.NewGenreRepository(database), tmdbClient)
	_, err := scheduler.Every(1).Monday().At("00:00").Tag(jobs.RunOnInitTag).Do(updateGenresJob)
	if err != nil {
		return nil, err
	}

	updateReleasesJob := jobs.NewUpdateReleasesJob(repositories.NewReleaseRepository(database), tmdbClient, traktClient)
	_, err = scheduler.Every(1).Day().At("00:05").Tag(jobs.RunOnInitTag).Do(updateReleasesJob)
	if err != nil {
		return nil, err
	}

	notifyUsersJob := jobs.NewNotifyUsersJob(repositories.NewUserRepository(database))
	_, err = scheduler.Every(1).Monday().At("00:10").Do(notifyUsersJob)
	if err != nil {
		return nil, err
	}

	return scheduler, nil
}
