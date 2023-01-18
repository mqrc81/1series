package repositories

import (
	"github.com/mqrc81/1series/domain"
	"time"
)

func MockReleasesRepository(releases ...domain.ReleaseRef) ReleasesRepository {
	data := make(map[int]*domain.ReleaseRef)
	for _, release := range releases {
		data[release.ShowId] = &release
	}
	return &mockReleasesRepository{data}
}

type mockReleasesRepository struct {
	data map[int]*domain.ReleaseRef
}

func (mock *mockReleasesRepository) FindAllInRange(amount int, offset int) ([]domain.ReleaseRef, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockReleasesRepository) FindAllAiringBetween(
	startDate time.Time,
	endDate time.Time,
) ([]domain.ReleaseRef, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockReleasesRepository) ReplaceAll(releases []domain.ReleaseRef, pastReleasesCount int) error {
	// TODO implement me
	panic("implement me")
}

func (mock *mockReleasesRepository) CountPastReleases() (int, error) {
	// TODO implement me
	panic("implement me")
}
