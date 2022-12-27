package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"github.com/mqrc81/zeries/usecases/shows"
	"sort"
)

const (
	defaultRating                = 6
	trackedShowRatingWeight      = 1
	recommendationPositionWeight = 1
)

func (job notifyUsersAboutRecommendationsJob) name() string {
	return "NOTIFY-USERS-ABOUT-RECOMMENDATIONS job"
}

func (job notifyUsersAboutRecommendationsJob) execute() error {
	logger.Info("Running %v", job.name())

	var usersNotified int

	users, err := job.userRepository.FindAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if !user.NotificationOptions.Recommendations {
			continue
		}
		trackedShows, err := job.trackedShowRepository.FindAllByUser(user)
		if err != nil {
			return err
		}
		recommendedShowsMap := make(map[int]int)
		for _, trackedShow := range trackedShows {
			recommendations, err := job.tmdbClient.GetTVRecommendations(trackedShow.ShowId, nil)
			if err != nil {
				return err
			}
			rating := trackedShow.Rating
			if rating == 0 {
				rating = defaultRating
			}
			for i := 0; i < 3 && i < len(recommendations.Results); i++ {
				recommendation := recommendations.Results[i]
				recommendedShowsMap[int(recommendation.ID)] += trackedShowRatingWeight*rating + recommendationPositionWeight*(3-i)
			}
		}
		for _, trackedShow := range trackedShows {
			delete(recommendedShowsMap, trackedShow.ShowId)
		}
		if len(recommendedShowsMap) < 3 {
			continue
		}
		recommendedShowsPairs := mapToSlice(recommendedShowsMap)
		sort.Slice(recommendedShowsPairs, func(i, j int) bool {
			return recommendedShowsPairs[i].val > recommendedShowsPairs[j].val
		})
		var recommendedShows []domain.Show
		for _, recommendedShowPair := range recommendedShowsPairs[:3] {
			tmdbShow, err := job.tmdbClient.GetTVDetails(recommendedShowPair.key, nil)
			if err != nil {
				return err
			}
			recommendedShows = append(recommendedShows, shows.ShowFromTmdbShow(tmdbShow))
		}
		emailData := email.ShowRecommendationsEmail{
			Recipient: user,
			Shows:     recommendedShows,
		}
		if err = job.emailClient.Send(emailData); err != nil {
			return err
		}
		usersNotified++
	}

	logger.Info("Completed %v with %d users notified", job.name(), usersNotified)
	return nil
}

func mapToSlice(in map[int]int) []pair {
	pairs := make([]pair, len(in))
	i := 0
	for k, v := range in {
		pairs[i].key = k
		pairs[i].val = v
		i++
	}
	return pairs
}

type pair struct {
	key int
	val int
}

type notifyUsersAboutRecommendationsJob struct {
	userRepository        repositories.UserRepository
	trackedShowRepository repositories.TrackedShowRepository
	tmdbClient            *tmdb.Client
	emailClient           *email.Client
}
