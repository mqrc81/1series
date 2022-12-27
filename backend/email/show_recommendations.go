package email

import (
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ShowRecommendationsEmail struct {
	Recipient domain.User
	Shows     []domain.Show
}

func (email ShowRecommendationsEmail) create(sender *mail.Email) *mail.SGMailV3 {
	logger.Warning("ShowRecommendationsEmail not implemented yet")
	return nil
}
