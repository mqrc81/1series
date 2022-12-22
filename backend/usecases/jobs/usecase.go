package jobs

import (
	"github.com/go-co-op/gocron"
)

type UseCase interface {
	RunJobsByTag(tag string) (int, error)
}

type useCase struct {
	scheduler *gocron.Scheduler
}

func NewUseCase(
	scheduler *gocron.Scheduler,
) UseCase {
	return &useCase{
		scheduler,
	}
}
