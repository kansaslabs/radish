// Code generated by protoc-gen-go. DO NOT EDIT.
// source: radish.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueueRequest struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Params               []byte   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Success              []byte   `protobuf:"bytes,3,opt,name=success,proto3" json:"success,omitempty"`
	Failure              []byte   `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueRequest) Reset()         { *m = QueueRequest{} }
func (m *QueueRequest) String() string { return proto.CompactTextString(m) }
func (*QueueRequest) ProtoMessage()    {}
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{0}
}

func (m *QueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRequest.Unmarshal(m, b)
}
func (m *QueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRequest.Marshal(b, m, deterministic)
}
func (m *QueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRequest.Merge(m, src)
}
func (m *QueueRequest) XXX_Size() int {
	return xxx_messageInfo_QueueRequest.Size(m)
}
func (m *QueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRequest proto.InternalMessageInfo

func (m *QueueRequest) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *QueueRequest) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *QueueRequest) GetSuccess() []byte {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *QueueRequest) GetFailure() []byte {
	if m != nil {
		return m.Failure
	}
	return nil
}

type QueueReply struct {
	Uuid                 []byte   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueReply) Reset()         { *m = QueueReply{} }
func (m *QueueReply) String() string { return proto.CompactTextString(m) }
func (*QueueReply) ProtoMessage()    {}
func (*QueueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{1}
}

func (m *QueueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueReply.Unmarshal(m, b)
}
func (m *QueueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueReply.Marshal(b, m, deterministic)
}
func (m *QueueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueReply.Merge(m, src)
}
func (m *QueueReply) XXX_Size() int {
	return xxx_messageInfo_QueueReply.Size(m)
}
func (m *QueueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueueReply proto.InternalMessageInfo

func (m *QueueReply) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *QueueReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *QueueReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ScaleRequest struct {
	Workers              int32    `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleRequest) Reset()         { *m = ScaleRequest{} }
func (m *ScaleRequest) String() string { return proto.CompactTextString(m) }
func (*ScaleRequest) ProtoMessage()    {}
func (*ScaleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{2}
}

func (m *ScaleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleRequest.Unmarshal(m, b)
}
func (m *ScaleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleRequest.Marshal(b, m, deterministic)
}
func (m *ScaleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleRequest.Merge(m, src)
}
func (m *ScaleRequest) XXX_Size() int {
	return xxx_messageInfo_ScaleRequest.Size(m)
}
func (m *ScaleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleRequest proto.InternalMessageInfo

func (m *ScaleRequest) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

type ScaleReply struct {
	Workers              int32    `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScaleReply) Reset()         { *m = ScaleReply{} }
func (m *ScaleReply) String() string { return proto.CompactTextString(m) }
func (*ScaleReply) ProtoMessage()    {}
func (*ScaleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{3}
}

func (m *ScaleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScaleReply.Unmarshal(m, b)
}
func (m *ScaleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScaleReply.Marshal(b, m, deterministic)
}
func (m *ScaleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScaleReply.Merge(m, src)
}
func (m *ScaleReply) XXX_Size() int {
	return xxx_messageInfo_ScaleReply.Size(m)
}
func (m *ScaleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ScaleReply.DiscardUnknown(m)
}

var xxx_messageInfo_ScaleReply proto.InternalMessageInfo

func (m *ScaleReply) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *ScaleReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ScaleReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{4}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusReply struct {
	Workers              int32    `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	Queue                uint64   `protobuf:"varint,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Tasks                []string `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{5}
}

func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *StatusReply) GetQueue() uint64 {
	if m != nil {
		return m.Queue
	}
	return 0
}

func (m *StatusReply) GetTasks() []string {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec93cfcc38d8076b, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueRequest)(nil), "api.QueueRequest")
	proto.RegisterType((*QueueReply)(nil), "api.QueueReply")
	proto.RegisterType((*ScaleRequest)(nil), "api.ScaleRequest")
	proto.RegisterType((*ScaleReply)(nil), "api.ScaleReply")
	proto.RegisterType((*StatusRequest)(nil), "api.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "api.StatusReply")
	proto.RegisterType((*Error)(nil), "api.Error")
}

func init() { proto.RegisterFile("radish.proto", fileDescriptor_ec93cfcc38d8076b) }

