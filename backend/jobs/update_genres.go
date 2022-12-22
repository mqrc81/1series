package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
)

func (job updateGenresJob) name() string {
	return "UPDATE-GENRES job"
}

func (job updateGenresJob) execute() error {
	logger.Info("Yet to implement " + job.name())

	// TODO

	return nil
}

type updateGenresJob struct {
	genreRepository repositories.GenreRepository
	tmdbClient      *tmdb.Client
}
