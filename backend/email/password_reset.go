package email

import (
	"fmt"
	"github.com/mqrc81/1series/domain"
	"github.com/mqrc81/1series/env"
	"github.com/mqrc81/1series/logger"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/url"
)

type PasswordResetEmail struct {
	Recipient domain.User
	Token     domain.Token
}

func (e PasswordResetEmail) create(sender *mail.Email) *mail.SGMailV3 {
	logger.Warning("PasswordResetEmail only implemented temporarily")
	return mail.NewV3MailInit(sender, e.subject(), e.recipient(), e.content())
}

func (e PasswordResetEmail) subject() string {
	return "Reset your password"
}

func (e PasswordResetEmail) recipient() *mail.Email {
	return &mail.Email{
		Name:    e.Recipient.Username,
		Address: e.Recipient.Email,
	}
}

func (e PasswordResetEmail) content() *mail.Content {
	params := url.Values{}
	params.Add("token", e.Token.Id)
	return &mail.Content{
		Type: "text/plain",
		Value: fmt.Sprintf("Hello %v,\n"+
			"Use the following link to reset your password:\n"+
			"%v",
			e.Recipient.Username, env.Config().FrontendUrl+"/users/resetPassword?"+params.Encode()),
	}
}
