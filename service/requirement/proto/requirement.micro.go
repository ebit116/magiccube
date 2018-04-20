// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: requirement.proto

/*
Package requirement is a generated protocol buffer package.

It is generated from these files:
	requirement.proto

It has these top-level messages:
	PublishRequest
	PublishResponse
	QueryRequest
	QueryResponse
	QueryData
	QueryRow
*/
package requirement

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Requirement service

type RequirementClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type requirementClient struct {
	c           client.Client
	serviceName string
}

func NewRequirementClient(serviceName string, c client.Client) RequirementClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "requirement"
	}
	return &requirementClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *requirementClient) Publish(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Requirement.Publish", in)
	out := new(PublishResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requirementClient) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Requirement.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Requirement service

type RequirementHandler interface {
	Publish(context.Context, *PublishRequest, *PublishResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
}

func RegisterRequirementHandler(s server.Server, hdlr RequirementHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Requirement{hdlr}, opts...))
}

type Requirement struct {
	RequirementHandler
}

func (h *Requirement) Publish(ctx context.Context, in *PublishRequest, out *PublishResponse) error {
	return h.RequirementHandler.Publish(ctx, in, out)
}

func (h *Requirement) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.RequirementHandler.Query(ctx, in, out)
}