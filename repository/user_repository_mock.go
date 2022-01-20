package repository

import (
	"go-testify/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindById(id int) *models.User {
	// get args
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		// return args of 0 index
		user := args.Get(0).(models.User)
		return &user
	}
}
