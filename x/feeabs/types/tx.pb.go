// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Msg fund module account
type MsgFundFeeAbsModuleAccount struct {
	FromAddress string                                   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgFundFeeAbsModuleAccount) Reset()         { *m = MsgFundFeeAbsModuleAccount{} }
func (m *MsgFundFeeAbsModuleAccount) String() string { return proto.CompactTextString(m) }
func (*MsgFundFeeAbsModuleAccount) ProtoMessage()    {}
func (*MsgFundFeeAbsModuleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{0}
}
func (m *MsgFundFeeAbsModuleAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundFeeAbsModuleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundFeeAbsModuleAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundFeeAbsModuleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundFeeAbsModuleAccount.Merge(m, src)
}
func (m *MsgFundFeeAbsModuleAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundFeeAbsModuleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundFeeAbsModuleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundFeeAbsModuleAccount proto.InternalMessageInfo

func (m *MsgFundFeeAbsModuleAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgFundFeeAbsModuleAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgFundFeeAbsModuleAccountResponse struct {
}

func (m *MsgFundFeeAbsModuleAccountResponse) Reset()         { *m = MsgFundFeeAbsModuleAccountResponse{} }
func (m *MsgFundFeeAbsModuleAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundFeeAbsModuleAccountResponse) ProtoMessage()    {}
func (*MsgFundFeeAbsModuleAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{1}
}
func (m *MsgFundFeeAbsModuleAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundFeeAbsModuleAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundFeeAbsModuleAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundFeeAbsModuleAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundFeeAbsModuleAccountResponse.Merge(m, src)
}
func (m *MsgFundFeeAbsModuleAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundFeeAbsModuleAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundFeeAbsModuleAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundFeeAbsModuleAccountResponse proto.InternalMessageInfo

// Params defines the parameters for the feeabs module.
type MsgSendQueryIbcDenomTWAP struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSendQueryIbcDenomTWAP) Reset()         { *m = MsgSendQueryIbcDenomTWAP{} }
func (m *MsgSendQueryIbcDenomTWAP) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryIbcDenomTWAP) ProtoMessage()    {}
func (*MsgSendQueryIbcDenomTWAP) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{2}
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryIbcDenomTWAP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAP.Merge(m, src)
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAP.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryIbcDenomTWAP proto.InternalMessageInfo

func (m *MsgSendQueryIbcDenomTWAP) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type MsgSendQueryIbcDenomTWAPResponse struct {
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Reset()         { *m = MsgSendQueryIbcDenomTWAPResponse{} }
func (m *MsgSendQueryIbcDenomTWAPResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryIbcDenomTWAPResponse) ProtoMessage()    {}
func (*MsgSendQueryIbcDenomTWAPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{3}
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.Merge(m, src)
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse proto.InternalMessageInfo

// Params defines the parameters for the feeabs module.
type MsgSwapCrossChain struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	IbcDenom    string `protobuf:"bytes,2,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
}

func (m *MsgSwapCrossChain) Reset()         { *m = MsgSwapCrossChain{} }
func (m *MsgSwapCrossChain) String() string { return proto.CompactTextString(m) }
func (*MsgSwapCrossChain) ProtoMessage()    {}
func (*MsgSwapCrossChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{4}
}
func (m *MsgSwapCrossChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapCrossChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapCrossChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapCrossChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapCrossChain.Merge(m, src)
}
func (m *MsgSwapCrossChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapCrossChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapCrossChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapCrossChain proto.InternalMessageInfo

func (m *MsgSwapCrossChain) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSwapCrossChain) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

type MsgSwapCrossChainResponse struct {
}

func (m *MsgSwapCrossChainResponse) Reset()         { *m = MsgSwapCrossChainResponse{} }
func (m *MsgSwapCrossChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapCrossChainResponse) ProtoMessage()    {}
func (*MsgSwapCrossChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{5}
}
func (m *MsgSwapCrossChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapCrossChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapCrossChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapCrossChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapCrossChainResponse.Merge(m, src)
}
func (m *MsgSwapCrossChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapCrossChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapCrossChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapCrossChainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFundFeeAbsModuleAccount)(nil), "feeabstraction.absfee.v1beta1.MsgFundFeeAbsModuleAccount")
	proto.RegisterType((*MsgFundFeeAbsModuleAccountResponse)(nil), "feeabstraction.absfee.v1beta1.MsgFundFeeAbsModuleAccountResponse")
	proto.RegisterType((*MsgSendQueryIbcDenomTWAP)(nil), "feeabstraction.absfee.v1beta1.MsgSendQueryIbcDenomTWAP")
	proto.RegisterType((*MsgSendQueryIbcDenomTWAPResponse)(nil), "feeabstraction.absfee.v1beta1.MsgSendQueryIbcDenomTWAPResponse")
	proto.RegisterType((*MsgSwapCrossChain)(nil), "feeabstraction.absfee.v1beta1.MsgSwapCrossChain")
	proto.RegisterType((*MsgSwapCrossChainResponse)(nil), "feeabstraction.absfee.v1beta1.MsgSwapCrossChainResponse")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/tx.proto", fileDescriptor_84c172c34645b936)
}

var fileDescriptor_84c172c34645b936 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x8e, 0x1b, 0xa9, 0xa2, 0x53, 0x40, 0x22, 0x02, 0x91, 0xba, 0xe0, 0x04, 0x0b, 0xa1, 0x80,
	0x14, 0x0f, 0x0d, 0x0b, 0xa0, 0x12, 0x42, 0x6e, 0x50, 0x25, 0x16, 0x91, 0x20, 0xad, 0x84, 0xc4,
	0xa6, 0x9a, 0xb1, 0x27, 0x8e, 0x45, 0x3c, 0xcf, 0xf8, 0x8d, 0x4b, 0x23, 0x71, 0x08, 0x58, 0x71,
	0x07, 0x76, 0x2c, 0xb8, 0x43, 0x97, 0x5d, 0xb2, 0x2a, 0x28, 0xb9, 0x01, 0x27, 0x40, 0xfe, 0x8b,
	0x52, 0xc0, 0x21, 0x74, 0x65, 0x7b, 0xe6, 0xfb, 0x7b, 0xcf, 0x6f, 0x86, 0xdc, 0x19, 0x08, 0xc1,
	0x38, 0xaa, 0x88, 0x39, 0xca, 0x07, 0x49, 0x19, 0xc7, 0x81, 0x10, 0xf4, 0x70, 0x8b, 0x0b, 0xc5,
	0xb6, 0xa8, 0x3a, 0xb2, 0xc2, 0x08, 0x14, 0xd4, 0x6e, 0x9e, 0xc5, 0x59, 0x19, 0xce, 0xca, 0x71,
	0xfa, 0x55, 0x0f, 0x3c, 0x48, 0x91, 0x34, 0x79, 0xcb, 0x48, 0xfa, 0x0d, 0x0f, 0xc0, 0x1b, 0x09,
	0xca, 0x42, 0x9f, 0x32, 0x29, 0x41, 0xb1, 0x84, 0x8b, 0xf9, 0xee, 0x3d, 0x07, 0x30, 0x00, 0xa4,
	0x9c, 0xa1, 0xa0, 0x6f, 0x63, 0x11, 0x8d, 0x67, 0xb6, 0x21, 0xf3, 0x7c, 0x99, 0x82, 0x0b, 0xec,
	0xe2, 0x98, 0x21, 0x8b, 0x58, 0x50, 0xe8, 0xde, 0x5d, 0x8c, 0x15, 0x21, 0x38, 0xc3, 0x1c, 0x6a,
	0xcc, 0x47, 0x28, 0x00, 0x0e, 0xf8, 0x85, 0x6d, 0x23, 0x2f, 0x20, 0xfd, 0xe2, 0xf1, 0x80, 0x2a,
	0x3f, 0x10, 0xa8, 0x58, 0x10, 0x66, 0x00, 0xf3, 0xab, 0x46, 0xf4, 0x1e, 0x7a, 0xbb, 0xb1, 0x74,
	0x77, 0x85, 0xb0, 0x39, 0xf6, 0xc0, 0x8d, 0x47, 0xc2, 0x76, 0x1c, 0x88, 0xa5, 0xaa, 0xdd, 0x22,
	0x17, 0x07, 0x11, 0x04, 0x07, 0xcc, 0x75, 0x23, 0x81, 0x58, 0xd7, 0x9a, 0x5a, 0x6b, 0xad, 0xbf,
	0x9e, 0xac, 0xd9, 0xd9, 0x52, 0x4d, 0x91, 0x55, 0x16, 0x24, 0xe0, 0xfa, 0x4a, 0xb3, 0xda, 0x5a,
	0xef, 0x6c, 0x58, 0x59, 0x26, 0x2b, 0xc9, 0x54, 0xf4, 0xd7, 0xea, 0x82, 0x2f, 0x77, 0xec, 0xe3,
	0xd3, 0x46, 0xe5, 0xe7, 0x69, 0xe3, 0xd2, 0x98, 0x05, 0xa3, 0x6d, 0x33, 0xa3, 0x99, 0x9f, 0xbf,
	0x37, 0x5a, 0x9e, 0xaf, 0x86, 0x31, 0xb7, 0x1c, 0x08, 0x68, 0x5e, 0x51, 0xf6, 0x68, 0xa3, 0xfb,
	0x86, 0xaa, 0x71, 0x28, 0x30, 0x55, 0xc0, 0x7e, 0xee, 0x65, 0xde, 0x26, 0x66, 0x79, 0xec, 0xbe,
	0xc0, 0x10, 0x24, 0x0a, 0xf3, 0x09, 0xa9, 0xf7, 0xd0, 0xdb, 0x13, 0xd2, 0x7d, 0x99, 0xfc, 0x9e,
	0xe7, 0xdc, 0x79, 0x26, 0x24, 0x04, 0xfb, 0xaf, 0xec, 0x17, 0x4b, 0x94, 0x66, 0x9a, 0xa4, 0x59,
	0x46, 0x9f, 0x59, 0xec, 0x91, 0x2b, 0x09, 0xe6, 0x1d, 0x0b, 0xbb, 0x11, 0x20, 0x76, 0x87, 0xcc,
	0x97, 0xcb, 0xb4, 0x6d, 0x93, 0xac, 0xf9, 0xdc, 0x39, 0x70, 0x13, 0xc1, 0xfa, 0x4a, 0xba, 0x7f,
	0xc1, 0xcf, 0x0d, 0xcc, 0x4d, 0xb2, 0xf1, 0x87, 0x68, 0xe1, 0xd8, 0xf9, 0x52, 0x25, 0xd5, 0x1e,
	0x7a, 0xb5, 0x8f, 0x1a, 0xb9, 0xf6, 0xf7, 0xd2, 0x1e, 0x5a, 0x0b, 0x87, 0xdd, 0x2a, 0x2b, 0x4a,
	0x7f, 0x7a, 0x4e, 0x62, 0x91, 0xad, 0xf6, 0x9e, 0x5c, 0xfe, 0xad, 0x15, 0xf7, 0x97, 0x90, 0x3c,
	0xc3, 0xd0, 0x1f, 0xfd, 0x2f, 0x63, 0xe6, 0xfe, 0x49, 0x23, 0xd7, 0xcb, 0x26, 0xf9, 0xf1, 0xbf,
	0x55, 0x4b, 0xa8, 0xba, 0x7d, 0x6e, 0x6a, 0x91, 0x6c, 0x67, 0xff, 0x78, 0x62, 0x68, 0x27, 0x13,
	0x43, 0xfb, 0x31, 0x31, 0xb4, 0x0f, 0x53, 0xa3, 0x72, 0x32, 0x35, 0x2a, 0xdf, 0xa6, 0x46, 0xe5,
	0xf5, 0xf6, 0xdc, 0xe8, 0x4b, 0x48, 0xe4, 0xd9, 0xa8, 0x3d, 0x62, 0x1c, 0xe9, 0x40, 0x88, 0xf6,
	0xfc, 0x35, 0x70, 0xd8, 0xa1, 0x47, 0x34, 0x4b, 0x92, 0x1d, 0x09, 0xbe, 0x9a, 0x9e, 0xe1, 0x07,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xbc, 0x21, 0xe7, 0x04, 0x05, 0x00, 0x00,
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
	SendQueryIbcDenomTWAP(ctx context.Context, in *MsgSendQueryIbcDenomTWAP, opts ...grpc.CallOption) (*MsgSendQueryIbcDenomTWAPResponse, error)
	SwapCrossChain(ctx context.Context, in *MsgSwapCrossChain, opts ...grpc.CallOption) (*MsgSwapCrossChainResponse, error)
	FundFeeAbsModuleAccount(ctx context.Context, in *MsgFundFeeAbsModuleAccount, opts ...grpc.CallOption) (*MsgFundFeeAbsModuleAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendQueryIbcDenomTWAP(ctx context.Context, in *MsgSendQueryIbcDenomTWAP, opts ...grpc.CallOption) (*MsgSendQueryIbcDenomTWAPResponse, error) {
	out := new(MsgSendQueryIbcDenomTWAPResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Msg/SendQueryIbcDenomTWAP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapCrossChain(ctx context.Context, in *MsgSwapCrossChain, opts ...grpc.CallOption) (*MsgSwapCrossChainResponse, error) {
	out := new(MsgSwapCrossChainResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Msg/SwapCrossChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FundFeeAbsModuleAccount(ctx context.Context, in *MsgFundFeeAbsModuleAccount, opts ...grpc.CallOption) (*MsgFundFeeAbsModuleAccountResponse, error) {
	out := new(MsgFundFeeAbsModuleAccountResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Msg/FundFeeAbsModuleAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendQueryIbcDenomTWAP(context.Context, *MsgSendQueryIbcDenomTWAP) (*MsgSendQueryIbcDenomTWAPResponse, error)
	SwapCrossChain(context.Context, *MsgSwapCrossChain) (*MsgSwapCrossChainResponse, error)
	FundFeeAbsModuleAccount(context.Context, *MsgFundFeeAbsModuleAccount) (*MsgFundFeeAbsModuleAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendQueryIbcDenomTWAP(ctx context.Context, req *MsgSendQueryIbcDenomTWAP) (*MsgSendQueryIbcDenomTWAPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQueryIbcDenomTWAP not implemented")
}
func (*UnimplementedMsgServer) SwapCrossChain(ctx context.Context, req *MsgSwapCrossChain) (*MsgSwapCrossChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapCrossChain not implemented")
}
func (*UnimplementedMsgServer) FundFeeAbsModuleAccount(ctx context.Context, req *MsgFundFeeAbsModuleAccount) (*MsgFundFeeAbsModuleAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundFeeAbsModuleAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendQueryIbcDenomTWAP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendQueryIbcDenomTWAP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendQueryIbcDenomTWAP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Msg/SendQueryIbcDenomTWAP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendQueryIbcDenomTWAP(ctx, req.(*MsgSendQueryIbcDenomTWAP))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapCrossChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapCrossChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapCrossChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Msg/SwapCrossChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapCrossChain(ctx, req.(*MsgSwapCrossChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundFeeAbsModuleAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundFeeAbsModuleAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundFeeAbsModuleAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Msg/FundFeeAbsModuleAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundFeeAbsModuleAccount(ctx, req.(*MsgFundFeeAbsModuleAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feeabstraction.absfee.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendQueryIbcDenomTWAP",
			Handler:    _Msg_SendQueryIbcDenomTWAP_Handler,
		},
		{
			MethodName: "SwapCrossChain",
			Handler:    _Msg_SwapCrossChain_Handler,
		},
		{
			MethodName: "FundFeeAbsModuleAccount",
			Handler:    _Msg_FundFeeAbsModuleAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feeabstraction/absfee/v1beta1/tx.proto",
}

func (m *MsgFundFeeAbsModuleAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundFeeAbsModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundFeeAbsModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundFeeAbsModuleAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundFeeAbsModuleAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundFeeAbsModuleAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryIbcDenomTWAP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryIbcDenomTWAP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryIbcDenomTWAP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryIbcDenomTWAPResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryIbcDenomTWAPResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapCrossChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapCrossChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapCrossChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapCrossChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapCrossChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapCrossChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgFundFeeAbsModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgFundFeeAbsModuleAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendQueryIbcDenomTWAP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapCrossChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapCrossChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFundFeeAbsModuleAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFundFeeAbsModuleAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundFeeAbsModuleAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgFundFeeAbsModuleAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFundFeeAbsModuleAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundFeeAbsModuleAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgSendQueryIbcDenomTWAP) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSendQueryIbcDenomTWAPResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAPResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAPResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgSwapCrossChain) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapCrossChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapCrossChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
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
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
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
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSwapCrossChainResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSwapCrossChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapCrossChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
