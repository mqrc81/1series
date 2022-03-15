package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type NetworkStore struct {
	*sqlx.DB
}

func (s *NetworkStore) GetNetworks() (networks []domain.Network, err error) {

	if err = s.Select(
		&networks,
		`SELECT n.* FROM networks n`,
	); err != nil {
		err = fmt.Errorf("error getting networks: %w", err)
	}

	return networks, err
}

func (s *NetworkStore) AddNetwork(network domain.Network) (err error) {

	if _, err = s.Exec(
		`INSERT INTO networks(id, name, logo) VALUES ($1, $2, $3)`,
		network.Id,
		network.Name,
		network.Logo,
	); err != nil {
		err = fmt.Errorf("error adding network [%v, %v, %v]: %w", network.Id, network.Name, network.Logo, err)
	}

	return err
}
