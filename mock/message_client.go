// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/message_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockMessageSvcClient is a mock of MessageSvcClient interface.
type MockMessageSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockMessageSvcClientMockRecorder
}

// MockMessageSvcClientMockRecorder is the mock recorder for MockMessageSvcClient.
type MockMessageSvcClientMockRecorder struct {
	mock *MockMessageSvcClient
}

// NewMockMessageSvcClient creates a new mock instance.
func NewMockMessageSvcClient(ctrl *gomock.Controller) *MockMessageSvcClient {
	mock := &MockMessageSvcClient{ctrl: ctrl}
	mock.recorder = &MockMessageSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageSvcClient) EXPECT() *MockMessageSvcClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageSvcClient) Create(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.MessageReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*pb.MessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMessageSvcClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageSvcClient)(nil).Create), varargs...)
}

// CreateActionMessage mocks base method.
func (m *MockMessageSvcClient) CreateActionMessage(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateActionMessage", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActionMessage indicates an expected call of CreateActionMessage.
func (mr *MockMessageSvcClientMockRecorder) CreateActionMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActionMessage", reflect.TypeOf((*MockMessageSvcClient)(nil).CreateActionMessage), varargs...)
}

// Delete mocks base method.
func (m *MockMessageSvcClient) Delete(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageSvcClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageSvcClient)(nil).Delete), varargs...)
}

// DeleteWorkflowMessage mocks base method.
func (m *MockMessageSvcClient) DeleteWorkflowMessage(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWorkflowMessage", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWorkflowMessage indicates an expected call of DeleteWorkflowMessage.
func (mr *MockMessageSvcClientMockRecorder) DeleteWorkflowMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowMessage", reflect.TypeOf((*MockMessageSvcClient)(nil).DeleteWorkflowMessage), varargs...)
}

// Get mocks base method.
func (m *MockMessageSvcClient) Get(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.GetMessageReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*pb.GetMessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMessageSvcClientMockRecorder) Get(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMessageSvcClient)(nil).Get), varargs...)
}

// GetActionMessages mocks base method.
func (m *MockMessageSvcClient) GetActionMessages(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.ActionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActionMessages", varargs...)
	ret0, _ := ret[0].(*pb.ActionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionMessages indicates an expected call of GetActionMessages.
func (mr *MockMessageSvcClientMockRecorder) GetActionMessages(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionMessages", reflect.TypeOf((*MockMessageSvcClient)(nil).GetActionMessages), varargs...)
}

// LastByGroup mocks base method.
func (m *MockMessageSvcClient) LastByGroup(ctx context.Context, in *pb.LastByGroupRequest, opts ...grpc.CallOption) (*pb.LastByGroupReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LastByGroup", varargs...)
	ret0, _ := ret[0].(*pb.LastByGroupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastByGroup indicates an expected call of LastByGroup.
func (mr *MockMessageSvcClientMockRecorder) LastByGroup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastByGroup", reflect.TypeOf((*MockMessageSvcClient)(nil).LastByGroup), varargs...)
}

// LastInbox mocks base method.
func (m *MockMessageSvcClient) LastInbox(ctx context.Context, in *pb.InboxRequest, opts ...grpc.CallOption) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LastInbox", varargs...)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastInbox indicates an expected call of LastInbox.
func (mr *MockMessageSvcClientMockRecorder) LastInbox(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastInbox", reflect.TypeOf((*MockMessageSvcClient)(nil).LastInbox), varargs...)
}

// List mocks base method.
func (m *MockMessageSvcClient) List(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.MessagesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*pb.MessagesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMessageSvcClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMessageSvcClient)(nil).List), varargs...)
}

