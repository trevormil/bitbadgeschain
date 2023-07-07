// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges/collections.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// A BadgeCollection is the top level object for a collection of badges.
// It defines everything about the collection, such as the manager, metadata, etc.
//
// All collections are identified by a collectionId assigned by the blockchain, which is a uint64 that increments (i.e. first collection has ID 1).
//
// All collections also have a manager who is responsible for managing the collection.
// They can be granted certain permissions, such as the ability to mint new badges.
//
// Certain fields are timeline-based, which means they may have different values at different block heights.
// We fetch the value according to the current time.
// For example, we may set the manager to be Alice from Time1 to Time2, and then set the manager to be Bob from Time2 to Time3.
//
// Collections may have different balance types: standard vs off-chain vs inherited. See documentation for differences.
type BadgeCollection struct {
	// The collectionId is the unique identifier for this collection.
	CollectionId Uint `protobuf:"bytes,1,opt,name=collectionId,proto3,customtype=Uint" json:"collectionId"`
	// The collection metadata is the metadata for the collection itself.
	CollectionMetadataTimeline []*CollectionMetadataTimeline `protobuf:"bytes,2,rep,name=collectionMetadataTimeline,proto3" json:"collectionMetadataTimeline,omitempty"`
	// The badge metadata is the metadata for each badge in the collection.
	BadgeMetadataTimeline []*BadgeMetadataTimeline `protobuf:"bytes,3,rep,name=badgeMetadataTimeline,proto3" json:"badgeMetadataTimeline,omitempty"`
	// The balancesType is the type of balances this collection uses (standard, off-chain, or inherited).
	BalancesType string `protobuf:"bytes,4,opt,name=balancesType,proto3" json:"balancesType,omitempty"`
	// The off-chain balances metadata defines where to fetch the balances for collections with off-chain balances.
	OffChainBalancesMetadataTimeline []*OffChainBalancesMetadataTimeline `protobuf:"bytes,5,rep,name=offChainBalancesMetadataTimeline,proto3" json:"offChainBalancesMetadataTimeline,omitempty"`
	// The inherited balances metadata defines the parent balances for collections with inherited balances.
	InheritedBalancesTimeline []*InheritedBalancesTimeline `protobuf:"bytes,6,rep,name=inheritedBalancesTimeline,proto3" json:"inheritedBalancesTimeline,omitempty"`
	// The custom data field is an arbitrary field that can be used to store any data.
	CustomDataTimeline []*CustomDataTimeline `protobuf:"bytes,7,rep,name=customDataTimeline,proto3" json:"customDataTimeline,omitempty"`
	// The manager is the address of the manager of this collection.
	ManagerTimeline []*ManagerTimeline `protobuf:"bytes,8,rep,name=managerTimeline,proto3" json:"managerTimeline,omitempty"`
	// The permissions define what the manager of the collection can do or not do.
	CollectionPermissions *CollectionPermissions `protobuf:"bytes,9,opt,name=collectionPermissions,proto3" json:"collectionPermissions,omitempty"`
	// The approved transfers timeline defines the transferability of the collection for collections with standard balances.
	// This defines it on a collection-level. All transfers must be explicitly allowed on the collection-level, or else, they will fail.
	//
	// Collection approved transfers can optionally specify to override the user approvals for a transfer (e.g. forcefully revoke a badge).
	// If user approvals are not overriden, then a transfer must also satisfy the From user's approved outgoing transfers and the To user's approved incoming transfers.
	CollectionApprovedTransfersTimeline []*CollectionApprovedTransferTimeline `protobuf:"bytes,10,rep,name=collectionApprovedTransfersTimeline,proto3" json:"collectionApprovedTransfersTimeline,omitempty"`
	// Standards allow us to define a standard for the collection. This lets others know how to interpret the fields of the collection.
	StandardsTimeline []*StandardsTimeline `protobuf:"bytes,11,rep,name=standardsTimeline,proto3" json:"standardsTimeline,omitempty"`
	// The isArchivedTimeline defines whether the collection is archived or not.
	// When a collection is archived, it is read-only and no transactions can be processed.
	IsArchivedTimeline []*IsArchivedTimeline `protobuf:"bytes,12,rep,name=isArchivedTimeline,proto3" json:"isArchivedTimeline,omitempty"`
	// The contractAddressTimeline defines the contract address for the collection (if it has a corresponding contract).
	ContractAddressTimeline []*ContractAddressTimeline `protobuf:"bytes,13,rep,name=contractAddressTimeline,proto3" json:"contractAddressTimeline,omitempty"`
	// The defaultUserApprovedOutgoingTransfersTimeline defines the default user approved outgoing transfers for an uninitialized user balance.
	// The user can change this value at any time.
	DefaultUserApprovedOutgoingTransfersTimeline []*UserApprovedOutgoingTransferTimeline `protobuf:"bytes,14,rep,name=defaultUserApprovedOutgoingTransfersTimeline,proto3" json:"defaultUserApprovedOutgoingTransfersTimeline,omitempty"`
	// The defaultUserApprovedIncomingTransfersTimeline defines the default user approved incoming transfers for an uninitialized user balance.
	// The user can change this value at any time.
	//
	// Ex: Set this to disallow all incoming transfers by default, making the user have to opt-in to receiving the badge.
	DefaultUserApprovedIncomingTransfersTimeline []*UserApprovedIncomingTransferTimeline `protobuf:"bytes,15,rep,name=defaultUserApprovedIncomingTransfersTimeline,proto3" json:"defaultUserApprovedIncomingTransfersTimeline,omitempty"`
}

