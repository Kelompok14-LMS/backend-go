// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mentees "github.com/Kelompok14-LMS/backend-go/businesses/mentees"
	mock "github.com/stretchr/testify/mock"

	pkg "github.com/Kelompok14-LMS/backend-go/pkg"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *Usecase) FindAll() (*[]mentees.Domain, error) {
	ret := _m.Called()

	var r0 *[]mentees.Domain
	if rf, ok := ret.Get(0).(func() *[]mentees.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]mentees.Domain)
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

// FindByCourse provides a mock function with given fields: courseId, pagination
func (_m *Usecase) FindByCourse(courseId string, pagination pkg.Pagination) (*pkg.Pagination, error) {
	ret := _m.Called(courseId, pagination)

	var r0 *pkg.Pagination
	if rf, ok := ret.Get(0).(func(string, pkg.Pagination) *pkg.Pagination); ok {
		r0 = rf(courseId, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pkg.Pagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, pkg.Pagination) error); ok {
		r1 = rf(courseId, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *Usecase) FindById(id string) (*mentees.Domain, error) {
	ret := _m.Called(id)

	var r0 *mentees.Domain
	if rf, ok := ret.Get(0).(func(string) *mentees.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentees.Domain)
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

// ForgotPassword provides a mock function with given fields: forgotPassword
func (_m *Usecase) ForgotPassword(forgotPassword *mentees.MenteeForgotPassword) error {
	ret := _m.Called(forgotPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentees.MenteeForgotPassword) error); ok {
		r0 = rf(forgotPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: menteeAuth
func (_m *Usecase) Login(menteeAuth *mentees.MenteeAuth) (interface{}, error) {
	ret := _m.Called(menteeAuth)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*mentees.MenteeAuth) interface{}); ok {
		r0 = rf(menteeAuth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*mentees.MenteeAuth) error); ok {
		r1 = rf(menteeAuth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: menteeAuth
func (_m *Usecase) Register(menteeAuth *mentees.MenteeAuth) error {
	ret := _m.Called(menteeAuth)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentees.MenteeAuth) error); ok {
		r0 = rf(menteeAuth)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: id, menteeDomain
func (_m *Usecase) Update(id string, menteeDomain *mentees.Domain) error {
	ret := _m.Called(id, menteeDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *mentees.Domain) error); ok {
		r0 = rf(id, menteeDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyRegister provides a mock function with given fields: menteeRegister
func (_m *Usecase) VerifyRegister(menteeRegister *mentees.MenteeRegister) error {
	ret := _m.Called(menteeRegister)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentees.MenteeRegister) error); ok {
		r0 = rf(menteeRegister)
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
