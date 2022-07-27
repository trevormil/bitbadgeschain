// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgNewBadge struct {
	Creator               string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Uri                   string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Manager               string `protobuf:"bytes,4,opt,name=manager,proto3" json:"manager,omitempty"`
	Permissions           uint64 `protobuf:"varint,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	FreezeAddressesDigest string `protobuf:"bytes,6,opt,name=freezeAddressesDigest,proto3" json:"freezeAddressesDigest,omitempty"`
	SubassetUris          string `protobuf:"bytes,7,opt,name=subassetUris,proto3" json:"subassetUris,omitempty"`
}

func (m *MsgNewBadge) Reset()         { *m = MsgNewBadge{} }
func (m *MsgNewBadge) String() string { return proto.CompactTextString(m) }
func (*MsgNewBadge) ProtoMessage()    {}
func (*MsgNewBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc897b33479788c9, []int{0}
}
func (m *MsgNewBadge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewBadge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewBadge.Merge(m, src)
}
func (m *MsgNewBadge) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewBadge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewBadge proto.InternalMessageInfo

func (m *MsgNewBadge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgNewBadge) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *MsgNewBadge) GetManager() string {
	if m != nil {
		return m.Manager
	}
	return ""
}

func (m *MsgNewBadge) GetPermissions() uint64 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *MsgNewBadge) GetFreezeAddressesDigest() string {
	if m != nil {
		return m.FreezeAddressesDigest
	}
	return ""
}

func (m *MsgNewBadge) GetSubassetUris() string {
	if m != nil {
		return m.SubassetUris
	}
	return ""
}

type MsgNewBadgeResponse struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MsgNewBadgeResponse) Reset()         { *m = MsgNewBadgeResponse{} }
func (m *MsgNewBadgeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewBadgeResponse) ProtoMessage()    {}
func (*MsgNewBadgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc897b33479788c9, []int{1}
}
func (m *MsgNewBadgeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewBadgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewBadgeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewBadgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewBadgeResponse.Merge(m, src)
}
func (m *MsgNewBadgeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewBadgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewBadgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewBadgeResponse proto.InternalMessageInfo

func (m *MsgNewBadgeResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgNewBadgeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MsgNewSubBadge struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Supply  uint64 `protobuf:"varint,3,opt,name=supply,proto3" json:"supply,omitempty"`
}

func (m *MsgNewSubBadge) Reset()         { *m = MsgNewSubBadge{} }
func (m *MsgNewSubBadge) String() string { return proto.CompactTextString(m) }
func (*MsgNewSubBadge) ProtoMessage()    {}
func (*MsgNewSubBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc897b33479788c9, []int{2}
}
func (m *MsgNewSubBadge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewSubBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewSubBadge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewSubBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewSubBadge.Merge(m, src)
}
func (m *MsgNewSubBadge) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewSubBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewSubBadge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewSubBadge proto.InternalMessageInfo

func (m *MsgNewSubBadge) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgNewSubBadge) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgNewSubBadge) GetSupply() uint64 {
	if m != nil {
		return m.Supply
	}
	return 0
}

type MsgNewSubBadgeResponse struct {
	SubassetId uint64 `protobuf:"varint,1,opt,name=subassetId,proto3" json:"subassetId,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *MsgNewSubBadgeResponse) Reset()         { *m = MsgNewSubBadgeResponse{} }
func (m *MsgNewSubBadgeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewSubBadgeResponse) ProtoMessage()    {}
func (*MsgNewSubBadgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc897b33479788c9, []int{3}
}
func (m *MsgNewSubBadgeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewSubBadgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewSubBadgeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewSubBadgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewSubBadgeResponse.Merge(m, src)
}
func (m *MsgNewSubBadgeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewSubBadgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewSubBadgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewSubBadgeResponse proto.InternalMessageInfo

func (m *MsgNewSubBadgeResponse) GetSubassetId() uint64 {
	if m != nil {
		return m.SubassetId
	}
	return 0
}

func (m *MsgNewSubBadgeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgNewBadge)(nil), "trevormil.bitbadgeschain.badges.MsgNewBadge")
	proto.RegisterType((*MsgNewBadgeResponse)(nil), "trevormil.bitbadgeschain.badges.MsgNewBadgeResponse")
	proto.RegisterType((*MsgNewSubBadge)(nil), "trevormil.bitbadgeschain.badges.MsgNewSubBadge")
	proto.RegisterType((*MsgNewSubBadgeResponse)(nil), "trevormil.bitbadgeschain.badges.MsgNewSubBadgeResponse")
}