// ListByGroup mocks base method.
func (m *MockMessageSvcClient) ListByGroup(ctx context.Context, in *pb.GetMessagesRequest, opts ...grpc.CallOption) (*pb.GetMessagesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByGroup", varargs...)
	ret0, _ := ret[0].(*pb.GetMessagesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByGroup indicates an expected call of ListByGroup.
func (mr *MockMessageSvcClientMockRecorder) ListByGroup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByGroup", reflect.TypeOf((*MockMessageSvcClient)(nil).ListByGroup), varargs...)
}

// ListInbox mocks base method.
func (m *MockMessageSvcClient) ListInbox(ctx context.Context, in *pb.InboxRequest, opts ...grpc.CallOption) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInbox", varargs...)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInbox indicates an expected call of ListInbox.
func (mr *MockMessageSvcClientMockRecorder) ListInbox(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInbox", reflect.TypeOf((*MockMessageSvcClient)(nil).ListInbox), varargs...)
}

// MarkReadInbox mocks base method.
func (m *MockMessageSvcClient) MarkReadInbox(ctx context.Context, in *pb.InboxRequest, opts ...grpc.CallOption) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkReadInbox", varargs...)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkReadInbox indicates an expected call of MarkReadInbox.
func (mr *MockMessageSvcClientMockRecorder) MarkReadInbox(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReadInbox", reflect.TypeOf((*MockMessageSvcClient)(nil).MarkReadInbox), varargs...)
}

// MarkSendInbox mocks base method.
func (m *MockMessageSvcClient) MarkSendInbox(ctx context.Context, in *pb.InboxRequest, opts ...grpc.CallOption) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkSendInbox", varargs...)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkSendInbox indicates an expected call of MarkSendInbox.
func (mr *MockMessageSvcClientMockRecorder) MarkSendInbox(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkSendInbox", reflect.TypeOf((*MockMessageSvcClient)(nil).MarkSendInbox), varargs...)
}

// Run mocks base method.
func (m *MockMessageSvcClient) Run(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockMessageSvcClientMockRecorder) Run(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMessageSvcClient)(nil).Run), varargs...)
}

// Save mocks base method.
func (m *MockMessageSvcClient) Save(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.MessageReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(*pb.MessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMessageSvcClientMockRecorder) Save(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMessageSvcClient)(nil).Save), varargs...)
}

// Send mocks base method.
func (m *MockMessageSvcClient) Send(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockMessageSvcClientMockRecorder) Send(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessageSvcClient)(nil).Send), varargs...)
}

// MockMessageSvcServer is a mock of MessageSvcServer interface.
type MockMessageSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageSvcServerMockRecorder
}

// MockMessageSvcServerMockRecorder is the mock recorder for MockMessageSvcServer.
type MockMessageSvcServerMockRecorder struct {
	mock *MockMessageSvcServer
}

// NewMockMessageSvcServer creates a new mock instance.
func NewMockMessageSvcServer(ctrl *gomock.Controller) *MockMessageSvcServer {
	mock := &MockMessageSvcServer{ctrl: ctrl}
	mock.recorder = &MockMessageSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageSvcServer) EXPECT() *MockMessageSvcServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageSvcServer) Create(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.MessageReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*pb.MessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMessageSvcServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageSvcServer)(nil).Create), arg0, arg1)
}

// CreateActionMessage mocks base method.
func (m *MockMessageSvcServer) CreateActionMessage(arg0 context.Context, arg1 *pb.TextRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActionMessage", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActionMessage indicates an expected call of CreateActionMessage.
func (mr *MockMessageSvcServerMockRecorder) CreateActionMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActionMessage", reflect.TypeOf((*MockMessageSvcServer)(nil).CreateActionMessage), arg0, arg1)
}

// Delete mocks base method.
func (m *MockMessageSvcServer) Delete(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageSvcServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageSvcServer)(nil).Delete), arg0, arg1)
}

// DeleteWorkflowMessage mocks base method.
func (m *MockMessageSvcServer) DeleteWorkflowMessage(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowMessage", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWorkflowMessage indicates an expected call of DeleteWorkflowMessage.
func (mr *MockMessageSvcServerMockRecorder) DeleteWorkflowMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowMessage", reflect.TypeOf((*MockMessageSvcServer)(nil).DeleteWorkflowMessage), arg0, arg1)
}

// Get mocks base method.
func (m *MockMessageSvcServer) Get(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.GetMessageReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetMessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMessageSvcServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMessageSvcServer)(nil).Get), arg0, arg1)
}

// GetActionMessages mocks base method.
func (m *MockMessageSvcServer) GetActionMessages(arg0 context.Context, arg1 *pb.TextRequest) (*pb.ActionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionMessages", arg0, arg1)
	ret0, _ := ret[0].(*pb.ActionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionMessages indicates an expected call of GetActionMessages.
func (mr *MockMessageSvcServerMockRecorder) GetActionMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionMessages", reflect.TypeOf((*MockMessageSvcServer)(nil).GetActionMessages), arg0, arg1)
}

// LastByGroup mocks base method.
func (m *MockMessageSvcServer) LastByGroup(arg0 context.Context, arg1 *pb.LastByGroupRequest) (*pb.LastByGroupReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastByGroup", arg0, arg1)
	ret0, _ := ret[0].(*pb.LastByGroupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastByGroup indicates an expected call of LastByGroup.
func (mr *MockMessageSvcServerMockRecorder) LastByGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastByGroup", reflect.TypeOf((*MockMessageSvcServer)(nil).LastByGroup), arg0, arg1)
}

// LastInbox mocks base method.
func (m *MockMessageSvcServer) LastInbox(arg0 context.Context, arg1 *pb.InboxRequest) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastInbox", arg0, arg1)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastInbox indicates an expected call of LastInbox.
func (mr *MockMessageSvcServerMockRecorder) LastInbox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastInbox", reflect.TypeOf((*MockMessageSvcServer)(nil).LastInbox), arg0, arg1)
}

// List mocks base method.
func (m *MockMessageSvcServer) List(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.MessagesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*pb.MessagesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMessageSvcServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMessageSvcServer)(nil).List), arg0, arg1)
}

