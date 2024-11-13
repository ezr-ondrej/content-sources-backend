// Code generated by mockery. DO NOT EDIT.

package dao

import (
	context "context"

	models "github.com/content-services/content-sources-backend/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockDomainDao is an autogenerated mock type for the DomainDao type
type MockDomainDao struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, orgId
func (_m *MockDomainDao) Fetch(ctx context.Context, orgId string) (string, error) {
	ret := _m.Called(ctx, orgId)

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, orgId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchOrCreateDomain provides a mock function with given fields: ctx, orgId
func (_m *MockDomainDao) FetchOrCreateDomain(ctx context.Context, orgId string) (string, error) {
	ret := _m.Called(ctx, orgId)

	if len(ret) == 0 {
		panic("no return value specified for FetchOrCreateDomain")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, orgId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, orgId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *MockDomainDao) List(ctx context.Context) ([]models.Domain, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []models.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Domain, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDomainDao creates a new instance of MockDomainDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDomainDao(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDomainDao {
	mock := &MockDomainDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
