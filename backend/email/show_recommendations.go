package email

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ShowRecommendationsEmail struct {
	Recipient domain.User
	Shows     []domain.Show
}

func (e ShowRecommendationsEmail) create(sender *mail.Email) *mail.SGMailV3 {
	logger.Warning("ShowRecommendationsEmail only implemented temporarily")
	return mail.NewV3MailInit(sender, e.subject(), e.recipient(), e.content())
}

func (e ShowRecommendationsEmail) subject() string {
	return "TV-Series Recommendations based on your Watchlist"
}

func (e ShowRecommendationsEmail) recipient() *mail.Email {
	return &mail.Email{
		Name:    e.Recipient.Username,
		Address: e.Recipient.Email,
	}
}

func (e ShowRecommendationsEmail) content() *mail.Content {
	var recommendationsContent string
	for i, show := range e.Shows {
		recommendationsContent += fmt.Sprintf("\t%v. %v: %.1f/10 (%v)\n", i+1, show.Name, show.Rating, show.Poster)
	}
	return &mail.Content{
		Type: "text/plain",
		Value: fmt.Sprintf("Hello %v, this text content is only a temporary until I (the developer) make it prettier.\n"+
			"The following series are recommended for your next binge-watching session, based entirely on your watchlist and ratings:\n"+
			"%v",
			e.Recipient.Username, recommendationsContent),
	}
}
