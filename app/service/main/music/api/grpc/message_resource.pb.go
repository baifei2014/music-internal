// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_resource.proto

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

type MessageListReq struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageListReq) Reset()         { *m = MessageListReq{} }
func (m *MessageListReq) String() string { return proto.CompactTextString(m) }
func (*MessageListReq) ProtoMessage()    {}
func (*MessageListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c4fa190060ab9f, []int{0}
}

func (m *MessageListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageListReq.Unmarshal(m, b)
}
func (m *MessageListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageListReq.Marshal(b, m, deterministic)
}
func (m *MessageListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageListReq.Merge(m, src)
}
func (m *MessageListReq) XXX_Size() int {
	return xxx_messageInfo_MessageListReq.Size(m)
}
func (m *MessageListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageListReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageListReq proto.InternalMessageInfo

func (m *MessageListReq) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *MessageListReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *MessageListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type MessageListResp struct {
	More                 bool                    `protobuf:"varint,1,opt,name=more,proto3" json:"more,omitempty"`
	CurrentPage          int64                   `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	List                 []*MessageListResp_List `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MessageListResp) Reset()         { *m = MessageListResp{} }
func (m *MessageListResp) String() string { return proto.CompactTextString(m) }
func (*MessageListResp) ProtoMessage()    {}
func (*MessageListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c4fa190060ab9f, []int{1}
}

func (m *MessageListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageListResp.Unmarshal(m, b)
}
func (m *MessageListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageListResp.Marshal(b, m, deterministic)
}
func (m *MessageListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageListResp.Merge(m, src)
}
func (m *MessageListResp) XXX_Size() int {
	return xxx_messageInfo_MessageListResp.Size(m)
}
func (m *MessageListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageListResp.DiscardUnknown(m)
}

var xxx_messageInfo_MessageListResp proto.InternalMessageInfo

func (m *MessageListResp) GetMore() bool {
	if m != nil {
		return m.More
	}
	return false
}

func (m *MessageListResp) GetCurrentPage() int64 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *MessageListResp) GetList() []*MessageListResp_List {
	if m != nil {
		return m.List
	}
	return nil
}

