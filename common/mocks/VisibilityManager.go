// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import "github.com/uber/cadence/common/persistence"
import "github.com/stretchr/testify/mock"

// VisibilityManager is an autogenerated mock type for the VisibilityManager type
type VisibilityManager struct {
	mock.Mock
}

// ListClosedWorkflowExecutions provides a mock function with given fields: request
func (_m *VisibilityManager) ListClosedWorkflowExecutions(request *persistence.ListWorkflowExecutionsRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClosedWorkflowExecutionsByStatus provides a mock function with given fields: request
func (_m *VisibilityManager) ListClosedWorkflowExecutionsByStatus(request *persistence.ListClosedWorkflowExecutionsByStatusRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListClosedWorkflowExecutionsByStatusRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListClosedWorkflowExecutionsByStatusRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClosedWorkflowExecutionsByType provides a mock function with given fields: request
func (_m *VisibilityManager) ListClosedWorkflowExecutionsByType(request *persistence.ListWorkflowExecutionsByTypeRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsByTypeRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsByTypeRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClosedWorkflowExecutionsByWorkflowID provides a mock function with given fields: request
func (_m *VisibilityManager) ListClosedWorkflowExecutionsByWorkflowID(request *persistence.ListWorkflowExecutionsByWorkflowIDRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsByWorkflowIDRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsByWorkflowIDRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOpenWorkflowExecutions provides a mock function with given fields: request
func (_m *VisibilityManager) ListOpenWorkflowExecutions(request *persistence.ListWorkflowExecutionsRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOpenWorkflowExecutionsByType provides a mock function with given fields: request
func (_m *VisibilityManager) ListOpenWorkflowExecutionsByType(request *persistence.ListWorkflowExecutionsByTypeRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsByTypeRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsByTypeRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOpenWorkflowExecutionsByWorkflowID provides a mock function with given fields: request
func (_m *VisibilityManager) ListOpenWorkflowExecutionsByWorkflowID(request *persistence.ListWorkflowExecutionsByWorkflowIDRequest) (*persistence.ListWorkflowExecutionsResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(*persistence.ListWorkflowExecutionsByWorkflowIDRequest) *persistence.ListWorkflowExecutionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.ListWorkflowExecutionsByWorkflowIDRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordWorkflowExecutionClosed provides a mock function with given fields: request
func (_m *VisibilityManager) RecordWorkflowExecutionClosed(request *persistence.RecordWorkflowExecutionClosedRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.RecordWorkflowExecutionClosedRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RecordWorkflowExecutionStarted provides a mock function with given fields: request
func (_m *VisibilityManager) RecordWorkflowExecutionStarted(request *persistence.RecordWorkflowExecutionStartedRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.RecordWorkflowExecutionStartedRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
