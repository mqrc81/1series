package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type ReleaseStore struct {
	*sqlx.DB
}

func (s *ReleaseStore) GetReleases(amount int, offset int) (releases []domain.ReleaseRef, err error) {

	if err = s.Select(
		&releases,
		"SELECT r.show_id, r.season_number, r.air_date FROM releases r LIMIT $1 OFFSET $2",
		amount,
		offset,
	); err != nil {
		err = fmt.Errorf("error getting releases: %w", err)
	}

	return releases, err
}

func (s *ReleaseStore) SaveRelease(release domain.ReleaseRef, expiry time.Time) (err error) {

	if _, err = s.Exec(
		// Insert if (show_id, season_number) doesn't exist; else update
		"INSERT INTO releases(show_id, season_number, air_date, expiry) VALUES($1, $2, $3, $4) ON CONFLICT (show_id, season_number) DO UPDATE SET air_date = $3, expiry = $4",
		release.ShowId,
		release.SeasonNumber,
		release.AirDate,
		expiry,
	); err != nil {
		err = fmt.Errorf("error saving release: %w", err)
	}

	return err

}

func (s *ReleaseStore) ClearExpiredReleases(now time.Time) (err error) {

	if _, err = s.Exec(
		"DELETE FROM releases r WHERE r.expiry < $1",
		now,
	); err != nil {
		err = fmt.Errorf("error clearing expired releases: %w", err)
	}

	return err
}

func (s *ReleaseStore) SetPastReleasesCount(amount int) (err error) {

	if _, err = s.Exec(
		"UPDATE past_releases SET amount = $1",
		amount,
	); err != nil {
		err = fmt.Errorf("error setting past releases count: %w", err)
	}

	return err
}

func (s *ReleaseStore) GetPastReleasesCount() (amount int, err error) {

	if err = s.Get(
		&amount,
		"SELECT ps.amount FROM past_releases ps LIMIT 1",
	); err != nil || amount == 0 {
		err = fmt.Errorf("error getting past releases count: %w", err)
	}

	return amount, err
}
