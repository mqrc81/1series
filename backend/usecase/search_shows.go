package usecase

import (
	"github.com/mqrc81/zeries/domain"
)

func (uc *showUseCase) SearchShows(searchTerm string) ([]domain.Show, error) {

	tmdbShows, err := uc.tmdbClient.GetSearchTVShow(searchTerm, map[string]string{"language": "en-US"})
	if err != nil {
		return []domain.Show{}, err
	}

	return showsFromTmdbShowsSearch(tmdbShows, 8), nil
}
