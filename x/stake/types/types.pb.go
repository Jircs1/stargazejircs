// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/stake/types/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	ID       uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	VendorID uint64    `protobuf:"varint,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty" yaml:"vendor_id"`
	Body     string    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
	VoteEnd  time.Time `protobuf:"bytes,5,opt,name=vote_end,json=voteEnd,proto3,stdtime" json:"vote_end" yaml:"vote_end"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc33dc9db3f2c3c8, []int{0}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Post) GetVendorID() uint64 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetVoteEnd() time.Time {
	if m != nil {
		return m.VoteEnd
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Post)(nil), "rocket_protocol.stakebird.x.stake.v1.Post")
}

func init() { proto.RegisterFile("x/stake/types/types.proto", fileDescriptor_bc33dc9db3f2c3c8) }

var fileDescriptor_bc33dc9db3f2c3c8 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x4f, 0xfa, 0x40,
	0x1c, 0xc6, 0x7b, 0xfd, 0xf1, 0x53, 0x28, 0x03, 0xa6, 0x2e, 0xb5, 0x26, 0xbd, 0xa6, 0x38, 0xb0,
	0x78, 0x17, 0x71, 0x33, 0x71, 0x69, 0x74, 0x20, 0x71, 0x30, 0x8d, 0x71, 0x70, 0x21, 0x94, 0x3b,
	0xeb, 0x05, 0xca, 0x97, 0xb4, 0x07, 0x81, 0x77, 0xc1, 0xcb, 0x62, 0x64, 0x74, 0xaa, 0xa6, 0x7d,
	0x07, 0xbc, 0x02, 0xc3, 0x1d, 0x60, 0x5c, 0x2e, 0xdf, 0x3f, 0x9f, 0xe7, 0x2e, 0xcf, 0x73, 0xd6,
	0xc5, 0x82, 0xe6, 0x72, 0x30, 0xe2, 0x54, 0x2e, 0xa7, 0x3c, 0xd7, 0x27, 0x99, 0x66, 0x20, 0xc1,
	0xbe, 0xca, 0x60, 0x38, 0xe2, 0xb2, 0xaf, 0xba, 0x21, 0x8c, 0x89, 0x02, 0x63, 0x91, 0x31, 0xb2,
	0xd0, 0x35, 0x99, 0xdf, 0xb8, 0xae, 0x5a, 0xd3, 0x04, 0x12, 0xf8, 0xad, 0xf4, 0x0d, 0x2e, 0x4e,
	0x00, 0x92, 0x31, 0xa7, 0xaa, 0x8b, 0x67, 0xef, 0x54, 0x8a, 0x94, 0xe7, 0x72, 0x90, 0x4e, 0x35,
	0x10, 0x54, 0xc8, 0xaa, 0x3d, 0x43, 0x2e, 0xed, 0xb6, 0x65, 0x0a, 0xe6, 0x20, 0x1f, 0x75, 0x6a,
	0xe1, 0x79, 0x59, 0x60, 0xb3, 0xf7, 0xb0, 0x2d, 0x70, 0x63, 0x39, 0x48, 0xc7, 0x77, 0x81, 0x60,
	0x41, 0x64, 0x0a, 0x66, 0xdf, 0x5b, 0x8d, 0x39, 0x9f, 0x30, 0xc8, 0xfa, 0x82, 0x39, 0xa6, 0x62,
	0xfd, 0xb2, 0xc0, 0xf5, 0x57, 0x35, 0x54, 0x8a, 0x33, 0xad, 0x38, 0x62, 0x41, 0x54, 0xd7, 0x75,
	0x8f, 0xd9, 0x6d, 0xab, 0x16, 0x03, 0x5b, 0x3a, 0xff, 0x7c, 0xd4, 0x69, 0x84, 0xad, 0x6d, 0x81,
	0x9b, 0x9a, 0xde, 0x4d, 0x83, 0x48, 0x2d, 0xed, 0xc8, 0xaa, 0xcf, 0x41, 0xf2, 0x3e, 0x9f, 0x30,
	0xe7, 0xbf, 0x8f, 0x3a, 0xcd, 0xae, 0x4b, 0xb4, 0x0b, 0x72, 0x70, 0x41, 0x5e, 0x0e, 0x2e, 0xc2,
	0xcb, 0x75, 0x81, 0x8d, 0x6d, 0x81, 0x5b, 0xfb, 0x67, 0xf7, 0xca, 0x60, 0xf5, 0x85, 0x51, 0x74,
	0xba, 0x6b, 0x1f, 0x27, 0x2c, 0x7c, 0x5a, 0x97, 0x1e, 0xda, 0x94, 0x1e, 0xfa, 0x2e, 0x3d, 0xb4,
	0xaa, 0x3c, 0x63, 0x53, 0x79, 0xc6, 0x67, 0xe5, 0x19, 0x6f, 0xdd, 0x44, 0xc8, 0x8f, 0x59, 0x4c,
	0x86, 0x90, 0x52, 0x9d, 0xf6, 0xf5, 0x21, 0x6d, 0x7a, 0x4c, 0x9b, 0xfe, 0xf9, 0xa2, 0xf8, 0x44,
	0x11, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x46, 0xfa, 0xfb, 0xba, 0x01, 0x00, 0x00,
}

func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VoteEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VoteEnd):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VendorID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTypes(uint64(m.ID))
	}
	if m.VendorID != 0 {
		n += 1 + sovTypes(uint64(m.VendorID))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VoteEnd)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VoteEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
