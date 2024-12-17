package repo

import (
	"w4/d2/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (m *ProductRepoMock) FindByID(id int) *entity.Product {
	res := m.Mock.Called(id)

	if res.Get(0) == nil {
		return nil
	}

	product := res.Get(0).(entity.Product)
	return &product
}