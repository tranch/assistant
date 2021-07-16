// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: finance.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BillRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillRequest) Reset()         { *m = BillRequest{} }
func (m *BillRequest) String() string { return proto.CompactTextString(m) }
func (*BillRequest) ProtoMessage()    {}
func (*BillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{0}
}
func (m *BillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillRequest.Unmarshal(m, b)
}
func (m *BillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillRequest.Marshal(b, m, deterministic)
}
func (m *BillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillRequest.Merge(m, src)
}
func (m *BillRequest) XXX_Size() int {
	return xxx_messageInfo_BillRequest.Size(m)
}
func (m *BillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BillRequest proto.InternalMessageInfo

func (m *BillRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BillRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BillReply struct {
	Bill                 *Bill    `protobuf:"bytes,1,opt,name=bill,proto3" json:"bill,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillReply) Reset()         { *m = BillReply{} }
func (m *BillReply) String() string { return proto.CompactTextString(m) }
func (*BillReply) ProtoMessage()    {}
func (*BillReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{1}
}
func (m *BillReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillReply.Unmarshal(m, b)
}
func (m *BillReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillReply.Marshal(b, m, deterministic)
}
func (m *BillReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillReply.Merge(m, src)
}
func (m *BillReply) XXX_Size() int {
	return xxx_messageInfo_BillReply.Size(m)
}
func (m *BillReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BillReply.DiscardUnknown(m)
}

var xxx_messageInfo_BillReply proto.InternalMessageInfo

func (m *BillReply) GetBill() *Bill {
	if m != nil {
		return m.Bill
	}
	return nil
}

type BillsReply struct {
	Bills                []*Bill  `protobuf:"bytes,1,rep,name=bills,proto3" json:"bills,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillsReply) Reset()         { *m = BillsReply{} }
func (m *BillsReply) String() string { return proto.CompactTextString(m) }
func (*BillsReply) ProtoMessage()    {}
func (*BillsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{2}
}
func (m *BillsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillsReply.Unmarshal(m, b)
}
func (m *BillsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillsReply.Marshal(b, m, deterministic)
}
func (m *BillsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillsReply.Merge(m, src)
}
func (m *BillsReply) XXX_Size() int {
	return xxx_messageInfo_BillsReply.Size(m)
}
func (m *BillsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BillsReply.DiscardUnknown(m)
}

var xxx_messageInfo_BillsReply proto.InternalMessageInfo

func (m *BillsReply) GetBills() []*Bill {
	if m != nil {
		return m.Bills
	}
	return nil
}

type Bill struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Payee                string   `protobuf:"bytes,3,opt,name=payee,proto3" json:"payee,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Amount               float32  `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bill) Reset()         { *m = Bill{} }
func (m *Bill) String() string { return proto.CompactTextString(m) }
func (*Bill) ProtoMessage()    {}
func (*Bill) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{3}
}
func (m *Bill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bill.Unmarshal(m, b)
}
func (m *Bill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bill.Marshal(b, m, deterministic)
}
func (m *Bill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bill.Merge(m, src)
}
func (m *Bill) XXX_Size() int {
	return xxx_messageInfo_Bill.Size(m)
}
func (m *Bill) XXX_DiscardUnknown() {
	xxx_messageInfo_Bill.DiscardUnknown(m)
}

var xxx_messageInfo_Bill proto.InternalMessageInfo

func (m *Bill) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bill) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Bill) GetPayee() string {
	if m != nil {
		return m.Payee
	}
	return ""
}

func (m *Bill) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Bill) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Bill) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Bill) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type BillRecord struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BillId               int64    `protobuf:"varint,2,opt,name=billId,proto3" json:"billId,omitempty"`
	Posting              string   `protobuf:"bytes,3,opt,name=posting,proto3" json:"posting,omitempty"`
	Amount               float32  `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BillRecord) Reset()         { *m = BillRecord{} }
func (m *BillRecord) String() string { return proto.CompactTextString(m) }
func (*BillRecord) ProtoMessage()    {}
func (*BillRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{4}
}
func (m *BillRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillRecord.Unmarshal(m, b)
}
func (m *BillRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillRecord.Marshal(b, m, deterministic)
}
func (m *BillRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillRecord.Merge(m, src)
}
func (m *BillRecord) XXX_Size() int {
	return xxx_messageInfo_BillRecord.Size(m)
}
func (m *BillRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_BillRecord.DiscardUnknown(m)
}

var xxx_messageInfo_BillRecord proto.InternalMessageInfo

func (m *BillRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BillRecord) GetBillId() int64 {
	if m != nil {
		return m.BillId
	}
	return 0
}

func (m *BillRecord) GetPosting() string {
	if m != nil {
		return m.Posting
	}
	return ""
}

func (m *BillRecord) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Assets struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Balance              float32  `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assets) Reset()         { *m = Assets{} }
func (m *Assets) String() string { return proto.CompactTextString(m) }
func (*Assets) ProtoMessage()    {}
func (*Assets) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{5}
}
func (m *Assets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assets.Unmarshal(m, b)
}
func (m *Assets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assets.Marshal(b, m, deterministic)
}
func (m *Assets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assets.Merge(m, src)
}
func (m *Assets) XXX_Size() int {
	return xxx_messageInfo_Assets.Size(m)
}
func (m *Assets) XXX_DiscardUnknown() {
	xxx_messageInfo_Assets.DiscardUnknown(m)
}

