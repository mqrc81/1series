package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/sql"

	"github.com/mqrc81/zeries/domain"
)

type usersRepository struct {
	*sql.Database
}

func (r *usersRepository) Find(userId int) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.* FROM users u WHERE u.id = $1`,
		userId,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", userId, err)
	}

	return user, err
}

func (r *usersRepository) FindAll() (users []domain.User, err error) {

	if err = r.Select(
		&users,
		`SELECT u.* FROM users u`,
	); err != nil {
		err = fmt.Errorf("error finding users: %w", err)
	}

	return users, err
}

func (r *usersRepository) FindByUsername(username string) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.* FROM users u WHERE u.username = $1`,
		username,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", username, err)
	}

	return user, err
}

func (r *usersRepository) FindByEmail(email string) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.* FROM users u WHERE u.email = $1`,
		email,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", email, err)
	}

	return user, err
}

func (r *usersRepository) Save(user domain.User) (id int, err error) {

	if res, err := r.Exec(`INSERT INTO users(username, password, email) VALUES ($1, $2, $3)`,
		user.Username,
		user.Email,
		user.Password,
	); err != nil {
		err = fmt.Errorf("error saving user [%v]: %w", user, err)
	} else {
		id = newId(res)
	}

	return id, err
}
