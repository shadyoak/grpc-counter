// Code generated by protoc-gen-go.
// source: service/counter.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service/counter.proto

It has these top-level messages:
	CounterValue
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CounterValue struct {
	Count int32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *CounterValue) Reset()                    { *m = CounterValue{} }
func (m *CounterValue) String() string            { return proto.CompactTextString(m) }
func (*CounterValue) ProtoMessage()               {}
func (*CounterValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CounterValue) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*CounterValue)(nil), "service.CounterValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Counter service

type CounterClient interface {
	IncrementCounter(ctx context.Context, opts ...grpc.CallOption) (Counter_IncrementCounterClient, error)
}

type counterClient struct {
	cc *grpc.ClientConn
}

func NewCounterClient(cc *grpc.ClientConn) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) IncrementCounter(ctx context.Context, opts ...grpc.CallOption) (Counter_IncrementCounterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Counter_serviceDesc.Streams[0], c.cc, "/service.Counter/IncrementCounter", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterIncrementCounterClient{stream}
	return x, nil
}

type Counter_IncrementCounterClient interface {
	Send(*CounterValue) error
	Recv() (*CounterValue, error)
	grpc.ClientStream
}

type counterIncrementCounterClient struct {
	grpc.ClientStream
}

func (x *counterIncrementCounterClient) Send(m *CounterValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *counterIncrementCounterClient) Recv() (*CounterValue, error) {
	m := new(CounterValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Counter service

type CounterServer interface {
	IncrementCounter(Counter_IncrementCounterServer) error
}

func RegisterCounterServer(s *grpc.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_IncrementCounter_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CounterServer).IncrementCounter(&counterIncrementCounterServer{stream})
}

type Counter_IncrementCounterServer interface {
	Send(*CounterValue) error
	Recv() (*CounterValue, error)
	grpc.ServerStream
}

type counterIncrementCounterServer struct {
	grpc.ServerStream
}

func (x *counterIncrementCounterServer) Send(m *CounterValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *counterIncrementCounterServer) Recv() (*CounterValue, error) {
	m := new(CounterValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IncrementCounter",
			Handler:       _Counter_IncrementCounter_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/counter.proto",
}

func init() { proto.RegisterFile("service/counter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x87, 0x0a, 0x2b, 0xa9, 0x70, 0xf1, 0x38, 0x43, 0x64, 0xc2, 0x12, 0x73, 0x4a,
	0x53, 0x85, 0x44, 0xb8, 0x58, 0xc1, 0x2a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c,
	0xa3, 0x40, 0x2e, 0x76, 0xa8, 0x2a, 0x21, 0x37, 0x2e, 0x01, 0xcf, 0xbc, 0xe4, 0xa2, 0xd4, 0xdc,
	0xd4, 0xbc, 0x12, 0x98, 0x98, 0xa8, 0x1e, 0xd4, 0x38, 0x3d, 0x64, 0xb3, 0xa4, 0xb0, 0x0b, 0x2b,
	0x31, 0x68, 0x30, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x1d, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x66, 0x3f, 0x2c, 0xd2, 0xa1, 0x00, 0x00, 0x00,
}
