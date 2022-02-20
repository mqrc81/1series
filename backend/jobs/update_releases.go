package jobs

import (
	"fmt"
	"log"
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/trakt"
)

const (
	negativeTwoWeeks    = -2 * 7 * 24 * time.Hour
	EightMonthsInDays   = 8 * 30
	defaultErrorMessage = "error executing update-releases job"
)

func (e UpdateReleasesJobExecutor) Execute() error {
	e.logStart()
	var releasesUpdated, pastReleases int
	now := time.Now()
	expiry := now.Add(20 * time.Hour)

	traktReleases, err := e.trakt.GetSeasonPremieres(now.Add(negativeTwoWeeks), EightMonthsInDays)
	if err != nil {
		return fmt.Errorf("%v whilst fetching trakt season-premieres: %w", defaultErrorMessage, err)
	}

	for _, traktRelease := range traktReleases {
		if hasRelevantIds(traktRelease) {
			tmdbShow, err := e.tmdb.GetTVDetails(traktRelease.TmdbId(),
				map[string]string{"append_to_response": "translations"})
			if err != nil {
				return fmt.Errorf("%v whilst fetching tmdb show-details for [%v]: %w", defaultErrorMessage,
					traktRelease.SlugId(), err)
			}

			if hasRelevantInfo(tmdbShow) {
				if err = e.store.SaveRelease(domain.ReleaseRef{
					ShowId:       traktRelease.TmdbId(),
					SeasonNumber: traktRelease.SeasonNumber(),
					AirDate:      traktRelease.AirDate(),
				}, expiry); err != nil {
					return fmt.Errorf("%v whilst saving release [%v, %d, %v]: %w", defaultErrorMessage,
						traktRelease.TmdbId(), traktRelease.SeasonNumber(), traktRelease.AirDate(), err)
				}

				if traktRelease.AirDate().Before(now) {
					pastReleases++
				}
				releasesUpdated++
			}
		}
	}

	if err = e.store.SetPastReleasesCount(pastReleases); err != nil {
		return fmt.Errorf("%v: %w", defaultErrorMessage, err)
	}

	if err = e.store.ClearExpiredReleases(now); err != nil {
		return fmt.Errorf("%v: %w", defaultErrorMessage, err)
	}

	return e.logEnd(releasesUpdated)
}

func hasRelevantIds(release trakt.SeasonPremieresDto) bool {
	ids := release.Show.Ids
	return ids.Tmdb != 0 && ids.Tvdb != 0 && ids.Imdb != "" && ids.Slug != ""
}

func hasRelevantInfo(show *tmdb.TVDetails) bool {
	return len(show.Genres) > 0 &&
		len(show.Networks) > 0 &&
		len(show.Overview) > 0 &&
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
	if t != "Reality" && t != "News" && t != "Talk Show" {
		log.Printf("Unknown type [%v] detected for show [%d, %v]\n", show.Type, show.ID, show.Name)
	}
	return false
}

type UpdateReleasesJobExecutor struct {
	store domain.Store
	tmdb  *tmdb.Client
	trakt *trakt.Client
}

func (e UpdateReleasesJobExecutor) logStart() {
	log.Println("Running update-releases job")
}

func (e UpdateReleasesJobExecutor) logEnd(actions int) error {
	log.Printf("Completed update-releases job with %d releases updated\n", actions)
	return nil
}
