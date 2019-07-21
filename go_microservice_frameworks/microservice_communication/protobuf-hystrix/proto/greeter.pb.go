// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/proto/greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
    github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/proto/greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package greeter

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
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
		serviceName = "greeter"
	}
	return &greeterClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Greeter.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Greeter{hdlr}, opts...))
}

type Greeter struct {
	GreeterHandler
}

func (h *Greeter) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/ArturoTarinVillaescusa/go_cloud_orchestration/go_microservice_frameworks/microservice_communication/protobuf-hystrix/proto/greeter.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x3d, 0x4b, 0xc6, 0x30,
	0x14, 0x85, 0x7d, 0xc5, 0xcf, 0xe0, 0xeb, 0x90, 0xa9, 0x74, 0x92, 0x4c, 0x05, 0x49, 0x02, 0xf6,
	0x17, 0x88, 0x60, 0x5d, 0x94, 0xd2, 0xd5, 0x29, 0x4d, 0x2f, 0x69, 0x68, 0x9a, 0x5b, 0xf3, 0x51,
	0xff, 0xbe, 0x10, 0x45, 0x74, 0xbb, 0x07, 0xce, 0x7d, 0x9e, 0x43, 0xde, 0x8d, 0x4d, 0x73, 0x1e,
	0x85, 0xc6, 0x55, 0xf6, 0x4a, 0x2f, 0xa9, 0xcf, 0xa3, 0xb3, 0x71, 0xb6, 0xde, 0xc8, 0xc7, 0x69,
	0x57, 0x5e, 0xc3, 0xc4, 0x9f, 0x1c, 0xe6, 0x89, 0xbf, 0xa9, 0x64, 0x77, 0xe0, 0x1d, 0xca, 0xe7,
	0xa0, 0x56, 0xf8, 0xc4, 0xb0, 0x44, 0xd9, 0x21, 0x7f, 0xb5, 0x3a, 0xa0, 0xdc, 0x02, 0x26, 0x94,
	0x26, 0x00, 0x24, 0x08, 0xa2, 0x24, 0xc6, 0xc8, 0xcd, 0x0b, 0x38, 0x87, 0x03, 0x7c, 0x64, 0x88,
	0x89, 0x52, 0x72, 0xe6, 0xd5, 0x0a, 0xd5, 0xe1, 0xee, 0xd0, 0x5c, 0x0f, 0xe5, 0x66, 0xf7, 0xe4,
	0xf8, 0xd3, 0x89, 0x1b, 0xfa, 0x08, 0xb4, 0x26, 0x57, 0x85, 0x62, 0xbd, 0xa9, 0x4e, 0x4b, 0xf1,
	0x37, 0x3f, 0xb4, 0xe4, 0xb2, 0xfb, 0x36, 0xd0, 0x86, 0x9c, 0x97, 0x3f, 0x7a, 0x14, 0x7f, 0x1d,
	0xf5, 0xad, 0xf8, 0x87, 0x63, 0x27, 0xe3, 0x45, 0x19, 0xd3, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xb8, 0x0c, 0x3b, 0x06, 0xeb, 0x00, 0x00, 0x00,
}
