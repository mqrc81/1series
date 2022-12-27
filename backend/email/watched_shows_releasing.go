package email

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type WatchedShowsReleasingEmail struct {
	Recipient domain.User
	Releases  []domain.Release
}

func (e WatchedShowsReleasingEmail) create(sender *mail.Email) *mail.SGMailV3 {
	logger.Warning("WatchedShowsReleasingEmail only implemented temporarily")
	return mail.NewV3MailInit(sender, e.subject(), e.recipient(), e.content())
}

func (e WatchedShowsReleasingEmail) subject() string {
	var showsSubject string
	for i := 0; i < 3 && i < len(e.Releases); i++ {
		showsSubject += e.Releases[i].Show.Name + ", "
	}
	return fmt.Sprintf("New Seasons from %v and more!", showsSubject)
}

func (e WatchedShowsReleasingEmail) recipient() *mail.Email {
	return &mail.Email{
		Name:    e.Recipient.Username,
		Address: e.Recipient.Email,
	}
}

func (e WatchedShowsReleasingEmail) content() *mail.Content {
	var releasesContent string
	for _, release := range e.Releases {
		releasesContent += fmt.Sprintf("\t- %v: %v (airing %v)\n",
			release.Show.Name, release.Season.Name, release.AirDate.Format("Monday at 15:04 UTC"))
	}
	return &mail.Content{
		Type: "text/plain",
		Value: fmt.Sprintf("Hello %v, this text content is only a temporary until I (the developer) make it prettier.\n"+
			"The following series on your watchlist are releasing a new season withing the next week:\n"+
			"%v",
			e.Recipient.Username, releasesContent),
	}
}
