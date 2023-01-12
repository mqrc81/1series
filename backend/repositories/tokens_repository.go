package repositories

import (
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type tokensRepository struct {
	*sql.Database
}

func (r *tokensRepository) FindByToken(token string) (domain.Token, error) {
	// TODO implement me
	panic("implement me")
}

func (r *tokensRepository) Save(token domain.Token) error {
	// TODO implement me
	panic("implement me")
}

func (r *tokensRepository) Delete(token domain.Token) error {
	// TODO implement me
	panic("implement me")
}
