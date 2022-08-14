// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: badges/ranges.proto

package types

import (
	fmt "fmt"
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

//Id ranges define a range of IDs from start to end. Can be used for subbadgeIds, nonces, addresses anything. If end == 0, we assume end == start. Start must be >= end.
type IdRange struct {
	Start uint64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *IdRange) Reset()         { *m = IdRange{} }
func (m *IdRange) String() string { return proto.CompactTextString(m) }
func (*IdRange) ProtoMessage()    {}
func (*IdRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4071aa00cc93a07, []int{0}
}
func (m *IdRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRange.Merge(m, src)
}
func (m *IdRange) XXX_Size() int {
	return m.Size()
}
func (m *IdRange) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRange.DiscardUnknown(m)
}

var xxx_messageInfo_IdRange proto.InternalMessageInfo

func (m *IdRange) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *IdRange) GetEnd() uint64 {
	if m != nil {
		return m.End
	}
	return 0
}

//Defines a balance object. The specified balance holds for all ids specified within the id ranges array.
type BalanceObject struct {
	Balance  uint64     `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	IdRanges []*IdRange `protobuf:"bytes,2,rep,name=id_ranges,json=idRanges,proto3" json:"id_ranges,omitempty"`
}

func (m *BalanceObject) Reset()         { *m = BalanceObject{} }
func (m *BalanceObject) String() string { return proto.CompactTextString(m) }
func (*BalanceObject) ProtoMessage()    {}
func (*BalanceObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4071aa00cc93a07, []int{1}
}
func (m *BalanceObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceObject.Merge(m, src)
}
func (m *BalanceObject) XXX_Size() int {
	return m.Size()
}
func (m *BalanceObject) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceObject.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceObject proto.InternalMessageInfo

func (m *BalanceObject) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *BalanceObject) GetIdRanges() []*IdRange {
	if m != nil {
		return m.IdRanges
	}
	return nil
}

func init() {
	proto.RegisterType((*IdRange)(nil), "trevormil.bitbadgeschain.badges.IdRange")
	proto.RegisterType((*BalanceObject)(nil), "trevormil.bitbadgeschain.badges.BalanceObject")
}

func init() { proto.RegisterFile("badges/ranges.proto", fileDescriptor_c4071aa00cc93a07) }

var fileDescriptor_c4071aa00cc93a07 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4a, 0x4c, 0x49,
	0x4f, 0x2d, 0xd6, 0x2f, 0x4a, 0xcc, 0x4b, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x2f, 0x29, 0x4a, 0x2d, 0xcb, 0x2f, 0xca, 0xcd, 0xcc, 0xd1, 0x4b, 0xca, 0x2c, 0x81, 0xa8,
	0x48, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xb0, 0x95, 0x0c, 0xb9, 0xd8, 0x3d, 0x53, 0x82, 0x40,
	0x5a, 0x84, 0x44, 0xb8, 0x58, 0x8b, 0x4b, 0x12, 0x8b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x82, 0x20, 0x1c, 0x21, 0x01, 0x2e, 0xe6, 0xd4, 0xbc, 0x14, 0x09, 0x26, 0xb0, 0x18, 0x88, 0xa9,
	0x54, 0xc0, 0xc5, 0xeb, 0x94, 0x98, 0x93, 0x98, 0x97, 0x9c, 0xea, 0x9f, 0x94, 0x95, 0x9a, 0x5c,
	0x22, 0x24, 0xc1, 0xc5, 0x9e, 0x04, 0x11, 0x80, 0x6a, 0x85, 0x71, 0x85, 0x5c, 0xb9, 0x38, 0x33,
	0x53, 0xe2, 0x21, 0x2e, 0x92, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xd0, 0x23, 0xe0, 0x24,
	0x3d, 0xa8, 0x7b, 0x82, 0x38, 0x32, 0x21, 0x8c, 0x62, 0x27, 0x9f, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x87, 0x9b, 0xab, 0x8f, 0x6a, 0xae, 0x7e, 0x85, 0x3e, 0x34, 0x68, 0x4a, 0x2a, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0x41, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x9b, 0xea, 0xf0,
	0x31, 0x01, 0x00, 0x00,
}

func (m *IdRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.End != 0 {
		i = encodeVarintRanges(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintRanges(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BalanceObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdRanges) > 0 {
		for iNdEx := len(m.IdRanges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IdRanges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRanges(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Balance != 0 {
		i = encodeVarintRanges(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRanges(dAtA []byte, offset int, v uint64) int {
	offset -= sovRanges(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovRanges(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovRanges(uint64(m.End))
	}
	return n
}

func (m *BalanceObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Balance != 0 {
		n += 1 + sovRanges(uint64(m.Balance))
	}
	if len(m.IdRanges) > 0 {
		for _, e := range m.IdRanges {
			l = e.Size()
			n += 1 + l + sovRanges(uint64(l))
		}
	}
	return n
}

func sovRanges(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRanges(x uint64) (n int) {
	return sovRanges(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRanges
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
			return fmt.Errorf("proto: IdRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRanges
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRanges
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRanges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRanges
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
func (m *BalanceObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRanges
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
			return fmt.Errorf("proto: BalanceObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRanges
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdRanges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRanges
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
				return ErrInvalidLengthRanges
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRanges
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdRanges = append(m.IdRanges, &IdRange{})
			if err := m.IdRanges[len(m.IdRanges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRanges(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRanges
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
func skipRanges(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRanges
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
					return 0, ErrIntOverflowRanges
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
					return 0, ErrIntOverflowRanges
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
				return 0, ErrInvalidLengthRanges
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRanges
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRanges
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRanges        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRanges          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRanges = fmt.Errorf("proto: unexpected end of group")
)
