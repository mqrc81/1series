package repositories

import (
	"errors"
	"github.com/mqrc81/zeries/domain"
)

func MockUsersRepository(users ...domain.User) UsersRepository {
	data := make(map[int]*domain.User)
	for _, user := range users {
		data[user.Id] = &user
	}
	return &mockUsersRepository{data}
}

type mockUsersRepository struct {
	data map[int]*domain.User
}

func (mock *mockUsersRepository) Find(userId int) (domain.User, error) {
	user := mock.data[userId]
	if user == nil {
		return domain.User{}, errors.New("user not found")
	}
	return *user, nil
}

func (mock *mockUsersRepository) FindAll() ([]domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockUsersRepository) FindByUsername(username string) (domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockUsersRepository) FindByEmail(email string) (domain.User, error) {
	// TODO implement me
	panic("implement me")
}

func (mock *mockUsersRepository) Save(user domain.User) error {
	// TODO implement me
	panic("implement me")
}
