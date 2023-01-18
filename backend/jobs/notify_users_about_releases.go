package jobs

import (
	"errors"
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/1series/controllers/shows"
	"github.com/mqrc81/1series/domain"
	"github.com/mqrc81/1series/email"
	"github.com/mqrc81/1series/logger"
	"github.com/mqrc81/1series/repositories"
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

	users, err := job.usersRepository.FindAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if !user.NotificationOptions.Releases {
			continue
		}
		upcomingTrackedShowReleases, err := job.findUpcomingTrackedShowReleasesOfUser(user, upcomingReleasesMap)
		if err != nil {
			return err
		}
		if len(upcomingTrackedShowReleases) > 0 {
			emailData := email.TrackedShowsReleasingEmail{
				Recipient: user,
				Releases:  upcomingTrackedShowReleases,
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

func (job notifyUsersAboutReleasesJob) fetchReleasesAiringWithinTheNextWeek() (map[int]*domain.Release, error) {
	var (
		today               = atBeginningOfDay(time.Now().UTC())
		nextWeek            = today.Add(sevenDays)
		upcomingReleasesMap = make(map[int]*domain.Release)
	)

	releasesInTheNextWeek, err := job.releasesRepository.FindAllAiringBetween(today, nextWeek)
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

func (job notifyUsersAboutReleasesJob) findUpcomingTrackedShowReleasesOfUser(
	user domain.User, upcomingReleasesMap map[int]*domain.Release,
) ([]domain.Release, error) {
	var upcomingTrackedShowReleases []domain.Release

	trackedShows, err := job.trackedShowsRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}
	for _, trackedShow := range trackedShows {
		if upcomingReleasesMap[trackedShow.ShowId] != nil {
			upcomingTrackedShowReleases = append(upcomingTrackedShowReleases, *upcomingReleasesMap[trackedShow.ShowId])
		}
	}
	return upcomingTrackedShowReleases, nil
}

func atBeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

type notifyUsersAboutReleasesJob struct {
	usersRepository        repositories.UsersRepository
	releasesRepository     repositories.ReleasesRepository
	trackedShowsRepository repositories.TrackedShowsRepository
	tmdbClient             *tmdb.Client
	emailClient            *email.Client
}
