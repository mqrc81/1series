package usecase

import (
	"github.com/mqrc81/zeries/domain"
)

func (uc *showUseCase) GetPopularShows(page int) ([]domain.Show, error) {
	traktShows, err := uc.traktClient.GetShowsWatchedWeekly(page, 20)
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
