package jobs

import (
	"errors"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/usecases/shows"
	"time"
)

const (
	sevenDays = time.Hour * 24 * 7
)

func (job notifyUsersAboutReleasesJob) name() string {
	return "NOTIFY-USERS-ABOUT-RELEASES job"
}

func (job notifyUsersAboutReleasesJob) execute() error {
	logger.Info("Running %v", job.name())

	var usersNotified int

	upcomingReleasesMap, err := job.fetchReleasesAiringWithinTheNextWeek()
	if err != nil {
		return err
	}

	users, err := job.userRepository.FindAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		upcomingWatchedShowReleases, err := job.findUpcomingWatchedShowReleasesOfUser(user, upcomingReleasesMap)
		if err != nil {
			return err
		}
		if len(upcomingWatchedShowReleases) > 0 {
			emailData := email.WatchedShowsReleasingEmail{
				Recipient: user,
				Releases:  upcomingWatchedShowReleases,
			}
			if err = job.emailClient.Send(emailData); err != nil {
				return err
			}
			usersNotified++
		}
	}

	logger.Info("Completed %v with %d users notified", job.name(), usersNotified)
	return nil
}

func (job notifyUsersAboutReleasesJob) findUpcomingWatchedShowReleasesOfUser(
	user domain.User, upcomingReleasesMap map[int]*domain.Release,
) ([]domain.Release, error) {
	var upcomingWatchedShowReleases []domain.Release

	watchedShows, err := job.watchedShowRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}
	for _, watchedShow := range watchedShows {
		if upcomingReleasesMap[watchedShow.ShowId] != nil {
			upcomingWatchedShowReleases = append(upcomingWatchedShowReleases, *upcomingReleasesMap[watchedShow.ShowId])
		}
	}
	return upcomingWatchedShowReleases, nil
}

func (job notifyUsersAboutReleasesJob) fetchReleasesAiringWithinTheNextWeek() (map[int]*domain.Release, error) {
	var (
		today               = atBeginningOfDay(time.Now().UTC())
		nextWeek            = today.Add(sevenDays)
		upcomingReleasesMap = make(map[int]*domain.Release)
	)

	releasesInTheNextWeek, err := job.releaseRepository.FindAllAiringBetween(today, nextWeek)
	if err != nil {
		return nil, err
	} else if len(releasesInTheNextWeek) == 0 {
		return nil, errors.New("no releases airing within the next week found in the database")
	}
	for _, releaseRef := range releasesInTheNextWeek {
		tmdbRelease, err := job.tmdbClient.GetTVDetails(releaseRef.ShowId, map[string]string{"append_to_response": "translations"})
		if err != nil {
			return nil, err
		}
		show := shows.ReleaseFromTmdbShow(tmdbRelease, releaseRef.SeasonNumber, releaseRef.AirDate, releaseRef.AnticipationLevel)
		upcomingReleasesMap[releaseRef.ShowId] = &show
	}
	return upcomingReleasesMap, nil
}

func atBeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

type notifyUsersAboutReleasesJob struct {
	userRepository        repositories.UserRepository
	releaseRepository     repositories.ReleaseRepository
	watchedShowRepository repositories.WatchedShowRepository
	tmdbClient            *tmdb.Client
	emailClient           *email.Client
}
