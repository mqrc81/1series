package registry

import (
	"github.com/go-co-op/gocron"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/env"
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

func NewDatabase() *sql.Database {
	database, err := sqlx.Connect("postgres", env.Config.DatabaseUrl)
	logger.FatalOnError(err)
	goose.SetLogger(logger.DefaultLogger)
	err = goose.SetDialect("postgres")
	logger.FatalOnError(err)
	return &sql.Database{DB: database}
}

func NewTmdbClient() *tmdb.Client {
	tmdbClient, err := tmdb.Init(env.Config.TmdbKey)
	logger.FatalOnError(err)
	return tmdbClient
}

func NewTraktClient() *trakt.Client {
	traktClient, err := trakt.Init(env.Config.TraktKey)
	logger.FatalOnError(err)
	return traktClient
}

func NewEmailClient() *email.Client {
	emailClient, err := email.NewEmailClient(env.Config.SendGridKey, env.Config.SendGridSenderEmail)
	logger.FatalOnError(err)
	return emailClient
}

func NewScheduler(
	database *sql.Database,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
	emailClient *email.Client,
) *gocron.Scheduler {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.SetMaxConcurrentJobs(1, gocron.WaitMode)

	err := jobs.RegisterUpdateGenresJob(scheduler, repositories.NewGenresRepository(database), tmdbClient)
	logger.FatalOnError(err)

	err = jobs.RegisterUpdateReleasesJob(scheduler, repositories.NewReleasesRepository(database), tmdbClient, traktClient)
	logger.FatalOnError(err)

	err = jobs.RegisterNotifyUsersAboutReleasesJob(scheduler, repositories.NewUsersRepository(database), repositories.NewReleasesRepository(database), repositories.NewTrackedShowsRepository(database), tmdbClient, emailClient)
	logger.FatalOnError(err)

	err = jobs.RegisterNotifyUsersAboutRecommendationsJob(scheduler, repositories.NewUsersRepository(database), repositories.NewTrackedShowsRepository(database), tmdbClient, emailClient)
	logger.FatalOnError(err)

	return scheduler
}

func NewController(
	database *sql.Database,
	tmdbClient *tmdb.Client,
	traktClient *trakt.Client,
	emailClient *email.Client,
	scheduler *gocron.Scheduler,
) controllers.Controller {
	controller, err := controllers.NewController(database, tmdbClient, traktClient, emailClient, scheduler)
	logger.FatalOnError(err)
	return controller
}
