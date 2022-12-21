package job

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
)

const (
	RunOnInitTag = "INIT"
)

type Executor interface {
	Execute() error
}

func NewUpdateReleasesJob(
	releaseRepository repository.ReleaseRepository, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) Executor {
	return updateReleasesJob{
		releaseRepository,
		tmdbClient,
		traktClient,
	}
}

func NewRefreshGenresAndNetworksJob(
	genreRepository repository.GenreRepository, networkRepository repository.NetworkRepository, tmdbClient *tmdb.Client,
) Executor {
	return refreshGenresAndNetworksJob{
		genreRepository,
		networkRepository,
		tmdbClient,
	}
}

func NewNotifyUsersJob(userRepository repository.UserRepository) Executor {
	return notifyUsersJob{
		userRepository,
	}
}
