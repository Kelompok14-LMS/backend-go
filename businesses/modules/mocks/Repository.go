// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	modules "github.com/Kelompok14-LMS/backend-go/businesses/modules"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: moduleDomain
func (_m *Repository) Create(moduleDomain *modules.Domain) error {
	ret := _m.Called(moduleDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*modules.Domain) error); ok {
		r0 = rf(moduleDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: moduleId
func (_m *Repository) Delete(moduleId string) error {
	ret := _m.Called(moduleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(moduleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByCourse provides a mock function with given fields: courseId
func (_m *Repository) FindByCourse(courseId string) ([]modules.Domain, error) {
	ret := _m.Called(courseId)

	var r0 []modules.Domain
	if rf, ok := ret.Get(0).(func(string) []modules.Domain); ok {
		r0 = rf(courseId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]modules.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(courseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: moduleId
func (_m *Repository) FindById(moduleId string) (*modules.Domain, error) {
	ret := _m.Called(moduleId)

	var r0 *modules.Domain
	if rf, ok := ret.Get(0).(func(string) *modules.Domain); ok {
		r0 = rf(moduleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modules.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(moduleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: moduleId, moduleDomain
func (_m *Repository) Update(moduleId string, moduleDomain *modules.Domain) error {
	ret := _m.Called(moduleId, moduleDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *modules.Domain) error); ok {
		r0 = rf(moduleId, moduleDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
