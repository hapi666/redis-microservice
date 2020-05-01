// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redis.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PushQueueRequest struct {
	Urls                 []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushQueueRequest) Reset()         { *m = PushQueueRequest{} }
func (m *PushQueueRequest) String() string { return proto.CompactTextString(m) }
func (*PushQueueRequest) ProtoMessage()    {}
func (*PushQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{0}
}

func (m *PushQueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushQueueRequest.Unmarshal(m, b)
}
func (m *PushQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushQueueRequest.Marshal(b, m, deterministic)
}
func (m *PushQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushQueueRequest.Merge(m, src)
}
func (m *PushQueueRequest) XXX_Size() int {
	return xxx_messageInfo_PushQueueRequest.Size(m)
}
func (m *PushQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushQueueRequest proto.InternalMessageInfo

func (m *PushQueueRequest) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type PopQueueReply struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PopQueueReply) Reset()         { *m = PopQueueReply{} }
func (m *PopQueueReply) String() string { return proto.CompactTextString(m) }
func (*PopQueueReply) ProtoMessage()    {}
func (*PopQueueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{1}
}

func (m *PopQueueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PopQueueReply.Unmarshal(m, b)
}
func (m *PopQueueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PopQueueReply.Marshal(b, m, deterministic)
}
func (m *PopQueueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PopQueueReply.Merge(m, src)
}
func (m *PopQueueReply) XXX_Size() int {
	return xxx_messageInfo_PopQueueReply.Size(m)
}
func (m *PopQueueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PopQueueReply.DiscardUnknown(m)
}

var xxx_messageInfo_PopQueueReply proto.InternalMessageInfo

func (m *PopQueueReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type RangeQueueReply struct {
	Urls                 []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeQueueReply) Reset()         { *m = RangeQueueReply{} }
func (m *RangeQueueReply) String() string { return proto.CompactTextString(m) }
func (*RangeQueueReply) ProtoMessage()    {}
func (*RangeQueueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{2}
}

func (m *RangeQueueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangeQueueReply.Unmarshal(m, b)
}
func (m *RangeQueueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangeQueueReply.Marshal(b, m, deterministic)
}
func (m *RangeQueueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeQueueReply.Merge(m, src)
}
func (m *RangeQueueReply) XXX_Size() int {
	return xxx_messageInfo_RangeQueueReply.Size(m)
}
func (m *RangeQueueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeQueueReply.DiscardUnknown(m)
}

var xxx_messageInfo_RangeQueueReply proto.InternalMessageInfo

func (m *RangeQueueReply) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type SisMemberRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SisMemberRequest) Reset()         { *m = SisMemberRequest{} }
func (m *SisMemberRequest) String() string { return proto.CompactTextString(m) }
func (*SisMemberRequest) ProtoMessage()    {}
func (*SisMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{3}
}

