// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff5ab49d8a5fcc4, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptor_6ff5ab49d8a5fcc4) }

var fileDescriptor_6ff5ab49d8a5fcc4 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62,
	0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x54, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x4a,
	0x60, 0x5c, 0x23, 0x67, 0x2e, 0x56, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x29, 0x2e, 0xe6, 0xe0,
	0xc4, 0x4a, 0x21, 0x0e, 0x3d, 0xa8, 0x99, 0x52, 0x9c, 0x7a, 0x70, 0xed, 0x72, 0x5c, 0xec, 0xee,
	0xf9, 0xf9, 0x29, 0x49, 0x95, 0xa9, 0x58, 0xe5, 0x93, 0xd8, 0xc0, 0x0e, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0xd5, 0x14, 0x5f, 0xa4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hello service

type HelloClient interface {
	Say(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Goodbye(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type helloClient struct {
	c           client.Client
	serviceName string
}

func NewHelloClient(serviceName string, c client.Client) HelloClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "hello"
	}
	return &helloClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *helloClient) Say(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Hello.Say", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Goodbye(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Hello.Goodbye", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloHandler interface {
	Say(context.Context, *Request, *Response) error
	Goodbye(context.Context, *Request, *Response) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Hello{hdlr}, opts...))
}

type Hello struct {
	HelloHandler
}

func (h *Hello) Say(ctx context.Context, in *Request, out *Response) error {
	return h.HelloHandler.Say(ctx, in, out)
}

func (h *Hello) Goodbye(ctx context.Context, in *Request, out *Response) error {
	return h.HelloHandler.Goodbye(ctx, in, out)
}
