package jobs

import (
	"fmt"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
	"github.com/mqrc81/zeries/util"
)

const (
	thirtyDays          = traktDaysCap * 24 * time.Hour
	traktDaysCap        = 30
	defaultErrorMessage = "error executing update-releases job"
)

func (e UpdateReleasesJobExecutor) Execute() error {
	e.logStart()
	var releasesUpdated int
	now := time.Now()
	// Start at 30 days in the past to allow users to view past releases
	startDate := now.Add(-thirtyDays)
	expiry := now.Add(3 * time.Hour)

	// Trakt's limit is 30 days per request, but we want 9 * 30 days
	for i := 0; i < 9; i++ {

		traktReleases, err := e.trakt.GetSeasonPremieres(startDate, traktDaysCap)
		if err != nil {
			return fmt.Errorf("%v whilst fetching trakt season-premieres: %w", defaultErrorMessage, err)
		}

		for _, traktRelease := range traktReleases {
			if e.hasRelevantIds(traktRelease) {
				tmdbShow, err := e.tmdb.GetTVDetails(traktRelease.TmdbId(),
					map[string]string{"append_to_response": "translations"})
				if err != nil {
					// On rare occasions trakt's tmdb-id might be incorrect
					// We treat this case as if jobs#hasRelevantIds was false
					e.Warn("Incorrect tmdb-id [%v]: %v", traktRelease.Ids(), err)
					continue
				}

				if e.hasRelevantInfo(tmdbShow) {
					if err = e.store.SaveRelease(domain.ReleaseRef{
						ShowId:       traktRelease.TmdbId(),
						SeasonNumber: traktRelease.SeasonNumber(),
						AirDate:      traktRelease.AirDate(),
					}, expiry); err != nil {
						return fmt.Errorf("%v whilst saving release [%v, %d, %v]: %w", defaultErrorMessage,
							traktRelease.TmdbId(), traktRelease.SeasonNumber(), traktRelease.AirDate(), err)
					}
					releasesUpdated++
				}
			}
		}

		if i == 0 {
			// The first iteration takes care of all past releases
			if err := e.store.SetPastReleasesCount(releasesUpdated); err != nil {
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

func (e UpdateReleasesJobExecutor) hasRelevantIds(release trakt.SeasonPremieresDto) bool {
	ids := release.Show.Ids
	return ids.Tmdb != 0 && ids.Tvdb != 0 && ids.Imdb != "" && ids.Slug != ""
}

func (e UpdateReleasesJobExecutor) hasRelevantInfo(show *tmdb.TVDetails) bool {
	return len(show.Genres) > 0 &&
		len(show.Networks) > 0 &&
		len(show.Overview) > 0 &&
		len(show.Seasons) > 0 &&
		e.hasEnglishTranslation(show.Translations) &&
		e.hasRelevantType(show)
}

func (e UpdateReleasesJobExecutor) hasEnglishTranslation(translations *tmdb.TVTranslations) bool {
	for _, translation := range translations.Translations {
		if translation.Iso639_1 == "en" {
			return true
		}
	}
	return false
}

func (e UpdateReleasesJobExecutor) hasRelevantType(show *tmdb.TVDetails) bool {
	t := show.Type
	if t == "Scripted" || t == "Miniseries" || t == "Documentary" {
		return true
	}
	if t != "Reality" && t != "News" && t != "Talk Show" && t != "Video" {
		e.Warn("Unknown type [%v] detected for show [%d, %v]", show.Type, show.ID, show.Name)
	}
	return false
}

type UpdateReleasesJobExecutor struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
	util.Logger
}

func (e UpdateReleasesJobExecutor) logStart() {
	e.Info("Running update-releases job")
}

func (e UpdateReleasesJobExecutor) logEnd(actions int) error {
	e.Info("Completed update-releases job with %d releases updated", actions)
	return nil
}
