package job

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/repository"
	. "github.com/mqrc81/zeries/util"
)

func (e refreshGenresAndNetworksJob) execute() error {
	LogInfo("Yet to implement refresh-genres-and-networks job")

	// TODO

	return nil
}

type refreshGenresAndNetworksJob struct {
	genreRepository   repository.GenreRepository
	networkRepository repository.NetworkRepository
	tmdbClient        *tmdb.Client
}
