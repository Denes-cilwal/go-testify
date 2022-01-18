package repository

import (
	"go-testify/infrastructure"
	"go-testify/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	r.db.DB = trxHandle
	return r
}

func (u UserRepository) CreateUser(user models.User) (models.User, error) {
	return user, u.db.DB.Create(user).Error
}

