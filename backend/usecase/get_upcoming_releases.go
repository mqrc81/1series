package usecase

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (uc *showUseCase) GetUpcomingReleases(page int) ([]domain.Release, error) {

	pastReleases, err := uc.releaseRepository.CountPastReleases()
	if err != nil {
		return []domain.Release{}, err
	}

	amount, offset := calculateRange(page, pastReleases)

	releasesRef, err := uc.releaseRepository.FindAllInRange(amount, offset)
	if err != nil {
		return []domain.Release{}, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	var releases []domain.Release
	for _, releaseRef := range releasesRef {
		tmdbRelease, err := uc.tmdbClient.GetTVDetails(
			releaseRef.ShowId, map[string]string{"append_to_response": "translations"},
		)
		if err != nil {
			return []domain.Release{}, echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		releases = append(
			releases,
			releaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate, releaseRef.AnticipationLevel),
		)
	}
	return releases, nil
}

func calculateRange(page int, pastReleases int) (int, int) {
	if page < 0 {
		return calculateRangeForPastReleases(pastReleases, page)
	}
	return calculateRangeForUpcomingReleases(pastReleases, page)
}

func calculateRangeForUpcomingReleases(pastReleases int, page int) (int, int) {
	// For pages 0+ return 20 releases
	return releasesPerRequest, pastReleases + releasesPerRequest*page
}

func calculateRangeForPastReleases(pastReleases int, page int) (int, int) {
	// For negative pages return 20 releases or max releases left for last page
	offset := pastReleases + releasesPerRequest*page
	amount := releasesPerRequest
	if offset <= 0 {
		// The last possible page for past releases has been reached
		amount += offset
		offset = 0
	}
	return amount, offset
}
