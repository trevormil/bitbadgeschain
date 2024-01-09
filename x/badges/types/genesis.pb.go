// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the badges module's genesis state.
type GenesisState struct {
	Params                    Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId                    string              `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Collections               []*BadgeCollection  `protobuf:"bytes,3,rep,name=collections,proto3" json:"collections,omitempty"`
	NextCollectionId          Uint                `protobuf:"bytes,4,opt,name=nextCollectionId,proto3,customtype=Uint" json:"nextCollectionId"`
	Balances                  []*UserBalanceStore `protobuf:"bytes,5,rep,name=balances,proto3" json:"balances,omitempty"`
	BalanceStoreKeys          []string            `protobuf:"bytes,6,rep,name=balanceStoreKeys,proto3" json:"balanceStoreKeys,omitempty"`
	ChallengeTrackers         []Uint              `protobuf:"bytes,7,rep,name=challengeTrackers,proto3,customtype=Uint" json:"challengeTrackers"`
	ChallengeTrackerStoreKeys []string            `protobuf:"bytes,8,rep,name=challengeTrackerStoreKeys,proto3" json:"challengeTrackerStoreKeys,omitempty"`
	AddressLists              []*AddressList      `protobuf:"bytes,9,rep,name=addressLists,proto3" json:"addressLists,omitempty"`
	ApprovalTrackers          []*ApprovalTracker  `protobuf:"bytes,10,rep,name=approvalTrackers,proto3" json:"approvalTrackers,omitempty"`
	ApprovalTrackerStoreKeys  []string            `protobuf:"bytes,11,rep,name=approvalTrackerStoreKeys,proto3" json:"approvalTrackerStoreKeys,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_966136619b34853e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetCollections() []*BadgeCollection {
	if m != nil {
		return m.Collections
	}
	return nil
}

func (m *GenesisState) GetBalances() []*UserBalanceStore {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *GenesisState) GetBalanceStoreKeys() []string {
	if m != nil {
		return m.BalanceStoreKeys
	}
	return nil
}

func (m *GenesisState) GetChallengeTrackerStoreKeys() []string {
	if m != nil {
		return m.ChallengeTrackerStoreKeys
	}
	return nil
}

func (m *GenesisState) GetAddressLists() []*AddressList {
	if m != nil {
		return m.AddressLists
	}
	return nil
}

func (m *GenesisState) GetApprovalTrackers() []*ApprovalTracker {
	if m != nil {
		return m.ApprovalTrackers
	}
	return nil
}

func (m *GenesisState) GetApprovalTrackerStoreKeys() []string {
	if m != nil {
		return m.ApprovalTrackerStoreKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "badges.GenesisState")
}

func init() { proto.RegisterFile("badges/genesis.proto", fileDescriptor_966136619b34853e) }

