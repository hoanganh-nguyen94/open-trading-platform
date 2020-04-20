// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Side int32

const (
	Side_BUY  Side = 0
	Side_SELL Side = 1
)

var Side_name = map[int32]string{
	0: "BUY",
	1: "SELL",
}

var Side_value = map[string]int32{
	"BUY":  0,
	"SELL": 1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}

func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

type OrderStatus int32

const (
	OrderStatus_NONE      OrderStatus = 0
	OrderStatus_LIVE      OrderStatus = 1
	OrderStatus_FILLED    OrderStatus = 2
	OrderStatus_CANCELLED OrderStatus = 3
)

var OrderStatus_name = map[int32]string{
	0: "NONE",
	1: "LIVE",
	2: "FILLED",
	3: "CANCELLED",
}

var OrderStatus_value = map[string]int32{
	"NONE":      0,
	"LIVE":      1,
	"FILLED":    2,
	"CANCELLED": 3,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}

func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

type Order struct {
	Version              int32       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id                   string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Side                 Side        `protobuf:"varint,3,opt,name=side,proto3,enum=model.Side" json:"side,omitempty"`
	Quantity             *Decimal64  `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                *Decimal64  `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	ListingId            int32       `protobuf:"varint,6,opt,name=listingId,proto3" json:"listingId,omitempty"`
	RemainingQuantity    *Decimal64  `protobuf:"bytes,7,opt,name=remainingQuantity,proto3" json:"remainingQuantity,omitempty"`
	TradedQuantity       *Decimal64  `protobuf:"bytes,8,opt,name=tradedQuantity,proto3" json:"tradedQuantity,omitempty"`
	AvgTradePrice        *Decimal64  `protobuf:"bytes,9,opt,name=avgTradePrice,proto3" json:"avgTradePrice,omitempty"`
	Status               OrderStatus `protobuf:"varint,10,opt,name=status,proto3,enum=model.OrderStatus" json:"status,omitempty"`
	TargetStatus         OrderStatus `protobuf:"varint,11,opt,name=targetStatus,proto3,enum=model.OrderStatus" json:"targetStatus,omitempty"`
	Created              *Timestamp  `protobuf:"bytes,12,opt,name=created,proto3" json:"created,omitempty"`
	OwnerId              string      `protobuf:"bytes,13,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	OriginatorId         string      `protobuf:"bytes,14,opt,name=originatorId,proto3" json:"originatorId,omitempty"`
	OriginatorRef        string      `protobuf:"bytes,15,opt,name=originatorRef,proto3" json:"originatorRef,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_BUY
}

func (m *Order) GetQuantity() *Decimal64 {
	if m != nil {
		return m.Quantity
	}
	return nil
}

func (m *Order) GetPrice() *Decimal64 {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Order) GetListingId() int32 {
	if m != nil {
		return m.ListingId
	}
	return 0
}

func (m *Order) GetRemainingQuantity() *Decimal64 {
	if m != nil {
		return m.RemainingQuantity
	}
	return nil
}

func (m *Order) GetTradedQuantity() *Decimal64 {
	if m != nil {
		return m.TradedQuantity
	}
	return nil
}

func (m *Order) GetAvgTradePrice() *Decimal64 {
	if m != nil {
		return m.AvgTradePrice
	}
	return nil
}

func (m *Order) GetStatus() OrderStatus {
	if m != nil {
		return m.Status
	}
	return OrderStatus_NONE
}

func (m *Order) GetTargetStatus() OrderStatus {
	if m != nil {
		return m.TargetStatus
	}
	return OrderStatus_NONE
}

func (m *Order) GetCreated() *Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Order) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Order) GetOriginatorId() string {
	if m != nil {
		return m.OriginatorId
	}
	return ""
}

func (m *Order) GetOriginatorRef() string {
	if m != nil {
		return m.OriginatorRef
	}
	return ""
}

