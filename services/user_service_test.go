package services

import (
	"go-testify/models"
	"go-testify/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{
	Mock: mock.Mock{}}
var userService = UserService{
	repository: userRepository,
}

func TestUserService_GetUserNotFound(t *testing.T) {

	// program mock
	userRepository.Mock.On("FindById", 1).Return(nil)

	category, err := userService.Get(1)
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestUserService_GetUserSuccess(t *testing.T) {
	user := models.User{
		Id:        1,
		Name:      "dinesh",
		Email:     "test@gmail.com",
		Age:       25,
		CreatedAt: time.Now(),
	}
	userRepository.Mock.On("FindById", 2).Return(user)

	result, err := userService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.Age, result.Age)
	assert.Equal(t, user.CreatedAt, result.CreatedAt)
}