var fileDescriptor_966136619b34853e = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0x5a, 0xb2, 0xd5, 0xad, 0x50, 0xf1, 0x06, 0x33, 0x3d, 0x64, 0x15, 0xa7, 0x0a,
	0xa1, 0x46, 0x2a, 0x48, 0xc0, 0xc4, 0x65, 0xd9, 0x01, 0x4d, 0xec, 0x80, 0x3c, 0x76, 0xe1, 0x32,
	0x39, 0xc9, 0x4b, 0x6a, 0x91, 0xd9, 0x91, 0x6d, 0xd0, 0xf6, 0x2d, 0xf8, 0x58, 0x3b, 0xee, 0x88,
	0x38, 0x4c, 0x53, 0xfb, 0x45, 0xa6, 0x24, 0x76, 0xff, 0xaa, 0xa7, 0xba, 0xcf, 0xef, 0x79, 0xfd,
	0x3e, 0xce, 0xfb, 0xa2, 0xfd, 0x98, 0xa5, 0x19, 0xe8, 0x30, 0x03, 0x01, 0x9a, 0xeb, 0x51, 0xa1,
	0xa4, 0x91, 0xd8, 0xaf, 0xd5, 0xfe, 0x7e, 0x26, 0x33, 0x59, 0x49, 0x61, 0x79, 0xaa, 0x69, 0x7f,
	0xcf, 0xd6, 0x14, 0x4c, 0xb1, 0x2b, 0x5b, 0xd2, 0x27, 0x56, 0x4c, 0x64, 0x9e, 0x43, 0x62, 0xb8,
	0x14, 0x8e, 0xbc, 0xb0, 0x24, 0x66, 0x39, 0x13, 0x09, 0x38, 0xf9, 0xa5, 0x95, 0x8d, 0x62, 0x42,
	0xff, 0x04, 0xe5, 0xf4, 0xbe, 0xd5, 0x59, 0x9a, 0x2a, 0xd0, 0xfa, 0x32, 0xe7, 0xda, 0x58, 0xf6,
	0xfa, 0xa1, 0x85, 0xba, 0x5f, 0xea, 0xa4, 0xe7, 0x86, 0x19, 0xc0, 0x6f, 0x91, 0x5f, 0xa7, 0x20,
	0xde, 0xc0, 0x1b, 0x76, 0xc6, 0xcf, 0x46, 0x75, 0xf5, 0xe8, 0x5b, 0xa5, 0x46, 0xad, 0xdb, 0xfb,
	0xc3, 0x06, 0xb5, 0x1e, 0x7c, 0x80, 0x76, 0x0a, 0xa9, 0xcc, 0x25, 0x4f, 0xc9, 0x93, 0x81, 0x37,
	0x6c, 0x53, 0xbf, 0xfc, 0x7b, 0x9a, 0xe2, 0x4f, 0xa8, 0xb3, 0x94, 0x9b, 0x34, 0x07, 0xcd, 0x61,
	0x67, 0x7c, 0xe0, 0xee, 0x8a, 0xca, 0x9f, 0x93, 0x39, 0xa7, 0xcb, 0x5e, 0xfc, 0x11, 0xf5, 0x04,
	0x5c, 0x9b, 0x05, 0x3e, 0x4d, 0x49, 0xab, 0xbc, 0x3c, 0xea, 0x96, 0xbd, 0xff, 0xdf, 0x1f, 0xb6,
	0x2e, 0xb8, 0x30, 0x74, 0xc3, 0x85, 0xdf, 0xa3, 0x5d, 0xf7, 0x49, 0xc8, 0xd3, 0xaa, 0x23, 0x71,
	0x1d, 0x2f, 0x34, 0xa8, 0xa8, 0x66, 0xe7, 0x46, 0x2a, 0xa0, 0x73, 0x27, 0x7e, 0x83, 0x7a, 0xf1,
	0x12, 0xf9, 0x0a, 0x37, 0x9a, 0xf8, 0x83, 0xe6, 0xb0, 0x4d, 0x37, 0x74, 0x7c, 0x84, 0x9e, 0x27,
	0x13, 0x96, 0xe7, 0x20, 0x32, 0xf8, 0xae, 0x58, 0xf2, 0x0b, 0x94, 0x26, 0x3b, 0xa5, 0x79, 0x2d,
	0xdc, 0xa6, 0x0d, 0x7f, 0x46, 0xaf, 0xd6, 0xc5, 0x45, 0xc3, 0xdd, 0xaa, 0xe1, 0x76, 0x03, 0xfe,
	0x80, 0xba, 0x76, 0x7e, 0x67, 0xe5, 0xf8, 0x48, 0xbb, 0x7a, 0xdf, 0x9e, 0x7b, 0xdf, 0xf1, 0x82,
	0xd1, 0x15, 0x23, 0x3e, 0x41, 0x3d, 0x56, 0x14, 0x4a, 0xfe, 0x61, 0xf9, 0x3c, 0x31, 0x5a, 0x1d,
	0xc7, 0xf1, 0x2a, 0xa7, 0x1b, 0x05, 0xf8, 0x08, 0x91, 0x35, 0x6d, 0x11, 0xbd, 0x53, 0x45, 0xdf,
	0xca, 0xa3, 0xb3, 0xdb, 0x69, 0xe0, 0xdd, 0x4d, 0x03, 0xef, 0x61, 0x1a, 0x78, 0x7f, 0x67, 0x41,
	0xe3, 0x6e, 0x16, 0x34, 0xfe, 0xcd, 0x82, 0xc6, 0x8f, 0x71, 0xc6, 0xcd, 0xe4, 0x77, 0x3c, 0x4a,
	0xe4, 0x55, 0x18, 0x73, 0xe3, 0xb6, 0xda, 0x9d, 0x92, 0x09, 0xe3, 0x22, 0xbc, 0x0e, 0xdd, 0x5a,
	0xdf, 0x14, 0xa0, 0x63, 0xbf, 0xda, 0xdb, 0x77, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x23, 0x9f,
	0xd4, 0x79, 0x67, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApprovalTrackerStoreKeys) > 0 {
		for iNdEx := len(m.ApprovalTrackerStoreKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ApprovalTrackerStoreKeys[iNdEx])
			copy(dAtA[i:], m.ApprovalTrackerStoreKeys[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.ApprovalTrackerStoreKeys[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.ApprovalTrackers) > 0 {
		for iNdEx := len(m.ApprovalTrackers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApprovalTrackers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.AddressLists) > 0 {
		for iNdEx := len(m.AddressLists) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AddressLists[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ChallengeTrackerStoreKeys) > 0 {
		for iNdEx := len(m.ChallengeTrackerStoreKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChallengeTrackerStoreKeys[iNdEx])
			copy(dAtA[i:], m.ChallengeTrackerStoreKeys[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChallengeTrackerStoreKeys[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ChallengeTrackers) > 0 {
		for iNdEx := len(m.ChallengeTrackers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.ChallengeTrackers[iNdEx].Size()
				i -= size
				if _, err := m.ChallengeTrackers[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.BalanceStoreKeys) > 0 {
		for iNdEx := len(m.BalanceStoreKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BalanceStoreKeys[iNdEx])
			copy(dAtA[i:], m.BalanceStoreKeys[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BalanceStoreKeys[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.NextCollectionId.Size()
		i -= size
		if _, err := m.NextCollectionId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Collections) > 0 {
		for iNdEx := len(m.Collections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Collections) > 0 {
		for _, e := range m.Collections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.NextCollectionId.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BalanceStoreKeys) > 0 {
		for _, s := range m.BalanceStoreKeys {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChallengeTrackers) > 0 {
		for _, e := range m.ChallengeTrackers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChallengeTrackerStoreKeys) > 0 {
		for _, s := range m.ChallengeTrackerStoreKeys {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AddressLists) > 0 {
		for _, e := range m.AddressLists {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ApprovalTrackers) > 0 {
		for _, e := range m.ApprovalTrackers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ApprovalTrackerStoreKeys) > 0 {
		for _, s := range m.ApprovalTrackerStoreKeys {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collections = append(m.Collections, &BadgeCollection{})
			if err := m.Collections[len(m.Collections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextCollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NextCollectionId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, &UserBalanceStore{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceStoreKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalanceStoreKeys = append(m.BalanceStoreKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeTrackers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Uint
			m.ChallengeTrackers = append(m.ChallengeTrackers, v)
			if err := m.ChallengeTrackers[len(m.ChallengeTrackers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeTrackerStoreKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengeTrackerStoreKeys = append(m.ChallengeTrackerStoreKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressLists", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressLists = append(m.AddressLists, &AddressList{})
			if err := m.AddressLists[len(m.AddressLists)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalTrackers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalTrackers = append(m.ApprovalTrackers, &ApprovalTracker{})
			if err := m.ApprovalTrackers[len(m.ApprovalTrackers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApprovalTrackerStoreKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApprovalTrackerStoreKeys = append(m.ApprovalTrackerStoreKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
