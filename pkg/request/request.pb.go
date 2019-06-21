// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package request

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

type Request struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
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

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Results struct {
	// repeated google.protobuf.Any response = 1;
	RawString            []string `protobuf:"bytes,2,rep,name=rawString,proto3" json:"rawString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetRawString() []string {
	if m != nil {
		return m.RawString
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Results)(nil), "Results")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0x62, 0x0f, 0x82, 0x08,
	0x08, 0x89, 0x70, 0xb1, 0x16, 0x96, 0xa6, 0x16, 0x55, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x38, 0x4a, 0xea, 0x20, 0x05, 0xc5, 0xa5, 0x39, 0x25, 0xc5, 0x42, 0x32, 0x5c, 0x9c, 0x45,
	0x89, 0xe5, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x08,
	0x01, 0x23, 0x4f, 0x2e, 0x8e, 0x80, 0x9c, 0xc4, 0xcc, 0xbc, 0xe0, 0x40, 0x1f, 0x21, 0x19, 0x2e,
	0x16, 0xd7, 0xe4, 0x8c, 0x7c, 0x21, 0x0e, 0x3d, 0xa8, 0xe1, 0x52, 0x70, 0x96, 0x12, 0x83, 0x90,
	0x3c, 0x17, 0xbb, 0x6b, 0x45, 0x6a, 0x72, 0x69, 0x49, 0x2a, 0x9a, 0x02, 0xb0, 0x35, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0xb7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0xef, 0xb3, 0x3a, 0xac,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlainSQLClient is the client API for PlainSQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlainSQLClient interface {
	Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error)
	Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Results, error)
}

type plainSQLClient struct {
	cc *grpc.ClientConn
}

func NewPlainSQLClient(cc *grpc.ClientConn) PlainSQLClient {
	return &plainSQLClient{cc}
}

func (c *plainSQLClient) Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Request, error) {
	out := new(Request)
	err := c.cc.Invoke(ctx, "/PlainSQL/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plainSQLClient) Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/PlainSQL/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlainSQLServer is the server API for PlainSQL service.
type PlainSQLServer interface {
	Echo(context.Context, *Request) (*Request, error)
	Execute(context.Context, *Request) (*Results, error)
}

// UnimplementedPlainSQLServer can be embedded to have forward compatible implementations.
type UnimplementedPlainSQLServer struct {
}

func (*UnimplementedPlainSQLServer) Echo(ctx context.Context, req *Request) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedPlainSQLServer) Execute(ctx context.Context, req *Request) (*Results, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterPlainSQLServer(s *grpc.Server, srv PlainSQLServer) {
	s.RegisterService(&_PlainSQL_serviceDesc, srv)
}

func _PlainSQL_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlainSQLServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlainSQL/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlainSQLServer).Echo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlainSQL_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlainSQLServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlainSQL/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlainSQLServer).Execute(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlainSQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PlainSQL",
	HandlerType: (*PlainSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _PlainSQL_Echo_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _PlainSQL_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "request.proto",
}
