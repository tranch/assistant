// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: middle.proto

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

type PageRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{0}
}
func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PageRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type PageReply struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageReply) Reset()         { *m = PageReply{} }
func (m *PageReply) String() string { return proto.CompactTextString(m) }
func (*PageReply) ProtoMessage()    {}
func (*PageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{1}
}
func (m *PageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageReply.Unmarshal(m, b)
}
func (m *PageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageReply.Marshal(b, m, deterministic)
}
func (m *PageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageReply.Merge(m, src)
}
func (m *PageReply) XXX_Size() int {
	return xxx_messageInfo_PageReply.Size(m)
}
func (m *PageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PageReply.DiscardUnknown(m)
}

var xxx_messageInfo_PageReply proto.InternalMessageInfo

func (m *PageReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PageReply) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PageReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type AppReply struct {
	Apps                 []*App   `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReply) Reset()         { *m = AppReply{} }
func (m *AppReply) String() string { return proto.CompactTextString(m) }
func (*AppReply) ProtoMessage()    {}
func (*AppReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{2}
}
func (m *AppReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppReply.Unmarshal(m, b)
}
func (m *AppReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppReply.Marshal(b, m, deterministic)
}
func (m *AppReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReply.Merge(m, src)
}
func (m *AppReply) XXX_Size() int {
	return xxx_messageInfo_AppReply.Size(m)
}
func (m *AppReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppReply proto.InternalMessageInfo

func (m *AppReply) GetApps() []*App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type App struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	IsAuthorized         bool     `protobuf:"varint,2,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{3}
}
func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *App) GetIsAuthorized() bool {
	if m != nil {
		return m.IsAuthorized
	}
	return false
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{4}
}
func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SettingReply struct {
	Items                []*KV    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingReply) Reset()         { *m = SettingReply{} }
func (m *SettingReply) String() string { return proto.CompactTextString(m) }
func (*SettingReply) ProtoMessage()    {}
func (*SettingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{5}
}
func (m *SettingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingReply.Unmarshal(m, b)
}
func (m *SettingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingReply.Marshal(b, m, deterministic)
}
func (m *SettingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingReply.Merge(m, src)
}
func (m *SettingReply) XXX_Size() int {
	return xxx_messageInfo_SettingReply.Size(m)
}
func (m *SettingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingReply.DiscardUnknown(m)
}

var xxx_messageInfo_SettingReply proto.InternalMessageInfo

func (m *SettingReply) GetItems() []*KV {
	if m != nil {
		return m.Items
	}
	return nil
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_492fc6d32fb115aa, []int{6}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "pb.PageRequest")
	proto.RegisterType((*PageReply)(nil), "pb.PageReply")
	proto.RegisterType((*AppReply)(nil), "pb.AppReply")
	proto.RegisterType((*App)(nil), "pb.App")
	proto.RegisterType((*Text)(nil), "pb.Text")
	proto.RegisterType((*SettingReply)(nil), "pb.SettingReply")
	proto.RegisterType((*KV)(nil), "pb.KV")
}

func init() { proto.RegisterFile("middle.proto", fileDescriptor_492fc6d32fb115aa) }