func (m *SisMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SisMemberRequest.Unmarshal(m, b)
}
func (m *SisMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SisMemberRequest.Marshal(b, m, deterministic)
}
func (m *SisMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SisMemberRequest.Merge(m, src)
}
func (m *SisMemberRequest) XXX_Size() int {
	return xxx_messageInfo_SisMemberRequest.Size(m)
}
func (m *SisMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SisMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SisMemberRequest proto.InternalMessageInfo

func (m *SisMemberRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type SisMemberReply struct {
	IsExist              bool     `protobuf:"varint,1,opt,name=isExist,proto3" json:"isExist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SisMemberReply) Reset()         { *m = SisMemberReply{} }
func (m *SisMemberReply) String() string { return proto.CompactTextString(m) }
func (*SisMemberReply) ProtoMessage()    {}
func (*SisMemberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{4}
}

func (m *SisMemberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SisMemberReply.Unmarshal(m, b)
}
func (m *SisMemberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SisMemberReply.Marshal(b, m, deterministic)
}
func (m *SisMemberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SisMemberReply.Merge(m, src)
}
func (m *SisMemberReply) XXX_Size() int {
	return xxx_messageInfo_SisMemberReply.Size(m)
}
func (m *SisMemberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SisMemberReply.DiscardUnknown(m)
}

var xxx_messageInfo_SisMemberReply proto.InternalMessageInfo

func (m *SisMemberReply) GetIsExist() bool {
	if m != nil {
		return m.IsExist
	}
	return false
}

type SadDRequest struct {
	CrawledURL           []string `protobuf:"bytes,1,rep,name=crawledURL,proto3" json:"crawledURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SadDRequest) Reset()         { *m = SadDRequest{} }
func (m *SadDRequest) String() string { return proto.CompactTextString(m) }
func (*SadDRequest) ProtoMessage()    {}
func (*SadDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d954120a2319ff8f, []int{5}
}

func (m *SadDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SadDRequest.Unmarshal(m, b)
}
func (m *SadDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SadDRequest.Marshal(b, m, deterministic)
}
func (m *SadDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SadDRequest.Merge(m, src)
}
func (m *SadDRequest) XXX_Size() int {
	return xxx_messageInfo_SadDRequest.Size(m)
}
func (m *SadDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SadDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SadDRequest proto.InternalMessageInfo

func (m *SadDRequest) GetCrawledURL() []string {
	if m != nil {
		return m.CrawledURL
	}
	return nil
}

func init() {
	proto.RegisterType((*PushQueueRequest)(nil), "proto.PushQueueRequest")
	proto.RegisterType((*PopQueueReply)(nil), "proto.PopQueueReply")
	proto.RegisterType((*RangeQueueReply)(nil), "proto.RangeQueueReply")
	proto.RegisterType((*SisMemberRequest)(nil), "proto.SisMemberRequest")
	proto.RegisterType((*SisMemberReply)(nil), "proto.SisMemberReply")
	proto.RegisterType((*SadDRequest)(nil), "proto.SadDRequest")
}

func init() { proto.RegisterFile("redis.proto", fileDescriptor_d954120a2319ff8f) }

var fileDescriptor_d954120a2319ff8f = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x5b, 0x5b, 0xb5, 0x99, 0xa2, 0x96, 0x41, 0x6b, 0x89, 0x20, 0x75, 0x51, 0x29, 0x82,
	0x29, 0xa8, 0x27, 0xc1, 0x1e, 0xa4, 0xbd, 0x29, 0xd4, 0x2d, 0x3e, 0x40, 0xd3, 0x8c, 0x31, 0x90,
	0x9a, 0xb8, 0x9b, 0x55, 0xfb, 0x6e, 0x3e, 0x9c, 0xe4, 0xcf, 0xa6, 0x21, 0x35, 0xa7, 0xcc, 0x4c,
	0xbe, 0xf9, 0x4d, 0xf2, 0x41, 0x5b, 0x90, 0xe3, 0x49, 0x2b, 0x14, 0x41, 0x14, 0xe0, 0x76, 0xf2,
	0x30, 0x4f, 0xdc, 0x20, 0x70, 0x7d, 0x1a, 0x26, 0x9d, 0xad, 0xde, 0x86, 0xb4, 0x0c, 0xa3, 0x55,
	0xca, 0xb0, 0x4b, 0xe8, 0x4c, 0x95, 0x7c, 0x7f, 0x51, 0xa4, 0x88, 0xd3, 0xa7, 0x22, 0x19, 0x21,
	0x42, 0x53, 0x09, 0x5f, 0xf6, 0xea, 0xfd, 0xc6, 0xc0, 0xe0, 0x49, 0xcd, 0xce, 0x60, 0x6f, 0x1a,
	0x84, 0x19, 0x16, 0xfa, 0x2b, 0xec, 0x40, 0x43, 0x09, 0xbf, 0x57, 0xef, 0xd7, 0x07, 0x06, 0x8f,
	0x4b, 0x76, 0x01, 0x07, 0x7c, 0xfe, 0xe1, 0x52, 0x01, 0xfa, 0x2f, 0xe9, 0x1c, 0x3a, 0x33, 0x4f,
	0x3e, 0xd3, 0xd2, 0x26, 0xa1, 0x2f, 0x6e, 0x86, 0x5d, 0xc1, 0x7e, 0x81, 0x8a, 0xb3, 0x7a, 0xb0,
	0xeb, 0xc9, 0xc9, 0x8f, 0x27, 0xa3, 0x84, 0x6b, 0x71, 0xdd, 0xb2, 0x6b, 0x68, 0xcf, 0xe6, 0xce,
	0x58, 0x87, 0x9d, 0x02, 0x2c, 0xc4, 0xfc, 0xdb, 0x27, 0xe7, 0x95, 0x3f, 0x65, 0xa7, 0x0b, 0x93,
	0x9b, 0xdf, 0x2d, 0x30, 0xc6, 0x8f, 0x33, 0x12, 0x5f, 0xde, 0x82, 0x70, 0x04, 0x46, 0x2e, 0x00,
	0x8f, 0x53, 0x2b, 0x56, 0x59, 0x89, 0xd9, 0xb5, 0x52, 0x89, 0x96, 0x96, 0x68, 0x4d, 0x62, 0x89,
	0xac, 0x86, 0xf7, 0xd0, 0xd2, 0x62, 0xb0, 0x82, 0x32, 0x0f, 0x75, 0x6c, 0xd1, 0x20, 0xab, 0xe1,
	0x08, 0x60, 0x6d, 0xac, 0x72, 0xbb, 0x9b, 0x6d, 0x97, 0xe4, 0xb2, 0x1a, 0x3e, 0x80, 0x91, 0x4b,
	0xca, 0xbf, 0xbd, 0x2c, 0xd7, 0x3c, 0xda, 0x7c, 0x91, 0xae, 0xdf, 0x41, 0x33, 0xf6, 0x86, 0xa8,
	0x81, 0xb5, 0xc4, 0xea, 0x1f, 0xb6, 0x77, 0x92, 0xc9, 0xed, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x64, 0x75, 0x85, 0xa7, 0x6b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBServiceClient interface {
	PushQueue(ctx context.Context, in *PushQueueRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	PopQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PopQueueReply, error)
	RangeQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RangeQueueReply, error)
	SisMember(ctx context.Context, in *SisMemberRequest, opts ...grpc.CallOption) (*SisMemberReply, error)
	SadD(ctx context.Context, in *SadDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dBServiceClient struct {
	cc *grpc.ClientConn
}

func NewDBServiceClient(cc *grpc.ClientConn) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) PushQueue(ctx context.Context, in *PushQueueRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.DBService/PushQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) PopQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PopQueueReply, error) {
	out := new(PopQueueReply)
	err := c.cc.Invoke(ctx, "/proto.DBService/PopQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) RangeQueue(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*RangeQueueReply, error) {
	out := new(RangeQueueReply)
	err := c.cc.Invoke(ctx, "/proto.DBService/RangeQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) SisMember(ctx context.Context, in *SisMemberRequest, opts ...grpc.CallOption) (*SisMemberReply, error) {
	out := new(SisMemberReply)
	err := c.cc.Invoke(ctx, "/proto.DBService/SisMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) SadD(ctx context.Context, in *SadDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.DBService/SadD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
type DBServiceServer interface {
	PushQueue(context.Context, *PushQueueRequest) (*empty.Empty, error)
	PopQueue(context.Context, *empty.Empty) (*PopQueueReply, error)
	RangeQueue(context.Context, *empty.Empty) (*RangeQueueReply, error)
	SisMember(context.Context, *SisMemberRequest) (*SisMemberReply, error)
	SadD(context.Context, *SadDRequest) (*empty.Empty, error)
}

// UnimplementedDBServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (*UnimplementedDBServiceServer) PushQueue(ctx context.Context, req *PushQueueRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushQueue not implemented")
}
func (*UnimplementedDBServiceServer) PopQueue(ctx context.Context, req *empty.Empty) (*PopQueueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PopQueue not implemented")
}
func (*UnimplementedDBServiceServer) RangeQueue(ctx context.Context, req *empty.Empty) (*RangeQueueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RangeQueue not implemented")
}
func (*UnimplementedDBServiceServer) SisMember(ctx context.Context, req *SisMemberRequest) (*SisMemberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SisMember not implemented")
}
func (*UnimplementedDBServiceServer) SadD(ctx context.Context, req *SadDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SadD not implemented")
}

func RegisterDBServiceServer(s *grpc.Server, srv DBServiceServer) {
	s.RegisterService(&_DBService_serviceDesc, srv)
}

func _DBService_PushQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).PushQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBService/PushQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).PushQueue(ctx, req.(*PushQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_PopQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).PopQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBService/PopQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).PopQueue(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_RangeQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).RangeQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBService/RangeQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).RangeQueue(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_SisMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SisMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).SisMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBService/SisMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).SisMember(ctx, req.(*SisMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_SadD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SadDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).SadD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBService/SadD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).SadD(ctx, req.(*SadDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushQueue",
			Handler:    _DBService_PushQueue_Handler,
		},
		{
			MethodName: "PopQueue",
			Handler:    _DBService_PopQueue_Handler,
		},
		{
			MethodName: "RangeQueue",
			Handler:    _DBService_RangeQueue_Handler,
		},
		{
			MethodName: "SisMember",
			Handler:    _DBService_SisMember_Handler,
		},
		{
			MethodName: "SadD",
			Handler:    _DBService_SadD_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redis.proto",
}
