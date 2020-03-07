// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_resource.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ListUserInfoReq struct {
	Uid                  []int64  `protobuf:"varint,1,rep,packed,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserInfoReq) Reset()         { *m = ListUserInfoReq{} }
func (m *ListUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*ListUserInfoReq) ProtoMessage()    {}
func (*ListUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f23b72a6a2080599, []int{0}
}

func (m *ListUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserInfoReq.Unmarshal(m, b)
}
func (m *ListUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserInfoReq.Marshal(b, m, deterministic)
}
func (m *ListUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserInfoReq.Merge(m, src)
}
func (m *ListUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_ListUserInfoReq.Size(m)
}
func (m *ListUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserInfoReq proto.InternalMessageInfo

func (m *ListUserInfoReq) GetUid() []int64 {
	if m != nil {
		return m.Uid
	}
	return nil
}

type ListUserInfoResp struct {
	List                 []*ListUserInfoResp_List `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListUserInfoResp) Reset()         { *m = ListUserInfoResp{} }
func (m *ListUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*ListUserInfoResp) ProtoMessage()    {}
func (*ListUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f23b72a6a2080599, []int{1}
}

func (m *ListUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserInfoResp.Unmarshal(m, b)
}
func (m *ListUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserInfoResp.Marshal(b, m, deterministic)
}
func (m *ListUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserInfoResp.Merge(m, src)
}
func (m *ListUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_ListUserInfoResp.Size(m)
}
func (m *ListUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserInfoResp proto.InternalMessageInfo

func (m *ListUserInfoResp) GetList() []*ListUserInfoResp_List {
	if m != nil {
		return m.List
	}
	return nil
}

type ListUserInfoResp_List struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserInfoResp_List) Reset()         { *m = ListUserInfoResp_List{} }
func (m *ListUserInfoResp_List) String() string { return proto.CompactTextString(m) }
func (*ListUserInfoResp_List) ProtoMessage()    {}
func (*ListUserInfoResp_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_f23b72a6a2080599, []int{1, 0}
}

func (m *ListUserInfoResp_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserInfoResp_List.Unmarshal(m, b)
}
func (m *ListUserInfoResp_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserInfoResp_List.Marshal(b, m, deterministic)
}
func (m *ListUserInfoResp_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserInfoResp_List.Merge(m, src)
}
func (m *ListUserInfoResp_List) XXX_Size() int {
	return xxx_messageInfo_ListUserInfoResp_List.Size(m)
}
func (m *ListUserInfoResp_List) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserInfoResp_List.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserInfoResp_List proto.InternalMessageInfo

func (m *ListUserInfoResp_List) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ListUserInfoResp_List) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *ListUserInfoResp_List) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*ListUserInfoReq)(nil), "main.music.ListUserInfoReq")
	proto.RegisterType((*ListUserInfoResp)(nil), "main.music.ListUserInfoResp")
	proto.RegisterType((*ListUserInfoResp_List)(nil), "main.music.ListUserInfoResp.List")
}

func init() { proto.RegisterFile("user_resource.proto", fileDescriptor_f23b72a6a2080599) }

var fileDescriptor_f23b72a6a2080599 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0xca, 0x4d, 0xcc, 0xcc, 0xd3, 0xcb, 0x2d, 0x2d, 0xce, 0x4c, 0x96, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x49,
	0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x55, 0x49, 0x87, 0x8b, 0xdf, 0x27, 0xb3,
	0xb8, 0x24, 0xb4, 0x38, 0xb5, 0xc8, 0x33, 0x2f, 0x2d, 0x3f, 0x28, 0xb5, 0x50, 0x48, 0x92, 0x8b,
	0xb9, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x59, 0x83, 0xd9, 0x89, 0xfd, 0xd5, 0x3d, 0x79, 0x10,
	0x37, 0x08, 0x44, 0x28, 0x9d, 0x67, 0xe4, 0x12, 0x40, 0x55, 0x5e, 0x5c, 0x20, 0x64, 0xcf, 0xc5,
	0x92, 0x93, 0x59, 0x5c, 0x02, 0xd6, 0xc0, 0x6d, 0xa4, 0xa8, 0x87, 0x70, 0x8c, 0x1e, 0xba, 0x5a,
	0xb0, 0x80, 0x13, 0xc7, 0xab, 0x7b, 0xf2, 0x60, 0x2d, 0x41, 0x60, 0x52, 0xaa, 0x8a, 0x8b, 0x05,
	0x24, 0x8e, 0xb0, 0x98, 0x11, 0xdd, 0x62, 0x21, 0x0d, 0x2e, 0x8e, 0xbc, 0xcc, 0xe4, 0xec, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x9e, 0x57, 0xf7, 0xe4, 0xe1, 0x62,
	0x41, 0x70, 0x96, 0x90, 0x2e, 0x17, 0x57, 0x62, 0x59, 0x62, 0x49, 0x62, 0x51, 0x7c, 0x69, 0x51,
	0x8e, 0x04, 0x33, 0x58, 0x2d, 0xdf, 0xab, 0x7b, 0xf2, 0x48, 0xa2, 0x41, 0x9c, 0x10, 0x76, 0x68,
	0x51, 0x8e, 0x51, 0x24, 0x17, 0x0f, 0xc8, 0x81, 0x41, 0xd0, 0x00, 0x15, 0xf2, 0xe4, 0xe2, 0x41,
	0x76, 0xb4, 0x90, 0x34, 0x6e, 0xef, 0x14, 0x4a, 0xc9, 0xe0, 0xf3, 0xab, 0x13, 0x4b, 0x14, 0x53,
	0x99, 0x61, 0x12, 0x1b, 0x38, 0x9c, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x7d, 0x76,
	0x37, 0xb9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserResourceClient is the client API for UserResource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserResourceClient interface {
	ListUserInfo(ctx context.Context, in *ListUserInfoReq, opts ...grpc.CallOption) (*ListUserInfoResp, error)
}

type userResourceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserResourceClient(cc grpc.ClientConnInterface) UserResourceClient {
	return &userResourceClient{cc}
}

func (c *userResourceClient) ListUserInfo(ctx context.Context, in *ListUserInfoReq, opts ...grpc.CallOption) (*ListUserInfoResp, error) {
	out := new(ListUserInfoResp)
	err := c.cc.Invoke(ctx, "/main.music.UserResource/ListUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserResourceServer is the server API for UserResource service.
type UserResourceServer interface {
	ListUserInfo(context.Context, *ListUserInfoReq) (*ListUserInfoResp, error)
}

// UnimplementedUserResourceServer can be embedded to have forward compatible implementations.
type UnimplementedUserResourceServer struct {
}

func (*UnimplementedUserResourceServer) ListUserInfo(ctx context.Context, req *ListUserInfoReq) (*ListUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserInfo not implemented")
}

func RegisterUserResourceServer(s *grpc.Server, srv UserResourceServer) {
	s.RegisterService(&_UserResource_serviceDesc, srv)
}

func _UserResource_ListUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserResourceServer).ListUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.music.UserResource/ListUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserResourceServer).ListUserInfo(ctx, req.(*ListUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserResource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.music.UserResource",
	HandlerType: (*UserResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUserInfo",
			Handler:    _UserResource_ListUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_resource.proto",
}
