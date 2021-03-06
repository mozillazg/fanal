// Code generated by mockery v1.0.0. DO NOT EDIT.

package image

import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"
import types "github.com/containers/image/v5/types"

// MockImageSource is an autogenerated mock type for the ImageSource type
type MockImageSource struct {
	mock.Mock
}

type ImageSourceCloseReturns struct {
	_a0 error
}

type ImageSourceCloseExpectation struct {
	Returns ImageSourceCloseReturns
}

func (_m *MockImageSource) ApplyCloseExpectation(e ImageSourceCloseExpectation) {
	var args []interface{}
	_m.On("Close", args...).Return(e.Returns._a0)
}

func (_m *MockImageSource) ApplyCloseExpectations(expectations []ImageSourceCloseExpectation) {
	for _, e := range expectations {
		_m.ApplyCloseExpectation(e)
	}
}

// Close provides a mock function with given fields:
func (_m *MockImageSource) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ImageSourceGetBlobArgs struct {
	Ctx           context.Context
	CtxAnything   bool
	Info          types.BlobInfo
	InfoAnything  bool
	Cache         types.BlobInfoCache
	CacheAnything bool
}

type ImageSourceGetBlobReturns struct {
	Reader io.ReadCloser
	N      int64
	Err    error
}

type ImageSourceGetBlobExpectation struct {
	Args    ImageSourceGetBlobArgs
	Returns ImageSourceGetBlobReturns
}

func (_m *MockImageSource) ApplyGetBlobExpectation(e ImageSourceGetBlobExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.InfoAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Info)
	}
	if e.Args.CacheAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Cache)
	}
	_m.On("GetBlob", args...).Return(e.Returns.Reader, e.Returns.N, e.Returns.Err)
}

func (_m *MockImageSource) ApplyGetBlobExpectations(expectations []ImageSourceGetBlobExpectation) {
	for _, e := range expectations {
		_m.ApplyGetBlobExpectation(e)
	}
}

// GetBlob provides a mock function with given fields: ctx, info, cache
func (_m *MockImageSource) GetBlob(ctx context.Context, info types.BlobInfo, cache types.BlobInfoCache) (io.ReadCloser, int64, error) {
	ret := _m.Called(ctx, info, cache)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, types.BlobInfo, types.BlobInfoCache) io.ReadCloser); ok {
		r0 = rf(ctx, info, cache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, types.BlobInfo, types.BlobInfoCache) int64); ok {
		r1 = rf(ctx, info, cache)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, types.BlobInfo, types.BlobInfoCache) error); ok {
		r2 = rf(ctx, info, cache)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
