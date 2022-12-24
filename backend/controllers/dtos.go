package controllers

import (
	"github.com/mqrc81/zeries/domain"
)

type popularShowsDto struct {
	NextPage int           `json:"nextPage"`
	Shows    []domain.Show `json:"shows"`
}

type upcomingReleasesDto struct {
	PreviousPage int              `json:"previousPage"`
	NextPage     int              `json:"nextPage"`
	Releases     []domain.Release `json:"releases"`
}