func (m *BadgeCollection) Reset()         { *m = BadgeCollection{} }
func (m *BadgeCollection) String() string { return proto.CompactTextString(m) }
func (*BadgeCollection) ProtoMessage()    {}
func (*BadgeCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eac0b7495c54217, []int{0}
}
func (m *BadgeCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadgeCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadgeCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadgeCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadgeCollection.Merge(m, src)
}
func (m *BadgeCollection) XXX_Size() int {
	return m.Size()
}
func (m *BadgeCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_BadgeCollection.DiscardUnknown(m)
}

var xxx_messageInfo_BadgeCollection proto.InternalMessageInfo

func (m *BadgeCollection) GetCollectionMetadataTimeline() []*CollectionMetadataTimeline {
	if m != nil {
		return m.CollectionMetadataTimeline
	}
	return nil
}

func (m *BadgeCollection) GetBadgeMetadataTimeline() []*BadgeMetadataTimeline {
	if m != nil {
		return m.BadgeMetadataTimeline
	}
	return nil
}

func (m *BadgeCollection) GetBalancesType() string {
	if m != nil {
		return m.BalancesType
	}
	return ""
}

func (m *BadgeCollection) GetOffChainBalancesMetadataTimeline() []*OffChainBalancesMetadataTimeline {
	if m != nil {
		return m.OffChainBalancesMetadataTimeline
	}
	return nil
}

func (m *BadgeCollection) GetInheritedBalancesTimeline() []*InheritedBalancesTimeline {
	if m != nil {
		return m.InheritedBalancesTimeline
	}
	return nil
}

func (m *BadgeCollection) GetCustomDataTimeline() []*CustomDataTimeline {
	if m != nil {
		return m.CustomDataTimeline
	}
	return nil
}

func (m *BadgeCollection) GetManagerTimeline() []*ManagerTimeline {
	if m != nil {
		return m.ManagerTimeline
	}
	return nil
}

func (m *BadgeCollection) GetCollectionPermissions() *CollectionPermissions {
	if m != nil {
		return m.CollectionPermissions
	}
	return nil
}

func (m *BadgeCollection) GetCollectionApprovedTransfersTimeline() []*CollectionApprovedTransferTimeline {
	if m != nil {
		return m.CollectionApprovedTransfersTimeline
	}
	return nil
}

func (m *BadgeCollection) GetStandardsTimeline() []*StandardsTimeline {
	if m != nil {
		return m.StandardsTimeline
	}
	return nil
}

func (m *BadgeCollection) GetIsArchivedTimeline() []*IsArchivedTimeline {
	if m != nil {
		return m.IsArchivedTimeline
	}
	return nil
}

func (m *BadgeCollection) GetContractAddressTimeline() []*ContractAddressTimeline {
	if m != nil {
		return m.ContractAddressTimeline
	}
	return nil
}

