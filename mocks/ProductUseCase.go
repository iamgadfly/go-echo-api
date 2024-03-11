// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/iamgadfly/go-echo-api/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductUseCase is an autogenerated mock type for the ProductUseCase type
type ProductUseCase struct {
	mock.Mock
}

// GetById provides a mock function with given fields: ctx, id
func (_m *ProductUseCase) GetById(ctx context.Context, id string) (*models.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseByLink provides a mock function with given fields: link
func (_m *ProductUseCase) ParseByLink(link string) (models.Product, error) {
	ret := _m.Called(link)

	if len(ret) == 0 {
		panic("no return value specified for ParseByLink")
	}

	var r0 models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Product, error)); ok {
		return rf(link)
	}
	if rf, ok := ret.Get(0).(func(string) models.Product); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Get(0).(models.Product)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseWbCat provides a mock function with given fields: urls
func (_m *ProductUseCase) ParseWbCat(urls []string) error {
	ret := _m.Called(urls)

	if len(ret) == 0 {
		panic("no return value specified for ParseWbCat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(urls)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: ctx, string1
func (_m *ProductUseCase) Search(ctx context.Context, string1 string) ([]models.Product, error) {
	ret := _m.Called(ctx, string1)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Product, error)); ok {
		return rf(ctx, string1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Product); ok {
		r0 = rf(ctx, string1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, string1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteCsv provides a mock function with given fields:
func (_m *ProductUseCase) WriteCsv() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for WriteCsv")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductUseCase creates a new instance of ProductUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductUseCase {
	mock := &ProductUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
