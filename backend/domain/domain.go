// Package domain
// Domain models
package domain

import (
	"time"
)

type Show struct {
	Id            int // Tmdb Id
	Name          string
	Description   string
	Year          int
	Poster        string
	Rating        int
	Homepage      string
	SeasonsCount  int
	EpisodesCount int
	Genres        []Genre
	Networks      []Network
}

type Season struct {
	Id            int // Tmdb Id
	Number        string
	Name          string
	Description   string
	Poster        string
	EpisodesCount int
	ShowId        int // Tmdb Id
}

type Release struct {
	Show         Show
	SeasonNumber int
	AirDate      time.Time
}

type Genre struct {
	Id   int // Tmdb Id
	Name string
}

type Network struct {
	Id   int // Tmdb Id
	Name string
	Logo string
}

type User struct {
	Username         string
	Email            string
	Password         string
	NotificationType NotificationType
}

type NotificationType int

const (
	None NotificationType = iota
	ReleasesOnly
	RecommendationsOnly
	All
)
