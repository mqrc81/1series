package controllers

import (
	"github.com/mqrc81/zeries/domain"
)

type popularShowsDto struct {
	NextPage int           `json:"nextPage"`
	Shows    []domain.Show `json:"shows"`
}