func (m *BadgeCollection) GetDefaultUserApprovedOutgoingTransfersTimeline() []*UserApprovedOutgoingTransferTimeline {
	if m != nil {
		return m.DefaultUserApprovedOutgoingTransfersTimeline
	}
	return nil
}

func (m *BadgeCollection) GetDefaultUserApprovedIncomingTransfersTimeline() []*UserApprovedIncomingTransferTimeline {
	if m != nil {
		return m.DefaultUserApprovedIncomingTransfersTimeline
	}
	return nil
}

func init() {
	proto.RegisterType((*BadgeCollection)(nil), "bitbadges.bitbadgeschain.badges.BadgeCollection")
}

func init() { proto.RegisterFile("badges/collections.proto", fileDescriptor_9eac0b7495c54217) }

var fileDescriptor_9eac0b7495c54217 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0x7f, 0xfb, 0x2f, 0x74, 0x9a, 0x52, 0x31, 0xe2, 0x23, 0xcd, 0xc2, 0x89, 0xca,
	0xa6, 0x0b, 0x64, 0x57, 0xa9, 0x84, 0x10, 0xac, 0x92, 0xc0, 0x22, 0x12, 0x55, 0xd1, 0x90, 0x6e,
	0xba, 0x62, 0x6c, 0x4f, 0x9c, 0x91, 0xec, 0x19, 0x6b, 0x3c, 0x41, 0x8d, 0x78, 0x06, 0x24, 0x36,
	0x3c, 0x04, 0x2f, 0xc0, 0x33, 0x74, 0xd9, 0x25, 0x62, 0x51, 0xa1, 0xe4, 0x45, 0x90, 0x27, 0xfe,
	0x20, 0xb1, 0x1d, 0xa7, 0xec, 0x26, 0x33, 0x73, 0xce, 0xf9, 0xdd, 0xeb, 0x1b, 0x1b, 0x34, 0x2c,
	0xec, 0xb8, 0x24, 0x34, 0x6d, 0xee, 0x79, 0xc4, 0x96, 0x94, 0xb3, 0xd0, 0x08, 0x04, 0x97, 0x1c,
	0xb6, 0x2c, 0x2a, 0x17, 0x87, 0x46, 0xba, 0xb2, 0xc7, 0x98, 0x32, 0x63, 0xb1, 0x6e, 0x1e, 0xba,
	0x9c, 0xbb, 0x1e, 0x31, 0xd5, 0x75, 0x6b, 0x32, 0x32, 0x31, 0x9b, 0x2e, 0xb4, 0xcd, 0x27, 0xb1,
	0xab, 0x14, 0x98, 0x85, 0x23, 0x22, 0x62, 0xcf, 0xe6, 0xe3, 0x78, 0xdf, 0xc2, 0x1e, 0x66, 0x36,
	0x49, 0xb6, 0x13, 0x88, 0x80, 0x08, 0x9f, 0x86, 0x61, 0x06, 0x91, 0x0a, 0x7c, 0x22, 0xb1, 0x83,
	0x25, 0x5e, 0xf5, 0xa7, 0x3e, 0xf1, 0x28, 0x4b, 0x8d, 0x1e, 0xb9, 0xdc, 0xe5, 0x6a, 0x69, 0x46,
	0xab, 0xc5, 0xee, 0xd1, 0x8f, 0x7d, 0x70, 0xd0, 0x8b, 0x04, 0xfd, 0xb4, 0x48, 0x78, 0x02, 0xea,
	0x59, 0xc9, 0x03, 0xa7, 0xa1, 0xb5, 0xb5, 0xe3, 0xdd, 0x5e, 0xfd, 0xfa, 0xb6, 0x55, 0xfb, 0x75,
	0xdb, 0xda, 0xbe, 0xa0, 0x4c, 0xa2, 0xa5, 0x1b, 0xf0, 0x33, 0x68, 0x66, 0xbf, 0xcf, 0x62, 0x9e,
	0x61, 0x0c, 0xd0, 0xf8, 0xaf, 0xbd, 0x75, 0xbc, 0xd7, 0x79, 0x6d, 0x54, 0x34, 0xcd, 0xe8, 0x97,
	0x5a, 0xa0, 0x35, 0xf6, 0xd0, 0x03, 0x8b, 0x4e, 0xe4, 0x72, 0xb7, 0x54, 0xee, 0x8b, 0xca, 0xdc,
	0x5e, 0x91, 0x1a, 0x15, 0x9b, 0xc2, 0x23, 0x50, 0x4f, 0x9e, 0xd0, 0x70, 0x1a, 0x90, 0xc6, 0x76,
	0xd4, 0x1c, 0xb4, 0xb4, 0x07, 0xbf, 0x68, 0xa0, 0xcd, 0x47, 0xa3, 0x7e, 0x14, 0xd2, 0x8b, 0x0f,
	0x72, 0x74, 0xff, 0x2b, 0xba, 0x6e, 0x25, 0xdd, 0x79, 0x85, 0x11, 0xaa, 0x8c, 0x82, 0x57, 0xe0,
	0x90, 0xb2, 0x31, 0x11, 0x54, 0x12, 0x27, 0xb9, 0x94, 0x72, 0xec, 0x28, 0x8e, 0x57, 0x95, 0x1c,
	0x83, 0x32, 0x07, 0x54, 0x6e, 0x0e, 0x6d, 0x00, 0xed, 0x49, 0x28, 0xb9, 0xff, 0xe6, 0xef, 0xd2,
	0xef, 0xa9, 0xc8, 0xd3, 0xea, 0x81, 0xc8, 0x49, 0x51, 0x81, 0x1d, 0xbc, 0x04, 0x07, 0x3e, 0x66,
	0xd8, 0x25, 0x22, 0x4d, 0xb8, 0xaf, 0x12, 0x4e, 0x2a, 0x13, 0xce, 0x96, 0x75, 0x68, 0xd5, 0x28,
	0x1a, 0xae, 0x6c, 0xf4, 0xde, 0x67, 0xff, 0xc1, 0xc6, 0x6e, 0x5b, 0xdb, 0x68, 0xb8, 0xfa, 0x45,
	0x6a, 0x54, 0x6c, 0x0a, 0xbf, 0x69, 0xe0, 0x59, 0x76, 0xd2, 0x0d, 0x02, 0xc1, 0x3f, 0x11, 0x67,
	0x98, 0xbc, 0x29, 0xd2, 0xf2, 0x80, 0x2a, 0xaf, 0x7f, 0x87, 0xf0, 0x55, 0xaf, 0xb4, 0xe2, 0x4d,
	0xf2, 0xe0, 0x47, 0xf0, 0x30, 0x94, 0x98, 0x39, 0x58, 0x38, 0x19, 0xc4, 0x9e, 0x82, 0xe8, 0x54,
	0x42, 0x7c, 0x58, 0x55, 0xa2, 0xbc, 0x59, 0x34, 0x28, 0x34, 0xec, 0x0a, 0x7b, 0x4c, 0x23, 0x80,
	0x24, 0xa2, 0xbe, 0xe1, 0xa0, 0x0c, 0x72, 0x52, 0x54, 0x60, 0x07, 0x05, 0x78, 0x6a, 0x73, 0x26,
	0x05, 0xb6, 0x65, 0xd7, 0x71, 0x04, 0x09, 0xb3, 0x62, 0xf6, 0x55, 0xd2, 0xcb, 0x0d, 0x3a, 0x5a,
	0xa8, 0x47, 0x65, 0xc6, 0xf0, 0xbb, 0x06, 0x9e, 0x3b, 0x64, 0x84, 0x27, 0x9e, 0xbc, 0x08, 0x89,
	0x48, 0x7a, 0x7c, 0x3e, 0x91, 0x2e, 0xa7, 0xcc, 0xcd, 0x3f, 0xdb, 0x07, 0x8a, 0xe4, 0x6d, 0x25,
	0xc9, 0x3a, 0xb7, 0x14, 0xeb, 0x4e, 0xd1, 0x65, 0xac, 0x03, 0x66, 0x73, 0xbf, 0x90, 0xf5, 0xe0,
	0x1f, 0x58, 0x57, 0xdd, 0xd6, 0xb2, 0x96, 0x46, 0xf7, 0xde, 0x5d, 0xcf, 0x74, 0xed, 0x66, 0xa6,
	0x6b, 0xbf, 0x67, 0xba, 0xf6, 0x75, 0xae, 0xd7, 0x6e, 0xe6, 0x7a, 0xed, 0xe7, 0x5c, 0xaf, 0x5d,
	0x76, 0x5c, 0x2a, 0xc7, 0x13, 0xcb, 0xb0, 0xb9, 0x6f, 0xa6, 0x38, 0xe6, 0x32, 0x98, 0x79, 0x65,
	0x26, 0x9f, 0xc9, 0x69, 0x40, 0x42, 0x6b, 0x47, 0x7d, 0x0d, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0xcf, 0x09, 0xf5, 0xf3, 0x07, 0x00, 0x00,
}