func init() {
	proto.RegisterEnum("model.Side", Side_name, Side_value)
	proto.RegisterEnum("model.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterType((*Order)(nil), "model.Order")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x14, 0xc5, 0x4d, 0x9b, 0x34, 0xed, 0x4d, 0x5b, 0xb3, 0xf7, 0x69, 0x14, 0xc1, 0xb2, 0x88, 0x84,
	0x20, 0x7d, 0x58, 0xa5, 0xf8, 0x20, 0x82, 0xbb, 0x1b, 0x21, 0x10, 0xba, 0x9a, 0xae, 0x82, 0xbe,
	0x8d, 0x99, 0x31, 0x0c, 0x34, 0x33, 0x75, 0x32, 0xbb, 0xe2, 0x77, 0xf6, 0x43, 0x48, 0x26, 0xcd,
	0xae, 0x51, 0xf2, 0x14, 0xee, 0x39, 0xbf, 0x33, 0x99, 0x3f, 0x07, 0x02, 0xa5, 0x19, 0xd7, 0xeb,
	0x83, 0x56, 0x46, 0xa1, 0x57, 0x29, 0xc6, 0xf7, 0x8f, 0x4f, 0xec, 0xa7, 0x50, 0x55, 0xa5, 0x64,
	0xeb, 0x9c, 0xfe, 0x76, 0xc1, 0xbb, 0x6a, 0x48, 0x24, 0xe0, 0xdf, 0x72, 0x5d, 0x0b, 0x25, 0x89,
	0xb3, 0x72, 0x22, 0x2f, 0xef, 0x46, 0x5c, 0xc2, 0x48, 0x30, 0x32, 0x5a, 0x39, 0xd1, 0x2c, 0x1f,
	0x09, 0x86, 0x4f, 0xc1, 0xad, 0x05, 0xe3, 0x64, 0xbc, 0x72, 0xa2, 0xe5, 0x59, 0xb0, 0xb6, 0xab,
	0xae, 0x77, 0x82, 0xf1, 0xdc, 0x1a, 0xf8, 0x02, 0xa6, 0x3f, 0x6e, 0xa8, 0x34, 0xc2, 0xfc, 0x22,
	0xee, 0xca, 0x89, 0x82, 0xb3, 0xf0, 0x08, 0x5d, 0xf2, 0x42, 0x54, 0x74, 0xbf, 0x79, 0x95, 0xdf,
	0x11, 0xf8, 0x1c, 0xbc, 0x83, 0x16, 0x05, 0x27, 0xde, 0x00, 0xda, 0xda, 0xf8, 0x04, 0x66, 0x7b,
	0x51, 0x1b, 0x21, 0xcb, 0x94, 0x91, 0x89, 0xdd, 0xe2, 0xbd, 0x80, 0x6f, 0xe1, 0x44, 0xf3, 0x8a,
	0x0a, 0x29, 0x64, 0xf9, 0xb1, 0xfb, 0xb9, 0x3f, 0xb0, 0xe2, 0xff, 0x28, 0xbe, 0x86, 0xa5, 0xd1,
	0x94, 0x71, 0x76, 0x17, 0x9e, 0x0e, 0x84, 0xff, 0xe1, 0x70, 0x03, 0x0b, 0x7a, 0x5b, 0x5e, 0x37,
	0xe2, 0x07, 0x7b, 0x8e, 0xd9, 0x40, 0xb0, 0x8f, 0x61, 0x0c, 0x93, 0xda, 0x50, 0x73, 0x53, 0x13,
	0xb0, 0x17, 0x89, 0xc7, 0x80, 0x7d, 0x8e, 0x9d, 0x75, 0xf2, 0x23, 0x81, 0x1b, 0x98, 0x1b, 0xaa,
	0x4b, 0x6e, 0x5a, 0x9d, 0x04, 0x83, 0x89, 0x1e, 0x87, 0x31, 0xf8, 0x85, 0xe6, 0xd4, 0x70, 0x46,
	0xe6, 0xbd, 0x5d, 0x5d, 0x8b, 0x8a, 0xd7, 0x86, 0x56, 0x87, 0xbc, 0x03, 0x9a, 0x02, 0xa8, 0x9f,
	0x92, 0xeb, 0x94, 0x91, 0x85, 0x7d, 0xeb, 0x6e, 0xc4, 0x53, 0x98, 0x2b, 0x2d, 0x4a, 0x21, 0xa9,
	0x51, 0x8d, 0xbd, 0xb4, 0x76, 0x4f, 0xc3, 0x67, 0xb0, 0xb8, 0x9f, 0x73, 0xfe, 0x9d, 0x3c, 0xb4,
	0x50, 0x5f, 0x8c, 0x1f, 0x81, 0xdb, 0xf4, 0x04, 0x7d, 0x18, 0x9f, 0x7f, 0xfa, 0x12, 0x3e, 0xc0,
	0x29, 0xb8, 0xbb, 0x24, 0xcb, 0x42, 0x27, 0x7e, 0x03, 0xc1, 0x5f, 0xe7, 0x68, 0x8c, 0xed, 0xd5,
	0x36, 0x69, 0x91, 0x2c, 0xfd, 0x9c, 0x84, 0x0e, 0x02, 0x4c, 0xde, 0xa7, 0x59, 0x96, 0x5c, 0x86,
	0x23, 0x5c, 0xc0, 0xec, 0xe2, 0xdd, 0xf6, 0x22, 0xb1, 0xe3, 0xf8, 0xdc, 0xff, 0xda, 0x76, 0xfc,
	0xdb, 0xc4, 0xf6, 0xfa, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xfd, 0xd0, 0x44, 0x00,
	0x03, 0x00, 0x00,
}
