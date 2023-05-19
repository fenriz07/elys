// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/pool_asset.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type PoolAsset struct {
	Token  types.Coin `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	Weight string     `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *PoolAsset) Reset()         { *m = PoolAsset{} }
func (m *PoolAsset) String() string { return proto.CompactTextString(m) }
func (*PoolAsset) ProtoMessage()    {}
func (*PoolAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_152994f419c8cb00, []int{0}
}
func (m *PoolAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAsset.Merge(m, src)
}
func (m *PoolAsset) XXX_Size() int {
	return m.Size()
}
func (m *PoolAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAsset.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAsset proto.InternalMessageInfo

func (m *PoolAsset) GetToken() types.Coin {
	if m != nil {
		return m.Token
	}
	return types.Coin{}
}

func (m *PoolAsset) GetWeight() string {
	if m != nil {
		return m.Weight
	}
	return ""
}

func init() {
	proto.RegisterType((*PoolAsset)(nil), "elysnetwork.elys.amm.PoolAsset")
}

func init() { proto.RegisterFile("elys/amm/pool_asset.proto", fileDescriptor_152994f419c8cb00) }

var fileDescriptor_152994f419c8cb00 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0x31, 0x4e, 0xc3, 0x40,
	0x10, 0x45, 0xbd, 0x08, 0x22, 0xc5, 0x74, 0x56, 0x84, 0x92, 0x14, 0x4b, 0x44, 0xe5, 0x86, 0x19,
	0x05, 0xc4, 0x01, 0x30, 0x17, 0x40, 0x29, 0xd3, 0xa0, 0xb5, 0x35, 0x72, 0xac, 0x78, 0x3d, 0x56,
	0x76, 0x20, 0xe4, 0x16, 0x1c, 0x2b, 0xa5, 0x4b, 0x2a, 0x84, 0xec, 0x8b, 0x44, 0x6b, 0xbb, 0xfb,
	0x5f, 0xff, 0xff, 0xd1, 0xbc, 0x70, 0x41, 0xe5, 0xc9, 0xa1, 0xb1, 0x16, 0x6b, 0xe6, 0xf2, 0xc3,
	0x38, 0x47, 0x02, 0xf5, 0x81, 0x85, 0xa3, 0x99, 0x8f, 0x2a, 0x92, 0x23, 0x1f, 0xf6, 0xe0, 0x35,
	0x18, 0x6b, 0x97, 0xb3, 0x9c, 0x73, 0xee, 0x0b, 0xe8, 0xd5, 0xd0, 0x5d, 0xea, 0x8c, 0x9d, 0x65,
	0x87, 0xa9, 0x71, 0x84, 0x5f, 0xeb, 0x94, 0xc4, 0xac, 0x31, 0xe3, 0xa2, 0x1a, 0xf2, 0x87, 0x6d,
	0x38, 0x7d, 0x67, 0x2e, 0x5f, 0xfd, 0xf9, 0xe8, 0x25, 0xbc, 0x11, 0xde, 0x53, 0x35, 0x57, 0x2b,
	0x15, 0xdf, 0x3e, 0x2d, 0x60, 0x18, 0x83, 0x1f, 0xc3, 0x38, 0x86, 0x37, 0x2e, 0xaa, 0xe4, 0xfa,
	0xfc, 0x77, 0x1f, 0x6c, 0x86, 0x76, 0x74, 0x17, 0x4e, 0x8e, 0x54, 0xe4, 0x3b, 0x99, 0x5f, 0xad,
	0x54, 0x3c, 0xdd, 0x8c, 0x2e, 0x49, 0xce, 0xad, 0x56, 0x4d, 0xab, 0xd5, 0x7f, 0xab, 0xd5, 0x4f,
	0xa7, 0x83, 0xa6, 0xd3, 0xc1, 0x6f, 0xa7, 0x83, 0x6d, 0x9c, 0x17, 0xb2, 0xfb, 0x4c, 0x21, 0x63,
	0x8b, 0x1e, 0xe0, 0x71, 0xa4, 0xe9, 0x0d, 0x7e, 0xf7, 0xd8, 0x72, 0xaa, 0xc9, 0xa5, 0x93, 0xfe,
	0xcd, 0xe7, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x32, 0x53, 0x23, 0x0f, 0x01, 0x00, 0x00,
}

func (m *PoolAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weight) > 0 {
		i -= len(m.Weight)
		copy(dAtA[i:], m.Weight)
		i = encodeVarintPoolAsset(dAtA, i, uint64(len(m.Weight)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPoolAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPoolAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Token.Size()
	n += 1 + l + sovPoolAsset(uint64(l))
	l = len(m.Weight)
	if l > 0 {
		n += 1 + l + sovPoolAsset(uint64(l))
	}
	return n
}

func sovPoolAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolAsset(x uint64) (n int) {
	return sovPoolAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolAsset
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
			return fmt.Errorf("proto: PoolAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolAsset
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
				return ErrInvalidLengthPoolAsset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPoolAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolAsset
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
				return ErrInvalidLengthPoolAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Weight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolAsset
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
func skipPoolAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolAsset
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
					return 0, ErrIntOverflowPoolAsset
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
					return 0, ErrIntOverflowPoolAsset
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
				return 0, ErrInvalidLengthPoolAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolAsset = fmt.Errorf("proto: unexpected end of group")
)