// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Params defines the parameters for the feeabs module.
type Params struct {
	// native ibced in osmosis
	NativeIbcedInOsmosis string `protobuf:"bytes,1,opt,name=native_ibced_in_osmosis,json=nativeIbcedInOsmosis,proto3" json:"native_ibced_in_osmosis,omitempty"`
	// we'll update the fee rate each `osmosis_exchange_rate_update_period`
	OsmosisExchangeRateUpdatePeriod time.Duration `protobuf:"bytes,2,opt,name=osmosis_exchange_rate_update_period,json=osmosisExchangeRateUpdatePeriod,proto3,stdduration" json:"osmosis_exchange_rate_update_period"`
	// we'll swap our accumulated osmosis fee to native token each `accumulated_osmosis_fee_swap_period`
	AccumulatedOsmosisFeeSwapPeriod time.Duration `protobuf:"bytes,3,opt,name=accumulated_osmosis_fee_swap_period,json=accumulatedOsmosisFeeSwapPeriod,proto3,stdduration" json:"accumulated_osmosis_fee_swap_period"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e66a0978c84086, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetNativeIbcedInOsmosis() string {
	if m != nil {
		return m.NativeIbcedInOsmosis
	}
	return ""
}

func (m *Params) GetOsmosisExchangeRateUpdatePeriod() time.Duration {
	if m != nil {
		return m.OsmosisExchangeRateUpdatePeriod
	}
	return 0
}

func (m *Params) GetAccumulatedOsmosisFeeSwapPeriod() time.Duration {
	if m != nil {
		return m.AccumulatedOsmosisFeeSwapPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "feeabstraction.absfee.v1beta1.Params")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/params.proto", fileDescriptor_64e66a0978c84086)
}

var fileDescriptor_64e66a0978c84086 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0xed, 0x14, 0x42, 0xeb, 0xee, 0x42, 0xa0, 0x69, 0x68, 0x9d, 0xd0, 0x6e, 0x42, 0xa0,
	0x16, 0x69, 0xe9, 0xa2, 0xdb, 0xd0, 0x16, 0xb2, 0x6a, 0x48, 0xe8, 0xa6, 0x1b, 0x71, 0x65, 0x5f,
	0x3b, 0x02, 0x5b, 0x52, 0x2c, 0x39, 0x3f, 0x6f, 0xd1, 0x65, 0x5f, 0xa0, 0xef, 0x92, 0x65, 0x96,
	0xb3, 0x9a, 0x19, 0x92, 0x17, 0x19, 0x2c, 0x2b, 0x90, 0x99, 0xdd, 0xac, 0xf4, 0x73, 0xce, 0xfd,
	0xce, 0x11, 0x28, 0x18, 0xa7, 0x88, 0xc0, 0xb4, 0x29, 0x21, 0x36, 0x5c, 0x0a, 0x02, 0x4c, 0xa7,
	0x88, 0x64, 0x33, 0x61, 0x68, 0x60, 0x42, 0x14, 0x94, 0x50, 0xe8, 0x48, 0x95, 0xd2, 0xc8, 0xce,
	0xfb, 0xc7, 0xde, 0xa8, 0xf1, 0x46, 0xce, 0xdb, 0xef, 0x66, 0x32, 0x93, 0xd6, 0x49, 0xea, 0x5d,
	0x33, 0xd4, 0x7f, 0x97, 0x49, 0x99, 0xe5, 0x48, 0x40, 0x71, 0x02, 0x42, 0x48, 0x03, 0xf5, 0xac,
	0x43, 0xf6, 0xc7, 0xb1, 0xd4, 0x85, 0xd4, 0x84, 0x81, 0x46, 0xb2, 0xae, 0xb0, 0xdc, 0x5f, 0x45,
	0x67, 0x5c, 0x58, 0xb3, 0xf3, 0x86, 0x8e, 0x64, 0x4f, 0xac, 0x4a, 0x49, 0x52, 0x95, 0x57, 0xfa,
	0x87, 0xff, 0xad, 0xa0, 0x3d, 0xb7, 0x7d, 0x3b, 0x5f, 0x83, 0x37, 0xf5, 0xe8, 0x06, 0x29, 0x67,
	0x31, 0x26, 0x94, 0x0b, 0x6a, 0x73, 0xb8, 0xee, 0xf9, 0x43, 0x7f, 0xf4, 0x6a, 0xd1, 0x6d, 0xe4,
	0x59, 0xad, 0xce, 0xc4, 0xaf, 0x46, 0xeb, 0xac, 0x83, 0x8f, 0xce, 0x46, 0x71, 0x17, 0xaf, 0x40,
	0x64, 0x48, 0x4b, 0x30, 0x48, 0x2b, 0x95, 0xd4, 0x8b, 0xc2, 0x92, 0xcb, 0xa4, 0xd7, 0x1a, 0xfa,
	0xa3, 0xd7, 0x9f, 0xdf, 0x46, 0x4d, 0x9f, 0xe8, 0xd2, 0x27, 0xfa, 0xee, 0xfa, 0x4c, 0x5f, 0x1e,
	0x6e, 0x07, 0xde, 0xbf, 0xbb, 0x81, 0xbf, 0x18, 0x38, 0xde, 0x0f, 0x87, 0x5b, 0x80, 0xc1, 0xdf,
	0x16, 0x36, 0xb7, 0xac, 0x3a, 0x12, 0xe2, 0xb8, 0x2a, 0xaa, 0x1c, 0x0c, 0x26, 0x97, 0x96, 0x34,
	0x45, 0xa4, 0x7a, 0x0b, 0xea, 0x12, 0xf9, 0xe2, 0x19, 0x91, 0x57, 0x3c, 0xf7, 0xb0, 0x9f, 0x88,
	0xcb, 0x2d, 0xa8, 0x26, 0x72, 0xba, 0x3c, 0x9c, 0x42, 0xff, 0x78, 0x0a, 0xfd, 0xfb, 0x53, 0xe8,
	0xff, 0x3d, 0x87, 0xde, 0xf1, 0x1c, 0x7a, 0x37, 0xe7, 0xd0, 0xfb, 0xf3, 0x2d, 0xe3, 0x66, 0x55,
	0xb1, 0x28, 0x96, 0x05, 0x11, 0xb2, 0xe6, 0x42, 0xfe, 0x29, 0x07, 0xa6, 0xc9, 0x93, 0x5f, 0xb2,
	0x99, 0x90, 0x9d, 0xbb, 0x23, 0x66, 0xaf, 0x50, 0xb3, 0xb6, 0xad, 0xf4, 0xe5, 0x21, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x2b, 0x79, 0x66, 0x50, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.AccumulatedOsmosisFeeSwapPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.AccumulatedOsmosisFeeSwapPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.OsmosisExchangeRateUpdatePeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.OsmosisExchangeRateUpdatePeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.NativeIbcedInOsmosis) > 0 {
		i -= len(m.NativeIbcedInOsmosis)
		copy(dAtA[i:], m.NativeIbcedInOsmosis)
		i = encodeVarintParams(dAtA, i, uint64(len(m.NativeIbcedInOsmosis)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NativeIbcedInOsmosis)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.OsmosisExchangeRateUpdatePeriod)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.AccumulatedOsmosisFeeSwapPeriod)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeIbcedInOsmosis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NativeIbcedInOsmosis = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsmosisExchangeRateUpdatePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.OsmosisExchangeRateUpdatePeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulatedOsmosisFeeSwapPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.AccumulatedOsmosisFeeSwapPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
