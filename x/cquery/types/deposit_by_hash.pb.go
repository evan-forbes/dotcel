// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cquery/deposit_by_hash.proto

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

type DepositByHash struct {
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *DepositByHash) Reset()         { *m = DepositByHash{} }
func (m *DepositByHash) String() string { return proto.CompactTextString(m) }
func (*DepositByHash) ProtoMessage()    {}
func (*DepositByHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8c3a7b73b5913f, []int{0}
}
func (m *DepositByHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositByHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositByHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositByHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositByHash.Merge(m, src)
}
func (m *DepositByHash) XXX_Size() int {
	return m.Size()
}
func (m *DepositByHash) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositByHash.DiscardUnknown(m)
}

var xxx_messageInfo_DepositByHash proto.InternalMessageInfo

func (m *DepositByHash) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DepositByHash)(nil), "celestiaorg.dotcel.cquery.DepositByHash")
}

func init() { proto.RegisterFile("cquery/deposit_by_hash.proto", fileDescriptor_3b8c3a7b73b5913f) }

var fileDescriptor_3b8c3a7b73b5913f = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2e, 0x2c, 0x4d,
	0x2d, 0xaa, 0xd4, 0x4f, 0x49, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0x89, 0x4f, 0xaa, 0x8c, 0xcf, 0x48,
	0x2c, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0x4e, 0xcd, 0x49, 0x2d, 0x2e,
	0xc9, 0x4c, 0xcc, 0x2f, 0x4a, 0xd7, 0x4b, 0xc9, 0x2f, 0x49, 0x4e, 0xcd, 0xd1, 0x83, 0x68, 0x50,
	0x52, 0xe6, 0xe2, 0x75, 0x81, 0xe8, 0x71, 0xaa, 0xf4, 0x48, 0x2c, 0xce, 0x10, 0x12, 0xe2, 0x62,
	0x01, 0xe9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0xdc, 0x4e, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x27, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x1f, 0xc9, 0x12, 0x7d, 0x88, 0x25, 0xfa, 0x15, 0xfa, 0x50, 0x77, 0x95, 0x54,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x9d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x53, 0x94,
	0x1d, 0x17, 0xae, 0x00, 0x00, 0x00,
}

func (m *DepositByHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositByHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositByHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintDepositByHash(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDepositByHash(dAtA []byte, offset int, v uint64) int {
	offset -= sovDepositByHash(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositByHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDepositByHash(uint64(l))
	}
	return n
}

func sovDepositByHash(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDepositByHash(x uint64) (n int) {
	return sovDepositByHash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositByHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDepositByHash
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
			return fmt.Errorf("proto: DepositByHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositByHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDepositByHash
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
				return ErrInvalidLengthDepositByHash
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDepositByHash
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDepositByHash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDepositByHash
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
func skipDepositByHash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDepositByHash
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
					return 0, ErrIntOverflowDepositByHash
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
					return 0, ErrIntOverflowDepositByHash
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
				return 0, ErrInvalidLengthDepositByHash
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDepositByHash
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDepositByHash
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDepositByHash        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDepositByHash          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDepositByHash = fmt.Errorf("proto: unexpected end of group")
)
