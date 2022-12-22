// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	materials "github.com/Kelompok14-LMS/backend-go/businesses/materials"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, materialDomain
func (_m *Usecase) Create(ctx context.Context, materialDomain *materials.Domain) error {
	ret := _m.Called(ctx, materialDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *materials.Domain) error); ok {
		r0 = rf(ctx, materialDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: materialId
func (_m *Usecase) Delete(materialId string) error {
	ret := _m.Called(materialId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(materialId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deletes provides a mock function with given fields: moduleId
func (_m *Usecase) Deletes(moduleId string) error {
	ret := _m.Called(moduleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(moduleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: materialId
func (_m *Usecase) FindById(materialId string) (*materials.Domain, error) {
	ret := _m.Called(materialId)

	var r0 *materials.Domain
	if rf, ok := ret.Get(0).(func(string) *materials.Domain); ok {
		r0 = rf(materialId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*materials.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(materialId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, materialId, materialDomain
func (_m *Usecase) Update(ctx context.Context, materialId string, materialDomain *materials.Domain) error {
	ret := _m.Called(ctx, materialId, materialDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *materials.Domain) error); ok {
		r0 = rf(ctx, materialId, materialDomain)
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
