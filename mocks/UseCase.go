// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/iamgadfly/go-echo-api/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, link
func (_m *UseCase) Create(ctx context.Context, link string) (*models.Vacancy, error) {
	ret := _m.Called(ctx, link)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Vacancy, error)); ok {
		return rf(ctx, link)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Vacancy); ok {
		r0 = rf(ctx, link)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *UseCase) GetById(ctx context.Context, id string) (*models.Vacancy, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.Vacancy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Vacancy, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Vacancy); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Vacancy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
