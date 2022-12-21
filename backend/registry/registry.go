package registry

import (
	"github.com/go-co-op/gocron"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/cyruzin/golang-tmdb"
	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/controller"
	"github.com/mqrc81/zeries/job"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
)

func NewDatabase(
	dataSourceName string,
) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dataSourceName)
}

func NewSessionManager(
	db *sqlx.DB,
) (*scs.SessionManager, error) {
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.NewWithCleanupInterval(db.DB, 1*time.Hour)
	sessionManager.Lifetime = 24 * time.Hour
	return sessionManager, nil
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
	database *sqlx.DB, sessionManager *scs.SessionManager, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (controller.Controller, error) {
	return controller.NewController(database, sessionManager, tmdbClient, traktClient)
}

func NewScheduler(
	database *sqlx.DB, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) (*gocron.Scheduler, error) {
	scheduler := gocron.NewScheduler(time.UTC)

	refreshGenresAndNetworksJob := job.NewRefreshGenresAndNetworksJob(repository.NewGenreRepository(database), repository.NewNetworkRepository(database), tmdbClient)
	_, err := scheduler.Tag(job.RunOnInitTag).Every(1).Day().At("00:00").Do(refreshGenresAndNetworksJob.Execute)
	if err != nil {
		return nil, err
	}

	updateReleasesJob := job.NewUpdateReleasesJob(repository.NewReleaseRepository(database), tmdbClient, traktClient)
	_, err = scheduler.Tag(job.RunOnInitTag).Every(1).Day().At("00:30").Do(updateReleasesJob.Execute)
	if err != nil {
		return nil, err
	}

	notifyUsersJob := job.NewNotifyUsersJob(repository.NewUserRepository(database))
	_, err = scheduler.Every(3).Days().At("01:00").Do(notifyUsersJob.Execute)
	if err != nil {
		return nil, err
	}

	return scheduler, nil
}
