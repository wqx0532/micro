// Code generated by protoc-gen-go.
// source: github.com/micro/micro/examples/greeter/api/rpc/proto/hello/hello.proto
// DO NOT EDIT!

/*
Package go_micro_api_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/micro/examples/greeter/api/rpc/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_api_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.api.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type greeterClient struct {
	c           client.Client
	serviceName string
}

func NewGreeterClient(serviceName string, c client.Client) GreeterClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.api.greeter"
	}
	return &greeterClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *greeterClient) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Greeter.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Greeter{hdlr}, opts...))
}

type Greeter struct {
	GreeterHandler
}

func (h *Greeter) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/micro/micro/examples/greeter/api/rpc/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x8e, 0xcd, 0x0a, 0xc2, 0x30,
	0x10, 0x84, 0x2d, 0xfe, 0x54, 0xf7, 0x24, 0xc1, 0x83, 0x88, 0x15, 0xe9, 0xc9, 0xd3, 0x06, 0xf4,
	0x21, 0xda, 0x73, 0x7d, 0x82, 0xb6, 0x2c, 0x69, 0x21, 0x69, 0x62, 0x92, 0x82, 0x8f, 0xaf, 0x4d,
	0x73, 0xd4, 0xcb, 0x30, 0xec, 0xec, 0xec, 0xb7, 0x50, 0x88, 0xde, 0x77, 0x63, 0x83, 0xad, 0x56,
	0x5c, 0xf5, 0xad, 0xd5, 0x51, 0xe9, 0x5d, 0x2b, 0x23, 0xc9, 0x71, 0x61, 0x89, 0x3c, 0x59, 0x5e,
	0x9b, 0x9e, 0x5b, 0xd3, 0x72, 0x63, 0xb5, 0xd7, 0xbc, 0x23, 0x29, 0xa3, 0x62, 0x98, 0xb0, 0x83,
	0xd0, 0x18, 0xaa, 0xf8, 0xdd, 0xc4, 0xd8, 0xca, 0x33, 0x48, 0x2b, 0x7a, 0x8d, 0xe4, 0x3c, 0x63,
	0xb0, 0x1a, 0x6a, 0x45, 0xc7, 0xe4, 0x9a, 0xdc, 0x76, 0x55, 0xf0, 0xf9, 0x19, 0xb6, 0x15, 0x39,
	0xa3, 0x07, 0x47, 0x6c, 0x0f, 0x4b, 0xe5, 0x44, 0x8c, 0x27, 0x7b, 0x7f, 0x42, 0x5a, 0xcc, 0x77,
	0x58, 0x09, 0xeb, 0x72, 0x82, 0xb1, 0x0c, 0x7f, 0x71, 0x30, 0x42, 0x4e, 0x97, 0x7f, 0xf1, 0x0c,
	0xc9, 0x17, 0xcd, 0x26, 0xbc, 0xfb, 0xf8, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x9c, 0xe7, 0x65,
	0xf9, 0x00, 0x00, 0x00,
}
