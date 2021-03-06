// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sync "sync"
)

// receiver is an autogenerated mock type for the receiver type
type receiver struct {
	mock.Mock
}

// StartReceive provides a mock function with given fields: ctx, wg
func (_m *receiver) StartReceive(ctx context.Context, wg *sync.WaitGroup) {
	_m.Called(ctx, wg)
}

func NewReceiver() *receiver {
	return &receiver{}
}
