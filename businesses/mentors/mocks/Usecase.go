// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	mentors "github.com/Kelompok14-LMS/backend-go/businesses/mentors"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *Usecase) FindAll() (*[]mentors.Domain, error) {
	ret := _m.Called()

	var r0 *[]mentors.Domain
	if rf, ok := ret.Get(0).(func() *[]mentors.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]mentors.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *Usecase) FindById(id string) (*mentors.Domain, error) {
	ret := _m.Called(id)

	var r0 *mentors.Domain
	if rf, ok := ret.Get(0).(func(string) *mentors.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentors.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: mentorAuth
func (_m *Usecase) Login(mentorAuth *mentors.MentorAuth) (*string, error) {
	ret := _m.Called(mentorAuth)

	var r0 *string
	if rf, ok := ret.Get(0).(func(*mentors.MentorAuth) *string); ok {
		r0 = rf(mentorAuth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*mentors.MentorAuth) error); ok {
		r1 = rf(mentorAuth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: mentorAuth
func (_m *Usecase) Register(mentorAuth *mentors.MentorRegister) error {
	ret := _m.Called(mentorAuth)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentors.MentorRegister) error); ok {
		r0 = rf(mentorAuth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: updateMentor
func (_m *Usecase) Update(updateMentor *mentors.MentorUpdateProfile) error {
	ret := _m.Called(updateMentor)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentors.MentorUpdateProfile) error); ok {
		r0 = rf(updateMentor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: updatePassword
func (_m *Usecase) UpdatePassword(updatePassword *mentors.MentorUpdatePassword) error {
	ret := _m.Called(updatePassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentors.MentorUpdatePassword) error); ok {
		r0 = rf(updatePassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
