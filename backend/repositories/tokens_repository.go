package repositories

import (
	"fmt"
	"github.com/mqrc81/1series/domain"
	"github.com/mqrc81/1series/sql"
)

type tokensRepository struct {
	*sql.Database
}

func (r *tokensRepository) Find(tokenId string) (token domain.Token, err error) {

	if err = r.Get(
		&token,
		`SELECT t.id, t.user_id, t.purpose, t.expires_at FROM tokens t WHERE t.id = $1`,
		tokenId,
	); err != nil {
		err = fmt.Errorf("error finding token [%v]: %w", tokenId, err)
	}

	return token, err
}

func (r *tokensRepository) SaveOrReplace(token domain.Token) (err error) {

	if err = r.DeleteByUserIdAndPurpose(token.UserId, token.Purpose); err != nil {
		return err
	}

	if _, err = r.Exec(
		`INSERT INTO tokens(id, user_id, purpose, expires_at) VALUES ($1, $2, $3, $4);`,
		token.Id,
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
		`DELETE FROM tokens t WHERE t.id = $1`,
		token.Id,
	); err != nil {
		err = fmt.Errorf("error deleting token [%v]: %w", token, err)
	}

	return err
}

func (r *tokensRepository) DeleteByUserIdAndPurpose(userId int, purpose domain.TokenPurpose) (err error) {

	if _, err = r.Exec(
		`DELETE FROM tokens t WHERE t.user_id = $1 AND t.purpose = $2`,
		userId,
		purpose,
	); err != nil {
		err = fmt.Errorf("error deleting token [%v, %v]: %w", userId, purpose, err)
	}

	return err
}
