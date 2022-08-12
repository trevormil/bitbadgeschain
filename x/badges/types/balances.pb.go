// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges/balances.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type UserBalanceInfo struct {
	BalanceAmounts []*BalanceObject   `protobuf:"bytes,2,rep,name=balanceAmounts,proto3" json:"balanceAmounts,omitempty"`
	PendingNonce   uint64             `protobuf:"varint,3,opt,name=pendingNonce,proto3" json:"pendingNonce,omitempty"`
	Pending        []*PendingTransfer `protobuf:"bytes,4,rep,name=pending,proto3" json:"pending,omitempty"`
	Approvals      []*Approval        `protobuf:"bytes,5,rep,name=approvals,proto3" json:"approvals,omitempty"`
}

func (m *UserBalanceInfo) Reset()         { *m = UserBalanceInfo{} }
func (m *UserBalanceInfo) String() string { return proto.CompactTextString(m) }
func (*UserBalanceInfo) ProtoMessage()    {}
func (*UserBalanceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_233d29a167e739f0, []int{0}
}
func (m *UserBalanceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBalanceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBalanceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBalanceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBalanceInfo.Merge(m, src)
}
func (m *UserBalanceInfo) XXX_Size() int {
	return m.Size()
}
func (m *UserBalanceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBalanceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserBalanceInfo proto.InternalMessageInfo

func (m *UserBalanceInfo) GetBalanceAmounts() []*BalanceObject {
	if m != nil {
		return m.BalanceAmounts
	}
	return nil
}

func (m *UserBalanceInfo) GetPendingNonce() uint64 {
	if m != nil {
		return m.PendingNonce
	}
	return 0
}

func (m *UserBalanceInfo) GetPending() []*PendingTransfer {
	if m != nil {
		return m.Pending
	}
	return nil
}

func (m *UserBalanceInfo) GetApprovals() []*Approval {
	if m != nil {
		return m.Approvals
	}
	return nil
}

type Approval struct {
	Address uint64 `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	// uint64 expirationTime = 2; Can maybe add this in the future but gets tricky with which ones you want to deduct / add from, etc. Probably best to just leave it as is
	ApprovalAmounts []*BalanceObject `protobuf:"bytes,3,rep,name=approvalAmounts,proto3" json:"approvalAmounts,omitempty"`
}

func (m *Approval) Reset()         { *m = Approval{} }
func (m *Approval) String() string { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()    {}
func (*Approval) Descriptor() ([]byte, []int) {
	return fileDescriptor_233d29a167e739f0, []int{1}
}
func (m *Approval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Approval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Approval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Approval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Approval.Merge(m, src)
}
func (m *Approval) XXX_Size() int {
	return m.Size()
}
func (m *Approval) XXX_DiscardUnknown() {
	xxx_messageInfo_Approval.DiscardUnknown(m)
}

var xxx_messageInfo_Approval proto.InternalMessageInfo

func (m *Approval) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *Approval) GetApprovalAmounts() []*BalanceObject {
	if m != nil {
		return m.ApprovalAmounts
	}
	return nil
}

type PendingTransfer struct {
	SubbadgeRange        *IdRange `protobuf:"bytes,1,opt,name=subbadgeRange,proto3" json:"subbadgeRange,omitempty"`
	ThisPendingNonce     uint64   `protobuf:"varint,2,opt,name=thisPendingNonce,proto3" json:"thisPendingNonce,omitempty"`
	OtherPendingNonce    uint64   `protobuf:"varint,3,opt,name=otherPendingNonce,proto3" json:"otherPendingNonce,omitempty"`
	Amount               uint64   `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Sent                 bool     `protobuf:"varint,5,opt,name=sent,proto3" json:"sent,omitempty"`
	To                   uint64   `protobuf:"varint,6,opt,name=to,proto3" json:"to,omitempty"`
	From                 uint64   `protobuf:"varint,7,opt,name=from,proto3" json:"from,omitempty"`
	ApprovedBy           uint64   `protobuf:"varint,9,opt,name=approvedBy,proto3" json:"approvedBy,omitempty"`
	MarkedAsAccepted     bool     `protobuf:"varint,10,opt,name=markedAsAccepted,proto3" json:"markedAsAccepted,omitempty"`
	ExpirationTime       uint64   `protobuf:"varint,11,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	CantCancelBeforeTime uint64   `protobuf:"varint,12,opt,name=cantCancelBeforeTime,proto3" json:"cantCancelBeforeTime,omitempty"`
}

func (m *PendingTransfer) Reset()         { *m = PendingTransfer{} }
func (m *PendingTransfer) String() string { return proto.CompactTextString(m) }
func (*PendingTransfer) ProtoMessage()    {}
func (*PendingTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_233d29a167e739f0, []int{2}
}
func (m *PendingTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingTransfer.Merge(m, src)
}
func (m *PendingTransfer) XXX_Size() int {
	return m.Size()
}
func (m *PendingTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_PendingTransfer proto.InternalMessageInfo

func (m *PendingTransfer) GetSubbadgeRange() *IdRange {
	if m != nil {
		return m.SubbadgeRange
	}
	return nil
}

func (m *PendingTransfer) GetThisPendingNonce() uint64 {
	if m != nil {
		return m.ThisPendingNonce
	}
	return 0
}

func (m *PendingTransfer) GetOtherPendingNonce() uint64 {
	if m != nil {
		return m.OtherPendingNonce
	}
	return 0
}

func (m *PendingTransfer) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PendingTransfer) GetSent() bool {
	if m != nil {
		return m.Sent
	}
	return false
}

func (m *PendingTransfer) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *PendingTransfer) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *PendingTransfer) GetApprovedBy() uint64 {
	if m != nil {
		return m.ApprovedBy
	}
	return 0
}

func (m *PendingTransfer) GetMarkedAsAccepted() bool {
	if m != nil {
		return m.MarkedAsAccepted
	}
	return false
}

func (m *PendingTransfer) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *PendingTransfer) GetCantCancelBeforeTime() uint64 {
	if m != nil {
		return m.CantCancelBeforeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserBalanceInfo)(nil), "trevormil.bitbadgeschain.badges.UserBalanceInfo")
	proto.RegisterType((*Approval)(nil), "trevormil.bitbadgeschain.badges.Approval")
	proto.RegisterType((*PendingTransfer)(nil), "trevormil.bitbadgeschain.badges.PendingTransfer")
}

func init() { proto.RegisterFile("badges/balances.proto", fileDescriptor_233d29a167e739f0) }

var fileDescriptor_233d29a167e739f0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x8e, 0x93, 0x34, 0x69, 0x27, 0xfd, 0x93, 0x9f, 0xa5, 0xa0, 0x55, 0x0f, 0x26, 0xca, 0x01,
	0x05, 0x84, 0x6c, 0x14, 0x9e, 0x20, 0xe1, 0x80, 0x8a, 0x50, 0x89, 0xac, 0x82, 0x10, 0xb7, 0xb5,
	0x3d, 0x49, 0x0c, 0xf1, 0xae, 0xb5, 0xbb, 0xa9, 0xda, 0x0b, 0xcf, 0xc0, 0x33, 0xf0, 0x0a, 0xbc,
	0x04, 0xc7, 0x1e, 0x39, 0xa2, 0xe4, 0x45, 0x50, 0xc6, 0x4e, 0x69, 0x1c, 0xa4, 0x48, 0x9c, 0x3c,
	0xfb, 0xed, 0xf7, 0x7d, 0x33, 0xb3, 0x9f, 0x0c, 0x0f, 0x42, 0x11, 0x4f, 0xd1, 0xf8, 0xa1, 0x98,
	0x0b, 0x19, 0xa1, 0xf1, 0x32, 0xad, 0xac, 0x62, 0x8f, 0xac, 0xc6, 0x4b, 0xa5, 0xd3, 0x64, 0xee,
	0x85, 0x89, 0xcd, 0x39, 0xd1, 0x4c, 0x24, 0xd2, 0xcb, 0xeb, 0xd3, 0x93, 0xa9, 0x9a, 0x2a, 0xe2,
	0xfa, 0xeb, 0x2a, 0x97, 0x9d, 0xde, 0x2f, 0xdc, 0x32, 0xa1, 0x45, 0x6a, 0x4a, 0x60, 0xfe, 0x29,
	0x81, 0x5a, 0xc8, 0x5b, 0xb0, 0xf7, 0xad, 0x0a, 0x9d, 0x77, 0x06, 0xf5, 0x28, 0x1f, 0xe6, 0x4c,
	0x4e, 0x14, 0x7b, 0x0f, 0xed, 0x62, 0xb6, 0x61, 0xaa, 0x16, 0xd2, 0x1a, 0x5e, 0xed, 0xd6, 0xfa,
	0xad, 0x81, 0xe7, 0xed, 0x19, 0xd1, 0x2b, 0x5c, 0xde, 0x86, 0x9f, 0x30, 0xb2, 0x41, 0xc9, 0x85,
	0xf5, 0xe0, 0x38, 0x43, 0x19, 0x27, 0x72, 0x7a, 0xae, 0x64, 0x84, 0xbc, 0xd6, 0x75, 0xfa, 0xf5,
	0x60, 0x0b, 0x63, 0xaf, 0xa1, 0x59, 0x9c, 0x79, 0x9d, 0x9a, 0x3e, 0xdf, 0xdb, 0x74, 0x9c, 0xf3,
	0x2f, 0xb4, 0x90, 0x66, 0x82, 0x3a, 0xd8, 0x18, 0xb0, 0x57, 0x70, 0x24, 0xb2, 0x4c, 0xab, 0x4b,
	0x31, 0x37, 0xfc, 0x80, 0xdc, 0x9e, 0xec, 0x75, 0x1b, 0x16, 0x8a, 0xe0, 0x8f, 0xb6, 0xf7, 0x05,
	0x0e, 0x37, 0x30, 0xe3, 0xd0, 0x14, 0x71, 0xac, 0xd1, 0x18, 0xee, 0xd0, 0xfc, 0x9b, 0x23, 0xfb,
	0x00, 0x9d, 0x8d, 0x64, 0xf3, 0x6e, 0xb5, 0x7f, 0x7a, 0xb7, 0xb2, 0x4d, 0xef, 0x7b, 0x0d, 0x3a,
	0xa5, 0x2d, 0xd9, 0x39, 0xfc, 0x67, 0x16, 0x21, 0xe9, 0x83, 0x75, 0xa0, 0x34, 0x4d, 0x6b, 0xd0,
	0xdf, 0xdb, 0xeb, 0x2c, 0x26, 0x7e, 0xb0, 0x2d, 0x67, 0x4f, 0xe1, 0x7f, 0x3b, 0x4b, 0xcc, 0xf8,
	0x6e, 0x40, 0x55, 0x5a, 0x70, 0x07, 0x67, 0xcf, 0xe0, 0x9e, 0xb2, 0x33, 0xd4, 0xe3, 0xdd, 0x34,
	0x77, 0x2f, 0xd8, 0x43, 0x68, 0x08, 0x5a, 0x84, 0xd7, 0x89, 0x52, 0x9c, 0x18, 0x83, 0xba, 0x41,
	0x69, 0xf9, 0x41, 0xd7, 0xe9, 0x1f, 0x06, 0x54, 0xb3, 0x36, 0x54, 0xad, 0xe2, 0x0d, 0xe2, 0x55,
	0xad, 0x5a, 0x73, 0x26, 0x5a, 0xa5, 0xbc, 0x49, 0x08, 0xd5, 0xcc, 0x05, 0xc8, 0x1f, 0x08, 0xe3,
	0xd1, 0x35, 0x3f, 0xa2, 0x9b, 0x3b, 0xc8, 0x7a, 0x93, 0x54, 0xe8, 0xcf, 0x18, 0x0f, 0xcd, 0x30,
	0x8a, 0x30, 0xb3, 0x18, 0x73, 0xa0, 0x1e, 0x3b, 0x38, 0x7b, 0x0c, 0x6d, 0xbc, 0xca, 0x12, 0x2d,
	0x6c, 0xa2, 0xe4, 0x45, 0x92, 0x22, 0x6f, 0x91, 0x5f, 0x09, 0x65, 0x03, 0x38, 0x89, 0x84, 0xb4,
	0x2f, 0xd7, 0x29, 0xcd, 0x47, 0x38, 0x51, 0x1a, 0x89, 0x7d, 0x4c, 0xec, 0xbf, 0xde, 0x8d, 0xde,
	0xfc, 0x58, 0xba, 0xce, 0xcd, 0xd2, 0x75, 0x7e, 0x2d, 0x5d, 0xe7, 0xeb, 0xca, 0xad, 0xdc, 0xac,
	0xdc, 0xca, 0xcf, 0x95, 0x5b, 0xf9, 0x38, 0x98, 0x26, 0x76, 0xb6, 0x08, 0xbd, 0x48, 0xa5, 0xfe,
	0x6d, 0x5c, 0xfe, 0x76, 0x5c, 0xfe, 0x55, 0xf1, 0xf7, 0xfa, 0xf6, 0x3a, 0x43, 0x13, 0x36, 0xe8,
	0x7f, 0x7d, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x08, 0xbc, 0xbf, 0x3e, 0x04, 0x00, 0x00,
}

func (m *UserBalanceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBalanceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBalanceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalances(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Pending) > 0 {
		for iNdEx := len(m.Pending) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pending[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalances(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.PendingNonce != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.PendingNonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BalanceAmounts) > 0 {
		for iNdEx := len(m.BalanceAmounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BalanceAmounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalances(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *Approval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Approval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Approval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApprovalAmounts) > 0 {
		for iNdEx := len(m.ApprovalAmounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApprovalAmounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalances(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Address != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.Address))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PendingTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CantCancelBeforeTime != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.CantCancelBeforeTime))
		i--
		dAtA[i] = 0x60
	}
	if m.ExpirationTime != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.ExpirationTime))
		i--
		dAtA[i] = 0x58
	}
	if m.MarkedAsAccepted {
		i--
		if m.MarkedAsAccepted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.ApprovedBy != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.ApprovedBy))
		i--
		dAtA[i] = 0x48
	}
	if m.From != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.From))
		i--
		dAtA[i] = 0x38
	}
	if m.To != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.To))
		i--
		dAtA[i] = 0x30
	}
	if m.Sent {
		i--
		if m.Sent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Amount != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if m.OtherPendingNonce != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.OtherPendingNonce))
		i--
		dAtA[i] = 0x18
	}
	if m.ThisPendingNonce != 0 {
		i = encodeVarintBalances(dAtA, i, uint64(m.ThisPendingNonce))
		i--
		dAtA[i] = 0x10
	}
	if m.SubbadgeRange != nil {
		{
			size, err := m.SubbadgeRange.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBalances(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBalances(dAtA []byte, offset int, v uint64) int {
	offset -= sovBalances(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserBalanceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BalanceAmounts) > 0 {
		for _, e := range m.BalanceAmounts {
			l = e.Size()
			n += 1 + l + sovBalances(uint64(l))
		}
	}
	if m.PendingNonce != 0 {
		n += 1 + sovBalances(uint64(m.PendingNonce))
	}
	if len(m.Pending) > 0 {
		for _, e := range m.Pending {
			l = e.Size()
			n += 1 + l + sovBalances(uint64(l))
		}
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovBalances(uint64(l))
		}
	}
	return n
}

func (m *Approval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != 0 {
		n += 1 + sovBalances(uint64(m.Address))
	}
	if len(m.ApprovalAmounts) > 0 {
		for _, e := range m.ApprovalAmounts {
			l = e.Size()
			n += 1 + l + sovBalances(uint64(l))
		}
	}
	return n
}

func (m *PendingTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubbadgeRange != nil {
		l = m.SubbadgeRange.Size()
		n += 1 + l + sovBalances(uint64(l))
	}
	if m.ThisPendingNonce != 0 {
		n += 1 + sovBalances(uint64(m.ThisPendingNonce))
	}
	if m.OtherPendingNonce != 0 {
		n += 1 + sovBalances(uint64(m.OtherPendingNonce))
	}
	if m.Amount != 0 {
		n += 1 + sovBalances(uint64(m.Amount))
	}
	if m.Sent {
		n += 2
	}
	if m.To != 0 {
		n += 1 + sovBalances(uint64(m.To))
	}
	if m.From != 0 {
		n += 1 + sovBalances(uint64(m.From))
	}
	if m.ApprovedBy != 0 {
		n += 1 + sovBalances(uint64(m.ApprovedBy))
	}
	if m.MarkedAsAccepted {
		n += 2
	}
	if m.ExpirationTime != 0 {
		n += 1 + sovBalances(uint64(m.ExpirationTime))
	}
	if m.CantCancelBeforeTime != 0 {
		n += 1 + sovBalances(uint64(m.CantCancelBeforeTime))
	}
	return n
}

func sovBalances(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBalances(x uint64) (n int) {
	return sovBalances(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserBalanceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalances
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
			return fmt.Errorf("proto: UserBalanceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBalanceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceAmounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalanceAmounts = append(m.BalanceAmounts, &BalanceObject{})
			if err := m.BalanceAmounts[len(m.BalanceAmounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingNonce", wireType)
			}
			m.PendingNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PendingNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pending", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pending = append(m.Pending, &PendingTransfer{})
			if err := m.Pending[len(m.Pending)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Approval{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalances(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalances
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
func (m *Approval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalances
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
			return fmt.Errorf("proto: Approval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			m.Address = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Address |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalAmounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalAmounts = append(m.ApprovalAmounts, &BalanceObject{})
			if err := m.ApprovalAmounts[len(m.ApprovalAmounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalances(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalances
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
func (m *PendingTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalances
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
			return fmt.Errorf("proto: PendingTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubbadgeRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBalances
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalances
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubbadgeRange == nil {
				m.SubbadgeRange = &IdRange{}
			}
			if err := m.SubbadgeRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThisPendingNonce", wireType)
			}
			m.ThisPendingNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThisPendingNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherPendingNonce", wireType)
			}
			m.OtherPendingNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherPendingNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sent = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			m.To = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.To |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovedBy", wireType)
			}
			m.ApprovedBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApprovedBy |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarkedAsAccepted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MarkedAsAccepted = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			m.ExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CantCancelBeforeTime", wireType)
			}
			m.CantCancelBeforeTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalances
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CantCancelBeforeTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBalances(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalances
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
func skipBalances(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalances
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
					return 0, ErrIntOverflowBalances
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
					return 0, ErrIntOverflowBalances
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
				return 0, ErrInvalidLengthBalances
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBalances
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBalances
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBalances        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalances          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBalances = fmt.Errorf("proto: unexpected end of group")
)
