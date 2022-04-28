// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: todo.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type TodoRequest struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (m *TodoRequest) Reset()         { *m = TodoRequest{} }
func (m *TodoRequest) String() string { return proto.CompactTextString(m) }
func (*TodoRequest) ProtoMessage()    {}
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}
func (m *TodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoRequest.Unmarshal(m, b)
}
func (m *TodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoRequest.Marshal(b, m, deterministic)
}
func (m *TodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoRequest.Merge(m, src)
}
func (m *TodoRequest) XXX_Size() int {
	return xxx_messageInfo_TodoRequest.Size(m)
}
func (m *TodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TodoRequest proto.InternalMessageInfo

func (m *TodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type TodoReply struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (m *TodoReply) Reset()         { *m = TodoReply{} }
func (m *TodoReply) String() string { return proto.CompactTextString(m) }
func (*TodoReply) ProtoMessage()    {}
func (*TodoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}
func (m *TodoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoReply.Unmarshal(m, b)
}
func (m *TodoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoReply.Marshal(b, m, deterministic)
}
func (m *TodoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoReply.Merge(m, src)
}
func (m *TodoReply) XXX_Size() int {
	return xxx_messageInfo_TodoReply.Size(m)
}
func (m *TodoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoReply.DiscardUnknown(m)
}

var xxx_messageInfo_TodoReply proto.InternalMessageInfo

func (m *TodoReply) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type TodosReply struct {
	Todos []*Todo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (m *TodosReply) Reset()         { *m = TodosReply{} }
func (m *TodosReply) String() string { return proto.CompactTextString(m) }
func (*TodosReply) ProtoMessage()    {}
func (*TodosReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}
func (m *TodosReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodosReply.Unmarshal(m, b)
}
func (m *TodosReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodosReply.Marshal(b, m, deterministic)
}
func (m *TodosReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodosReply.Merge(m, src)
}
func (m *TodosReply) XXX_Size() int {
	return xxx_messageInfo_TodosReply.Size(m)
}
func (m *TodosReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TodosReply.DiscardUnknown(m)
}

var xxx_messageInfo_TodosReply proto.InternalMessageInfo

func (m *TodosReply) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type Todo struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id"`
	// @inject_tag: db:"sequence"
	Sequence int64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty" db:"sequence"`
	// @inject_tag: db:"content"
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty" db:"content"`
	// @inject_tag: db:"category"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty" db:"category"`
	// @inject_tag: db:"remark"
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty" db:"remark"`
	// @inject_tag: db:"priority"
	Priority int64 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty" db:"priority"`
	// @inject_tag: db:"is_remind_at_time"
	IsRemindAtTime bool `protobuf:"varint,8,opt,name=is_remind_at_time,json=isRemindAtTime,proto3" json:"is_remind_at_time,omitempty" db:"is_remind_at_time"`
	// @inject_tag: db:"remind_at"
	RemindAt int64 `protobuf:"varint,9,opt,name=remind_at,json=remindAt,proto3" json:"remind_at,omitempty" db:"remind_at"`
	// @inject_tag: db:"repeat_method"
	RepeatMethod string `protobuf:"bytes,10,opt,name=repeat_method,json=repeatMethod,proto3" json:"repeat_method,omitempty" db:"repeat_method"`
	// @inject_tag: db:"repeat_rule"
	RepeatRule string `protobuf:"bytes,11,opt,name=repeat_rule,json=repeatRule,proto3" json:"repeat_rule,omitempty" db:"repeat_rule"`
	// @inject_tag: db:"repeat_end_at"
	RepeatEndAt int64 `protobuf:"varint,12,opt,name=repeat_end_at,json=repeatEndAt,proto3" json:"repeat_end_at,omitempty" db:"repeat_end_at"`
	// @inject_tag: db:"complete"
	Complete bool `protobuf:"varint,13,opt,name=complete,proto3" json:"complete,omitempty" db:"complete"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Todo) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Todo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Todo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Todo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Todo) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Todo) GetIsRemindAtTime() bool {
	if m != nil {
		return m.IsRemindAtTime
	}
	return false
}

func (m *Todo) GetRemindAt() int64 {
	if m != nil {
		return m.RemindAt
	}
	return 0
}

func (m *Todo) GetRepeatMethod() string {
	if m != nil {
		return m.RepeatMethod
	}
	return ""
}

func (m *Todo) GetRepeatRule() string {
	if m != nil {
		return m.RepeatRule
	}
	return ""
}

func (m *Todo) GetRepeatEndAt() int64 {
	if m != nil {
		return m.RepeatEndAt
	}
	return 0
}

func (m *Todo) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *Todo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Todo) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*TodoRequest)(nil), "pb.TodoRequest")
	proto.RegisterType((*TodoReply)(nil), "pb.TodoReply")
	proto.RegisterType((*TodosReply)(nil), "pb.TodosReply")
	proto.RegisterType((*Todo)(nil), "pb.Todo")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x93, 0x34, 0x3f, 0x93, 0x9f, 0x8a, 0x15, 0x82, 0x55, 0x28, 0x26, 0x32, 0x97, 0x54,
	0x85, 0x44, 0xb4, 0x4f, 0x10, 0x0a, 0xaa, 0x38, 0x70, 0x71, 0xcb, 0x85, 0x4b, 0xe4, 0xc4, 0xa3,
	0xb0, 0x22, 0xce, 0x9a, 0xf5, 0x18, 0x29, 0x6f, 0xc1, 0x2b, 0x71, 0xe3, 0xd8, 0x23, 0x47, 0xaa,
	0xbc, 0x08, 0xda, 0x59, 0xc7, 0x15, 0x12, 0x88, 0xdc, 0xf6, 0xfb, 0x9b, 0x59, 0xcf, 0xac, 0x01,
	0x48, 0x27, 0x7a, 0x92, 0x19, 0x4d, 0x5a, 0xd4, 0xb2, 0xc5, 0x10, 0x16, 0x71, 0x8e, 0x0e, 0x0f,
	0x1f, 0xae, 0xf4, 0x4a, 0xf3, 0x71, 0x6a, 0x4f, 0x8e, 0x0d, 0xcf, 0xa0, 0x7b, 0xa3, 0x13, 0x1d,
	0xe1, 0x97, 0x02, 0x73, 0x12, 0x27, 0xd0, 0xb0, 0x25, 0xa4, 0x3f, 0xf2, 0xc7, 0xdd, 0xf3, 0xf6,
	0x24, 0x5b, 0x4c, 0x58, 0x66, 0x36, 0x3c, 0x85, 0x8e, 0x33, 0x67, 0xeb, 0xed, 0x7f, 0xac, 0x2f,
	0x00, 0x2c, 0xca, 0x9d, 0x37, 0x80, 0x23, 0xcb, 0xe6, 0xd2, 0x1f, 0xd5, 0xff, 0x30, 0x3b, 0x3a,
	0xfc, 0x5e, 0x87, 0x86, 0xc5, 0x62, 0x00, 0x35, 0x95, 0x70, 0xc9, 0x7a, 0x54, 0x53, 0x89, 0x78,
	0x0c, 0xad, 0x22, 0x47, 0x33, 0x57, 0x89, 0xac, 0x31, 0xd9, 0xb4, 0xf0, 0x5d, 0x22, 0x86, 0xd0,
	0xce, 0xed, 0x9d, 0x37, 0x4b, 0x94, 0x75, 0x56, 0x2a, 0x2c, 0x24, 0xb4, 0x96, 0x7a, 0x43, 0xb8,
	0x21, 0xd9, 0x18, 0xf9, 0xe3, 0x4e, 0xb4, 0x87, 0x36, 0xb5, 0x8c, 0x09, 0x57, 0xda, 0x6c, 0xe5,
	0x11, 0x4b, 0x15, 0x16, 0x8f, 0xa0, 0x69, 0x30, 0x8d, 0xcd, 0x67, 0xd9, 0x64, 0xa5, 0x44, 0x36,
	0x93, 0x19, 0xa5, 0x8d, 0xa2, 0xad, 0x6c, 0xb9, 0x4e, 0x7b, 0x2c, 0x4e, 0xe1, 0x81, 0xca, 0xe7,
	0x06, 0x53, 0xb5, 0x49, 0xe6, 0x31, 0xcd, 0x49, 0xa5, 0x28, 0xdb, 0x23, 0x7f, 0xdc, 0x8e, 0x06,
	0x2a, 0x8f, 0x98, 0x9f, 0xd1, 0x8d, 0x4a, 0x51, 0x3c, 0x81, 0x4e, 0xe5, 0x93, 0x1d, 0x57, 0xc7,
	0x94, 0x06, 0xf1, 0x1c, 0xfa, 0x06, 0x33, 0x8c, 0x69, 0x9e, 0x22, 0x7d, 0xd2, 0x89, 0x04, 0xbe,
	0x42, 0xcf, 0x91, 0xef, 0x99, 0x13, 0xcf, 0xa0, 0x5b, 0x9a, 0x4c, 0xb1, 0x46, 0xd9, 0x65, 0x0b,
	0x38, 0x2a, 0x2a, 0xd6, 0x28, 0xc2, 0xaa, 0x0a, 0xba, 0x36, 0x3d, 0x6e, 0x53, 0xa6, 0xde, 0x72,
	0x27, 0x3b, 0x01, 0x9d, 0x66, 0x6b, 0x24, 0x94, 0x7d, 0xbe, 0x68, 0x85, 0xc5, 0x53, 0x80, 0xa5,
	0xc1, 0x98, 0x90, 0xc3, 0x03, 0x0e, 0x77, 0x4a, 0x66, 0x46, 0x56, 0x2e, 0xb2, 0x64, 0x2f, 0x1f,
	0x3b, 0xb9, 0x64, 0x66, 0x74, 0xfe, 0xab, 0x06, 0x2d, 0xbb, 0xc3, 0xeb, 0xaf, 0x4b, 0x31, 0x05,
	0xb8, 0xe4, 0x1c, 0x2f, 0xf5, 0xb8, 0x5a, 0xb7, 0x7b, 0x65, 0xc3, 0x81, 0x25, 0xae, 0x29, 0x26,
	0xe4, 0xe7, 0x11, 0x7a, 0xe2, 0x0c, 0x5a, 0x57, 0x48, 0x7f, 0x77, 0xf7, 0xef, 0x09, 0x67, 0x7e,
	0x09, 0xed, 0xd2, 0x9c, 0xff, 0xa3, 0xf6, 0xfd, 0xd3, 0x0b, 0x3d, 0x71, 0x01, 0x83, 0x2b, 0x24,
	0xb7, 0x8c, 0x83, 0x43, 0x53, 0x80, 0x37, 0x68, 0xa7, 0x72, 0xe8, 0x17, 0x4c, 0x01, 0x3e, 0xf0,
	0x2c, 0x0e, 0x0d, 0xbc, 0x82, 0xde, 0x65, 0x39, 0xf9, 0x03, 0x23, 0xaf, 0x4f, 0x7e, 0xdc, 0x05,
	0xfe, 0xed, 0x5d, 0xe0, 0x7d, 0xdb, 0x05, 0xde, 0xed, 0x2e, 0xf0, 0x7e, 0xee, 0x02, 0xef, 0x63,
	0x33, 0xce, 0xd4, 0x34, 0x5b, 0x2c, 0x9a, 0xfc, 0x47, 0x5f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0x3e, 0x03, 0xf1, 0x05, 0x04, 0x00, 0x00,
}
