package users

import users "sabasy/users/domain"

type UserRepositories interface {
	GetUsers() ([][]string, error)
	CreateUser(*users.User) (*users.User, error)
	GetUserByID(id string) ([][]string, error)
	DeleteUserByID(id string) (bool, error)
}
