// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// product is an autogenerated mock type for the product type
type product struct {
	mock.Mock
}

// Back provides a mock function with given fields:
func (_m *product) Back() {
	_m.Called()
}

// Forward provides a mock function with given fields:
func (_m *product) Forward() {
	_m.Called()
}

func NewProduct() *product {
	return &product{}
}