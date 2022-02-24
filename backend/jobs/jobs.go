// Package jobs defines the jobs the daily scheduler must execute
package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/util"
)

type JobExecutor interface {
	Execute() error
	logStart()
	logEnd(actions int) error
	util.Logger
}

func NewUpdateReleasesJob(store domain.Store, tmdbClient *tmdb.Client, traktClient *trakt.Client,
	logger util.Logger) JobExecutor {
	return UpdateReleasesJobExecutor{
		store:  store,
		tmdb:   tmdbClient,
		trakt:  traktClient,
		Logger: logger,
	}
}

func NewNotifyUsersJob(logger util.Logger) JobExecutor {
	return NotifyUsersJobExecutor{
		Logger: logger,
	}
}
