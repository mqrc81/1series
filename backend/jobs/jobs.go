// Package jobs defines the jobs the daily scheduler must execute
package jobs

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

type JobExecutor interface {
	Execute() error
	logStart()
	logEnd() error
}

func NewUpdateReleasesJob(store domain.Store, tmdbClient *tmdb.Client, traktClient *trakt.Client) JobExecutor {
	return UpdateReleasesJobExecutor{
		store: store,
		tmdb:  tmdbClient,
		trakt: traktClient,
	}
}

func NewNotifyUsersJob() JobExecutor {
	return NotifyUsersJobExecutor{}
}
