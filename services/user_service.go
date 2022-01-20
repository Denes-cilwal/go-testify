package services

import (
	"errors"
	"go-testify/models"
	"go-testify/repository"
)

type UserService struct {
	repository repository.UsersRepo
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(userrepo repository.UsersRepo) UserService {
	return UserService{
		repository: userrepo,
	}
}

func (us UserService) Get(id int) (*models.User, error) {
	user := us.repository.FindById(id)
	if user == nil {
		return nil, errors.New("user not found")
	} else {
		return user, nil
	}
}
