// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	mentee_assignments "github.com/Kelompok14-LMS/backend-go/businesses/menteeAssignments"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: assignmentmenteeDomain
func (_m *Repository) Create(assignmentmenteeDomain *mentee_assignments.Domain) error {
	ret := _m.Called(assignmentmenteeDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentee_assignments.Domain) error); ok {
		r0 = rf(assignmentmenteeDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: assignmentmenteeId
func (_m *Repository) Delete(assignmentmenteeId string) error {
	ret := _m.Called(assignmentmenteeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(assignmentmenteeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByAssignmentId provides a mock function with given fields: assignmentId
func (_m *Repository) FindByAssignmentId(assignmentId string) ([]mentee_assignments.Domain, error) {
	ret := _m.Called(assignmentId)

	var r0 []mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string) []mentee_assignments.Domain); ok {
		r0 = rf(assignmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee_assignments.Domain)
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

// FindById provides a mock function with given fields: assignmentmenteeId
func (_m *Repository) FindById(assignmentmenteeId string) (*mentee_assignments.Domain, error) {
	ret := _m.Called(assignmentmenteeId)

	var r0 *mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string) *mentee_assignments.Domain); ok {
		r0 = rf(assignmentmenteeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentee_assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(assignmentmenteeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByMenteeId provides a mock function with given fields: menteeId
func (_m *Repository) FindByMenteeId(menteeId string) (*mentee_assignments.Domain, error) {
	ret := _m.Called(menteeId)

	var r0 *mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string) *mentee_assignments.Domain); ok {
		r0 = rf(menteeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentee_assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(menteeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: assignmentmenteeId, assignmentmenteeDomain
func (_m *Repository) Update(assignmentmenteeId string, assignmentmenteeDomain *mentee_assignments.Domain) error {
	ret := _m.Called(assignmentmenteeId, assignmentmenteeDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *mentee_assignments.Domain) error); ok {
		r0 = rf(assignmentmenteeId, assignmentmenteeDomain)
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
