package mocks

import (
	"github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	"github.com/stretchr/testify/mock"
)

type MenteeUsecaseMock struct {
	Mock mock.Mock
}

func (m MenteeUsecaseMock) Register(menteeAuth *mentees.MenteeAuth) error {
	ret := m.Mock.Called(menteeAuth)

	return ret.Error(0)
}

func (m MenteeUsecaseMock) VerifyRegister(menteeRegister *mentees.MenteeRegister) error {
	ret := m.Mock.Called(menteeRegister)

	return ret.Error(0)
}

func (m MenteeUsecaseMock) ForgotPassword(forgotPassword *mentees.MenteeForgotPassword) error {
	ret := m.Mock.Called(forgotPassword)

	return ret.Error(0)
}

func (m MenteeUsecaseMock) Login(menteeAuth *mentees.MenteeAuth) (*string, error) {
	ret := m.Mock.Called(menteeAuth)

	return ret.Get(0).(*string), ret.Error(1)
}

func (m MenteeUsecaseMock) FindAll() (*[]mentees.Domain, error) {
	ret := m.Mock.Called()

	return ret.Get(0).(*[]mentees.Domain), ret.Error(1)
}

func (m MenteeUsecaseMock) FindById(id string) (*mentees.Domain, error) {
	ret := m.Mock.Called(id)

	return ret.Get(0).(*mentees.Domain), ret.Error(1)
}

func (m MenteeUsecaseMock) Update(id string, userDomain *mentees.Domain) error {
	ret := m.Mock.Called(id, userDomain)

	return ret.Error(0)
}
