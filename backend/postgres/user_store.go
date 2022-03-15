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
		`SELECT u.* FROM users u WHERE u.id = $1`,
		userId,
	); err != nil {
		err = fmt.Errorf("error getting user [%v]: %w", userId, err)
	}

	return user, err
}

func (s *UserStore) GetUserByUsername(username string) (user domain.User, err error) {

	if err = s.Get(
		&user,
		`SELECT u.* FROM users u WHERE u.username = $1`,
		username,
	); err != nil {
		err = fmt.Errorf("error getting user [%v]: %w", username, err)
	}

	return user, err
}

func (s *UserStore) GetUserByEmail(email string) (user domain.User, err error) {

	if err = s.Get(
		&user,
		`SELECT u.* FROM users u WHERE u.email = $1`,
		email,
	); err != nil {
		err = fmt.Errorf("error getting user [%v]: %w", email, err)
	}

	return user, err
}
