package registry

import (
	"github.com/go-co-op/gocron"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/sql"
	"github.com/pressly/goose/v3"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/controllers"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/repositories"
	_ "github.com/mqrc81/zeries/sql"
	"github.com/mqrc81/zeries/trakt"
)

func NewDatabase(
	dataSourceName string,
) (*sql.Database, error) {
	database, err := sqlx.Connect("postgres", dataSourceName)
	if err == nil {
		goose.SetLogger(logger.DefaultLogger)
		err = goose.SetDialect("postgres")
	}
	return &sql.Database{DB: database}, err
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
	database *sql.Database, tmdbClient *tmdb.Client, traktClient *trakt.Client, scheduler *gocron.Scheduler,
) (controllers.Controller, error) {
	return controllers.NewController(database, tmdbClient, traktClient, scheduler)
}

func NewScheduler(
	database *sql.Database, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (*gocron.Scheduler, error) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SetMaxConcurrentJobs(1, gocron.WaitMode)

	err := jobs.RegisterUpdateGenresJob(scheduler, repositories.NewGenreRepository(database), tmdbClient)
	if err != nil {
		return nil, err
	}

	err = jobs.RegisterUpdateReleasesJob(scheduler, repositories.NewReleaseRepository(database), tmdbClient, traktClient)
	if err != nil {
		return nil, err
	}

	err = jobs.RegisterNotifyUsersAboutReleasesJob(scheduler, repositories.NewUserRepository(database), repositories.NewWatchedShowRepository(database))
	if err != nil {
		return nil, err
	}

	return scheduler, nil
}
