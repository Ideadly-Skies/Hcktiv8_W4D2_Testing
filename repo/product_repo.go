package repo

import (
	"w4/d2/entity"
)

type ProductRepo interface {
	FindByID(id int) *entity.Product
	// CreateProduct() *entity.Product
	// FindAll() *entity.Product
	// UpdateProduct(id int) *entity.Product
	// DeleteProduct(id int) *entity.Product



}