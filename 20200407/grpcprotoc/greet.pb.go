// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package Greet

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GreetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GreetRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type GreetReplay struct {
	Ack                  string   `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetReplay) Reset()         { *m = GreetReplay{} }
func (m *GreetReplay) String() string { return proto.CompactTextString(m) }
func (*GreetReplay) ProtoMessage()    {}
func (*GreetReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetReplay.Unmarshal(m, b)
}
func (m *GreetReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetReplay.Marshal(b, m, deterministic)
}
func (m *GreetReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetReplay.Merge(m, src)
}
func (m *GreetReplay) XXX_Size() int {
	return xxx_messageInfo_GreetReplay.Size(m)
}
func (m *GreetReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetReplay.DiscardUnknown(m)
}

var xxx_messageInfo_GreetReplay proto.InternalMessageInfo

func (m *GreetReplay) GetAck() string {
	if m != nil {
		return m.Ack
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "Greet.GreetRequest")
	proto.RegisterType((*GreetReplay)(nil), "Greet.GreetReplay")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x75, 0x07, 0x71, 0x94, 0x6c, 0xb8, 0x78,
	0xc0, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x39, 0x3f,
	0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x2a, 0xc9, 0x73, 0x71, 0x43, 0x75,
	0x17, 0xe4, 0x24, 0x56, 0x0a, 0x09, 0x70, 0x31, 0x27, 0x26, 0x67, 0x43, 0xf5, 0x82, 0x98, 0x46,
	0x4e, 0x50, 0xe3, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x8c, 0xb8, 0x20, 0xf6, 0x0a,
	0x09, 0xeb, 0x81, 0x69, 0x3d, 0x64, 0xcb, 0xa5, 0x84, 0x50, 0x05, 0x41, 0x66, 0x2a, 0x31, 0x24,
	0xb1, 0x81, 0x1d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x51, 0x51, 0x9f, 0xbf, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReplay, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetReplay, error) {
	out := new(GreetReplay)
	err := c.cc.Invoke(ctx, "/Greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetReplay, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}
