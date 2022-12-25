package shows

import (
	"github.com/mqrc81/zeries/domain"
)

type PopularShowsDto struct {
	nextPage int
	shows    []domain.Show
}

func (uc *useCase) GetPopularShows(page int) ([]domain.Show, error) {
	traktShows, err := uc.traktClient.GetShowsWatchedWeekly(page, popularShowsPerRequest)
	if err != nil {
		return []domain.Show{}, err
	}

	var shows []domain.Show
	for _, traktShow := range traktShows {
		tmdbShow, err := uc.tmdbClient.GetTVDetails(traktShow.TmdbId(), nil)
		if err != nil {
			return []domain.Show{}, err
		}

		shows = append(shows, showFromTmdbShow(tmdbShow))
	}
	return shows, nil
}
