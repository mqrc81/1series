package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
	"time"
)

type networksRepository struct {
	*sql.Database
}

func (r *networksRepository) FindAll() (networks []domain.Network, err error) {

	if err = r.Select(
		&networks,
		`SELECT n.* FROM networks n`,
	); err != nil {
		err = fmt.Errorf("error finding networks: %w", err)
	}

	return networks, err
}

func (r *networksRepository) Save(network domain.Network) (err error) {

	if _, err = r.Exec(
		`INSERT INTO networks(network_id, name, logo, created_at) VALUES ($1, $2, $3, $4)`,
		network.NetworkId,
		network.Name,
		network.Logo,
		time.Now(),
	); err != nil {
		err = fmt.Errorf("error adding network [%v]: %w", network, err)
	}

	return err
}