var fileDescriptor_ec93cfcc38d8076b = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x4e, 0xc3, 0x30,
	0x0c, 0x5d, 0xd9, 0xda, 0x31, 0xaf, 0x68, 0x10, 0x21, 0x14, 0xed, 0x34, 0xe5, 0xd4, 0x0b, 0x13,
	0x1a, 0xe2, 0x13, 0xf8, 0x00, 0xb2, 0x2b, 0x97, 0xac, 0x0b, 0x50, 0xad, 0xa3, 0x59, 0xd2, 0x08,
	0xed, 0x4f, 0xf8, 0x5c, 0x64, 0x27, 0x65, 0xe5, 0x00, 0x17, 0x6e, 0x7e, 0xf6, 0x8b, 0x9f, 0xfd,
	0x62, 0xc8, 0xad, 0xda, 0x56, 0xee, 0x6d, 0x69, 0x6c, 0xd3, 0x36, 0x6c, 0xa8, 0x4c, 0x25, 0xde,
	0x21, 0x7f, 0xf2, 0xda, 0x6b, 0xa9, 0x0f, 0x5e, 0xbb, 0x96, 0x31, 0x18, 0xb5, 0xca, 0xed, 0x78,
	0xb2, 0x48, 0x8a, 0x89, 0xa4, 0x98, 0xdd, 0x40, 0x66, 0x94, 0x55, 0x7b, 0xc7, 0xcf, 0x16, 0x49,
	0x91, 0xcb, 0x88, 0x18, 0x87, 0xb1, 0xf3, 0x65, 0xa9, 0x9d, 0xe3, 0x43, 0x2a, 0x74, 0x10, 0x2b,
	0x2f, 0xaa, 0xaa, 0xbd, 0xd5, 0x7c, 0x14, 0x2a, 0x11, 0x8a, 0x67, 0x80, 0xa8, 0x67, 0xea, 0x23,
	0xaa, 0x79, 0x5f, 0x6d, 0x49, 0x2d, 0x97, 0x14, 0xf7, 0xbb, 0xa2, 0xdc, 0xf9, 0xa9, 0xeb, 0x02,
	0x52, 0x6d, 0x6d, 0x63, 0x49, 0x6d, 0xba, 0x82, 0xa5, 0x32, 0xd5, 0xf2, 0x11, 0x33, 0x32, 0x14,
	0x44, 0x01, 0xf9, 0xba, 0x54, 0xf5, 0xf7, 0x36, 0x1c, 0xc6, 0x1f, 0x8d, 0xdd, 0x69, 0xeb, 0x48,
	0x22, 0x95, 0x1d, 0x14, 0x1b, 0x80, 0xc8, 0xc4, 0x39, 0x7e, 0xe5, 0xfd, 0x6b, 0x9a, 0x19, 0x5c,
	0xac, 0x5b, 0xd5, 0x7a, 0x17, 0xc7, 0x11, 0x6b, 0x98, 0x76, 0x89, 0xbf, 0x55, 0xaf, 0x21, 0x3d,
	0xa0, 0x4b, 0xa4, 0x39, 0x92, 0x01, 0x60, 0x16, 0xff, 0x03, 0xdd, 0x1e, 0x16, 0x13, 0x19, 0x80,
	0x78, 0x80, 0x94, 0x54, 0xd1, 0xcc, 0xb2, 0xd9, 0xea, 0xd8, 0x8b, 0x62, 0x94, 0xd8, 0x6b, 0xe7,
	0xd4, 0x6b, 0x68, 0x35, 0x91, 0x1d, 0x5c, 0x7d, 0x26, 0x90, 0x49, 0x3a, 0x07, 0x76, 0x0b, 0x29,
	0xfd, 0x09, 0xbb, 0xa2, 0x1d, 0xfa, 0xf7, 0x30, 0x9f, 0xf5, 0x53, 0xa6, 0x3e, 0x8a, 0x01, 0xd2,
	0xc9, 0xba, 0x48, 0xef, 0x1b, 0x1e, 0xe9, 0x27, 0x67, 0xc5, 0x80, 0xdd, 0x41, 0x16, 0x96, 0x66,
	0x2c, 0x14, 0xfb, 0x96, 0xcc, 0x2f, 0x7f, 0xe4, 0xe8, 0xc5, 0x26, 0xa3, 0xfb, 0xbc, 0xff, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x77, 0x79, 0x34, 0x52, 0xaf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RadishClient is the client API for Radish service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RadishClient interface {
	Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueReply, error)
	Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type radishClient struct {
	cc *grpc.ClientConn
}

func NewRadishClient(cc *grpc.ClientConn) RadishClient {
	return &radishClient{cc}
}

func (c *radishClient) Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueReply, error) {
	out := new(QueueReply)
	err := c.cc.Invoke(ctx, "/api.Radish/Queue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radishClient) Scale(ctx context.Context, in *ScaleRequest, opts ...grpc.CallOption) (*ScaleReply, error) {
	out := new(ScaleReply)
	err := c.cc.Invoke(ctx, "/api.Radish/Scale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radishClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/api.Radish/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RadishServer is the server API for Radish service.
type RadishServer interface {
	Queue(context.Context, *QueueRequest) (*QueueReply, error)
	Scale(context.Context, *ScaleRequest) (*ScaleReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
}

func RegisterRadishServer(s *grpc.Server, srv RadishServer) {
	s.RegisterService(&_Radish_serviceDesc, srv)
}

func _Radish_Queue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadishServer).Queue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Radish/Queue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadishServer).Queue(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Radish_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadishServer).Scale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Radish/Scale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadishServer).Scale(ctx, req.(*ScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Radish_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadishServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Radish/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadishServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Radish_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Radish",
	HandlerType: (*RadishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Queue",
			Handler:    _Radish_Queue_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Radish_Scale_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Radish_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "radish.proto",
}
