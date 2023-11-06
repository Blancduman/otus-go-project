// Code generated by mockery v2.34.0. DO NOT EDIT.

package mocks

import (
	context "context"

	payload "github.com/Blancduman/banners-rotation/internal/reporter/payload"
	mock "github.com/stretchr/testify/mock"
)

// Producer is an autogenerated mock type for the Producer type
type Producer struct {
	mock.Mock
}

// Produce provides a mock function with given fields: ctx, _a1
func (_m *Producer) Produce(ctx context.Context, _a1 payload.Payload) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, payload.Payload) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProducer creates a new instance of Producer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProducer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Producer {
	mock := &Producer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}