package shows

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (uc *useCase) GetUpcomingReleases(page int) ([]domain.Release, bool, error) {

	pastReleases, err := uc.releaseRepository.CountPastReleases()
	if err != nil {
		return []domain.Release{}, false, err
	}

	amount, offset, possiblyHasMore := calculateRange(page, pastReleases)

	releasesRef, err := uc.releaseRepository.FindAllInRange(amount, offset)
	if err != nil {
		return []domain.Release{}, false, echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	var releases []domain.Release
	for _, releaseRef := range releasesRef {
		tmdbRelease, err := uc.tmdbClient.GetTVDetails(releaseRef.ShowId, map[string]string{"append_to_response": "translations"})
		if err != nil {
			return []domain.Release{}, false, echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		releases = append(
			releases,
			ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate, releaseRef.AnticipationLevel),
		)
	}
	return releases, possiblyHasMore && len(releases) >= 20, nil
}

func calculateRange(page int, pastReleases int) (int, int, bool) {
	if page < 0 {
		return calculateRangeForPastReleases(pastReleases, page)
	}
	return calculateRangeForUpcomingReleases(pastReleases, page)
}

func calculateRangeForUpcomingReleases(pastReleases int, page int) (int, int, bool) {
	// For pages > 0 return 20 releases
	return upcomingReleasesPerRequest, pastReleases + upcomingReleasesPerRequest*(page-1), true
}

func calculateRangeForPastReleases(pastReleases int, page int) (int, int, bool) {
	// For pages < 0 return 20 releases or max releases left for last page
	offset := pastReleases + upcomingReleasesPerRequest*page
	amount := upcomingReleasesPerRequest
	hasMore := true
	if offset <= 0 {
		// The last possible page for past releases has been reached
		hasMore = false
		amount += offset
		offset = 0
	}
	return amount, offset, hasMore
}
