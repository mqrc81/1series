package controllers

import (
	"github.com/mqrc81/zeries/domain"
)

type popularShowsDto struct {
	NextPage int           `json:"nextPage,omitempty"`
	Shows    []domain.Show `json:"shows"`
}

type upcomingReleasesDto struct {
	PreviousPage int              `json:"previousPage,omitempty"`
	NextPage     int              `json:"nextPage,omitempty"`
	Releases     []domain.Release `json:"releases"`
}
