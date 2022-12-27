package jobs

import (
	"time"

	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/trakt"
)

const (
	thirtyDays = 30 * 24 * time.Hour
)

func (job updateReleasesJob) name() string {
	return "UPDATE-RELEASES job"
}

func (job updateReleasesJob) execute() error {
	logger.Info("Running %v", job.name())

	var (
		releases          []domain.ReleaseRef
		pastReleasesCount int
		startDate         = time.Now().Add(-thirtyDays)
	)

	traktShowsAnticipated, err := job.traktClient.GetAnticipatedShows(1, 10)
	if err != nil {
		return err
	}

	// Trakt's limit is 33 days per request, but we want 9 * 30 days
	for i := 0; i < 9; i++ {

		traktReleases, err := job.traktClient.GetSeasonPremieres(startDate, 30)
		if err != nil {
			return err
		}

		releases, err = job.filterAndCollectReleases(releases, traktReleases, traktShowsAnticipated)
		if err != nil {
			return err
		}

		if i == 0 {
			// The first iteration takes care of all past releases
			pastReleasesCount = len(releases)
		}

		startDate = startDate.Add(thirtyDays)
	}

	if err = job.releaseRepository.ReplaceAll(releases, pastReleasesCount); err != nil {
		return err
	}

	logger.Info("Completed %v with %d releases saved", job.name(), len(releases))
	return nil
}

func (job updateReleasesJob) filterAndCollectReleases(
	releases []domain.ReleaseRef, traktReleases []trakt.SeasonPremieresDto, traktShows []trakt.ShowsAnticipatedDto,
) ([]domain.ReleaseRef, error) {
	for _, traktRelease := range traktReleases {
		if !hasRelevantIds(traktRelease) {
			continue
		}

		tmdbShow, err := job.tmdbClient.GetTVDetails(traktRelease.TmdbId(), map[string]string{"append_to_response": "translations"})

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
	switch show.Type {
	case "Scripted", "Miniseries", "Documentary":
		return true
	case "Reality", "News", "Talk Show", "Video":
		return false
	default:
		logger.Warning("Unknown type [%v] detected for show [%d, %v]", show.Type, show.ID, show.Name)
		return false
	}
}

func hasMatchingSeasons(release trakt.SeasonPremieresDto, show *tmdb.TVDetails) bool {
	return len(show.Seasons) >= release.SeasonNumber()
}

func anticipationLevelFor(releaseId int, showsAnticipated []trakt.ShowsAnticipatedDto) domain.AnticipationLevel {
	for i, showAnticipated := range showsAnticipated {
		if releaseId == showAnticipated.TmdbId() {
			if i == 0 {
				return domain.Extreme
			} else if i < 3 {
				return domain.High
			} else if i < 10 {
				return domain.Moderate
			}
		}
	}
	return domain.Zero
}

type updateReleasesJob struct {
	releaseRepository repositories.ReleaseRepository
	tmdbClient        *tmdb.Client
	traktClient       *trakt.Client
}
