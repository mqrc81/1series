package repositories

import (
	"github.com/mqrc81/zeries/domain"
)

func MockNetworksRepository(networks ...domain.Network) NetworksRepository {
	data := make(map[int]*domain.Network)
	for _, network := range networks {
		data[network.Id] = &network
	}
	return &mockNetworksRepository{data}
}

type mockNetworksRepository struct {
	data map[int]*domain.Network
}

func (mock *mockNetworksRepository) FindAll() ([]domain.Network, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockNetworksRepository) Save(network domain.Network) error {
	// TODO implement me
	panic("implement me")
}
