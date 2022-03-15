package jobs

import (
	"fmt"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	. "github.com/mqrc81/zeries/util"
)

const (
	thirtyDays          = 30 * 24 * time.Hour
	defaultErrorMessage = "error executing update-releases job"
)

func (e UpdateReleasesJobExecutor) Execute() error {
	e.logStart()

	var (
		releases          []domain.ReleaseRef
		pastReleasesCount int
		now               = time.Now()
		// Start at 30 days in the past to allow users to view past releases
		startDate = now.Add(-thirtyDays)
		expiry    = now.Add(3 * time.Hour)
	)

	traktShowsAnticipated, err := e.trakt.GetAnticipatedShows(1, 10)
	if err != nil {
		return fmt.Errorf("%v whilst fetching trakt season-premieres: %w", defaultErrorMessage, err)
	}

	// Trakt's limit is 33 days per request, but we want 9 * 30 days
	for i := 0; i < 9; i++ {

		traktReleases, err := e.trakt.GetSeasonPremieres(startDate, 30)
		if err != nil {
			return fmt.Errorf("%v whilst fetching trakt season-premieres: %w", defaultErrorMessage, err)
		}

		releases, err = e.filterAndCollectReleases(releases, traktReleases, traktShowsAnticipated)
		if err != nil {
			return err
		}

		if i == 0 {
			// The first iteration takes care of all past releases
			pastReleasesCount = len(releases)
		}

		startDate = startDate.Add(thirtyDays)
	}

	err = e.updateReleases(releases, expiry, pastReleasesCount, now)
	if err != nil {
		return err
	}

	return e.logEnd(len(releases))
}

func (e UpdateReleasesJobExecutor) filterAndCollectReleases(
	releases []domain.ReleaseRef, traktReleases []trakt.SeasonPremieresDto, traktShows []trakt.ShowsAnticipatedDto,
) ([]domain.ReleaseRef, error) {
	for _, traktRelease := range traktReleases {
		if !hasRelevantIds(traktRelease) {
			continue
		}

		tmdbShow, err := e.tmdb.GetTVDetails(traktRelease.TmdbId(),
			map[string]string{"append_to_response": "translations"})

		if err != nil || !hasRelevantInfo(tmdbShow) || !hasMatchingSeasons(traktRelease, tmdbShow) {
			continue
		}

		releases = append(releases, domain.ReleaseRef{
			ShowId:            traktRelease.TmdbId(),
			SeasonNumber:      traktRelease.SeasonNumber(),
			AirDate:           traktRelease.AirDate(),
			AnticipationLevel: anticipationLevelFor(traktRelease.TmdbId(), traktShows),
		})
	}
	return releases, nil
}

func (e UpdateReleasesJobExecutor) updateReleases(
	releases []domain.ReleaseRef, expiry time.Time, pastReleasesCount int, now time.Time,
) (err error) {
	var updatedPastReleasesCount bool
	for i, release := range releases {
		if err = e.store.SaveRelease(release, expiry); err != nil {
			return fmt.Errorf("%v: %w", defaultErrorMessage, err)
		}
		if err = e.store.ClearExpiredReleases(now, release.AirDate); err != nil {
			return fmt.Errorf("%v: %w", defaultErrorMessage, err)
		}
		if !updatedPastReleasesCount && i >= pastReleasesCount {
			if err = e.store.SetPastReleasesCount(pastReleasesCount); err != nil {
				return fmt.Errorf("%v: %w", defaultErrorMessage, err)
			}
		}
	}
	if err = e.store.ClearExpiredReleases(now, now.Add(12*30*24*time.Hour)); err != nil {
		return fmt.Errorf("%v: %w", defaultErrorMessage, err)
	}
	return nil
}

func hasRelevantIds(release trakt.SeasonPremieresDto) bool {
	ids := release.Show.Ids
	return ids.Tmdb != 0 && ids.Tvdb != 0 && ids.Imdb != "" && ids.Slug != ""
}

func hasRelevantInfo(show *tmdb.TVDetails) bool {
	return len(show.Genres) > 0 &&
		len(show.Networks) > 0 &&
		len(show.Overview) > 0 &&
		len(show.Seasons) > 0 &&
		hasEnglishTranslation(show.Translations) &&
		hasRelevantType(show)
}

func hasEnglishTranslation(translations *tmdb.TVTranslations) bool {
	for _, translation := range translations.Translations {
		if translation.Iso639_1 == "en" {
			return true
		}
	}
	return false
}

func hasRelevantType(show *tmdb.TVDetails) bool {
	t := show.Type
	if t == "Scripted" || t == "Miniseries" || t == "Documentary" {
		return true
	}
	if t != "Reality" && t != "News" && t != "Talk Show" && t != "Video" {
		LogWarning("Unknown type [%v] detected for show [%d, %v]", show.Type, show.ID, show.Name)
	}
	return false
}

func hasMatchingSeasons(release trakt.SeasonPremieresDto, show *tmdb.TVDetails) bool {
	return len(show.Seasons) >= release.SeasonNumber()
}

func anticipationLevelFor(releaseId int, showsAnticipated []trakt.ShowsAnticipatedDto) domain.AnticipationLevel {
	for i, showAnticipated := range showsAnticipated {
		if releaseId == showAnticipated.TmdbId() {
			if i == 0 {
				return domain.Zamn
			} else if i < 3 {
				return domain.Bussin
			} else if i < 10 {
				return domain.Mid
			}
		}
	}
	return domain.Zero
}

type UpdateReleasesJobExecutor struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

func (e UpdateReleasesJobExecutor) logStart() {
	LogInfo("Running update-releases job")
}

func (e UpdateReleasesJobExecutor) logEnd(actions int) error {
	LogInfo("Completed update-releases job with %d releases updated", actions)
	return nil
}
