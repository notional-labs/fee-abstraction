// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/feepool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
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

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/feepool.proto", fileDescriptor_aeac12e01fe4cbff)
}

var fileDescriptor_aeac12e01fe4cbff = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0x31, 0x4e, 0xc5, 0x30,
	0x0c, 0x86, 0xdb, 0x85, 0x81, 0x11, 0x31, 0x3d, 0x41, 0x0e, 0x00, 0x22, 0x56, 0xc5, 0xc4, 0xca,
	0x15, 0xd8, 0xd8, 0x9c, 0xca, 0x0d, 0x91, 0xd2, 0x38, 0xc4, 0x79, 0x15, 0xbd, 0x05, 0xc7, 0x62,
	0xec, 0xc8, 0x88, 0xda, 0x8b, 0xa0, 0xb6, 0x01, 0x09, 0x86, 0xb7, 0x59, 0xf6, 0xf7, 0x7f, 0xfe,
	0xcf, 0x6f, 0x3b, 0x22, 0x34, 0x92, 0x13, 0xb6, 0xd9, 0x71, 0x00, 0x34, 0xd2, 0x11, 0xc1, 0xd0,
	0x18, 0xca, 0xd8, 0x40, 0x47, 0x14, 0x99, 0xbd, 0x8e, 0x89, 0x33, 0x5f, 0x5c, 0xff, 0x85, 0xf5,
	0x0e, 0xeb, 0x02, 0x1f, 0x2e, 0x2d, 0x5b, 0xde, 0x48, 0x58, 0xa7, 0x3d, 0x74, 0xb8, 0xb2, 0xcc,
	0xd6, 0x13, 0x60, 0x74, 0x80, 0x21, 0x70, 0xc6, 0x35, 0x2b, 0xe5, 0x7a, 0xd3, 0xb2, 0xf4, 0x2c,
	0x60, 0x50, 0x08, 0x5e, 0x8f, 0x94, 0xc6, 0xdf, 0xdf, 0x11, 0xad, 0x0b, 0x1b, 0xfc, 0xc3, 0x9e,
	0xee, 0x1a, 0x31, 0x61, 0x5f, 0xbc, 0x8f, 0x4f, 0x1f, 0xb3, 0xaa, 0xa7, 0x59, 0xd5, 0x5f, 0xb3,
	0xaa, 0xdf, 0x17, 0x55, 0x4d, 0x8b, 0xaa, 0x3e, 0x17, 0x55, 0x3d, 0x3f, 0x58, 0x97, 0x5f, 0x8e,
	0x46, 0xb7, 0xdc, 0x43, 0xe0, 0x55, 0x84, 0xfe, 0xce, 0xa3, 0x11, 0xf8, 0xa7, 0x1f, 0x1a, 0x78,
	0x2b, 0x3b, 0xc8, 0x63, 0x24, 0x31, 0x67, 0x9b, 0xfb, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x79,
	0x5e, 0x2e, 0x48, 0x35, 0x01, 0x00, 0x00,
}
