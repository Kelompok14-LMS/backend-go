package mocks

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/categories"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (c *CategoryRepositoryMock) Create(categoryDomain *categories.Domain) error {
	ret := c.Mock.Called(categoryDomain)

	return ret.Error(0)
}

func (c *CategoryRepositoryMock) FindAll() (*[]categories.Domain, error) {
	ret := c.Mock.Called()

	return ret.Get(0).(*[]categories.Domain), ret.Error(1)
}

func (c *CategoryRepositoryMock) FindById(id string) (*categories.Domain, error) {
	ret := c.Mock.Called(id)

	return ret.Get(0).(*categories.Domain), ret.Error(1)
}

func (c *CategoryRepositoryMock) Update(id string, categoryDomain *categories.Domain) error {
	ret := c.Mock.Called(id, categoryDomain)

	return ret.Error(0)
}

// func (c *CategoryRepositoryMock) Delete(id string) error {
// 	panic("implement me")
// }
