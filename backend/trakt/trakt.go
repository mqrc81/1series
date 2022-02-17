// Package trakt
// Interface to Trakt's API (see https://trakt.docs.apiary.io/)
// Unlike Tmdb there is no actively maintained Go Trakt library
package trakt

import (
	"errors"
	"net/http"
)

type Client struct {
	apiKey string
	http   http.Client
}

func Init(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("trakt api key is empty")
	}
	return &Client{
		apiKey: apiKey,
		http:   http.Client{},
	}, nil
}
