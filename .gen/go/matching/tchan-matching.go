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

// @generated Code generated by thrift-gen. Do not modify.

// Package matching is generated code used to make or handle TChannel calls using Thrift.
package matching

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"

	"github.com/uber/cadence/.gen/go/shared"
)

var _ = shared.GoUnusedProtection__

// Interfaces for the service and client for the services defined in the IDL.

// TChanMatchingService is the interface that defines the server handler and client interface.
type TChanMatchingService interface {
	AddActivityTask(ctx thrift.Context, addRequest *AddActivityTaskRequest) error
	AddDecisionTask(ctx thrift.Context, addRequest *AddDecisionTaskRequest) error
	PollForActivityTask(ctx thrift.Context, pollRequest *PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error)
	PollForDecisionTask(ctx thrift.Context, pollRequest *PollForDecisionTaskRequest) (*PollForDecisionTaskResponse, error)
}

// Implementation of a client and service handler.

type tchanMatchingServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanMatchingServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanMatchingServiceClient {
	return &tchanMatchingServiceClient{
		thriftService,
		client,
	}
}

// NewTChanMatchingServiceClient creates a client that can be used to make remote calls.
func NewTChanMatchingServiceClient(client thrift.TChanClient) TChanMatchingService {
	return NewTChanMatchingServiceInheritedClient("MatchingService", client)
}

func (c *tchanMatchingServiceClient) AddActivityTask(ctx thrift.Context, addRequest *AddActivityTaskRequest) error {
	var resp MatchingServiceAddActivityTaskResult
	args := MatchingServiceAddActivityTaskArgs{
		AddRequest: addRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "AddActivityTask", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.BadRequestError != nil:
			err = resp.BadRequestError
		case resp.InternalServiceError != nil:
			err = resp.InternalServiceError
		case resp.ServiceBusyError != nil:
			err = resp.ServiceBusyError
		default:
			err = fmt.Errorf("received no result or unknown exception for AddActivityTask")
		}
	}

	return err
}

func (c *tchanMatchingServiceClient) AddDecisionTask(ctx thrift.Context, addRequest *AddDecisionTaskRequest) error {
	var resp MatchingServiceAddDecisionTaskResult
	args := MatchingServiceAddDecisionTaskArgs{
		AddRequest: addRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "AddDecisionTask", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.BadRequestError != nil:
			err = resp.BadRequestError
		case resp.InternalServiceError != nil:
			err = resp.InternalServiceError
		case resp.ServiceBusyError != nil:
			err = resp.ServiceBusyError
		default:
			err = fmt.Errorf("received no result or unknown exception for AddDecisionTask")
		}
	}

	return err
}

func (c *tchanMatchingServiceClient) PollForActivityTask(ctx thrift.Context, pollRequest *PollForActivityTaskRequest) (*shared.PollForActivityTaskResponse, error) {
	var resp MatchingServicePollForActivityTaskResult
	args := MatchingServicePollForActivityTaskArgs{
		PollRequest: pollRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "PollForActivityTask", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.BadRequestError != nil:
			err = resp.BadRequestError
		case resp.InternalServiceError != nil:
			err = resp.InternalServiceError
		default:
			err = fmt.Errorf("received no result or unknown exception for PollForActivityTask")
		}
	}

	return resp.GetSuccess(), err
}

func (c *tchanMatchingServiceClient) PollForDecisionTask(ctx thrift.Context, pollRequest *PollForDecisionTaskRequest) (*PollForDecisionTaskResponse, error) {
	var resp MatchingServicePollForDecisionTaskResult
	args := MatchingServicePollForDecisionTaskArgs{
		PollRequest: pollRequest,
	}
	success, err := c.client.Call(ctx, c.thriftService, "PollForDecisionTask", &args, &resp)
	if err == nil && !success {
		switch {
		case resp.BadRequestError != nil:
			err = resp.BadRequestError
		case resp.InternalServiceError != nil:
			err = resp.InternalServiceError
		default:
			err = fmt.Errorf("received no result or unknown exception for PollForDecisionTask")
		}
	}

	return resp.GetSuccess(), err
}

type tchanMatchingServiceServer struct {
	handler TChanMatchingService
}

// NewTChanMatchingServiceServer wraps a handler for TChanMatchingService so it can be
// registered with a thrift.Server.
func NewTChanMatchingServiceServer(handler TChanMatchingService) thrift.TChanServer {
	return &tchanMatchingServiceServer{
		handler,
	}
}

func (s *tchanMatchingServiceServer) Service() string {
	return "MatchingService"
}

func (s *tchanMatchingServiceServer) Methods() []string {
	return []string{
		"AddActivityTask",
		"AddDecisionTask",
		"PollForActivityTask",
		"PollForDecisionTask",
	}
}

func (s *tchanMatchingServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "AddActivityTask":
		return s.handleAddActivityTask(ctx, protocol)
	case "AddDecisionTask":
		return s.handleAddDecisionTask(ctx, protocol)
	case "PollForActivityTask":
		return s.handlePollForActivityTask(ctx, protocol)
	case "PollForDecisionTask":
		return s.handlePollForDecisionTask(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanMatchingServiceServer) handleAddActivityTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServiceAddActivityTaskArgs
	var res MatchingServiceAddActivityTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.AddActivityTask(ctx, req.AddRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		case *shared.ServiceBusyError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceBusyError returned non-nil error type *shared.ServiceBusyError but nil value")
			}
			res.ServiceBusyError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handleAddDecisionTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServiceAddDecisionTaskArgs
	var res MatchingServiceAddDecisionTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	err :=
		s.handler.AddDecisionTask(ctx, req.AddRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		case *shared.ServiceBusyError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for serviceBusyError returned non-nil error type *shared.ServiceBusyError but nil value")
			}
			res.ServiceBusyError = v
		default:
			return false, nil, err
		}
	} else {
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handlePollForActivityTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServicePollForActivityTaskArgs
	var res MatchingServicePollForActivityTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.PollForActivityTask(ctx, req.PollRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}

func (s *tchanMatchingServiceServer) handlePollForDecisionTask(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req MatchingServicePollForDecisionTaskArgs
	var res MatchingServicePollForDecisionTaskResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.PollForDecisionTask(ctx, req.PollRequest)

	if err != nil {
		switch v := err.(type) {
		case *shared.BadRequestError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for badRequestError returned non-nil error type *shared.BadRequestError but nil value")
			}
			res.BadRequestError = v
		case *shared.InternalServiceError:
			if v == nil {
				return false, nil, fmt.Errorf("Handler for internalServiceError returned non-nil error type *shared.InternalServiceError but nil value")
			}
			res.InternalServiceError = v
		default:
			return false, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}
