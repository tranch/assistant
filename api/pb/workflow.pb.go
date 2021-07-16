// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: workflow.proto

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

type WorkflowRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowRequest) Reset()         { *m = WorkflowRequest{} }
func (m *WorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowRequest) ProtoMessage()    {}
func (*WorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{0}
}
func (m *WorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowRequest.Unmarshal(m, b)
}
func (m *WorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowRequest.Merge(m, src)
}
func (m *WorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowRequest.Size(m)
}
func (m *WorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowRequest proto.InternalMessageInfo

func (m *WorkflowRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *WorkflowRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type WorkflowReply struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowReply) Reset()         { *m = WorkflowReply{} }
func (m *WorkflowReply) String() string { return proto.CompactTextString(m) }
func (*WorkflowReply) ProtoMessage()    {}
func (*WorkflowReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{1}
}
func (m *WorkflowReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowReply.Unmarshal(m, b)
}
func (m *WorkflowReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowReply.Marshal(b, m, deterministic)
}
func (m *WorkflowReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowReply.Merge(m, src)
}
func (m *WorkflowReply) XXX_Size() int {
	return xxx_messageInfo_WorkflowReply.Size(m)
}
func (m *WorkflowReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowReply.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowReply proto.InternalMessageInfo

func (m *WorkflowReply) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type TriggerRequest struct {
	Trigger              *Trigger `protobuf:"bytes,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{2}
}
func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (m *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(m, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

func (m *TriggerRequest) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type Trigger struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Kind                 string   `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Flag                 string   `protobuf:"bytes,4,opt,name=flag,proto3" json:"flag,omitempty"`
	Secret               string   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	When                 string   `protobuf:"bytes,6,opt,name=when,proto3" json:"when,omitempty"`
	MessageId            int64    `protobuf:"varint,7,opt,name=messageId,proto3" json:"messageId,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MessageText          string   `protobuf:"bytes,9,opt,name=messageText,proto3" json:"messageText,omitempty"`
	Header               string   `protobuf:"bytes,10,opt,name=header,proto3" json:"header,omitempty"`
	Body                 string   `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{3}
}
func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Trigger) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Trigger) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Trigger) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *Trigger) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Trigger) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *Trigger) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Trigger) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Trigger) GetMessageText() string {
	if m != nil {
		return m.MessageText
	}
	return ""
}

func (m *Trigger) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Trigger) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowRequest)(nil), "pb.WorkflowRequest")
	proto.RegisterType((*WorkflowReply)(nil), "pb.WorkflowReply")
	proto.RegisterType((*TriggerRequest)(nil), "pb.TriggerRequest")
	proto.RegisterType((*Trigger)(nil), "pb.Trigger")
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor_892c7f566756b0be) }

