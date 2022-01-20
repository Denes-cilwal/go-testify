package repository

import (
	"go-testify/infrastructure"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db infrastructure.Database
}

func NewProductRepository(db infrastructure.Database) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

// WithTrx delegate transaction from user repository
func (p ProductRepository) WithTrx(trxHandle *gorm.DB) ProductRepository {
	p.db.DB = trxHandle
	return p
}

