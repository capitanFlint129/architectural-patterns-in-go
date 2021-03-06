// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// parser is an autogenerated mock type for the parser type
type parser struct {
	mock.Mock
}

// Parse provides a mock function with given fields: command
func (_m *parser) Parse(command string) []string {
	ret := _m.Called(command)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(command)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

func NewParserMock() *parser {
	return &parser{}
}
