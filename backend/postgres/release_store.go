package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type ReleaseStore struct {
	*sqlx.DB
}

func (s *ReleaseStore) GetReleases(amount int, offset int) (releases []domain.ReleaseRef, err error) {

	if err = s.Select(
		&releases,
		"SELECT r.* FROM releases r LIMIT $1 OFFSET $2",
		amount,
		offset,
	); err != nil {
		err = fmt.Errorf("error getting releases: %w", err)
	}

	return releases, err
}

func (s *ReleaseStore) SaveRelease(release domain.ReleaseRef) (err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ReleaseStore) ClearExpiredReleases() (err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ReleaseStore) SetPastReleasesCount(amount int) (err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ReleaseStore) GetPastReleasesCount() (amount int, err error) {
	// TODO implement me
	panic("implement me")
}
