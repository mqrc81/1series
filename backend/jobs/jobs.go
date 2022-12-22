package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

const (
	RunOnInitTag = "INIT"
)

func executor(job job) func() {
	return func() {
		err := job.
			execute()
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func errorMsg(job job) string {
	return "error executing " + job.name()
}

type job interface {
	execute() error
	name() string
}

func NewUpdateReleasesJob(
	releaseRepository repositories.ReleaseRepository, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) func() {
	return executor(updateReleasesJob{
		releaseRepository,
		tmdbClient,
		traktClient,
	})
}

func NewUpdateGenresJob(
	genreRepository repositories.GenreRepository, tmdbClient *tmdb.Client,
) func() {
	return executor(updateGenresJob{
		genreRepository,
		tmdbClient,
	})
}

func NewNotifyUsersJob(
	userRepository repositories.UserRepository,
) func() {
	return executor(notifyUsersJob{
		userRepository,
	})
}