// ListByGroup mocks base method.
func (m *MockMessageSvcServer) ListByGroup(arg0 context.Context, arg1 *pb.GetMessagesRequest) (*pb.GetMessagesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByGroup", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetMessagesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByGroup indicates an expected call of ListByGroup.
func (mr *MockMessageSvcServerMockRecorder) ListByGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByGroup", reflect.TypeOf((*MockMessageSvcServer)(nil).ListByGroup), arg0, arg1)
}

// ListInbox mocks base method.
func (m *MockMessageSvcServer) ListInbox(arg0 context.Context, arg1 *pb.InboxRequest) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInbox", arg0, arg1)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInbox indicates an expected call of ListInbox.
func (mr *MockMessageSvcServerMockRecorder) ListInbox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInbox", reflect.TypeOf((*MockMessageSvcServer)(nil).ListInbox), arg0, arg1)
}

// MarkReadInbox mocks base method.
func (m *MockMessageSvcServer) MarkReadInbox(arg0 context.Context, arg1 *pb.InboxRequest) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkReadInbox", arg0, arg1)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkReadInbox indicates an expected call of MarkReadInbox.
func (mr *MockMessageSvcServerMockRecorder) MarkReadInbox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkReadInbox", reflect.TypeOf((*MockMessageSvcServer)(nil).MarkReadInbox), arg0, arg1)
}

// MarkSendInbox mocks base method.
func (m *MockMessageSvcServer) MarkSendInbox(arg0 context.Context, arg1 *pb.InboxRequest) (*pb.InboxReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkSendInbox", arg0, arg1)
	ret0, _ := ret[0].(*pb.InboxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkSendInbox indicates an expected call of MarkSendInbox.
func (mr *MockMessageSvcServerMockRecorder) MarkSendInbox(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkSendInbox", reflect.TypeOf((*MockMessageSvcServer)(nil).MarkSendInbox), arg0, arg1)
}

// Run mocks base method.
func (m *MockMessageSvcServer) Run(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockMessageSvcServerMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMessageSvcServer)(nil).Run), arg0, arg1)
}

// Save mocks base method.
func (m *MockMessageSvcServer) Save(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.MessageReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*pb.MessageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMessageSvcServerMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMessageSvcServer)(nil).Save), arg0, arg1)
}

// Send mocks base method.
func (m *MockMessageSvcServer) Send(arg0 context.Context, arg1 *pb.MessageRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockMessageSvcServerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMessageSvcServer)(nil).Send), arg0, arg1)
}

// MockUnsafeMessageSvcServer is a mock of UnsafeMessageSvcServer interface.
type MockUnsafeMessageSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeMessageSvcServerMockRecorder
}

// MockUnsafeMessageSvcServerMockRecorder is the mock recorder for MockUnsafeMessageSvcServer.
type MockUnsafeMessageSvcServerMockRecorder struct {
	mock *MockUnsafeMessageSvcServer
}

// NewMockUnsafeMessageSvcServer creates a new mock instance.
func NewMockUnsafeMessageSvcServer(ctrl *gomock.Controller) *MockUnsafeMessageSvcServer {
	mock := &MockUnsafeMessageSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeMessageSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeMessageSvcServer) EXPECT() *MockUnsafeMessageSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedMessageSvcServer mocks base method.
func (m *MockUnsafeMessageSvcServer) mustEmbedUnimplementedMessageSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMessageSvcServer")
}

// mustEmbedUnimplementedMessageSvcServer indicates an expected call of mustEmbedUnimplementedMessageSvcServer.
func (mr *MockUnsafeMessageSvcServerMockRecorder) mustEmbedUnimplementedMessageSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMessageSvcServer", reflect.TypeOf((*MockUnsafeMessageSvcServer)(nil).mustEmbedUnimplementedMessageSvcServer))
}
