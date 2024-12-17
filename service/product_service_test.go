package service

import (
	"testing"
	"w4/d2/entity"
	"w4/d2/repo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"errors"
)

var productRepoMock = &repo.ProductRepoMock{Mock: mock.Mock{}}
var productService = ProductService{RepoProduct: productRepoMock}

func TestGetProductNotFound(t *testing.T) {
	productRepoMock.Mock.On("FindByID", 1).Return(nil, errors.New("Test Error Product not found"))

	product, err := productService.GetProduct(1)
	assert.Nil(t, product) // expected : product = nil
	assert.NotNil(t, err)  // expected : error = not nil
	assert.Equal(t, "Product not found", err.Error(), "Error message should match")
}

func TestGetProduct(t *testing.T){
	productRes := entity.Product {
		ID: 2,
		Name: "Pisang Goreng", 
	}

	productRepoMock.Mock.On("FindByID", 2).Return(productRes, nil)

	// actual is product
	product, err := productService.GetProduct(2)
	assert.Nil(t, err)		  // expected : error = nil
	assert.NotNil(t, product) // expected : product = not nil
	
	// this is the mock test
	assert.Equal(t, productRes.ID, product.ID, "Product not found")
	assert.Equal(t, productRes.Name, product.Name, "Product Name Should be the same")
	assert.Equal(t, &productRes, product, "Object Structure must be the same")
}