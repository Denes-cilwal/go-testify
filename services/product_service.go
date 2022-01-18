package services

import (
	"go-testify/models"
)

type ProductService struct {
	repository models.ProductRepo
}

func NewProductService(productrepo models.ProductRepo) ProductService {
	return ProductService{
		repository: productrepo,
	}
}

func (ps ProductService) IsProductReservable(id int) (models.Product, error) {
	return ps.repository.IsProductReservable(id)
}
