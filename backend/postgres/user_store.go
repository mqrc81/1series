package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mqrc81/zeries/domain"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) GetUser(userId int) (user domain.User, err error) {

	if err = s.Get(
		&user,
		"SELECT u.* FROM users u WHERE u.id = $1",
		userId,
	); err != nil {
		err = fmt.Errorf("error getting user: %w", err)
	}

	return user, err
}
