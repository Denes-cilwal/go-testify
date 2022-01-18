package services

import (
	"go-testify/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepoTestifyMock struct {
	mock.Mock
}

func (u UserRepoTestifyMock) CreateUser(user models.User) (models.User, error) {
	// user is args here ..
	args := u.Called()
	// get returns args at specific index
	result := args.Get(0)
	return result.(models.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {

	mockRepo := UserRepoTestifyMock{}

	user := models.User{
		Id:        1,
		Name:      "dinesh",
		Email:     "test@gmail.com",
		Age:       25,
		CreatedAt: time.Now(),
	}

	// setup expectations on mock repo
	mockRepo.On("CreateUser").Return(user, nil)
	testService := NewUserService(mockRepo)
	result, err := testService.CreateUser(user)

	mockRepo.AssertExpectations(t)

	// assert if expected == input
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "dinesh", result.Name)
	assert.Equal(t, "test@gmail.com", result.Email)
	assert.Equal(t, 35, result.Age)
	assert.Nil(t, err)
}
