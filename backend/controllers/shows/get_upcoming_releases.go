package shows

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mqrc81/zeries/domain"
)

func (c *showController) GetUpcomingReleases(ctx echo.Context) error {
	// Input
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "parameter page must be positive or negative")
	}

	// Use-Case
	pastReleases, err := c.releaseRepository.CountPastReleases()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	amount, offset, possiblyHasMore := calculateRange(page, pastReleases)

	releasesRef, err := c.releaseRepository.FindAllInRange(amount, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	var releases []domain.Release
	for _, releaseRef := range releasesRef {
		tmdbRelease, err := c.tmdbClient.GetTVDetails(releaseRef.ShowId, map[string]string{"append_to_response": "translations"})
		if err != nil {
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		releases = append(
			releases,
			ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate, releaseRef.AnticipationLevel),
		)
	}

	// Output
	previousPage, nextPage := paginationForUpcomingReleases(page, possiblyHasMore && len(releases) >= 20)
	return ctx.JSON(http.StatusOK, upcomingReleasesDto{
		PreviousPage: previousPage,
		NextPage:     nextPage,
		Releases:     releases,
	})
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

func paginationForUpcomingReleases(currentPage int, hasMoreReleases bool) (previousPage int, nextPage int) {
	if currentPage == 1 {
		previousPage = -1
	}
	if hasMoreReleases && currentPage > 0 {
		nextPage = currentPage + 1
	} else if hasMoreReleases && currentPage < 0 {
		previousPage = currentPage - 1
	}
	return previousPage, nextPage
}
