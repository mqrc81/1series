package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type releaseRepository struct {
	*sqlx.DB
}

func (r *releaseRepository) FindAllInRange(amount int, offset int) (releases []domain.ReleaseRef, err error) {

	if err = r.Select(
		&releases,
		`SELECT r.show_id, r.season_number, r.air_date, r.anticipation_level FROM releases r ORDER BY r.air_date LIMIT $1 OFFSET $2`,
		amount,
		offset,
	); err != nil {
		err = fmt.Errorf("error finding releases: %w", err)
	}

	return releases, err
}

func (r *releaseRepository) ReplaceAll(releases []domain.ReleaseRef, pastReleasesCount int) error {

	txn, err := r.Beginx()
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

	if err = r.deleteAllReleasesInTransaction(txn); err != nil {
		return err
	}

	for _, release := range releases {
		if err = r.saveReleaseInTransaction(txn, release); err != nil {
			return err
		}
	}

	if err = r.savePastReleasesCountInTransaction(txn, pastReleasesCount); err != nil {
		return err
	}

	return err
}

func (r *releaseRepository) saveReleaseInTransaction(txn *sqlx.Tx, release domain.ReleaseRef) (err error) {
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

func (r *releaseRepository) savePastReleasesCountInTransaction(txn *sqlx.Tx, pastReleasesCount int) (err error) {
	if _, err = txn.Exec(
		`UPDATE past_releases SET amount = $1 WHERE past_releases_id = 69`,
		pastReleasesCount,
	); err != nil {
		err = fmt.Errorf("error setting past releases count: %w", err)
	}
	return err
}

func (r *releaseRepository) deleteAllReleasesInTransaction(txn *sqlx.Tx) (err error) {
	//goland:noinspection SqlWithoutWhere
	if _, err = txn.Exec(`DELETE FROM releases`); err != nil {
		err = fmt.Errorf("error deleting releases: %w", err)
	}
	return err
}

func (r *releaseRepository) CountPastReleases() (amount int, err error) {

	if err = r.Get(
		&amount,
		`SELECT ps.amount FROM past_releases ps`,
	); err != nil || amount == 0 {
		err = fmt.Errorf("error counting past releases: %w", err)
	}

	return amount, err
}
