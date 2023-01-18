package repositories

import (
	"github.com/mqrc81/1series/domain"
)

func MockGenresRepository(genres ...domain.Genre) GenresRepository {
	data := make(map[int]*domain.Genre)
	for _, genre := range genres {
		data[genre.GenreId] = &genre
	}
	return &mockGenresRepository{data}
}

type mockGenresRepository struct {
	data map[int]*domain.Genre
}

func (mock *mockGenresRepository) FindAll() ([]domain.Genre, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockGenresRepository) Save(genre domain.Genre) error {
	// TODO implement me
	panic("implement me")
}

func (mock *mockGenresRepository) ReplaceAll(genres []domain.Genre) error {
	// TODO implement me
	panic("implement me")
}