func (m *BadgeCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadgeCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadgeCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DefaultUserApprovedIncomingTransfersTimeline) > 0 {
		for iNdEx := len(m.DefaultUserApprovedIncomingTransfersTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DefaultUserApprovedIncomingTransfersTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.DefaultUserApprovedOutgoingTransfersTimeline) > 0 {
		for iNdEx := len(m.DefaultUserApprovedOutgoingTransfersTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DefaultUserApprovedOutgoingTransfersTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.ContractAddressTimeline) > 0 {
		for iNdEx := len(m.ContractAddressTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractAddressTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.IsArchivedTimeline) > 0 {
		for iNdEx := len(m.IsArchivedTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IsArchivedTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.StandardsTimeline) > 0 {
		for iNdEx := len(m.StandardsTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StandardsTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.CollectionApprovedTransfersTimeline) > 0 {
		for iNdEx := len(m.CollectionApprovedTransfersTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionApprovedTransfersTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.CollectionPermissions != nil {
		{
			size, err := m.CollectionPermissions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCollections(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ManagerTimeline) > 0 {
		for iNdEx := len(m.ManagerTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ManagerTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CustomDataTimeline) > 0 {
		for iNdEx := len(m.CustomDataTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CustomDataTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.InheritedBalancesTimeline) > 0 {
		for iNdEx := len(m.InheritedBalancesTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InheritedBalancesTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.OffChainBalancesMetadataTimeline) > 0 {
		for iNdEx := len(m.OffChainBalancesMetadataTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OffChainBalancesMetadataTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BalancesType) > 0 {
		i -= len(m.BalancesType)
		copy(dAtA[i:], m.BalancesType)
		i = encodeVarintCollections(dAtA, i, uint64(len(m.BalancesType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BadgeMetadataTimeline) > 0 {
		for iNdEx := len(m.BadgeMetadataTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BadgeMetadataTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CollectionMetadataTimeline) > 0 {
		for iNdEx := len(m.CollectionMetadataTimeline) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionMetadataTimeline[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollections(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.CollectionId.Size()
		i -= size
		if _, err := m.CollectionId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCollections(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCollections(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollections(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BadgeCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CollectionId.Size()
	n += 1 + l + sovCollections(uint64(l))
	if len(m.CollectionMetadataTimeline) > 0 {
		for _, e := range m.CollectionMetadataTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.BadgeMetadataTimeline) > 0 {
		for _, e := range m.BadgeMetadataTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	l = len(m.BalancesType)
	if l > 0 {
		n += 1 + l + sovCollections(uint64(l))
	}
	if len(m.OffChainBalancesMetadataTimeline) > 0 {
		for _, e := range m.OffChainBalancesMetadataTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.InheritedBalancesTimeline) > 0 {
		for _, e := range m.InheritedBalancesTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.CustomDataTimeline) > 0 {
		for _, e := range m.CustomDataTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.ManagerTimeline) > 0 {
		for _, e := range m.ManagerTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if m.CollectionPermissions != nil {
		l = m.CollectionPermissions.Size()
		n += 1 + l + sovCollections(uint64(l))
	}
	if len(m.CollectionApprovedTransfersTimeline) > 0 {
		for _, e := range m.CollectionApprovedTransfersTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.StandardsTimeline) > 0 {
		for _, e := range m.StandardsTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.IsArchivedTimeline) > 0 {
		for _, e := range m.IsArchivedTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.ContractAddressTimeline) > 0 {
		for _, e := range m.ContractAddressTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.DefaultUserApprovedOutgoingTransfersTimeline) > 0 {
		for _, e := range m.DefaultUserApprovedOutgoingTransfersTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	if len(m.DefaultUserApprovedIncomingTransfersTimeline) > 0 {
		for _, e := range m.DefaultUserApprovedIncomingTransfersTimeline {
			l = e.Size()
			n += 1 + l + sovCollections(uint64(l))
		}
	}
	return n
}

func sovCollections(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollections(x uint64) (n int) {
	return sovCollections(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BadgeCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollections
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
			return fmt.Errorf("proto: BadgeCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadgeCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollectionId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionMetadataTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionMetadataTimeline = append(m.CollectionMetadataTimeline, &CollectionMetadataTimeline{})
			if err := m.CollectionMetadataTimeline[len(m.CollectionMetadataTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BadgeMetadataTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BadgeMetadataTimeline = append(m.BadgeMetadataTimeline, &BadgeMetadataTimeline{})
			if err := m.BadgeMetadataTimeline[len(m.BadgeMetadataTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancesType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalancesType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OffChainBalancesMetadataTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OffChainBalancesMetadataTimeline = append(m.OffChainBalancesMetadataTimeline, &OffChainBalancesMetadataTimeline{})
			if err := m.OffChainBalancesMetadataTimeline[len(m.OffChainBalancesMetadataTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InheritedBalancesTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InheritedBalancesTimeline = append(m.InheritedBalancesTimeline, &InheritedBalancesTimeline{})
			if err := m.InheritedBalancesTimeline[len(m.InheritedBalancesTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomDataTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustomDataTimeline = append(m.CustomDataTimeline, &CustomDataTimeline{})
			if err := m.CustomDataTimeline[len(m.CustomDataTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManagerTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManagerTimeline = append(m.ManagerTimeline, &ManagerTimeline{})
			if err := m.ManagerTimeline[len(m.ManagerTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionPermissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CollectionPermissions == nil {
				m.CollectionPermissions = &CollectionPermissions{}
			}
			if err := m.CollectionPermissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionApprovedTransfersTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionApprovedTransfersTimeline = append(m.CollectionApprovedTransfersTimeline, &CollectionApprovedTransferTimeline{})
			if err := m.CollectionApprovedTransfersTimeline[len(m.CollectionApprovedTransfersTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardsTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StandardsTimeline = append(m.StandardsTimeline, &StandardsTimeline{})
			if err := m.StandardsTimeline[len(m.StandardsTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsArchivedTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IsArchivedTimeline = append(m.IsArchivedTimeline, &IsArchivedTimeline{})
			if err := m.IsArchivedTimeline[len(m.IsArchivedTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddressTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddressTimeline = append(m.ContractAddressTimeline, &ContractAddressTimeline{})
			if err := m.ContractAddressTimeline[len(m.ContractAddressTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultUserApprovedOutgoingTransfersTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultUserApprovedOutgoingTransfersTimeline = append(m.DefaultUserApprovedOutgoingTransfersTimeline, &UserApprovedOutgoingTransferTimeline{})
			if err := m.DefaultUserApprovedOutgoingTransfersTimeline[len(m.DefaultUserApprovedOutgoingTransfersTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultUserApprovedIncomingTransfersTimeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollections
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
				return ErrInvalidLengthCollections
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollections
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultUserApprovedIncomingTransfersTimeline = append(m.DefaultUserApprovedIncomingTransfersTimeline, &UserApprovedIncomingTransferTimeline{})
			if err := m.DefaultUserApprovedIncomingTransfersTimeline[len(m.DefaultUserApprovedIncomingTransfersTimeline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollections(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollections
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
func skipCollections(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollections
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
					return 0, ErrIntOverflowCollections
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
					return 0, ErrIntOverflowCollections
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
				return 0, ErrInvalidLengthCollections
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollections
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollections
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollections        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollections          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollections = fmt.Errorf("proto: unexpected end of group")
)
