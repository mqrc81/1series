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
	var releasesUpdated int
	now := time.Now()
	// Start at 30 days in the past to allow users to view past releases
	startDate := now.Add(-thirtyDays)
	expiry := now.Add(3 * time.Hour)

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

		newReleasesUpdated, err := e.filterAndUpdateReleases(traktReleases, traktShowsAnticipated, expiry)
		if err != nil {
			return err
		}
		releasesUpdated += newReleasesUpdated

		if i == 0 {
			// The first iteration takes care of all past releases
			if err = e.store.SetPastReleasesCount(releasesUpdated); err != nil {
				return fmt.Errorf("%v: %w", defaultErrorMessage, err)
			}
		}

		startDate = startDate.Add(thirtyDays)
	}

	// TODO: there is a small window in which expired and updated releases can coexist
	if err := e.store.ClearExpiredReleases(now); err != nil {
		return fmt.Errorf("%v: %w", defaultErrorMessage, err)
	}

	return e.logEnd(releasesUpdated)
}

func (e UpdateReleasesJobExecutor) filterAndUpdateReleases(
	traktReleases []trakt.SeasonPremieresDto, traktShowsAnticipated []trakt.ShowsAnticipatedDto, expiry time.Time,
) (int, error) {
	var releasesUpdated int
	for _, traktRelease := range traktReleases {
		if !hasRelevantIds(traktRelease) {
			continue
		}

		tmdbShow, err := e.tmdb.GetTVDetails(traktRelease.TmdbId(),
			map[string]string{"append_to_response": "translations"})
		if err != nil {
			// On rare occasions trakt's tmdb-id might be incorrect
			// We treat this case as if jobs#hasRelevantIds was false
			LogWarning("Incorrect tmdb-id [%v]: %v", traktRelease.Ids(), err)
			continue
		}

		if !hasRelevantInfo(tmdbShow) || !hasMatchingSeasons(traktRelease, tmdbShow) {
			continue
		}

		if err = e.store.SaveRelease(domain.ReleaseRef{
			ShowId:            traktRelease.TmdbId(),
			SeasonNumber:      traktRelease.SeasonNumber(),
			AirDate:           traktRelease.AirDate(),
			AnticipationLevel: anticipationLevelFor(traktRelease.TmdbId(), traktShowsAnticipated),
		}, expiry); err != nil {
			return 0, fmt.Errorf("%v whilst saving release [%v, %d, %v]: %w", defaultErrorMessage,
				traktRelease.Ids(), traktRelease.SeasonNumber(), traktRelease.AirDate(), err)
		}

		releasesUpdated++
	}
	return releasesUpdated, nil
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
