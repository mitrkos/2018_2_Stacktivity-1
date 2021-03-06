// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session-server.proto

package session

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type SessionID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionID) Reset()         { *m = SessionID{} }
func (m *SessionID) String() string { return proto.CompactTextString(m) }
func (*SessionID) ProtoMessage()    {}
func (*SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
}

func (m *SessionID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionID.Unmarshal(m, b)
}
func (m *SessionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionID.Marshal(b, m, deterministic)
}
func (m *SessionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionID.Merge(m, src)
}
func (m *SessionID) XXX_Size() int {
	return xxx_messageInfo_SessionID.Size(m)
}
func (m *SessionID) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionID.DiscardUnknown(m)
}

var xxx_messageInfo_SessionID proto.InternalMessageInfo

func (m *SessionID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Session struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Nothing struct {
	Dummy                bool     `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func (m *Nothing) GetDummy() bool {
	if m != nil {
		return m.Dummy
	}
	return false
}

func init() {
	proto.RegisterType((*SessionID)(nil), "session-server.SessionID")
	proto.RegisterType((*Session)(nil), "session-server.Session")
	proto.RegisterType((*Nothing)(nil), "session-server.Nothing")
}

func init() { proto.RegisterFile("session-server.proto", fileDescriptor_3a6be1b361fa6f14) }

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb9,
	0x38, 0x83, 0x21, 0x4c, 0x4f, 0x17, 0x21, 0x3e, 0x2e, 0x26, 0x4f, 0x17, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x26, 0x4f, 0x17, 0x25, 0x49, 0x2e, 0x76, 0xa8, 0x24, 0x92, 0x14, 0x2b, 0x58,
	0x4a, 0x9e, 0x8b, 0xdd, 0x2f, 0xbf, 0x24, 0x23, 0x33, 0x2f, 0x5d, 0x48, 0x84, 0x8b, 0x35, 0xa5,
	0x34, 0x37, 0xb7, 0x12, 0x2c, 0xcb, 0x11, 0x04, 0xe1, 0x18, 0x2d, 0x65, 0xe4, 0xe2, 0x83, 0x6a,
	0xf6, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0x12, 0x32, 0xe0, 0x62, 0x73, 0x2e, 0x4a, 0x4d, 0x2c,
	0x49, 0x15, 0x12, 0xd0, 0x83, 0x39, 0x07, 0xaa, 0x44, 0x4a, 0x08, 0x5d, 0xc4, 0xd3, 0x45, 0x89,
	0x41, 0x48, 0x9f, 0x8b, 0xd5, 0x39, 0x23, 0x35, 0x39, 0x5b, 0x08, 0x8b, 0xb4, 0x14, 0x86, 0x21,
	0x4a, 0x0c, 0x20, 0x2b, 0x5c, 0x52, 0x73, 0x52, 0x4b, 0x52, 0x09, 0xe8, 0x80, 0xba, 0x5d, 0x89,
	0x21, 0x89, 0x0d, 0x1c, 0x20, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x0f, 0xd8, 0x25,
	0x21, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionManagerClient is the client API for SessionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionManagerClient interface {
	Create(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionID, error)
	Check(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error)
	Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error)
}

type sessionManagerClient struct {
	cc *grpc.ClientConn
}

func NewSessionManagerClient(cc *grpc.ClientConn) SessionManagerClient {
	return &sessionManagerClient{cc}
}

func (c *sessionManagerClient) Create(ctx context.Context, in *Session, opts ...grpc.CallOption) (*SessionID, error) {
	out := new(SessionID)
	err := c.cc.Invoke(ctx, "/session-server.SessionManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Check(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/session-server.SessionManager/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session-server.SessionManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionManagerServer is the public-api-server API for SessionManager service.
type SessionManagerServer interface {
	Create(context.Context, *Session) (*SessionID, error)
	Check(context.Context, *SessionID) (*Session, error)
	Delete(context.Context, *SessionID) (*Nothing, error)
}

func RegisterSessionManagerServer(s *grpc.Server, srv SessionManagerServer) {
	s.RegisterService(&_SessionManager_serviceDesc, srv)
}

func _SessionManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session-server.SessionManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Create(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session-server.SessionManager/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Check(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session-server.SessionManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Delete(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session-server.SessionManager",
	HandlerType: (*SessionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SessionManager_Create_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _SessionManager_Check_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SessionManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session-server.proto",
}