var fileDescriptor_892c7f566756b0be = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8f, 0x93, 0x40,
	0x14, 0xc6, 0x5d, 0xba, 0xd2, 0xe5, 0x91, 0x62, 0x1c, 0x13, 0x33, 0xd9, 0x78, 0xd8, 0x60, 0x4c,
	0x3c, 0xd5, 0x64, 0xb5, 0x9a, 0x1e, 0x6b, 0x7b, 0xf1, 0x4a, 0x9b, 0x34, 0xf1, 0x36, 0xc0, 0x2b,
	0x10, 0x90, 0x19, 0x87, 0xa9, 0x2d, 0x7f, 0x83, 0x37, 0xff, 0x62, 0x33, 0x33, 0x90, 0x56, 0x43,
	0x9a, 0xed, 0xed, 0xbd, 0xdf, 0xfb, 0x3e, 0x78, 0x7c, 0x2f, 0x40, 0x70, 0xe0, 0xb2, 0xdc, 0x55,
	0xfc, 0x30, 0x15, 0x92, 0x2b, 0x4e, 0x1c, 0x11, 0xdf, 0x43, 0xcc, 0x1a, 0xb4, 0x7d, 0x38, 0x87,
	0x17, 0xdb, 0x4e, 0x11, 0xe1, 0xcf, 0x3d, 0x36, 0x8a, 0x10, 0xb8, 0x55, 0x78, 0x54, 0xf4, 0xe6,
	0xe1, 0xe6, 0xbd, 0x17, 0x99, 0xda, 0xb0, 0x56, 0x20, 0x75, 0x3a, 0xd6, 0x0a, 0x0c, 0xdf, 0xc2,
	0xe4, 0x64, 0x15, 0x55, 0x3b, 0x64, 0x0c, 0xbf, 0x40, 0xb0, 0x91, 0x45, 0x96, 0xa1, 0xec, 0x1f,
	0xff, 0x0e, 0xc6, 0xca, 0x12, 0x23, 0xf4, 0x1f, 0xfd, 0xa9, 0x88, 0xa7, 0xbd, 0xa8, 0x9f, 0x85,
	0xbf, 0x1d, 0x18, 0x77, 0x90, 0x04, 0xe0, 0x14, 0xa9, 0x51, 0x8f, 0x22, 0xa7, 0x48, 0x87, 0xb6,
	0xd1, 0xac, 0x2c, 0xea, 0x94, 0x8e, 0x2c, 0xd3, 0xb5, 0x66, 0xbb, 0x8a, 0x65, 0xf4, 0xd6, 0x32,
	0x5d, 0x93, 0xd7, 0xe0, 0x36, 0x98, 0x48, 0x54, 0xf4, 0xb9, 0xa1, 0x5d, 0xa7, 0xb5, 0x87, 0x1c,
	0x6b, 0xea, 0x5a, 0xad, 0xae, 0xc9, 0x1b, 0xf0, 0x7e, 0x60, 0xd3, 0xb0, 0x0c, 0xbf, 0xa5, 0x74,
	0x6c, 0x5e, 0x7f, 0x02, 0x7a, 0x9a, 0x48, 0x64, 0x0a, 0xd3, 0x85, 0xa2, 0x77, 0xc6, 0x76, 0x02,
	0xe4, 0x01, 0xfc, 0x4e, 0xba, 0xd1, 0x99, 0x78, 0x66, 0x7e, 0x8e, 0xf4, 0x26, 0x39, 0xb2, 0x14,
	0x25, 0x05, 0xbb, 0x89, 0xed, 0xf4, 0x26, 0x31, 0x4f, 0x5b, 0xea, 0xdb, 0x4d, 0x74, 0xfd, 0xf8,
	0x67, 0x04, 0x7e, 0x1f, 0xf6, 0xfa, 0x57, 0x42, 0x3e, 0x81, 0xbf, 0x6e, 0x6b, 0xc5, 0x8e, 0xcb,
	0x1c, 0x93, 0x92, 0xbc, 0xd2, 0x11, 0xfe, 0x77, 0xc7, 0xfb, 0x40, 0xc3, 0xb5, 0x62, 0x0a, 0xcd,
	0x79, 0xc2, 0x67, 0x64, 0x06, 0x5e, 0xb4, 0xaf, 0x17, 0x89, 0x2a, 0x78, 0x3d, 0xec, 0x79, 0xf9,
	0x2f, 0xb4, 0xb6, 0x39, 0x04, 0x5b, 0x8c, 0x73, 0xce, 0xcb, 0xfe, 0x20, 0xe4, 0xfc, 0x64, 0x97,
	0xac, 0x9f, 0xc1, 0x5f, 0x4a, 0x5e, 0x5f, 0xed, 0x9b, 0xc1, 0x64, 0x69, 0xa2, 0xbc, 0xe4, 0x1c,
	0xfa, 0xc0, 0xc9, 0x0a, 0x2b, 0xbc, 0xde, 0xe6, 0xd9, 0x50, 0x56, 0x3c, 0x79, 0x7a, 0x2e, 0x5f,
	0xef, 0xbe, 0xbb, 0x4c, 0x14, 0x1f, 0x44, 0x1c, 0xbb, 0xe6, 0x67, 0xfa, 0xf8, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x7f, 0x31, 0x20, 0x6e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowSvcClient is the client API for WorkflowSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowSvcClient interface {
	SyntaxCheck(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*StateReply, error)
	RunAction(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowReply, error)
	WebhookTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*WorkflowReply, error)
	CronTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*WorkflowReply, error)
	CreateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*StateReply, error)
	DeleteTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*StateReply, error)
	ActionDoc(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowReply, error)
}

type workflowSvcClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowSvcClient(cc *grpc.ClientConn) WorkflowSvcClient {
	return &workflowSvcClient{cc}
}