var xxx_messageInfo_Assets proto.InternalMessageInfo

func (m *Assets) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Assets) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *Assets) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assets) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Assets) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Assets) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Account struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Balance              float32  `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{6}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*BillRequest)(nil), "pb.BillRequest")
	proto.RegisterType((*BillReply)(nil), "pb.BillReply")
	proto.RegisterType((*BillsReply)(nil), "pb.BillsReply")
	proto.RegisterType((*Bill)(nil), "pb.Bill")
	proto.RegisterType((*BillRecord)(nil), "pb.BillRecord")
	proto.RegisterType((*Assets)(nil), "pb.Assets")
	proto.RegisterType((*Account)(nil), "pb.Account")
}

func init() { proto.RegisterFile("finance.proto", fileDescriptor_c04e2e1c1ba79a81) }

var fileDescriptor_c04e2e1c1ba79a81 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x5d, 0x7e, 0x9a, 0xb6, 0x5f, 0xb5, 0x21, 0x59, 0x68, 0xb2, 0xaa, 0x0a, 0x45, 0xbe, 0x2a,
	0x02, 0x3a, 0x31, 0x9e, 0xa0, 0x03, 0x81, 0xb8, 0xcd, 0xee, 0xb8, 0x73, 0xec, 0x6f, 0x93, 0xa5,
	0x2c, 0x31, 0xb1, 0x8b, 0xd4, 0x97, 0xe1, 0x31, 0x78, 0x08, 0x9e, 0x0a, 0xf9, 0x67, 0x49, 0x60,
	0x63, 0xda, 0x9d, 0xcf, 0xf9, 0x8e, 0x73, 0x8e, 0xed, 0x13, 0x38, 0xbd, 0x51, 0x2d, 0x6f, 0x05,
	0xee, 0x74, 0xdf, 0xd9, 0x8e, 0xa4, 0xba, 0x5e, 0x43, 0xcd, 0x4d, 0xc4, 0xec, 0x3d, 0xac, 0xae,
	0x54, 0xd3, 0x54, 0xf8, 0xfd, 0x80, 0xc6, 0x92, 0x33, 0x48, 0x95, 0xa4, 0x49, 0x99, 0x6c, 0xb3,
	0x2a, 0x55, 0x92, 0x10, 0xc8, 0x5b, 0x7e, 0x87, 0x34, 0x2d, 0x93, 0xed, 0xb2, 0xf2, 0x6b, 0xf6,
	0x1a, 0x96, 0x61, 0x8b, 0x6e, 0x8e, 0x64, 0x03, 0x79, 0xad, 0x9a, 0xc6, 0x6f, 0x59, 0x5d, 0x2e,
	0x76, 0xba, 0xde, 0xf9, 0xa1, 0x67, 0xd9, 0x5b, 0x00, 0x87, 0x4c, 0xd0, 0xbe, 0x82, 0x99, 0x63,
	0x0d, 0x4d, 0xca, 0xec, 0x2f, 0x71, 0xa0, 0xd9, 0xaf, 0x04, 0x72, 0x87, 0x1f, 0x4b, 0x21, 0xb9,
	0x1d, 0x52, 0xb8, 0x35, 0x79, 0x09, 0x33, 0xcd, 0x8f, 0x88, 0x34, 0xf3, 0x64, 0x00, 0xa4, 0x84,
	0x95, 0x44, 0x23, 0x7a, 0xa5, 0xad, 0xea, 0x5a, 0x9a, 0xfb, 0xd9, 0x94, 0x22, 0xe7, 0x50, 0xf0,
	0xbb, 0xee, 0xd0, 0x5a, 0x3a, 0x2b, 0x93, 0x6d, 0x5a, 0x45, 0x44, 0x36, 0xb0, 0x14, 0x3d, 0x72,
	0x8b, 0x72, 0x6f, 0x69, 0xe1, 0xf7, 0x8d, 0x84, 0x9b, 0x1e, 0xb4, 0x8c, 0xd3, 0x79, 0x98, 0x0e,
	0x04, 0xbb, 0x09, 0xc7, 0xac, 0x50, 0x74, 0xbd, 0x7c, 0x90, 0xfe, 0x1c, 0x0a, 0x77, 0xbe, 0xaf,
	0xd2, 0xe7, 0xcf, 0xaa, 0x88, 0x08, 0x85, 0xb9, 0xee, 0x8c, 0x55, 0xed, 0x6d, 0x3c, 0xc3, 0x3d,
	0x9c, 0x64, 0xcc, 0xa7, 0x19, 0xd9, 0xcf, 0x04, 0x8a, 0xbd, 0x31, 0x68, 0xcd, 0x03, 0x93, 0x0d,
	0x2c, 0xb9, 0x10, 0x4e, 0x35, 0xf8, 0x8c, 0xc4, 0xf0, 0x8c, 0xd9, 0xf8, 0x8c, 0x64, 0x0d, 0x0b,
	0xc1, 0x2d, 0xde, 0x76, 0xfd, 0x31, 0xde, 0xd3, 0x80, 0x5d, 0xb4, 0x9a, 0x37, 0xae, 0x36, 0xf1,
	0x96, 0xee, 0xe1, 0xd3, 0xd7, 0xc4, 0x10, 0xe6, 0xfb, 0x60, 0xfa, 0x9c, 0x26, 0x4d, 0x6d, 0xb2,
	0x27, 0x6c, 0xf2, 0x7f, 0x6c, 0x2e, 0x7f, 0x27, 0x00, 0x9f, 0x43, 0xad, 0xaf, 0x7f, 0x08, 0x72,
	0x01, 0xf0, 0xd1, 0xcf, 0x7c, 0x79, 0x5e, 0x0c, 0xb5, 0x0a, 0x9d, 0x5e, 0x9f, 0x39, 0xe2, 0xda,
	0x72, 0x8b, 0xbe, 0x86, 0xec, 0x84, 0xbc, 0x81, 0xf9, 0x17, 0xb4, 0x8f, 0xab, 0x4f, 0x47, 0x22,
	0x88, 0xdf, 0xc1, 0x22, 0x8a, 0xcd, 0x7f, 0xbe, 0x3d, 0x56, 0x9c, 0x9d, 0xb8, 0x30, 0x9f, 0xb0,
	0xc1, 0x67, 0x87, 0xb9, 0x5a, 0x7c, 0x2b, 0xb8, 0x56, 0x17, 0xba, 0xae, 0x0b, 0xff, 0x4b, 0x7e,
	0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x35, 0x89, 0xd9, 0xb3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinanceSvcClient is the client API for FinanceSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinanceSvcClient interface {
	CreateBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*StateReply, error)
	GetBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillReply, error)
	GetBills(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillsReply, error)
	DeleteBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*StateReply, error)
}

type financeSvcClient struct {
	cc *grpc.ClientConn
}

func NewFinanceSvcClient(cc *grpc.ClientConn) FinanceSvcClient {
	return &financeSvcClient{cc}
}

func (c *financeSvcClient) CreateBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.FinanceSvc/CreateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeSvcClient) GetBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillReply, error) {
	out := new(BillReply)
	err := c.cc.Invoke(ctx, "/pb.FinanceSvc/GetBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeSvcClient) GetBills(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillsReply, error) {
	out := new(BillsReply)
	err := c.cc.Invoke(ctx, "/pb.FinanceSvc/GetBills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeSvcClient) DeleteBill(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.FinanceSvc/DeleteBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceSvcServer is the server API for FinanceSvc service.
type FinanceSvcServer interface {
	CreateBill(context.Context, *BillRequest) (*StateReply, error)
	GetBill(context.Context, *BillRequest) (*BillReply, error)
	GetBills(context.Context, *BillRequest) (*BillsReply, error)
	DeleteBill(context.Context, *BillRequest) (*StateReply, error)
}

// UnimplementedFinanceSvcServer can be embedded to have forward compatible implementations.
type UnimplementedFinanceSvcServer struct {
}

func (*UnimplementedFinanceSvcServer) CreateBill(ctx context.Context, req *BillRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (*UnimplementedFinanceSvcServer) GetBill(ctx context.Context, req *BillRequest) (*BillReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBill not implemented")
}
func (*UnimplementedFinanceSvcServer) GetBills(ctx context.Context, req *BillRequest) (*BillsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBills not implemented")
}
func (*UnimplementedFinanceSvcServer) DeleteBill(ctx context.Context, req *BillRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBill not implemented")
}

func RegisterFinanceSvcServer(s *grpc.Server, srv FinanceSvcServer) {
	s.RegisterService(&_FinanceSvc_serviceDesc, srv)
}

func _FinanceSvc_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceSvcServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FinanceSvc/CreateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceSvcServer).CreateBill(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceSvc_GetBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceSvcServer).GetBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FinanceSvc/GetBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceSvcServer).GetBill(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceSvc_GetBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceSvcServer).GetBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FinanceSvc/GetBills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceSvcServer).GetBills(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FinanceSvc_DeleteBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceSvcServer).DeleteBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FinanceSvc/DeleteBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceSvcServer).DeleteBill(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinanceSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FinanceSvc",
	HandlerType: (*FinanceSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBill",
			Handler:    _FinanceSvc_CreateBill_Handler,
		},
		{
			MethodName: "GetBill",
			Handler:    _FinanceSvc_GetBill_Handler,
		},
		{
			MethodName: "GetBills",
			Handler:    _FinanceSvc_GetBills_Handler,
		},
		{
			MethodName: "DeleteBill",
			Handler:    _FinanceSvc_DeleteBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "finance.proto",
}
