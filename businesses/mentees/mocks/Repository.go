package mocks

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/stretchr/testify/mock"
)

type MenteeRepositoryMock struct {
	Mock mock.Mock
}

func (m *MenteeRepositoryMock) Create(menteeDomain *mentees.Domain) error {
	ret := m.Mock.Called(menteeDomain)

	return ret.Error(0)
}

func (m *MenteeRepositoryMock) FindAll() (*[]mentees.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MenteeRepositoryMock) FindById(id string) (*mentees.Domain, error) {
	ret := m.Mock.Called(id)

	return ret.Get(0).(*mentees.Domain), ret.Error(1)
}

func (m *MenteeRepositoryMock) FindByIdUser(userId string) (*mentees.Domain, error) {
	ret := m.Mock.Called(userId)

	return ret.Get(0).(*mentees.Domain), ret.Error(1)
}

func (m *MenteeRepositoryMock) Update(id string, menteeDomain *mentees.Domain) error {
	ret := m.Mock.Called(id, menteeDomain)

	return ret.Error(0)
}
