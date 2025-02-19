// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: test.proto

package pbtest

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	TestGet(ctx context.Context, in *TestGetRequest, opts ...grpc.CallOption) (*Tuple, error)
	TestGetMany(ctx context.Context, in *TestGetManyRequest, opts ...grpc.CallOption) (*Tuples, error)
	TestPrefix(ctx context.Context, in *TestPrefixRequest, opts ...grpc.CallOption) (*Tuples, error)
	TestScan(ctx context.Context, in *TestScanRequest, opts ...grpc.CallOption) (*Tuples, error)
	TestSleep(ctx context.Context, in *TestSleepRequest, opts ...grpc.CallOption) (*Response, error)
	TestPanic(ctx context.Context, in *TestPanicRequest, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestGet(ctx context.Context, in *TestGetRequest, opts ...grpc.CallOption) (*Tuple, error) {
	out := new(Tuple)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestGetMany(ctx context.Context, in *TestGetManyRequest, opts ...grpc.CallOption) (*Tuples, error) {
	out := new(Tuples)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestGetMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestPrefix(ctx context.Context, in *TestPrefixRequest, opts ...grpc.CallOption) (*Tuples, error) {
	out := new(Tuples)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestScan(ctx context.Context, in *TestScanRequest, opts ...grpc.CallOption) (*Tuples, error) {
	out := new(Tuples)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestSleep(ctx context.Context, in *TestSleepRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestSleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestPanic(ctx context.Context, in *TestPanicRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/sf.test.v1.TestService/TestPanic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	TestGet(context.Context, *TestGetRequest) (*Tuple, error)
	TestGetMany(context.Context, *TestGetManyRequest) (*Tuples, error)
	TestPrefix(context.Context, *TestPrefixRequest) (*Tuples, error)
	TestScan(context.Context, *TestScanRequest) (*Tuples, error)
	TestSleep(context.Context, *TestSleepRequest) (*Response, error)
	TestPanic(context.Context, *TestPanicRequest) (*Response, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) TestGet(context.Context, *TestGetRequest) (*Tuple, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestGet not implemented")
}
func (UnimplementedTestServiceServer) TestGetMany(context.Context, *TestGetManyRequest) (*Tuples, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestGetMany not implemented")
}
func (UnimplementedTestServiceServer) TestPrefix(context.Context, *TestPrefixRequest) (*Tuples, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPrefix not implemented")
}
func (UnimplementedTestServiceServer) TestScan(context.Context, *TestScanRequest) (*Tuples, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestScan not implemented")
}
func (UnimplementedTestServiceServer) TestSleep(context.Context, *TestSleepRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestSleep not implemented")
}
func (UnimplementedTestServiceServer) TestPanic(context.Context, *TestPanicRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPanic not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestGet(ctx, req.(*TestGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestGetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestGetManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestGetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestGetMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestGetMany(ctx, req.(*TestGetManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestPrefix(ctx, req.(*TestPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestScan(ctx, req.(*TestScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestSleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestSleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestSleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestSleep(ctx, req.(*TestSleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestPanic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestPanicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestPanic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.test.v1.TestService/TestPanic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestPanic(ctx, req.(*TestPanicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.test.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestGet",
			Handler:    _TestService_TestGet_Handler,
		},
		{
			MethodName: "TestGetMany",
			Handler:    _TestService_TestGetMany_Handler,
		},
		{
			MethodName: "TestPrefix",
			Handler:    _TestService_TestPrefix_Handler,
		},
		{
			MethodName: "TestScan",
			Handler:    _TestService_TestScan_Handler,
		},
		{
			MethodName: "TestSleep",
			Handler:    _TestService_TestSleep_Handler,
		},
		{
			MethodName: "TestPanic",
			Handler:    _TestService_TestPanic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
