package users

import users "sabasy/users/domain"

type UserServices interface {
	Get() ([][]string, error)
}

type UserRepositories interface {
	GetUsers() ([][]string, error)

	CreateUser(*users.User) (*users.User, error)
}