var fileDescriptor_492fc6d32fb115aa = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0x3f, 0xb5, 0xdd, 0x89, 0x03, 0xd1, 0x8a, 0x83, 0x15, 0x38, 0x44, 0x8b, 0xa0,
	0x41, 0x54, 0x41, 0x2a, 0x0f, 0x80, 0x4c, 0x0f, 0x3d, 0x54, 0xa8, 0xd4, 0x45, 0x3d, 0x70, 0xb3,
	0xf1, 0xa8, 0xac, 0x70, 0xed, 0xc1, 0x1e, 0xa3, 0x84, 0x47, 0xe0, 0xa9, 0xd1, 0xae, 0xed, 0xc4,
	0x91, 0x88, 0xb8, 0xf4, 0x36, 0xf3, 0xed, 0xb7, 0x3f, 0x7f, 0x9e, 0xd1, 0x42, 0xf8, 0xa0, 0xf2,
	0xbc, 0xc0, 0x15, 0xd5, 0x15, 0x57, 0xc2, 0xa6, 0x4c, 0xde, 0xc0, 0xe4, 0x73, 0x7a, 0x8f, 0x09,
	0xfe, 0x6c, 0xb1, 0x61, 0x21, 0xc0, 0x6d, 0x5b, 0x95, 0x47, 0xd6, 0xc2, 0x5a, 0x9e, 0x24, 0xa6,
	0x16, 0xcf, 0xe0, 0x98, 0x15, 0x17, 0x18, 0xd9, 0x46, 0xec, 0x1a, 0x11, 0x81, 0xff, 0xad, 0x2a,
	0x19, 0x4b, 0x8e, 0x1c, 0xa3, 0x0f, 0xad, 0xbc, 0x86, 0x93, 0x0e, 0x49, 0xc5, 0xe6, 0x51, 0x80,
	0xa7, 0x10, 0xc4, 0x44, 0x1d, 0xef, 0x39, 0xb8, 0x29, 0x51, 0x13, 0x59, 0x0b, 0x67, 0x39, 0x39,
	0xf7, 0x57, 0x94, 0xad, 0xf4, 0x99, 0x11, 0xe5, 0x07, 0x70, 0x62, 0xa2, 0x1d, 0xdf, 0x1a, 0xf3,
	0x25, 0x84, 0xaa, 0x89, 0x5b, 0xfe, 0x5e, 0xd5, 0xea, 0x37, 0xe6, 0xe6, 0xe3, 0x41, 0xb2, 0xa7,
	0xc9, 0x39, 0xb8, 0x5f, 0x70, 0x6d, 0xc6, 0xc0, 0xb8, 0xe6, 0x21, 0xb5, 0xae, 0xe5, 0x19, 0x84,
	0xb7, 0xc8, 0xac, 0xca, 0xfb, 0x2e, 0xc9, 0x0b, 0x38, 0x56, 0x8c, 0x0f, 0x43, 0x14, 0x4f, 0x47,
	0xb9, 0xba, 0x4b, 0x3a, 0x51, 0x9e, 0x81, 0x7d, 0x75, 0x27, 0x66, 0xe0, 0xfc, 0xc0, 0x4d, 0x8f,
	0xd1, 0xa5, 0xce, 0xf6, 0x2b, 0x2d, 0xda, 0xed, 0xbf, 0x9b, 0xe6, 0xfc, 0x8f, 0x03, 0xde, 0x27,
	0xb3, 0x1a, 0xf1, 0x06, 0xe0, 0xa2, 0xc6, 0x94, 0x51, 0xcf, 0x50, 0x3c, 0xd5, 0xd4, 0xd1, 0x82,
	0xe6, 0x81, 0x16, 0x74, 0x46, 0x79, 0x24, 0xde, 0x82, 0x7f, 0x89, 0xfc, 0x6f, 0xdf, 0x74, 0x27,
	0x50, 0xb1, 0x91, 0x47, 0x62, 0x0e, 0xf6, 0x4d, 0x2d, 0xb6, 0xd7, 0xf7, 0x40, 0x0b, 0x70, 0x63,
	0xa2, 0x66, 0x74, 0x1a, 0x0e, 0x83, 0xed, 0x6f, 0xbf, 0x82, 0xe9, 0x2d, 0x57, 0x35, 0xc6, 0x44,
	0xd7, 0x7a, 0x5e, 0x07, 0x40, 0x2f, 0x61, 0x72, 0x51, 0x63, 0x8e, 0x25, 0xab, 0xb4, 0x68, 0x0e,
	0x98, 0x5e, 0xc3, 0x93, 0x4b, 0xe4, 0xff, 0xfb, 0x96, 0x30, 0xeb, 0x26, 0xb1, 0xb3, 0x1e, 0x70,
	0x9e, 0x82, 0xdf, 0xaf, 0x66, 0x64, 0x98, 0xe9, 0x6a, 0xbc, 0x31, 0x93, 0x6f, 0xda, 0x21, 0x07,
	0x7b, 0xbf, 0xb5, 0x31, 0xed, 0x63, 0xf0, 0xd5, 0x4b, 0x49, 0xbd, 0xa3, 0x2c, 0xf3, 0xcc, 0x3b,
	0x79, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x01, 0x2d, 0x75, 0xf0, 0x37, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MiddleClient is the client API for Middle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MiddleClient interface {
	CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*Text, error)
	GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageReply, error)
	Qr(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	Apps(ctx context.Context, in *Text, opts ...grpc.CallOption) (*AppReply, error)
	StoreAppOAuth(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	Credentials(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	GetCredentials(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	CreateCredential(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error)
	Setting(ctx context.Context, in *Text, opts ...grpc.CallOption) (*SettingReply, error)
	CreateSetting(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Text, error)
}

type middleClient struct {
	cc *grpc.ClientConn
}

func NewMiddleClient(cc *grpc.ClientConn) MiddleClient {
	return &middleClient{cc}
}

func (c *middleClient) CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreatePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PageReply, error) {
	out := new(PageReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Qr(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/Qr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Apps(ctx context.Context, in *Text, opts ...grpc.CallOption) (*AppReply, error) {
	out := new(AppReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/Apps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) StoreAppOAuth(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/StoreAppOAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Credentials(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/Credentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) GetCredentials(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/GetCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) CreateCredential(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) Setting(ctx context.Context, in *Text, opts ...grpc.CallOption) (*SettingReply, error) {
	out := new(SettingReply)
	err := c.cc.Invoke(ctx, "/pb.Middle/Setting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *middleClient) CreateSetting(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Text, error) {
	out := new(Text)
	err := c.cc.Invoke(ctx, "/pb.Middle/CreateSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiddleServer is the server API for Middle service.
type MiddleServer interface {
	CreatePage(context.Context, *PageRequest) (*Text, error)
	GetPage(context.Context, *PageRequest) (*PageReply, error)
	Qr(context.Context, *Text) (*Text, error)
	Apps(context.Context, *Text) (*AppReply, error)
	StoreAppOAuth(context.Context, *Text) (*Text, error)
	Credentials(context.Context, *Text) (*Text, error)
	GetCredentials(context.Context, *Text) (*Text, error)
	CreateCredential(context.Context, *Text) (*Text, error)
	Setting(context.Context, *Text) (*SettingReply, error)
	CreateSetting(context.Context, *KV) (*Text, error)
}

// UnimplementedMiddleServer can be embedded to have forward compatible implementations.
type UnimplementedMiddleServer struct {
}

func (*UnimplementedMiddleServer) CreatePage(ctx context.Context, req *PageRequest) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (*UnimplementedMiddleServer) GetPage(ctx context.Context, req *PageRequest) (*PageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (*UnimplementedMiddleServer) Qr(ctx context.Context, req *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Qr not implemented")
}
func (*UnimplementedMiddleServer) Apps(ctx context.Context, req *Text) (*AppReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apps not implemented")
}
func (*UnimplementedMiddleServer) StoreAppOAuth(ctx context.Context, req *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAppOAuth not implemented")
}
func (*UnimplementedMiddleServer) Credentials(ctx context.Context, req *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Credentials not implemented")
}
func (*UnimplementedMiddleServer) GetCredentials(ctx context.Context, req *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentials not implemented")
}
func (*UnimplementedMiddleServer) CreateCredential(ctx context.Context, req *Text) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredential not implemented")
}
func (*UnimplementedMiddleServer) Setting(ctx context.Context, req *Text) (*SettingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setting not implemented")
}
func (*UnimplementedMiddleServer) CreateSetting(ctx context.Context, req *KV) (*Text, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}

func RegisterMiddleServer(s *grpc.Server, srv MiddleServer) {
	s.RegisterService(&_Middle_serviceDesc, srv)
}

func _Middle_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreatePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreatePage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetPage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Qr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Qr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Qr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Qr(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Apps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Apps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Apps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Apps(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_StoreAppOAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).StoreAppOAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/StoreAppOAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).StoreAppOAuth(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Credentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Credentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Credentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Credentials(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_GetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).GetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/GetCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).GetCredentials(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_CreateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreateCredential(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_Setting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).Setting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/Setting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).Setting(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Middle_CreateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiddleServer).CreateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Middle/CreateSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiddleServer).CreateSetting(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

var _Middle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Middle",
	HandlerType: (*MiddleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePage",
			Handler:    _Middle_CreatePage_Handler,
		},
		{
			MethodName: "GetPage",
			Handler:    _Middle_GetPage_Handler,
		},
		{
			MethodName: "Qr",
			Handler:    _Middle_Qr_Handler,
		},
		{
			MethodName: "Apps",
			Handler:    _Middle_Apps_Handler,
		},
		{
			MethodName: "StoreAppOAuth",
			Handler:    _Middle_StoreAppOAuth_Handler,
		},
		{
			MethodName: "Credentials",
			Handler:    _Middle_Credentials_Handler,
		},
		{
			MethodName: "GetCredentials",
			Handler:    _Middle_GetCredentials_Handler,
		},
		{
			MethodName: "CreateCredential",
			Handler:    _Middle_CreateCredential_Handler,
		},
		{
			MethodName: "Setting",
			Handler:    _Middle_Setting_Handler,
		},
		{
			MethodName: "CreateSetting",
			Handler:    _Middle_CreateSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "middle.proto",
}
