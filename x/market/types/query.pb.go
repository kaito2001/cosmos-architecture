// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/market/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QuerySwapRequest is the request type for the Query/Swap RPC method.
type QuerySwapRequest struct {
	// offer_coin defines the coin being offered (i.e. 1000000uluna)
	OfferCoin string `protobuf:"bytes,1,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin,omitempty"`
	// ask_denom defines the denom of the coin to swap to
	AskDenom string `protobuf:"bytes,2,opt,name=ask_denom,json=askDenom,proto3" json:"ask_denom,omitempty"`
}

func (m *QuerySwapRequest) Reset()         { *m = QuerySwapRequest{} }
func (m *QuerySwapRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwapRequest) ProtoMessage()    {}
func (*QuerySwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{0}
}
func (m *QuerySwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapRequest.Merge(m, src)
}
func (m *QuerySwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapRequest proto.InternalMessageInfo

// QuerySwapResponse is the response type for the Query/Swap RPC method.
type QuerySwapResponse struct {
	// return_coin defines the coin returned as a result of the swap simulation.
	ReturnCoin types.Coin `protobuf:"bytes,1,opt,name=return_coin,json=returnCoin,proto3" json:"return_coin"`
}

func (m *QuerySwapResponse) Reset()         { *m = QuerySwapResponse{} }
func (m *QuerySwapResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwapResponse) ProtoMessage()    {}
func (*QuerySwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{1}
}
func (m *QuerySwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwapResponse.Merge(m, src)
}
func (m *QuerySwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwapResponse proto.InternalMessageInfo

func (m *QuerySwapResponse) GetReturnCoin() types.Coin {
	if m != nil {
		return m.ReturnCoin
	}
	return types.Coin{}
}

// QueryTerraPoolDeltaRequest is the request type for the Query/TerraPoolDelta RPC method.
type QueryTerraPoolDeltaRequest struct {
}

func (m *QueryTerraPoolDeltaRequest) Reset()         { *m = QueryTerraPoolDeltaRequest{} }
func (m *QueryTerraPoolDeltaRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTerraPoolDeltaRequest) ProtoMessage()    {}
func (*QueryTerraPoolDeltaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{2}
}
func (m *QueryTerraPoolDeltaRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTerraPoolDeltaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTerraPoolDeltaRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTerraPoolDeltaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTerraPoolDeltaRequest.Merge(m, src)
}
func (m *QueryTerraPoolDeltaRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTerraPoolDeltaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTerraPoolDeltaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTerraPoolDeltaRequest proto.InternalMessageInfo

// QueryTerraPoolDeltaResponse is the response type for the Query/TerraPoolDelta RPC method.
type QueryTerraPoolDeltaResponse struct {
	// terra_pool_delta defines the gap between the TerraPool and the TerraBasePool
	TerraPoolDelta github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=terra_pool_delta,json=terraPoolDelta,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"terra_pool_delta"`
}

func (m *QueryTerraPoolDeltaResponse) Reset()         { *m = QueryTerraPoolDeltaResponse{} }
func (m *QueryTerraPoolDeltaResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTerraPoolDeltaResponse) ProtoMessage()    {}
func (*QueryTerraPoolDeltaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{3}
}
func (m *QueryTerraPoolDeltaResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTerraPoolDeltaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTerraPoolDeltaResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTerraPoolDeltaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTerraPoolDeltaResponse.Merge(m, src)
}
func (m *QueryTerraPoolDeltaResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTerraPoolDeltaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTerraPoolDeltaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTerraPoolDeltaResponse proto.InternalMessageInfo

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c172d0f188bf2fb6, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QuerySwapRequest)(nil), "terra.market.v1beta1.QuerySwapRequest")
	proto.RegisterType((*QuerySwapResponse)(nil), "terra.market.v1beta1.QuerySwapResponse")
	proto.RegisterType((*QueryTerraPoolDeltaRequest)(nil), "terra.market.v1beta1.QueryTerraPoolDeltaRequest")
	proto.RegisterType((*QueryTerraPoolDeltaResponse)(nil), "terra.market.v1beta1.QueryTerraPoolDeltaResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "terra.market.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "terra.market.v1beta1.QueryParamsResponse")
}

func init() { proto.RegisterFile("terra/market/v1beta1/query.proto", fileDescriptor_c172d0f188bf2fb6) }

var fileDescriptor_c172d0f188bf2fb6 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0xfb, 0xf5, 0x8b, 0x9a, 0x29, 0xaa, 0xca, 0x90, 0x45, 0x71, 0x83, 0x53, 0x2c, 0x14,
	0x02, 0x52, 0x67, 0x48, 0xd9, 0x75, 0x85, 0x42, 0x1e, 0xa0, 0x0d, 0x20, 0x55, 0x6c, 0xa2, 0x49,
	0x32, 0x35, 0x56, 0x62, 0x5f, 0x77, 0x66, 0x42, 0x89, 0xd8, 0xc1, 0x86, 0x25, 0x12, 0x2f, 0xd0,
	0x0d, 0x2f, 0xc0, 0x53, 0x74, 0x59, 0x89, 0x0d, 0x62, 0x51, 0xa1, 0x84, 0x05, 0x8f, 0x81, 0xe6,
	0x27, 0xd0, 0x54, 0x56, 0x81, 0x55, 0xe2, 0x7b, 0xcf, 0x3d, 0xe7, 0xf8, 0xdc, 0x6b, 0xb4, 0xa5,
	0xb8, 0x10, 0x8c, 0xa6, 0x4c, 0x0c, 0xb9, 0xa2, 0x2f, 0x9b, 0x3d, 0xae, 0x58, 0x93, 0x1e, 0x8d,
	0xb9, 0x98, 0x90, 0x5c, 0x80, 0x02, 0x5c, 0x31, 0x08, 0x62, 0x11, 0xc4, 0x21, 0x82, 0x4a, 0x0c,
	0x31, 0x18, 0x00, 0xd5, 0xff, 0x2c, 0x36, 0xa8, 0xc6, 0x00, 0xf1, 0x88, 0x53, 0x96, 0x27, 0x94,
	0x65, 0x19, 0x28, 0xa6, 0x12, 0xc8, 0xa4, 0xeb, 0xde, 0x2e, 0xd4, 0x72, 0xc4, 0x16, 0x12, 0xf6,
	0x41, 0xa6, 0x20, 0x69, 0x8f, 0x49, 0xfe, 0x0b, 0xd1, 0x87, 0x24, 0xb3, 0xfd, 0xe8, 0x00, 0xad,
	0xef, 0x6b, 0x6f, 0x4f, 0x8e, 0x59, 0xde, 0xe1, 0x47, 0x63, 0x2e, 0x15, 0xbe, 0x85, 0x10, 0x1c,
	0x1e, 0x72, 0xd1, 0xd5, 0xb8, 0x0d, 0x7f, 0xcb, 0x6f, 0x94, 0x3b, 0x65, 0x53, 0x79, 0x0c, 0x49,
	0x86, 0x37, 0x51, 0x99, 0xc9, 0x61, 0x77, 0xc0, 0x33, 0x48, 0x37, 0x96, 0x4c, 0x77, 0x85, 0xc9,
	0x61, 0x5b, 0x3f, 0xef, 0xae, 0xbc, 0x3b, 0xa9, 0x79, 0x3f, 0x4e, 0x6a, 0x5e, 0xf4, 0x0c, 0x5d,
	0xbf, 0xc0, 0x2c, 0x73, 0xc8, 0x24, 0xc7, 0x8f, 0xd0, 0xaa, 0xe0, 0x6a, 0x2c, 0xb2, 0xdf, 0xdc,
	0xab, 0x3b, 0x37, 0x89, 0x35, 0x49, 0xb4, 0xc9, 0x79, 0x20, 0x44, 0x6b, 0xb5, 0x96, 0x4f, 0xcf,
	0x6b, 0x5e, 0x07, 0xd9, 0x19, 0x5d, 0x89, 0xaa, 0x28, 0x30, 0xb4, 0x4f, 0xf5, 0xab, 0xef, 0x01,
	0x8c, 0xda, 0x7c, 0xa4, 0x98, 0xb3, 0x1e, 0x1d, 0xa3, 0xcd, 0xc2, 0xae, 0x93, 0x3f, 0x40, 0xeb,
	0x26, 0xb2, 0x6e, 0x0e, 0x30, 0xea, 0x0e, 0x74, 0xcf, 0x78, 0xb8, 0xd6, 0x22, 0x5a, 0xe8, 0xeb,
	0x79, 0xad, 0x1e, 0x27, 0xea, 0xc5, 0xb8, 0x47, 0xfa, 0x90, 0x52, 0x17, 0x9d, 0xfd, 0xd9, 0x96,
	0x83, 0x21, 0x55, 0x93, 0x9c, 0x4b, 0xd2, 0xe6, 0xfd, 0xce, 0x9a, 0x5a, 0x50, 0x88, 0x2a, 0x08,
	0x1b, 0xe1, 0x3d, 0x26, 0x58, 0x2a, 0xe7, 0x76, 0xf6, 0xd1, 0x8d, 0x85, 0xaa, 0xb3, 0xb1, 0x8b,
	0x4a, 0xb9, 0xa9, 0xb8, 0x00, 0xaa, 0xa4, 0xe8, 0x24, 0x88, 0x9d, 0x72, 0x19, 0xb8, 0x89, 0x9d,
	0x4f, 0xff, 0xa1, 0xff, 0x0d, 0x27, 0x7e, 0x8d, 0x96, 0x75, 0xb6, 0xb8, 0x5e, 0x3c, 0x7d, 0x79,
	0xad, 0xc1, 0xdd, 0x3f, 0xe2, 0xac, 0xbd, 0x28, 0x7a, 0xf3, 0xf9, 0xfb, 0x87, 0xa5, 0x2a, 0x0e,
	0x68, 0xe1, 0x7d, 0x49, 0x2d, 0xfa, 0xd1, 0x47, 0x6b, 0x8b, 0x21, 0xe3, 0x07, 0x57, 0xf0, 0x17,
	0x6e, 0x2b, 0x68, 0xfe, 0xc3, 0x84, 0xf3, 0x46, 0x8c, 0xb7, 0x06, 0xae, 0x17, 0x7b, 0xbb, 0xbc,
	0x5d, 0xfc, 0xd6, 0x47, 0x25, 0x9b, 0x23, 0x6e, 0x5c, 0xa1, 0xb6, 0xb0, 0xb6, 0xe0, 0xde, 0x5f,
	0x20, 0x9d, 0x9f, 0x3b, 0xc6, 0x4f, 0x88, 0xab, 0xc5, 0x7e, 0xec, 0xd2, 0x5a, 0xed, 0xd3, 0x69,
	0xe8, 0x9f, 0x4d, 0x43, 0xff, 0xdb, 0x34, 0xf4, 0xdf, 0xcf, 0x42, 0xef, 0x6c, 0x16, 0x7a, 0x5f,
	0x66, 0xa1, 0xf7, 0xfc, 0xfe, 0x85, 0x7b, 0x33, 0x0c, 0xdb, 0x29, 0x64, 0x7c, 0x42, 0xfb, 0x20,
	0x38, 0x7d, 0x35, 0xa7, 0x33, 0x77, 0xd7, 0x2b, 0x99, 0x4f, 0xf6, 0xe1, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd8, 0x99, 0x62, 0x65, 0x63, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Swap returns simulated swap amount.
	Swap(ctx context.Context, in *QuerySwapRequest, opts ...grpc.CallOption) (*QuerySwapResponse, error)
	// TerraPoolDelta returns terra_pool_delta amount.
	TerraPoolDelta(ctx context.Context, in *QueryTerraPoolDeltaRequest, opts ...grpc.CallOption) (*QueryTerraPoolDeltaResponse, error)
	// Params queries all parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Swap(ctx context.Context, in *QuerySwapRequest, opts ...grpc.CallOption) (*QuerySwapResponse, error) {
	out := new(QuerySwapResponse)
	err := c.cc.Invoke(ctx, "/terra.market.v1beta1.Query/Swap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TerraPoolDelta(ctx context.Context, in *QueryTerraPoolDeltaRequest, opts ...grpc.CallOption) (*QueryTerraPoolDeltaResponse, error) {
	out := new(QueryTerraPoolDeltaResponse)
	err := c.cc.Invoke(ctx, "/terra.market.v1beta1.Query/TerraPoolDelta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/terra.market.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Swap returns simulated swap amount.
	Swap(context.Context, *QuerySwapRequest) (*QuerySwapResponse, error)
	// TerraPoolDelta returns terra_pool_delta amount.
	TerraPoolDelta(context.Context, *QueryTerraPoolDeltaRequest) (*QueryTerraPoolDeltaResponse, error)
	// Params queries all parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Swap(ctx context.Context, req *QuerySwapRequest) (*QuerySwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swap not implemented")
}
func (*UnimplementedQueryServer) TerraPoolDelta(ctx context.Context, req *QueryTerraPoolDeltaRequest) (*QueryTerraPoolDeltaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TerraPoolDelta not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Swap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Swap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.market.v1beta1.Query/Swap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Swap(ctx, req.(*QuerySwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TerraPoolDelta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTerraPoolDeltaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TerraPoolDelta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.market.v1beta1.Query/TerraPoolDelta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TerraPoolDelta(ctx, req.(*QueryTerraPoolDeltaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.market.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.market.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Swap",
			Handler:    _Query_Swap_Handler,
		},
		{
			MethodName: "TerraPoolDelta",
			Handler:    _Query_TerraPoolDelta_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/market/v1beta1/query.proto",
}

func (m *QuerySwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AskDenom) > 0 {
		i -= len(m.AskDenom)
		copy(dAtA[i:], m.AskDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AskDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OfferCoin) > 0 {
		i -= len(m.OfferCoin)
		copy(dAtA[i:], m.OfferCoin)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OfferCoin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ReturnCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryTerraPoolDeltaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTerraPoolDeltaRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTerraPoolDeltaRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTerraPoolDeltaResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTerraPoolDeltaResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTerraPoolDeltaResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TerraPoolDelta.Size()
		i -= size
		if _, err := m.TerraPoolDelta.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OfferCoin)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.AskDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ReturnCoin.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTerraPoolDeltaRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTerraPoolDeltaResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TerraPoolDelta.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferCoin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OfferCoin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AskDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReturnCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTerraPoolDeltaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTerraPoolDeltaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTerraPoolDeltaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTerraPoolDeltaResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTerraPoolDeltaResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTerraPoolDeltaResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerraPoolDelta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TerraPoolDelta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
