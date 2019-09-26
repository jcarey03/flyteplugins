// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import io "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/lyft/flytestdlib/storage"

// OutputWriter is an autogenerated mock type for the OutputWriter type
type OutputWriter struct {
	mock.Mock
}

// GetErrorPath provides a mock function with given fields:
func (_m *OutputWriter) GetErrorPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// GetOutputPath provides a mock function with given fields:
func (_m *OutputWriter) GetOutputPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// GetOutputPrefixPath provides a mock function with given fields:
func (_m *OutputWriter) GetOutputPrefixPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

// Put provides a mock function with given fields: ctx, reader
func (_m *OutputWriter) Put(ctx context.Context, reader io.OutputReader) error {
	ret := _m.Called(ctx, reader)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, io.OutputReader) error); ok {
		r0 = rf(ctx, reader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
