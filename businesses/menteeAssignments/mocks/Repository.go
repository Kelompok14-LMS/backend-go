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

// Create provides a mock function with given fields: assignmentMenteeDomain
func (_m *Repository) Create(assignmentMenteeDomain *mentee_assignments.Domain) error {
	ret := _m.Called(assignmentMenteeDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(*mentee_assignments.Domain) error); ok {
		r0 = rf(assignmentMenteeDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: assignmentMenteeId
func (_m *Repository) Delete(assignmentMenteeId string) error {
	ret := _m.Called(assignmentMenteeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(assignmentMenteeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByAssignmentId provides a mock function with given fields: assignmentId, limit, offset
func (_m *Repository) FindByAssignmentId(assignmentId string, limit int, offset int) ([]mentee_assignments.Domain, int, error) {
	ret := _m.Called(assignmentId, limit, offset)

	var r0 []mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string, int, int) []mentee_assignments.Domain); ok {
		r0 = rf(assignmentId, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee_assignments.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, int, int) int); ok {
		r1 = rf(assignmentId, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, int, int) error); ok {
		r2 = rf(assignmentId, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByCourse provides a mock function with given fields: menteeId, course
func (_m *Repository) FindByCourse(menteeId string, course string) (*mentee_assignments.Domain, error) {
	ret := _m.Called(menteeId, course)

	var r0 *mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string, string) *mentee_assignments.Domain); ok {
		r0 = rf(menteeId, course)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentee_assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(menteeId, course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: assignmentMenteeId
func (_m *Repository) FindById(assignmentMenteeId string) (*mentee_assignments.Domain, error) {
	ret := _m.Called(assignmentMenteeId)

	var r0 *mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string) *mentee_assignments.Domain); ok {
		r0 = rf(assignmentMenteeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentee_assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(assignmentMenteeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByMenteeId provides a mock function with given fields: menteeId
func (_m *Repository) FindByMenteeId(menteeId string) ([]mentee_assignments.Domain, error) {
	ret := _m.Called(menteeId)

	var r0 []mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string) []mentee_assignments.Domain); ok {
		r0 = rf(menteeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee_assignments.Domain)
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

// FindMenteeAssignmentEnrolled provides a mock function with given fields: menteeId, assignmentId
func (_m *Repository) FindMenteeAssignmentEnrolled(menteeId string, assignmentId string) (*mentee_assignments.Domain, error) {
	ret := _m.Called(menteeId, assignmentId)

	var r0 *mentee_assignments.Domain
	if rf, ok := ret.Get(0).(func(string, string) *mentee_assignments.Domain); ok {
		r0 = rf(menteeId, assignmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mentee_assignments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(menteeId, assignmentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: assignmentMenteeId, assignmentMenteeDomain
func (_m *Repository) Update(assignmentMenteeId string, assignmentMenteeDomain *mentee_assignments.Domain) error {
	ret := _m.Called(assignmentMenteeId, assignmentMenteeDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *mentee_assignments.Domain) error); ok {
		r0 = rf(assignmentMenteeId, assignmentMenteeDomain)
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
