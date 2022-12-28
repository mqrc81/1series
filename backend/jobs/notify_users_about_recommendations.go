package jobs

import (
	"github.com/cyruzin/golang-tmdb"
	"github.com/mqrc81/zeries/controllers/shows"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/email"
	"github.com/mqrc81/zeries/logger"
	"github.com/mqrc81/zeries/repositories"
	"sort"
)

const (
	trackedShowRatingMin             = 5
	trackedShowDefaultRating         = trackedShowRatingMin + 1
	trackedShowRatingWeight          = 1
	showRecommendationPositionWeight = 1
	showRecommendationDepth          = 3
	showRecommendationDepthMax       = showRecommendationDepth * 2
	totalRecommendationsAmount       = 4
)

func (job notifyUsersAboutRecommendationsJob) name() string {
	return "NOTIFY-USERS-ABOUT-RECOMMENDATIONS job"
}

func (job notifyUsersAboutRecommendationsJob) execute() error {
	logger.Info("Running %v", job.name())

	var usersNotified int

	users, err := job.usersRepository.FindAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if !user.NotificationOptions.Recommendations {
			continue
		}
		trackedShows, err := job.trackedShowsRepository.FindAllByUser(user)
		if err != nil {
			return err
		}
		recommendedShowsMap := make(map[int]int)
		for _, trackedShow := range trackedShows {
			recommendations, err := job.tmdbClient.GetTVRecommendations(trackedShow.ShowId, nil)
			if err != nil {
				return err
			}
			if trackedShow.Rating == 0 {
				trackedShow.Rating = trackedShowDefaultRating
			}
			for i := 0; i < showRecommendationDepth && i < len(recommendations.Results); i++ {
				recommendation := recommendations.Results[i]
				recommendedShowsMap[int(recommendation.ID)] +=
					trackedShowRatingWeight*(trackedShow.Rating-trackedShowRatingMin) +
						showRecommendationPositionWeight*(showRecommendationDepth-i)
			}
		}
		for _, trackedShow := range trackedShows {
			delete(recommendedShowsMap, trackedShow.ShowId)
		}
		for key, val := range recommendedShowsMap {
			if val <= 0 {
				delete(recommendedShowsMap, key)
			}
		}
		if len(recommendedShowsMap) < totalRecommendationsAmount {
			continue
		}
		recommendedShowsPairs := mapToSlice(recommendedShowsMap)
		sort.Slice(recommendedShowsPairs, func(i, j int) bool {
			return recommendedShowsPairs[i].val > recommendedShowsPairs[j].val
		})
		var recommendedShows []domain.Show
		for _, recommendedShowPair := range recommendedShowsPairs[:totalRecommendationsAmount] {
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

type pair struct {
	key int
	val int
}

func mapToSlice(m map[int]int) []pair {
	pairs := make([]pair, len(m))
	i := 0
	for k, v := range m {
		pairs[i].key = k
		pairs[i].val = v
		i++
	}
	return pairs
}

type notifyUsersAboutRecommendationsJob struct {
	usersRepository        repositories.UsersRepository
	trackedShowsRepository repositories.TrackedShowsRepository
	tmdbClient             *tmdb.Client
	emailClient            *email.Client
}