func (c *workflowSvcClient) SyntaxCheck(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/SyntaxCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) RunAction(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowReply, error) {
	out := new(WorkflowReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/RunAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) WebhookTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*WorkflowReply, error) {
	out := new(WorkflowReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/WebhookTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) CronTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*WorkflowReply, error) {
	out := new(WorkflowReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/CronTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) CreateTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/CreateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) DeleteTrigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/DeleteTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowSvcClient) ActionDoc(ctx context.Context, in *WorkflowRequest, opts ...grpc.CallOption) (*WorkflowReply, error) {
	out := new(WorkflowReply)
	err := c.cc.Invoke(ctx, "/pb.WorkflowSvc/ActionDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowSvcServer is the server API for WorkflowSvc service.
type WorkflowSvcServer interface {
	SyntaxCheck(context.Context, *WorkflowRequest) (*StateReply, error)
	RunAction(context.Context, *WorkflowRequest) (*WorkflowReply, error)
	WebhookTrigger(context.Context, *TriggerRequest) (*WorkflowReply, error)
	CronTrigger(context.Context, *TriggerRequest) (*WorkflowReply, error)
	CreateTrigger(context.Context, *TriggerRequest) (*StateReply, error)
	DeleteTrigger(context.Context, *TriggerRequest) (*StateReply, error)
	ActionDoc(context.Context, *WorkflowRequest) (*WorkflowReply, error)
}

// UnimplementedWorkflowSvcServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowSvcServer struct {
}

func (*UnimplementedWorkflowSvcServer) SyntaxCheck(ctx context.Context, req *WorkflowRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyntaxCheck not implemented")
}
func (*UnimplementedWorkflowSvcServer) RunAction(ctx context.Context, req *WorkflowRequest) (*WorkflowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunAction not implemented")
}
func (*UnimplementedWorkflowSvcServer) WebhookTrigger(ctx context.Context, req *TriggerRequest) (*WorkflowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebhookTrigger not implemented")
}
func (*UnimplementedWorkflowSvcServer) CronTrigger(ctx context.Context, req *TriggerRequest) (*WorkflowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CronTrigger not implemented")
}
func (*UnimplementedWorkflowSvcServer) CreateTrigger(ctx context.Context, req *TriggerRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrigger not implemented")
}
func (*UnimplementedWorkflowSvcServer) DeleteTrigger(ctx context.Context, req *TriggerRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrigger not implemented")
}
func (*UnimplementedWorkflowSvcServer) ActionDoc(ctx context.Context, req *WorkflowRequest) (*WorkflowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionDoc not implemented")
}

func RegisterWorkflowSvcServer(s *grpc.Server, srv WorkflowSvcServer) {
	s.RegisterService(&_WorkflowSvc_serviceDesc, srv)
}

func _WorkflowSvc_SyntaxCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).SyntaxCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/SyntaxCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).SyntaxCheck(ctx, req.(*WorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_RunAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).RunAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/RunAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).RunAction(ctx, req.(*WorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_WebhookTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).WebhookTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/WebhookTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).WebhookTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_CronTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).CronTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/CronTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).CronTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_CreateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).CreateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/CreateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).CreateTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_DeleteTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).DeleteTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/DeleteTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).DeleteTrigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowSvc_ActionDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowSvcServer).ActionDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkflowSvc/ActionDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowSvcServer).ActionDoc(ctx, req.(*WorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WorkflowSvc",
	HandlerType: (*WorkflowSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyntaxCheck",
			Handler:    _WorkflowSvc_SyntaxCheck_Handler,
		},
		{
			MethodName: "RunAction",
			Handler:    _WorkflowSvc_RunAction_Handler,
		},
		{
			MethodName: "WebhookTrigger",
			Handler:    _WorkflowSvc_WebhookTrigger_Handler,
		},
		{
			MethodName: "CronTrigger",
			Handler:    _WorkflowSvc_CronTrigger_Handler,
		},
		{
			MethodName: "CreateTrigger",
			Handler:    _WorkflowSvc_CreateTrigger_Handler,
		},
		{
			MethodName: "DeleteTrigger",
			Handler:    _WorkflowSvc_DeleteTrigger_Handler,
		},
		{
			MethodName: "ActionDoc",
			Handler:    _WorkflowSvc_ActionDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow.proto",
}
