// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	assignments "github.com/Kelompok14-LMS/backend-go/businesses/assignments"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: assignmentDomain
func (_m *Repository) Create(assignmentDomain *assignments.Domain) error {
	ret := _m.Called(assignmentDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*assignments.Domain) error); ok {
		r0 = rf(assignmentDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: assignmentId
func (_m *Repository) Delete(assignmentId string) error {
	ret := _m.Called(assignmentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(assignmentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByModuleId provides a mock function with given fields: moduleId
func (_m *Repository) DeleteByModuleId(moduleId string) error {
	ret := _m.Called(moduleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(moduleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: assignmentId
func (_m *Repository) FindById(assignmentId string) (*assignments.Domain, error) {
	ret := _m.Called(assignmentId)

	var r0 *assignments.Domain
	if rf, ok := ret.Get(0).(func(string) *assignments.Domain); ok {
		r0 = rf(assignmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(assignmentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByModuleId provides a mock function with given fields: moduleId
func (_m *Repository) FindByModuleId(moduleId string) (*assignments.Domain, error) {
	ret := _m.Called(moduleId)

	var r0 *assignments.Domain
	if rf, ok := ret.Get(0).(func(string) *assignments.Domain); ok {
		r0 = rf(moduleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assignments.Domain)
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

// Update provides a mock function with given fields: assignmentId, assignmentDomain
func (_m *Repository) Update(assignmentId string, assignmentDomain *assignments.Domain) error {
	ret := _m.Called(assignmentId, assignmentDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *assignments.Domain) error); ok {
		r0 = rf(assignmentId, assignmentDomain)
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
