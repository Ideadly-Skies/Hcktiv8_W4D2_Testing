package service

import (
	"errors"
	"w4/d2/entity"
	"w4/d2/repo"
)

type ProductService struct {
	RepoProduct repo.ProductRepo 
}

func (ps *ProductService) GetProduct(id int) (*entity.Product, error) {
	product := ps.RepoProduct.FindByID(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}
	return product, nil
}