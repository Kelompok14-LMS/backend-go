package mocks

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/users"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (u *UserRepositoryMock) Create(userDomain *users.Domain) error {
	ret := u.Mock.Called(userDomain)

	return ret.Error(0)
}

func (u *UserRepositoryMock) FindAll() (*[]users.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepositoryMock) FindByEmail(email string) (*users.Domain, error) {
	ret := u.Mock.Called(email)

	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}

	return ret.Get(0).(*users.Domain), ret.Error(1)
}

func (u *UserRepositoryMock) FindById(id string) (*users.Domain, error) {
	ret := u.Mock.Called(id)

	return ret.Get(0).(*users.Domain), ret.Error(1)
}

func (u *UserRepositoryMock) Update(id string, userDomain *users.Domain) error {
	ret := u.Mock.Called(id, userDomain)

	return ret.Error(0)
}
