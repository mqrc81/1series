package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
)

func (job updateGenresJob) name() string {
	return "UPDATE-GENRES job"
}

func (job updateGenresJob) execute() error {
	logger.Info("Running %v", job.name())

	tmdbGenres, err := job.tmdbClient.GetGenreTVList(nil)
	if err != nil {
		return err
	}

	var genres []domain.Genre
	for _, tmdbGenre := range tmdbGenres.Genres {
		genres = append(genres, domain.Genre{
			TmdbId: int(tmdbGenre.ID),
			Name:   tmdbGenre.Name,
		})
	}

	if err = job.genreRepository.ReplaceAll(genres); err != nil {
		return err
	}

	logger.Info("Completed %v with %d genres saved", job.name(), len(genres))
	return nil
}

type updateGenresJob struct {
	genreRepository repositories.GenreRepository
	tmdbClient      *tmdb.Client
}
