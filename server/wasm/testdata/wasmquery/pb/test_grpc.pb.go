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
	TestGet(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (TestService_TestGetClient, error)
	TestGetMany(ctx context.Context, in *TestGetManyRequest, opts ...grpc.CallOption) (TestService_TestGetManyClient, error)
	TestPrefix(ctx context.Context, in *TestPrefixRequest, opts ...grpc.CallOption) (TestService_TestPrefixClient, error)
	TestScan(ctx context.Context, in *TestScanRequest, opts ...grpc.CallOption) (TestService_TestScanClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestGet(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (TestService_TestGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/sf.test.v1.TestService/TestGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestGetClient interface {
	Recv() (*Tuple, error)
	grpc.ClientStream
}

type testServiceTestGetClient struct {
	grpc.ClientStream
}

func (x *testServiceTestGetClient) Recv() (*Tuple, error) {
	m := new(Tuple)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) TestGetMany(ctx context.Context, in *TestGetManyRequest, opts ...grpc.CallOption) (TestService_TestGetManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], "/sf.test.v1.TestService/TestGetMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestGetManyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestGetManyClient interface {
	Recv() (*OptionalTuples, error)
	grpc.ClientStream
}

type testServiceTestGetManyClient struct {
	grpc.ClientStream
}

func (x *testServiceTestGetManyClient) Recv() (*OptionalTuples, error) {
	m := new(OptionalTuples)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) TestPrefix(ctx context.Context, in *TestPrefixRequest, opts ...grpc.CallOption) (TestService_TestPrefixClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], "/sf.test.v1.TestService/TestPrefix", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestPrefixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestPrefixClient interface {
	Recv() (*Tuples, error)
	grpc.ClientStream
}

type testServiceTestPrefixClient struct {
	grpc.ClientStream
}

func (x *testServiceTestPrefixClient) Recv() (*Tuples, error) {
	m := new(Tuples)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) TestScan(ctx context.Context, in *TestScanRequest, opts ...grpc.CallOption) (TestService_TestScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[3], "/sf.test.v1.TestService/TestScan", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestScanClient interface {
	Recv() (*Tuples, error)
	grpc.ClientStream
}

type testServiceTestScanClient struct {
	grpc.ClientStream
}

func (x *testServiceTestScanClient) Recv() (*Tuples, error) {
	m := new(Tuples)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	TestGet(*GetTestRequest, TestService_TestGetServer) error
	TestGetMany(*TestGetManyRequest, TestService_TestGetManyServer) error
	TestPrefix(*TestPrefixRequest, TestService_TestPrefixServer) error
	TestScan(*TestScanRequest, TestService_TestScanServer) error
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) TestGet(*GetTestRequest, TestService_TestGetServer) error {
	return status.Errorf(codes.Unimplemented, "method TestGet not implemented")
}
func (UnimplementedTestServiceServer) TestGetMany(*TestGetManyRequest, TestService_TestGetManyServer) error {
	return status.Errorf(codes.Unimplemented, "method TestGetMany not implemented")
}
func (UnimplementedTestServiceServer) TestPrefix(*TestPrefixRequest, TestService_TestPrefixServer) error {
	return status.Errorf(codes.Unimplemented, "method TestPrefix not implemented")
}
func (UnimplementedTestServiceServer) TestScan(*TestScanRequest, TestService_TestScanServer) error {
	return status.Errorf(codes.Unimplemented, "method TestScan not implemented")
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

func _TestService_TestGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestGet(m, &testServiceTestGetServer{stream})
}

type TestService_TestGetServer interface {
	Send(*Tuple) error
	grpc.ServerStream
}

type testServiceTestGetServer struct {
	grpc.ServerStream
}

func (x *testServiceTestGetServer) Send(m *Tuple) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_TestGetMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestGetManyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestGetMany(m, &testServiceTestGetManyServer{stream})
}

type TestService_TestGetManyServer interface {
	Send(*OptionalTuples) error
	grpc.ServerStream
}

type testServiceTestGetManyServer struct {
	grpc.ServerStream
}

func (x *testServiceTestGetManyServer) Send(m *OptionalTuples) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_TestPrefix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestPrefixRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestPrefix(m, &testServiceTestPrefixServer{stream})
}

type TestService_TestPrefixServer interface {
	Send(*Tuples) error
	grpc.ServerStream
}

type testServiceTestPrefixServer struct {
	grpc.ServerStream
}

func (x *testServiceTestPrefixServer) Send(m *Tuples) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_TestScan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestScan(m, &testServiceTestScanServer{stream})
}

type TestService_TestScanServer interface {
	Send(*Tuples) error
	grpc.ServerStream
}

type testServiceTestScanServer struct {
	grpc.ServerStream
}

func (x *testServiceTestScanServer) Send(m *Tuples) error {
	return x.ServerStream.SendMsg(m)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.test.v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestGet",
			Handler:       _TestService_TestGet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TestGetMany",
			Handler:       _TestService_TestGetMany_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TestPrefix",
			Handler:       _TestService_TestPrefix_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TestScan",
			Handler:       _TestService_TestScan_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
