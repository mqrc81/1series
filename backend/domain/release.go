package domain

import (
	"time"
)

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

type AnticipationLevel int

const (
	Zero     AnticipationLevel = iota
	Moderate                   // top 10
	High                       // top 3
	Extreme                    // top 1
)
