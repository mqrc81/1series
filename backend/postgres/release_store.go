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
		`SELECT r.show_id, r.season_number, r.air_date, r.anticipation_level FROM releases r ORDER BY r.air_date LIMIT $1 OFFSET $2`,
		amount,
		offset,
	); err != nil {
		err = fmt.Errorf("error getting releases: %w", err)
	}

	return releases, err
}

func (s *ReleaseStore) SaveReleases(releases []domain.ReleaseRef, pastReleasesCount int) error {

	txn, err := s.Beginx()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	defer func() {
		if err == nil {
			err = txn.Commit()
		} else {
			_ = txn.Rollback()
		}
	}()

	if err = s.deleteAllReleasesInTransaction(txn); err != nil {
		return err
	}

	for _, release := range releases {
		if err = s.saveReleaseInTransaction(txn, release); err != nil {
			return err
		}
	}

	if err = s.savePastReleasesCountInTransaction(txn, pastReleasesCount); err != nil {
		return err
	}

	return err
}

func (s *ReleaseStore) saveReleaseInTransaction(txn *sqlx.Tx, release domain.ReleaseRef) (err error) {
	if _, err = txn.Exec(`INSERT INTO releases(show_id, season_number, air_date, anticipation_level) VALUES($1, $2, $3, $4)`,
		release.ShowId,
		release.SeasonNumber,
		release.AirDate,
		release.AnticipationLevel,
	); err != nil {
		err = fmt.Errorf("error saving release: %w", err)
	}
	return err
}

func (s *ReleaseStore) savePastReleasesCountInTransaction(txn *sqlx.Tx, pastReleasesCount int) (err error) {
	if _, err = txn.Exec(
		`UPDATE past_releases SET amount = $1 WHERE past_releases_id = 69`,
		pastReleasesCount,
	); err != nil {
		err = fmt.Errorf("error setting past releases count: %w", err)
	}
	return err
}

func (s *ReleaseStore) deleteAllReleasesInTransaction(txn *sqlx.Tx) (err error) {
	//goland:noinspection SqlWithoutWhere
	if _, err = txn.Exec(`DELETE FROM releases`); err != nil {
		err = fmt.Errorf("error deleting releases: %w", err)
	}
	return err
}

func (s *ReleaseStore) GetPastReleasesCount() (amount int, err error) {

	if err = s.Get(
		&amount,
		`SELECT ps.amount FROM past_releases ps`,
	); err != nil || amount == 0 {
		err = fmt.Errorf("error getting past releases count: %w", err)
	}

	return amount, err
}
