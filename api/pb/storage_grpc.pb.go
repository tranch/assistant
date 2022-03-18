// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageSvcClient is the client API for StorageSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageSvcClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageSvc_UploadFileClient, error)
	AbsolutePath(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error)
}

type storageSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageSvcClient(cc grpc.ClientConnInterface) StorageSvcClient {
	return &storageSvcClient{cc}
}

func (c *storageSvcClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageSvc_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StorageSvc_ServiceDesc.Streams[0], "/pb.StorageSvc/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageSvcUploadFileClient{stream}
	return x, nil
}

type StorageSvc_UploadFileClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileReply, error)
	grpc.ClientStream
}

type storageSvcUploadFileClient struct {
	grpc.ClientStream
}

func (x *storageSvcUploadFileClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageSvcUploadFileClient) CloseAndRecv() (*FileReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageSvcClient) AbsolutePath(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.StorageSvc/AbsolutePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageSvcServer is the server API for StorageSvc service.
// All implementations should embed UnimplementedStorageSvcServer
// for forward compatibility
type StorageSvcServer interface {
	UploadFile(StorageSvc_UploadFileServer) error
	AbsolutePath(context.Context, *TextRequest) (*TextReply, error)
}

// UnimplementedStorageSvcServer should be embedded to have forward compatible implementations.
type UnimplementedStorageSvcServer struct {
}

func (UnimplementedStorageSvcServer) UploadFile(StorageSvc_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedStorageSvcServer) AbsolutePath(context.Context, *TextRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbsolutePath not implemented")
}

// UnsafeStorageSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageSvcServer will
// result in compilation errors.
type UnsafeStorageSvcServer interface {
	mustEmbedUnimplementedStorageSvcServer()
}

func RegisterStorageSvcServer(s grpc.ServiceRegistrar, srv StorageSvcServer) {
	s.RegisterService(&StorageSvc_ServiceDesc, srv)
}

func _StorageSvc_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageSvcServer).UploadFile(&storageSvcUploadFileServer{stream})
}

type StorageSvc_UploadFileServer interface {
	SendAndClose(*FileReply) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type storageSvcUploadFileServer struct {
	grpc.ServerStream
}

func (x *storageSvcUploadFileServer) SendAndClose(m *FileReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageSvcUploadFileServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StorageSvc_AbsolutePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageSvcServer).AbsolutePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StorageSvc/AbsolutePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageSvcServer).AbsolutePath(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageSvc_ServiceDesc is the grpc.ServiceDesc for StorageSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StorageSvc",
	HandlerType: (*StorageSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AbsolutePath",
			Handler:    _StorageSvc_AbsolutePath_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _StorageSvc_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}
