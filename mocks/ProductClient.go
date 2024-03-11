// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	product "github.com/iamgadfly/go-echo-api/pkg/api/product_v1/proto"
)

// ProductClient is an autogenerated mock type for the ProductClient type
type ProductClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *ProductClient) Create(ctx context.Context, in *product.CreateRequest, opts ...grpc.CallOption) (*product.ProductResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *product.ProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *product.CreateRequest, ...grpc.CallOption) (*product.ProductResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *product.CreateRequest, ...grpc.CallOption) *product.ProductResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.ProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *product.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, in, opts
func (_m *ProductClient) GetById(ctx context.Context, in *product.GetByIdRequest, opts ...grpc.CallOption) (*product.ProductResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *product.ProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *product.GetByIdRequest, ...grpc.CallOption) (*product.ProductResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *product.GetByIdRequest, ...grpc.CallOption) *product.ProductResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.ProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *product.GetByIdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, in, opts
func (_m *ProductClient) Search(ctx context.Context, in *product.SearchRequest, opts ...grpc.CallOption) (*product.ProductListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *product.ProductListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *product.SearchRequest, ...grpc.CallOption) (*product.ProductListResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *product.SearchRequest, ...grpc.CallOption) *product.ProductListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.ProductListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *product.SearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductClient creates a new instance of ProductClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductClient {
	mock := &ProductClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
