package repositories

import (
	"github.com/mqrc81/1series/domain"
)

func MockTokensRepository(tokens ...domain.Token) TokensRepository {
	data := make(map[string]*domain.Token)
	for _, token := range tokens {
		data[token.Id] = &token
	}
	return &mockTokensRepository{data}
}

type mockTokensRepository struct {
	data map[string]*domain.Token
}

func (m *mockTokensRepository) Find(tokenId string) (domain.Token, error) {
	// TODO implement me
	panic("implement me")
}

func (m *mockTokensRepository) SaveOrReplace(token domain.Token) error {
	// TODO implement me
	panic("implement me")
}

func (m *mockTokensRepository) Delete(token domain.Token) error {
	// TODO implement me
	panic("implement me")
}

func (m *mockTokensRepository) DeleteByUserIdAndPurpose(userId int, purpose domain.TokenPurpose) error {
	// TODO implement me
	panic("implement me")
}
