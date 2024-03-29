package repository

import (
	"go-testify/infrastructure"
	"go-testify/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db infrastructure.Database
}

type UsersRepo interface {
	FindById(id int) *models.User
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
