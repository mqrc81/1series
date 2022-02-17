package domain

type UserStore interface {
	GetUser(userId int) (User, error)
}

type Store interface {
	UserStore
}
