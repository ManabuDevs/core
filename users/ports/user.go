package users

import users "sabasy/users/domain"

type UserRepositories interface {
	GetUsers() ([][]string, error)
	CreateUser(*users.User) (*users.User, error)
}
