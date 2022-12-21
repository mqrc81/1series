package job

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/repository"
	. "github.com/mqrc81/zeries/util"
)

func (e refreshGenresAndNetworksJob) Execute() error {
	LogInfo("Running refresh-genres-and-networks job")

	// TODO

	LogInfo("Completed refresh-genres-and-networks job with %d genres or networks changed", 0)
	return nil
}

type refreshGenresAndNetworksJob struct {
	genreRepository   repository.GenreRepository
	networkRepository repository.NetworkRepository
	tmdbClient        *tmdb.Client
}
