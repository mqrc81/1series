package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type networkRepository struct {
	*sqlx.DB
}

func (r *networkRepository) FindAll() (networks []domain.Network, err error) {

	if err = r.Select(
		&networks,
		`SELECT n.* FROM networks n`,
	); err != nil {
		err = fmt.Errorf("error finding networks: %w", err)
	}

	return networks, err
}

func (r *networkRepository) Save(network domain.Network) (err error) {

	if _, err = r.Exec(
		`INSERT INTO networks(id, name, logo) VALUES ($1, $2, $3)`,
		network.Id,
		network.Name,
		network.Logo,
	); err != nil {
		err = fmt.Errorf("error adding network [%v, %v, %v]: %w", network.Id, network.Name, network.Logo, err)
	}

	return err
}
