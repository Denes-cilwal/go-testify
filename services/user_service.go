package services

import "go-testify/repository"

type UserService struct {
	repository repository.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(userrepo repository.UserRepository) UserService {
	return UserService{
		repository: userrepo,
	}
}
