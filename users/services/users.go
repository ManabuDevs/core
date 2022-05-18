package users

import (
	domain "sabasy/users/domain"
	usersPorts "sabasy/users/ports"
)

type serviceUser struct {
	repoUser usersPorts.UserRepositories
}

func NewUserService(repo usersPorts.UserRepositories) *serviceUser {
	return &serviceUser{
		repoUser: repo,
	}
}

func (s *serviceUser) CreateUser(user *domain.User) (*domain.User, error) {
	return s.repoUser.CreateUser(user)
}

func (s *serviceUser) GetUsers() ([][]string, error) {
	return s.repoUser.GetUsers()
}

func (s *serviceUser) GetUserByID(id string) ([][]string, error) {
	return s.repoUser.GetUserByID(id)
}