type MessageListResp_List struct {
	FromUid              int64    `protobuf:"varint,1,opt,name=from_uid,json=fromUid,proto3" json:"from_uid,omitempty"`
	ToUid                int64    `protobuf:"varint,2,opt,name=to_uid,json=toUid,proto3" json:"to_uid,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageListResp_List) Reset()         { *m = MessageListResp_List{} }
func (m *MessageListResp_List) String() string { return proto.CompactTextString(m) }
func (*MessageListResp_List) ProtoMessage()    {}
func (*MessageListResp_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0c4fa190060ab9f, []int{1, 0}
}

func (m *MessageListResp_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageListResp_List.Unmarshal(m, b)
}
func (m *MessageListResp_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageListResp_List.Marshal(b, m, deterministic)
}
func (m *MessageListResp_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageListResp_List.Merge(m, src)
}
func (m *MessageListResp_List) XXX_Size() int {
	return xxx_messageInfo_MessageListResp_List.Size(m)
}
func (m *MessageListResp_List) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageListResp_List.DiscardUnknown(m)
}

var xxx_messageInfo_MessageListResp_List proto.InternalMessageInfo

func (m *MessageListResp_List) GetFromUid() int64 {
	if m != nil {
		return m.FromUid
	}
	return 0
}

func (m *MessageListResp_List) GetToUid() int64 {
	if m != nil {
		return m.ToUid
	}
	return 0
}

func (m *MessageListResp_List) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageListReq)(nil), "main.music.MessageListReq")
	proto.RegisterType((*MessageListResp)(nil), "main.music.MessageListResp")
	proto.RegisterType((*MessageListResp_List)(nil), "main.music.MessageListResp.List")
}

func init() { proto.RegisterFile("message_resource.proto", fileDescriptor_c0c4fa190060ab9f) }

var fileDescriptor_c0c4fa190060ab9f = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x25, 0x4d, 0x6c, 0xd3, 0x49, 0x7d, 0x30, 0x0b, 0x09, 0x51, 0x48, 0x2c, 0x88, 0x45, 0x30,
	0xc5, 0x76, 0xef, 0x22, 0x0b, 0x41, 0x50, 0x90, 0x11, 0x17, 0xba, 0x09, 0x6d, 0x3a, 0x8d, 0x03,
	0x26, 0x93, 0xce, 0x43, 0xb0, 0x5b, 0x3f, 0xc4, 0x3f, 0xcb, 0x07, 0xe4, 0x2b, 0x64, 0x26, 0xe9,
	0x4b, 0xd0, 0xcd, 0xcd, 0x39, 0x87, 0x3b, 0x77, 0xce, 0x3d, 0x19, 0x70, 0x9c, 0x61, 0xce, 0x27,
	0x29, 0x8e, 0x19, 0xe6, 0x54, 0xb2, 0x04, 0x87, 0x05, 0xa3, 0x82, 0x42, 0x90, 0x4d, 0x48, 0x1e,
	0x66, 0x92, 0x93, 0xc4, 0xbb, 0x4a, 0x89, 0x78, 0x93, 0xd3, 0x30, 0xa1, 0xd9, 0x30, 0xa5, 0x29,
	0x1d, 0xea, 0x96, 0xa9, 0x9c, 0x6b, 0xa6, 0x89, 0x46, 0xf5, 0xd1, 0xfe, 0x97, 0x01, 0x0e, 0x1e,
	0xea, 0xa9, 0xf7, 0x84, 0x0b, 0x84, 0x17, 0xf0, 0x02, 0xd8, 0x29, 0xa3, 0xb2, 0x88, 0xc9, 0xcc,
	0x35, 0x02, 0x63, 0xd0, 0x8d, 0x7a, 0x55, 0xe9, 0xaf, 0x35, 0xd4, 0xd1, 0xe8, 0x6e, 0x06, 0x4f,
	0x81, 0x55, 0x4c, 0x52, 0xec, 0xb6, 0x02, 0x63, 0x60, 0x46, 0x76, 0x55, 0xfa, 0x9a, 0x23, 0x5d,
	0xe1, 0x25, 0xe8, 0xaa, 0x6f, 0xcc, 0xc9, 0x12, 0xbb, 0xa6, 0x6e, 0xd9, 0xaf, 0x4a, 0x7f, 0x23,
	0x22, 0x5b, 0xc1, 0x27, 0xb2, 0xc4, 0xfd, 0xef, 0x16, 0x38, 0xdc, 0x71, 0xc1, 0x0b, 0x35, 0x3d,
	0xa3, 0x0c, 0x6b, 0x0b, 0x76, 0x3d, 0x5d, 0x71, 0xa4, 0x2b, 0x1c, 0x83, 0x5e, 0x22, 0x19, 0xc3,
	0xb9, 0x88, 0xb7, 0x3c, 0x1c, 0x55, 0xa5, 0xbf, 0xa3, 0x23, 0xa7, 0x61, 0x8f, 0xca, 0xd2, 0x0d,
	0xb0, 0xde, 0x09, 0x17, 0xae, 0x19, 0x98, 0x03, 0x67, 0x14, 0x84, 0x9b, 0xd8, 0xc2, 0x5f, 0xb7,
	0x87, 0x0a, 0xd4, 0x97, 0xaa, 0x13, 0x48, 0x57, 0xef, 0x13, 0x58, 0x4a, 0x57, 0x09, 0xcd, 0x19,
	0xcd, 0x62, 0xd9, 0x24, 0x64, 0xd6, 0x09, 0xad, 0x34, 0xd4, 0x51, 0xe8, 0x99, 0xcc, 0xe0, 0x19,
	0x68, 0x0b, 0xaa, 0xdb, 0x6a, 0x7f, 0xa0, 0x2a, 0xfd, 0x46, 0x41, 0x7b, 0x82, 0xaa, 0x96, 0x73,
	0xd0, 0x49, 0x68, 0x2e, 0x70, 0x2e, 0x74, 0x48, 0xdd, 0xc8, 0xa9, 0x4a, 0x7f, 0x25, 0xa1, 0x15,
	0x18, 0xbd, 0xac, 0x03, 0x42, 0xcd, 0xbf, 0x87, 0xb7, 0xc0, 0xd9, 0x72, 0x0d, 0xbd, 0x3f, 0xd7,
	0x59, 0x78, 0x27, 0xff, 0xac, 0x1a, 0x59, 0xaf, 0xad, 0x8f, 0xeb, 0x69, 0x5b, 0xbf, 0x87, 0xf1,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x79, 0x9d, 0x76, 0x64, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageResourceClient is the client API for MessageResource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageResourceClient interface {
	MessageList(ctx context.Context, in *MessageListReq, opts ...grpc.CallOption) (*MessageListResp, error)
}

type messageResourceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageResourceClient(cc grpc.ClientConnInterface) MessageResourceClient {
	return &messageResourceClient{cc}
}

func (c *messageResourceClient) MessageList(ctx context.Context, in *MessageListReq, opts ...grpc.CallOption) (*MessageListResp, error) {
	out := new(MessageListResp)
	err := c.cc.Invoke(ctx, "/main.music.MessageResource/MessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageResourceServer is the server API for MessageResource service.
type MessageResourceServer interface {
	MessageList(context.Context, *MessageListReq) (*MessageListResp, error)
}

// UnimplementedMessageResourceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageResourceServer struct {
}

func (*UnimplementedMessageResourceServer) MessageList(ctx context.Context, req *MessageListReq) (*MessageListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageList not implemented")
}

func RegisterMessageResourceServer(s *grpc.Server, srv MessageResourceServer) {
	s.RegisterService(&_MessageResource_serviceDesc, srv)
}

func _MessageResource_MessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageResourceServer).MessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.music.MessageResource/MessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageResourceServer).MessageList(ctx, req.(*MessageListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageResource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.music.MessageResource",
	HandlerType: (*MessageResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageList",
			Handler:    _MessageResource_MessageList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_resource.proto",
}