func init() { proto.RegisterFile("badges/tx.proto", fileDescriptor_bc897b33479788c9) }

var fileDescriptor_bc897b33479788c9 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x99, 0x52, 0x41, 0x0f, 0x06, 0xcd, 0x18, 0xc9, 0x84, 0x45, 0x25, 0x5d, 0xb1, 0x30,
	0x6d, 0x82, 0x24, 0x2e, 0x8d, 0xc4, 0x8d, 0x89, 0xb8, 0xa8, 0x71, 0xe3, 0xae, 0xa5, 0xc7, 0x32,
	0x09, 0xfd, 0x93, 0x39, 0x53, 0x01, 0x9f, 0xc2, 0xc7, 0x72, 0x27, 0x4b, 0x97, 0x37, 0xb0, 0xbf,
	0xcf, 0x70, 0x43, 0x4b, 0x9b, 0x92, 0xdc, 0x1b, 0x2e, 0xbb, 0x39, 0xdf, 0x9c, 0xf9, 0xcd, 0x37,
	0xdf, 0xe4, 0xc0, 0x8b, 0xc0, 0x0f, 0x23, 0x24, 0x57, 0x6f, 0x9c, 0x4c, 0xa5, 0x3a, 0xe5, 0x6f,
	0xb4, 0xc2, 0x5f, 0xa9, 0x8a, 0xe5, 0xca, 0x09, 0xa4, 0x2e, 0x77, 0x17, 0x4b, 0x5f, 0x26, 0x4e,
	0xb9, 0xb6, 0xff, 0x31, 0xe8, 0xcd, 0x29, 0xfa, 0x8a, 0xeb, 0xd9, 0x51, 0xe0, 0x02, 0xba, 0x0b,
	0x85, 0xbe, 0x4e, 0x95, 0x60, 0x23, 0x36, 0x7e, 0xe6, 0x55, 0x25, 0x7f, 0x09, 0xed, 0x5c, 0x49,
	0xd1, 0x2e, 0xd4, 0xe3, 0xf2, 0xd8, 0x1b, 0xfb, 0x89, 0x1f, 0xa1, 0x12, 0x66, 0xd9, 0x7b, 0x2a,
	0xf9, 0x08, 0x7a, 0x19, 0xaa, 0x58, 0x12, 0xc9, 0x34, 0x21, 0xf1, 0x64, 0xc4, 0xc6, 0xa6, 0xd7,
	0x94, 0xf8, 0x14, 0x5e, 0xff, 0x54, 0x88, 0xbf, 0xf1, 0x63, 0x18, 0x2a, 0x24, 0x42, 0xfa, 0x24,
	0x23, 0x24, 0x2d, 0x3a, 0x05, 0xe9, 0xfe, 0x4d, 0x6e, 0xc3, 0x73, 0xca, 0x03, 0x9f, 0x08, 0xf5,
	0x77, 0x25, 0x49, 0x74, 0x8b, 0xe6, 0x33, 0xcd, 0xfe, 0x00, 0xaf, 0x1a, 0x0f, 0xf2, 0x90, 0xb2,
	0x34, 0x21, 0xe4, 0x7d, 0x30, 0x64, 0x58, 0xbc, 0xc9, 0xf4, 0x0c, 0x19, 0x16, 0xe6, 0x91, 0xc8,
	0x8f, 0x50, 0x18, 0x27, 0xf3, 0x65, 0x69, 0x7b, 0xd0, 0x2f, 0x01, 0xdf, 0xf2, 0xe0, 0x52, 0x28,
	0x25, 0xd5, 0xa8, 0xa9, 0x03, 0xe8, 0x50, 0x9e, 0x65, 0xab, 0x6d, 0x91, 0x93, 0xe9, 0x9d, 0x2a,
	0xdb, 0x83, 0xc1, 0x39, 0xb3, 0xf6, 0x65, 0x01, 0x54, 0xf6, 0x3f, 0x57, 0xfe, 0x1a, 0xca, 0xc3,
	0x3e, 0x27, 0xb7, 0x0c, 0xda, 0x73, 0x8a, 0x78, 0x02, 0x4f, 0xeb, 0xef, 0x7b, 0xeb, 0x5c, 0xf8,
	0x70, 0xa7, 0x91, 0xcd, 0x70, 0x7a, 0x4d, 0x77, 0xed, 0x78, 0x0d, 0xbd, 0x66, 0x38, 0xee, 0x23,
	0x21, 0xd5, 0x81, 0xe1, 0xfb, 0x2b, 0x0f, 0x54, 0x17, 0xcf, 0xbe, 0xfc, 0xdd, 0x5b, 0x6c, 0xb7,
	0xb7, 0xd8, 0xcd, 0xde, 0x62, 0x7f, 0x0e, 0x56, 0x6b, 0x77, 0xb0, 0x5a, 0xff, 0x0f, 0x56, 0xeb,
	0xc7, 0x24, 0x92, 0x7a, 0x99, 0x07, 0xce, 0x22, 0x8d, 0xdd, 0x1a, 0xee, 0x9e, 0xc3, 0xdd, 0x8d,
	0x5b, 0x4d, 0xc7, 0x36, 0x43, 0x0a, 0x3a, 0xc5, 0x84, 0xbc, 0xbb, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x6e, 0xae, 0x79, 0x34, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	NewBadge(ctx context.Context, in *MsgNewBadge, opts ...grpc.CallOption) (*MsgNewBadgeResponse, error)
	NewSubBadge(ctx context.Context, in *MsgNewSubBadge, opts ...grpc.CallOption) (*MsgNewSubBadgeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) NewBadge(ctx context.Context, in *MsgNewBadge, opts ...grpc.CallOption) (*MsgNewBadgeResponse, error) {
	out := new(MsgNewBadgeResponse)
	err := c.cc.Invoke(ctx, "/trevormil.bitbadgeschain.badges.Msg/NewBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewSubBadge(ctx context.Context, in *MsgNewSubBadge, opts ...grpc.CallOption) (*MsgNewSubBadgeResponse, error) {
	out := new(MsgNewSubBadgeResponse)
	err := c.cc.Invoke(ctx, "/trevormil.bitbadgeschain.badges.Msg/NewSubBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	NewBadge(context.Context, *MsgNewBadge) (*MsgNewBadgeResponse, error)
	NewSubBadge(context.Context, *MsgNewSubBadge) (*MsgNewSubBadgeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) NewBadge(ctx context.Context, req *MsgNewBadge) (*MsgNewBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBadge not implemented")
}
func (*UnimplementedMsgServer) NewSubBadge(ctx context.Context, req *MsgNewSubBadge) (*MsgNewSubBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSubBadge not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_NewBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewBadge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trevormil.bitbadgeschain.badges.Msg/NewBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewBadge(ctx, req.(*MsgNewBadge))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewSubBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewSubBadge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewSubBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trevormil.bitbadgeschain.badges.Msg/NewSubBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewSubBadge(ctx, req.(*MsgNewSubBadge))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trevormil.bitbadgeschain.badges.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBadge",
			Handler:    _Msg_NewBadge_Handler,
		},
		{
			MethodName: "NewSubBadge",
			Handler:    _Msg_NewSubBadge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badges/tx.proto",
}

func (m *MsgNewBadge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewBadge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewBadge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubassetUris) > 0 {
		i -= len(m.SubassetUris)
		copy(dAtA[i:], m.SubassetUris)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SubassetUris)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FreezeAddressesDigest) > 0 {
		i -= len(m.FreezeAddressesDigest)
		copy(dAtA[i:], m.FreezeAddressesDigest)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FreezeAddressesDigest)))
		i--
		dAtA[i] = 0x32
	}
	if m.Permissions != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Permissions))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Manager) > 0 {
		i -= len(m.Manager)
		copy(dAtA[i:], m.Manager)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Manager)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewBadgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewBadgeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewBadgeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewSubBadge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewSubBadge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewSubBadge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Supply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Supply))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewSubBadgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewSubBadgeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewSubBadgeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubassetId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SubassetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgNewBadge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Manager)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Permissions != 0 {
		n += 1 + sovTx(uint64(m.Permissions))
	}
	l = len(m.FreezeAddressesDigest)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SubassetUris)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewBadgeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewSubBadge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	if m.Supply != 0 {
		n += 1 + sovTx(uint64(m.Supply))
	}
	return n
}

func (m *MsgNewSubBadgeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubassetId != 0 {
		n += 1 + sovTx(uint64(m.SubassetId))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgNewBadge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewBadge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewBadge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manager", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manager = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreezeAddressesDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FreezeAddressesDigest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubassetUris", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubassetUris = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNewBadgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewBadgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewBadgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNewSubBadge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewSubBadge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewSubBadge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			m.Supply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Supply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNewSubBadgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewSubBadgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewSubBadgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubassetId", wireType)
			}
			m.SubassetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubassetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
