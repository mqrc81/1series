package job

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
	. "github.com/mqrc81/zeries/util"
)

const (
	RunOnInitTag = "INIT"
)

func executor(job job) func() {
	return func() {
		err := job.execute()
		if err != nil {
			LogError(err.Error())
		}
	}
}

type job interface {
	execute() error
}

func NewUpdateReleasesJob(
	releaseRepository repository.ReleaseRepository, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) func() {
	return executor(updateReleasesJob{
		releaseRepository,
		tmdbClient,
		traktClient,
	})
}

func NewRefreshGenresAndNetworksJob(
	genreRepository repository.GenreRepository, networkRepository repository.NetworkRepository, tmdbClient *tmdb.Client,
) func() {
	return executor(refreshGenresAndNetworksJob{
		genreRepository,
		networkRepository,
		tmdbClient,
	})
}

func NewNotifyUsersJob(
	userRepository repository.UserRepository,
) func() {
	return executor(notifyUsersJob{
		userRepository,
	})
}
