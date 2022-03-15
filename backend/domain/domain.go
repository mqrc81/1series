// Package domain contains all domain models
package domain

import (
	"time"
)

type Show struct {
	Id            int
	Name          string
	Overview      string
	Year          int
	Poster        string
	Rating        float32
	Homepage      string
	SeasonsCount  int
	EpisodesCount int
	Genres        []Genre
	Networks      []Network
}

type Season struct {
	ShowId        int
	Number        int
	Name          string
	Overview      string
	Poster        string
	EpisodesCount int
}

type Release struct {
	Show              Show
	Season            Season
	AirDate           time.Time
	AnticipationLevel AnticipationLevel
}

type ReleaseRef struct {
	ShowId            int               `db:"show_id"`
	SeasonNumber      int               `db:"season_number"`
	AirDate           time.Time         `db:"air_date"`
	AnticipationLevel AnticipationLevel `db:"anticipation_level"`
}

type Genre struct {
	Id   int
	Name string
}

type Network struct {
	Id   int
	Name string
	Logo string
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	NotificationOptions
}

type NotificationOptions struct {
	Releases        bool `db:"notify_releases"`
	Recommendations bool `db:"notify_recommendations"`
}

type AnticipationLevel = int

const (
	Zero   AnticipationLevel = iota
	Mid                      // top 10
	Bussin                   // top 3
	Zamn                     // top 1
)
