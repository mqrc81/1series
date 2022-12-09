package job

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/repository"
	"github.com/mqrc81/zeries/trakt"
)

type JobExecutor interface {
	Execute() error
}

func NewUpdateReleasesJob(
	releaseRepository repository.ReleaseRepository, tmdbClient *tmdb.Client, traktClient *trakt.Client,
) JobExecutor {
	return updateReleasesJob{
		releaseRepository: releaseRepository,
		tmdbClient:        tmdbClient,
		traktClient:       traktClient,
	}
}

func NewNotifyUsersJob() JobExecutor {
	return notifyUsersJob{}
}
