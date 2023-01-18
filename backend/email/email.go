package email

import (
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Client struct {
	client *sendgrid.Client
	sender *mail.Email
}

func (c *Client) Send(data emailData) error {
	_, err := c.client.Send(data.create(c.sender))
	return err
}

func NewEmailClient(apiKey string, senderEmail string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("sendgrid api key is empty")
	}
	client := sendgrid.NewSendClient(apiKey)
	sender := &mail.Email{
		Name:    "1Series",
		Address: senderEmail,
	}
	return &Client{client, sender}, nil
}

type emailData interface {
	create(sender *mail.Email) *mail.SGMailV3
}
