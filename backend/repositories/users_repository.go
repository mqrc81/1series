package repositories

import (
	"fmt"
	"github.com/mqrc81/zeries/domain"
	"github.com/mqrc81/zeries/sql"
)

type usersRepository struct {
	*sql.Database
}

func (r *usersRepository) Find(userId int) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.id, u.username, u.email, u.password, u.email_verified, u.notify_releases, u.notify_recommendations FROM users u WHERE u.id = $1`,
		userId,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", userId, err)
	}

	return user, err
}

func (r *usersRepository) FindAll() (users []domain.User, err error) {

	if err = r.Select(
		&users,
		`SELECT u.id, u.username, u.email, u.password, u.email_verified, u.notify_releases, u.notify_recommendations FROM users u`,
	); err != nil {
		err = fmt.Errorf("error finding users: %w", err)
	}

	return users, err
}

func (r *usersRepository) FindByUsername(username string) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.id, u.username, u.email, u.password, u.email_verified, u.notify_releases, u.notify_recommendations FROM users u WHERE u.username = $1`,
		username,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", username, err)
	}

	return user, err
}

func (r *usersRepository) FindByEmail(email string) (user domain.User, err error) {

	if err = r.Get(
		&user,
		`SELECT u.id, u.username, u.email, u.password, u.email_verified, u.notify_releases, u.notify_recommendations FROM users u WHERE u.email = $1`,
		email,
	); err != nil {
		err = fmt.Errorf("error finding user [%v]: %w", email, err)
	}

	return user, err
}

func (r *usersRepository) Save(user domain.User) (err error) {

	if _, err = r.Exec(
		`INSERT INTO users(username, email, password, notify_releases, notify_recommendations) VALUES ($1, $2, $3, $4, $5)`,
		user.Username,
		user.Email,
		user.Password,
		user.NotificationOptions.Releases,
		user.NotificationOptions.Recommendations,
	); err != nil {
		err = fmt.Errorf("error saving user [%v]: %w", user, err)
	}

	return err
}

func (r *usersRepository) Update(user domain.User) (err error) {

	if _, err = r.Exec(
		`UPDATE users u SET u.username = $1, u.email = $2, u.password = $3, u.notify_releases = $4, u.notify_recommendations = $5 WHERE u.id = $6`,
		user.Username,
		user.Email,
		user.Password,
		user.NotificationOptions.Releases,
		user.NotificationOptions.Recommendations,
		user.Id,
	); err != nil {
		err = fmt.Errorf("error updating user [%v]: %w", user, err)
	}

	return err
}
