package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type tokensRepository struct {
	*sql.Database
}

func (r *tokensRepository) FindByTokenId(tokenId string) (token domain.Token, err error) {

	if err = r.Get(
		&token,
		`SELECT t.token_id, t.user_id, t.purpose, t.expires_at FROM tokens t WHERE t.token_id = $1`,
		tokenId,
	); err != nil {
		err = fmt.Errorf("error finding token [%v]: %w", tokenId, err)
	}

	return token, err
}

func (r *tokensRepository) Save(token domain.Token) (err error) {

	if _, err = r.Exec(
		`INSERT INTO tokens(token_id, user_id, purpose, expires_at) VALUES ($1, $2, $3, $4)`,
		token.TokenId,
		token.UserId,
		token.Purpose,
		token.ExpiresAt,
	); err != nil {
		err = fmt.Errorf("error saving token [%v]: %w", token, err)
	}

	return err
}

func (r *tokensRepository) Delete(token domain.Token) (err error) {

	if _, err = r.Exec(
		`DELETE FROM tokens t WHERE t.token_id = $1`,
		token.TokenId,
	); err != nil {
		err = fmt.Errorf("error deleting token [%v]: %w", token, err)
	}

	return err
}
