// Code generated by protoc-gen-go. DO NOT EDIT.
// source: viewservice.proto

package model

import (
	context "context"
	fmt "fmt"
	"github.com/ettec/otp-common/model"
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

type SubscribeToOrdersWithRootOriginatorIdArgs struct {
	RootOriginatorId     string   `protobuf:"bytes,2,opt,name=rootOriginatorId,proto3" json:"rootOriginatorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToOrdersWithRootOriginatorIdArgs) Reset() {
	*m = SubscribeToOrdersWithRootOriginatorIdArgs{}
}
func (m *SubscribeToOrdersWithRootOriginatorIdArgs) String() string { return proto.CompactTextString(m) }
func (*SubscribeToOrdersWithRootOriginatorIdArgs) ProtoMessage()    {}
func (*SubscribeToOrdersWithRootOriginatorIdArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec42da24455bcc25, []int{0}
}

func (m *SubscribeToOrdersWithRootOriginatorIdArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs.Unmarshal(m, b)
}
func (m *SubscribeToOrdersWithRootOriginatorIdArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs.Marshal(b, m, deterministic)
}
func (m *SubscribeToOrdersWithRootOriginatorIdArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs.Merge(m, src)
}
func (m *SubscribeToOrdersWithRootOriginatorIdArgs) XXX_Size() int {
	return xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs.Size(m)
}
func (m *SubscribeToOrdersWithRootOriginatorIdArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToOrdersWithRootOriginatorIdArgs proto.InternalMessageInfo

func (m *SubscribeToOrdersWithRootOriginatorIdArgs) GetRootOriginatorId() string {
	if m != nil {
		return m.RootOriginatorId
	}
	return ""
}

type GetOrderHistoryArgs struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ToVersion            int32    `protobuf:"varint,2,opt,name=toVersion,proto3" json:"toVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderHistoryArgs) Reset()         { *m = GetOrderHistoryArgs{} }
func (m *GetOrderHistoryArgs) String() string { return proto.CompactTextString(m) }
func (*GetOrderHistoryArgs) ProtoMessage()    {}
func (*GetOrderHistoryArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec42da24455bcc25, []int{1}
}

func (m *GetOrderHistoryArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderHistoryArgs.Unmarshal(m, b)
}
func (m *GetOrderHistoryArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderHistoryArgs.Marshal(b, m, deterministic)
}
func (m *GetOrderHistoryArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderHistoryArgs.Merge(m, src)
}
func (m *GetOrderHistoryArgs) XXX_Size() int {
	return xxx_messageInfo_GetOrderHistoryArgs.Size(m)
}
func (m *GetOrderHistoryArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderHistoryArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderHistoryArgs proto.InternalMessageInfo

func (m *GetOrderHistoryArgs) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *GetOrderHistoryArgs) GetToVersion() int32 {
	if m != nil {
		return m.ToVersion
	}
	return 0
}

type OrderUpdate struct {
	Order                *model.Order     `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Time                 *model.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OrderUpdate) Reset()         { *m = OrderUpdate{} }
func (m *OrderUpdate) String() string { return proto.CompactTextString(m) }
func (*OrderUpdate) ProtoMessage()    {}
func (*OrderUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec42da24455bcc25, []int{2}
}

func (m *OrderUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderUpdate.Unmarshal(m, b)
}
func (m *OrderUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderUpdate.Marshal(b, m, deterministic)
}
func (m *OrderUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderUpdate.Merge(m, src)
}
func (m *OrderUpdate) XXX_Size() int {
	return xxx_messageInfo_OrderUpdate.Size(m)
}
func (m *OrderUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_OrderUpdate proto.InternalMessageInfo

func (m *OrderUpdate) GetOrder() *model.Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OrderUpdate) GetTime() *model.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type OrderHistory struct {
	Updates              []*OrderUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OrderHistory) Reset()         { *m = OrderHistory{} }
func (m *OrderHistory) String() string { return proto.CompactTextString(m) }
func (*OrderHistory) ProtoMessage()    {}
func (*OrderHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec42da24455bcc25, []int{3}
}

func (m *OrderHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderHistory.Unmarshal(m, b)
}
func (m *OrderHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderHistory.Marshal(b, m, deterministic)
}
func (m *OrderHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderHistory.Merge(m, src)
}
func (m *OrderHistory) XXX_Size() int {
	return xxx_messageInfo_OrderHistory.Size(m)
}
func (m *OrderHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderHistory.DiscardUnknown(m)
}

var xxx_messageInfo_OrderHistory proto.InternalMessageInfo

func (m *OrderHistory) GetUpdates() []*OrderUpdate {
	if m != nil {
		return m.Updates
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeToOrdersWithRootOriginatorIdArgs)(nil), "viewservice.SubscribeToOrdersWithRootOriginatorIdArgs")
	proto.RegisterType((*GetOrderHistoryArgs)(nil), "viewservice.GetOrderHistoryArgs")
	proto.RegisterType((*OrderUpdate)(nil), "viewservice.OrderUpdate")
	proto.RegisterType((*OrderHistory)(nil), "viewservice.OrderHistory")
}

func init() { proto.RegisterFile("viewservice.proto", fileDescriptor_ec42da24455bcc25) }

var fileDescriptor_ec42da24455bcc25 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0x3a, 0x31,
	0x10, 0xc5, 0xd9, 0xff, 0x5f, 0x24, 0x4c, 0x49, 0x84, 0x7a, 0x59, 0x37, 0x1e, 0x48, 0xa3, 0x09,
	0x7a, 0xd8, 0x98, 0x35, 0xf1, 0x2e, 0x17, 0xe5, 0x60, 0x30, 0x0b, 0x42, 0xe2, 0x0d, 0xd8, 0x09,
	0x36, 0xb1, 0x0c, 0xb6, 0x05, 0xe2, 0xd7, 0xf4, 0x13, 0x19, 0x67, 0x21, 0x2e, 0xe0, 0x81, 0x63,
	0xdf, 0x7b, 0xfd, 0x4d, 0x67, 0xa6, 0xd0, 0x58, 0x6a, 0x5c, 0x39, 0xb4, 0x4b, 0x3d, 0xc1, 0x78,
	0x6e, 0xc9, 0x93, 0x14, 0x05, 0x29, 0x6a, 0x18, 0xca, 0xf0, 0x7d, 0x42, 0xc6, 0xd0, 0x2c, 0xf7,
	0x23, 0x41, 0x36, 0x43, 0x9b, 0x1f, 0xd4, 0x10, 0xae, 0x7a, 0x8b, 0xb1, 0x9b, 0x58, 0x3d, 0xc6,
	0x3e, 0x75, 0x7f, 0x1c, 0x37, 0xd4, 0xfe, 0x2d, 0x25, 0xf2, 0x5d, 0xab, 0xa7, 0x7a, 0x36, 0xf2,
	0x64, 0x3b, 0xd9, 0xbd, 0x9d, 0x3a, 0x79, 0x0d, 0x75, 0xbb, 0xa3, 0x87, 0xff, 0x9a, 0x41, 0xab,
	0x9a, 0xee, 0xe9, 0xea, 0x09, 0x4e, 0x1f, 0xd0, 0x33, 0xf0, 0x51, 0x3b, 0x4f, 0xf6, 0x93, 0x11,
	0x21, 0x54, 0xb8, 0x7c, 0x27, 0x0b, 0x03, 0xbe, 0xb9, 0x39, 0xca, 0x73, 0xa8, 0x7a, 0x1a, 0xa0,
	0x75, 0x9a, 0x66, 0x4c, 0x2d, 0xa7, 0xbf, 0x82, 0x1a, 0x82, 0x60, 0xd6, 0xcb, 0x3c, 0x1b, 0x79,
	0x94, 0x0a, 0xca, 0x7c, 0x8f, 0x21, 0x22, 0xa9, 0xc5, 0xdc, 0x66, 0xcc, 0x91, 0x34, 0xb7, 0xe4,
	0x05, 0x1c, 0x79, 0x6d, 0x90, 0x59, 0x22, 0xa9, 0xaf, 0x23, 0x7d, 0x6d, 0xd0, 0xf9, 0x91, 0x99,
	0xa7, 0xec, 0xaa, 0x36, 0xd4, 0x8a, 0x8f, 0x94, 0x09, 0x54, 0x16, 0x5c, 0xc3, 0x85, 0x41, 0xf3,
	0x7f, 0x4b, 0x24, 0x61, 0x5c, 0x1c, 0x71, 0xe1, 0x11, 0xe9, 0x26, 0x98, 0x7c, 0x05, 0x20, 0x06,
	0x1a, 0x57, 0xbd, 0x3c, 0x24, 0x3f, 0xe0, 0xf2, 0xa0, 0xa1, 0xca, 0xbb, 0x2d, 0xf6, 0xc1, 0x8b,
	0x88, 0xb6, 0xfa, 0x55, 0xa5, 0x9b, 0x40, 0x3e, 0xc3, 0xc9, 0xce, 0xb8, 0x65, 0x73, 0x0b, 0xfe,
	0xc7, 0x32, 0xa2, 0xb3, 0xfd, 0xd6, 0xd6, 0xb6, 0x2a, 0xb5, 0x2b, 0xaf, 0x65, 0x2e, 0x32, 0x3e,
	0xe6, 0x9f, 0x72, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xa9, 0x08, 0x40, 0x6b, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ViewServiceClient is the client API for ViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ViewServiceClient interface {
	SubscribeToOrdersWithRootOriginatorId(ctx context.Context, in *SubscribeToOrdersWithRootOriginatorIdArgs, opts ...grpc.CallOption) (ViewService_SubscribeToOrdersWithRootOriginatorIdClient, error)
	GetOrderHistory(ctx context.Context, in *GetOrderHistoryArgs, opts ...grpc.CallOption) (*OrderHistory, error)
}

type viewServiceClient struct {
	cc *grpc.ClientConn
}

func NewViewServiceClient(cc *grpc.ClientConn) ViewServiceClient {
	return &viewServiceClient{cc}
}

func (c *viewServiceClient) SubscribeToOrdersWithRootOriginatorId(ctx context.Context, in *SubscribeToOrdersWithRootOriginatorIdArgs, opts ...grpc.CallOption) (ViewService_SubscribeToOrdersWithRootOriginatorIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ViewService_serviceDesc.Streams[0], "/viewservice.ViewService/SubscribeToOrdersWithRootOriginatorId", opts...)
	if err != nil {
		return nil, err
	}
	x := &viewServiceSubscribeToOrdersWithRootOriginatorIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ViewService_SubscribeToOrdersWithRootOriginatorIdClient interface {
	Recv() (*model.Order, error)
	grpc.ClientStream
}

type viewServiceSubscribeToOrdersWithRootOriginatorIdClient struct {
	grpc.ClientStream
}

func (x *viewServiceSubscribeToOrdersWithRootOriginatorIdClient) Recv() (*model.Order, error) {
	m := new(model.Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *viewServiceClient) GetOrderHistory(ctx context.Context, in *GetOrderHistoryArgs, opts ...grpc.CallOption) (*OrderHistory, error) {
	out := new(OrderHistory)
	err := c.cc.Invoke(ctx, "/viewservice.ViewService/GetOrderHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewServiceServer is the server API for ViewService service.
type ViewServiceServer interface {
	SubscribeToOrdersWithRootOriginatorId(*SubscribeToOrdersWithRootOriginatorIdArgs, ViewService_SubscribeToOrdersWithRootOriginatorIdServer) error
	GetOrderHistory(context.Context, *GetOrderHistoryArgs) (*OrderHistory, error)
}

// UnimplementedViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedViewServiceServer struct {
}

func (*UnimplementedViewServiceServer) SubscribeToOrdersWithRootOriginatorId(req *SubscribeToOrdersWithRootOriginatorIdArgs, srv ViewService_SubscribeToOrdersWithRootOriginatorIdServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToOrdersWithRootOriginatorId not implemented")
}
func (*UnimplementedViewServiceServer) GetOrderHistory(ctx context.Context, req *GetOrderHistoryArgs) (*OrderHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderHistory not implemented")
}

func RegisterViewServiceServer(s *grpc.Server, srv ViewServiceServer) {
	s.RegisterService(&_ViewService_serviceDesc, srv)
}

func _ViewService_SubscribeToOrdersWithRootOriginatorId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToOrdersWithRootOriginatorIdArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ViewServiceServer).SubscribeToOrdersWithRootOriginatorId(m, &viewServiceSubscribeToOrdersWithRootOriginatorIdServer{stream})
}

type ViewService_SubscribeToOrdersWithRootOriginatorIdServer interface {
	Send(*model.Order) error
	grpc.ServerStream
}

type viewServiceSubscribeToOrdersWithRootOriginatorIdServer struct {
	grpc.ServerStream
}

func (x *viewServiceSubscribeToOrdersWithRootOriginatorIdServer) Send(m *model.Order) error {
	return x.ServerStream.SendMsg(m)
}

func _ViewService_GetOrderHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderHistoryArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetOrderHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viewservice.ViewService/GetOrderHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetOrderHistory(ctx, req.(*GetOrderHistoryArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _ViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "viewservice.ViewService",
	HandlerType: (*ViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderHistory",
			Handler:    _ViewService_GetOrderHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToOrdersWithRootOriginatorId",
			Handler:       _ViewService_SubscribeToOrdersWithRootOriginatorId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "viewservice.proto",
}
