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

	if err = s.Get(
		&networks,
		"SELECT n.* FROM networks n",
	); err != nil {
		err = fmt.Errorf("error getting network: %w", err)
	}

	return networks, err
}
