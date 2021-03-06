// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: exchange.proto

/*
Package exchange is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	PushRequest
	BuyAssetResponse
	QueryRequest
	GetTxListResponse
	QueryTxData
	TxRow
	TransferRequest
	TransferResponse
	QueryTransferRequest
	QueryTransferResponse
*/
package exchange

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

// Client API for Exchange service

type ExchangeClient interface {
	BuyAsset(ctx context.Context, in *PushRequest, opts ...client.CallOption) (*BuyAssetResponse, error)
}

type exchangeClient struct {
	c           client.Client
	serviceName string
}

func NewExchangeClient(serviceName string, c client.Client) ExchangeClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "exchange"
	}
	return &exchangeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exchangeClient) BuyAsset(ctx context.Context, in *PushRequest, opts ...client.CallOption) (*BuyAssetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Exchange.BuyAsset", in)
	out := new(BuyAssetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Exchange service

type ExchangeHandler interface {
	BuyAsset(context.Context, *PushRequest, *BuyAssetResponse) error
}

func RegisterExchangeHandler(s server.Server, hdlr ExchangeHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Exchange{hdlr}, opts...))
}

type Exchange struct {
	ExchangeHandler
}

func (h *Exchange) BuyAsset(ctx context.Context, in *PushRequest, out *BuyAssetResponse) error {
	return h.ExchangeHandler.BuyAsset(ctx, in, out)
}
