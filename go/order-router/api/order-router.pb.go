// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order-router.proto

package api

import (
	"github.com/ettec/open-trading-platform/go/model"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateAndRouteOrderParams struct {
	OrderSide            model.Side       `protobuf:"varint,1,opt,name=orderSide,proto3,enum=model.Side" json:"orderSide,omitempty"`
	Quantity             *model.Decimal64 `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                *model.Decimal64 `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Listing              *model.Listing   `protobuf:"bytes,4,opt,name=listing,proto3" json:"listing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateAndRouteOrderParams) Reset()         { *m = CreateAndRouteOrderParams{} }
func (m *CreateAndRouteOrderParams) String() string { return proto.CompactTextString(m) }
func (*CreateAndRouteOrderParams) ProtoMessage()    {}
func (*CreateAndRouteOrderParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eac200df04399a0, []int{0}
}

func (m *CreateAndRouteOrderParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAndRouteOrderParams.Unmarshal(m, b)
}
func (m *CreateAndRouteOrderParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAndRouteOrderParams.Marshal(b, m, deterministic)
}
func (m *CreateAndRouteOrderParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAndRouteOrderParams.Merge(m, src)
}
func (m *CreateAndRouteOrderParams) XXX_Size() int {
	return xxx_messageInfo_CreateAndRouteOrderParams.Size(m)
}
func (m *CreateAndRouteOrderParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAndRouteOrderParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAndRouteOrderParams proto.InternalMessageInfo

func (m *CreateAndRouteOrderParams) GetOrderSide() model.Side {
	if m != nil {
		return m.OrderSide
	}
	return model.Side_BUY
}

func (m *CreateAndRouteOrderParams) GetQuantity() *model.Decimal64 {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *CreateAndRouteOrderParams) GetPrice() *model.Decimal64 {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *CreateAndRouteOrderParams) GetListing() *model.Listing {
	if m != nil {
		return m.Listing
	}
	return nil
}

type OrderId struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderId) Reset()         { *m = OrderId{} }
func (m *OrderId) String() string { return proto.CompactTextString(m) }
func (*OrderId) ProtoMessage()    {}
func (*OrderId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eac200df04399a0, []int{1}
}

func (m *OrderId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderId.Unmarshal(m, b)
}
func (m *OrderId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderId.Marshal(b, m, deterministic)
}
func (m *OrderId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderId.Merge(m, src)
}
func (m *OrderId) XXX_Size() int {
	return xxx_messageInfo_OrderId.Size(m)
}
func (m *OrderId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderId.DiscardUnknown(m)
}

var xxx_messageInfo_OrderId proto.InternalMessageInfo

func (m *OrderId) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type CancelOrderParams struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Listing              *model.Listing `protobuf:"bytes,2,opt,name=listing,proto3" json:"listing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOrderParams) Reset()         { *m = CancelOrderParams{} }
func (m *CancelOrderParams) String() string { return proto.CompactTextString(m) }
func (*CancelOrderParams) ProtoMessage()    {}
func (*CancelOrderParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eac200df04399a0, []int{2}
}

func (m *CancelOrderParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOrderParams.Unmarshal(m, b)
}
func (m *CancelOrderParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOrderParams.Marshal(b, m, deterministic)
}
func (m *CancelOrderParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOrderParams.Merge(m, src)
}
func (m *CancelOrderParams) XXX_Size() int {
	return xxx_messageInfo_CancelOrderParams.Size(m)
}
func (m *CancelOrderParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOrderParams.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOrderParams proto.InternalMessageInfo

func (m *CancelOrderParams) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *CancelOrderParams) GetListing() *model.Listing {
	if m != nil {
		return m.Listing
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAndRouteOrderParams)(nil), "executionvenue.CreateAndRouteOrderParams")
	proto.RegisterType((*OrderId)(nil), "executionvenue.OrderId")
	proto.RegisterType((*CancelOrderParams)(nil), "executionvenue.CancelOrderParams")
}

func init() { proto.RegisterFile("order-router.proto", fileDescriptor_9eac200df04399a0) }

var fileDescriptor_9eac200df04399a0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x9b, 0x96, 0x52, 0x7a, 0x81, 0x8a, 0x9a, 0x81, 0x90, 0xa9, 0x04, 0x09, 0xb5, 0x12,
	0x64, 0x08, 0x88, 0x19, 0x28, 0x0c, 0x48, 0x48, 0xa0, 0x30, 0x20, 0xc1, 0x64, 0x92, 0x13, 0xb2,
	0x14, 0xdb, 0xc1, 0x75, 0x10, 0x7d, 0x27, 0x9e, 0x83, 0xe7, 0x42, 0x71, 0x5c, 0x08, 0x85, 0xb0,
	0xc5, 0xf7, 0x7f, 0xf7, 0xfb, 0xfe, 0x8b, 0x81, 0x48, 0x95, 0xa2, 0x3a, 0x54, 0xb2, 0xd0, 0xa8,
	0xc2, 0x5c, 0x49, 0x2d, 0xc9, 0x00, 0xdf, 0x30, 0x29, 0x34, 0x93, 0xe2, 0x15, 0x45, 0x81, 0xfe,
	0x46, 0xc6, 0x66, 0x9a, 0x89, 0xe7, 0x4a, 0xf6, 0x5d, 0xd3, 0x62, 0x0f, 0x43, 0x2e, 0x53, 0xcc,
	0x12, 0xc9, 0xb9, 0x14, 0x55, 0x29, 0xf8, 0x70, 0x60, 0x67, 0xaa, 0x90, 0x6a, 0x3c, 0x13, 0x69,
	0x5c, 0x1a, 0xdf, 0x94, 0x0d, 0xb7, 0x54, 0x51, 0x3e, 0x23, 0x13, 0xe8, 0x9b, 0xfe, 0x3b, 0x96,
	0xa2, 0xe7, 0x8c, 0x9c, 0xf1, 0x20, 0x72, 0x43, 0x63, 0x12, 0x96, 0xa5, 0xf8, 0x5b, 0x25, 0x07,
	0xb0, 0xf6, 0x52, 0x50, 0xa1, 0x99, 0x9e, 0x7b, 0xed, 0x91, 0x33, 0x76, 0xa3, 0x4d, 0x4b, 0x5e,
	0x60, 0xc2, 0x38, 0xcd, 0x4e, 0x8e, 0xe3, 0x2f, 0x82, 0xec, 0x43, 0x37, 0x57, 0x2c, 0x41, 0xaf,
	0xd3, 0x80, 0x56, 0x32, 0x19, 0x43, 0xcf, 0xe6, 0xf1, 0x56, 0x0c, 0x39, 0xb0, 0xe4, 0x75, 0x55,
	0x8d, 0x17, 0x72, 0xb0, 0x07, 0x3d, 0x33, 0xf9, 0x55, 0x4a, 0x3c, 0xe8, 0xc9, 0xea, 0xd3, 0xcc,
	0xdc, 0x8f, 0x17, 0xc7, 0xe0, 0x1e, 0x86, 0x53, 0x2a, 0x12, 0xcc, 0xea, 0x21, 0x1b, 0xf1, 0xfa,
	0xed, 0xed, 0x7f, 0x6f, 0x8f, 0xde, 0x1d, 0x70, 0x8d, 0xa7, 0x59, 0xa1, 0x22, 0x8f, 0xb0, 0xf5,
	0xc7, 0x56, 0xc9, 0x24, 0xfc, 0xf9, 0xb7, 0xc2, 0xc6, 0xd5, 0xfb, 0xdb, 0xcb, 0xa8, 0x4d, 0x17,
	0xb4, 0xc8, 0x29, 0xb8, 0xb5, 0x14, 0x64, 0xf7, 0x97, 0xe9, 0x72, 0x44, 0x7f, 0xdd, 0xce, 0x7d,
	0xc9, 0x73, 0x3d, 0x0f, 0x5a, 0xe7, 0xdd, 0x87, 0x0e, 0xcd, 0xd9, 0xd3, 0xaa, 0x79, 0x03, 0x47,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xa5, 0xea, 0x78, 0x58, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderRouterClient is the client API for OrderRouter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderRouterClient interface {
	CreateAndRouteOrder(ctx context.Context, in *CreateAndRouteOrderParams, opts ...grpc.CallOption) (*OrderId, error)
	CancelOrder(ctx context.Context, in *CancelOrderParams, opts ...grpc.CallOption) (*model.Empty, error)
}

type orderRouterClient struct {
	cc *grpc.ClientConn
}

func NewOrderRouterClient(cc *grpc.ClientConn) OrderRouterClient {
	return &orderRouterClient{cc}
}

func (c *orderRouterClient) CreateAndRouteOrder(ctx context.Context, in *CreateAndRouteOrderParams, opts ...grpc.CallOption) (*OrderId, error) {
	out := new(OrderId)
	err := c.cc.Invoke(ctx, "/executionvenue.OrderRouter/CreateAndRouteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRouterClient) CancelOrder(ctx context.Context, in *CancelOrderParams, opts ...grpc.CallOption) (*model.Empty, error) {
	out := new(model.Empty)
	err := c.cc.Invoke(ctx, "/executionvenue.OrderRouter/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderRouterServer is the server API for OrderRouter service.
type OrderRouterServer interface {
	CreateAndRouteOrder(context.Context, *CreateAndRouteOrderParams) (*OrderId, error)
	CancelOrder(context.Context, *CancelOrderParams) (*model.Empty, error)
}

// UnimplementedOrderRouterServer can be embedded to have forward compatible implementations.
type UnimplementedOrderRouterServer struct {
}

func (*UnimplementedOrderRouterServer) CreateAndRouteOrder(ctx context.Context, req *CreateAndRouteOrderParams) (*OrderId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAndRouteOrder not implemented")
}
func (*UnimplementedOrderRouterServer) CancelOrder(ctx context.Context, req *CancelOrderParams) (*model.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}

func RegisterOrderRouterServer(s *grpc.Server, srv OrderRouterServer) {
	s.RegisterService(&_OrderRouter_serviceDesc, srv)
}

func _OrderRouter_CreateAndRouteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndRouteOrderParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRouterServer).CreateAndRouteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executionvenue.OrderRouter/CreateAndRouteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRouterServer).CreateAndRouteOrder(ctx, req.(*CreateAndRouteOrderParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderRouter_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRouterServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/executionvenue.OrderRouter/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRouterServer).CancelOrder(ctx, req.(*CancelOrderParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderRouter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "executionvenue.OrderRouter",
	HandlerType: (*OrderRouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAndRouteOrder",
			Handler:    _OrderRouter_CreateAndRouteOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderRouter_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order-router.proto",
}
