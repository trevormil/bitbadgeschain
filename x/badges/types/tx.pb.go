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
	Id                    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *MsgNewBadge) GetId() string {
	if m != nil {
		return m.Id
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
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *MsgNewBadgeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgNewBadgeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgNewBadge)(nil), "trevormil.bitbadgeschain.badges.MsgNewBadge")
	proto.RegisterType((*MsgNewBadgeResponse)(nil), "trevormil.bitbadgeschain.badges.MsgNewBadgeResponse")
}

func init() { proto.RegisterFile("badges/tx.proto", fileDescriptor_bc897b33479788c9) }

var fileDescriptor_bc897b33479788c9 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x31, 0x4b, 0x33, 0x41,
	0x10, 0xcd, 0x26, 0xf9, 0x92, 0xcf, 0x8d, 0xa8, 0xac, 0x08, 0x8b, 0xc5, 0x1a, 0x52, 0xa5, 0x90,
	0x3d, 0x88, 0xe9, 0xc5, 0x60, 0x69, 0x2c, 0x02, 0x36, 0x76, 0x7b, 0xb9, 0x71, 0xb3, 0xe0, 0xdd,
	0x1e, 0x3b, 0x7b, 0x1a, 0xfd, 0x15, 0xfe, 0x2c, 0xcb, 0x94, 0x76, 0x4a, 0xf2, 0x47, 0xe4, 0x2e,
	0xb9, 0x70, 0x01, 0x41, 0xec, 0xe6, 0xbd, 0x99, 0x79, 0x33, 0x6f, 0x18, 0x7a, 0x18, 0xaa, 0x48,
	0x03, 0x06, 0x7e, 0x2e, 0x53, 0x67, 0xbd, 0x65, 0x67, 0xde, 0xc1, 0x93, 0x75, 0xb1, 0x79, 0x94,
	0xa1, 0xf1, 0xeb, 0xec, 0x74, 0xa6, 0x4c, 0x22, 0xd7, 0x71, 0xef, 0x93, 0xd0, 0xce, 0x18, 0xf5,
	0x2d, 0x3c, 0x8f, 0x72, 0x82, 0x71, 0xda, 0x9e, 0x3a, 0x50, 0xde, 0x3a, 0x4e, 0xba, 0xa4, 0xbf,
	0x37, 0x29, 0x21, 0x3b, 0xa0, 0x75, 0x13, 0xf1, 0x7a, 0x41, 0xd6, 0x4d, 0xc4, 0x8e, 0x68, 0x23,
	0x73, 0x86, 0x37, 0x0a, 0x22, 0x0f, 0xf3, 0xde, 0x58, 0x25, 0x4a, 0x83, 0xe3, 0xcd, 0x75, 0xef,
	0x06, 0xb2, 0x2e, 0xed, 0xa4, 0xe0, 0x62, 0x83, 0x68, 0x6c, 0x82, 0xfc, 0x5f, 0x97, 0xf4, 0x9b,
	0x93, 0x2a, 0xc5, 0x86, 0xf4, 0xe4, 0xc1, 0x01, 0xbc, 0xc2, 0x55, 0x14, 0x39, 0x40, 0x04, 0xbc,
	0x36, 0x1a, 0xd0, 0xf3, 0x56, 0xa1, 0xf4, 0x73, 0x92, 0xf5, 0xe8, 0x3e, 0x66, 0xa1, 0x42, 0x04,
	0x7f, 0xe7, 0x0c, 0xf2, 0x76, 0x51, 0xbc, 0xc3, 0xf5, 0x2e, 0xe9, 0x71, 0xc5, 0xe0, 0x04, 0x30,
	0xb5, 0x09, 0xc2, 0xc6, 0x0e, 0xd9, 0xda, 0xc9, 0x97, 0x07, 0x44, 0xa5, 0x61, 0xe3, 0xb1, 0x84,
	0x83, 0x8c, 0x36, 0xc6, 0xa8, 0x59, 0x42, 0xff, 0x6f, 0xaf, 0x74, 0x2e, 0x7f, 0xb9, 0xab, 0xac,
	0x8c, 0x3c, 0x1d, 0xfe, 0xa5, 0xba, 0x5c, 0x70, 0x74, 0xf3, 0xbe, 0x14, 0x64, 0xb1, 0x14, 0xe4,
	0x6b, 0x29, 0xc8, 0xdb, 0x4a, 0xd4, 0x16, 0x2b, 0x51, 0xfb, 0x58, 0x89, 0xda, 0xfd, 0x40, 0x1b,
	0x3f, 0xcb, 0x42, 0x39, 0xb5, 0x71, 0xb0, 0x55, 0x0e, 0x76, 0x95, 0x83, 0x79, 0x50, 0xfe, 0xc2,
	0x4b, 0x0a, 0x18, 0xb6, 0x8a, 0x7f, 0xb8, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xf9, 0xa2,
	0x6e, 0x22, 0x02, 0x00, 0x00,
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

// MsgServer is the server API for Msg service.
type MsgServer interface {
	NewBadge(context.Context, *MsgNewBadge) (*MsgNewBadgeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) NewBadge(ctx context.Context, req *MsgNewBadge) (*MsgNewBadgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBadge not implemented")
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

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trevormil.bitbadgeschain.badges.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBadge",
			Handler:    _Msg_NewBadge_Handler,
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
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
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
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
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
	l = len(m.Id)
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
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
