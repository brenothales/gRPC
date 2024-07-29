// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: proto/course/v1/course.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CourseService_CreateCourse_FullMethodName                    = "/pb.CourseService/CreateCourse"
	CourseService_CreateCourseStream_FullMethodName              = "/pb.CourseService/CreateCourseStream"
	CourseService_CreateCourseStreamBidirectional_FullMethodName = "/pb.CourseService/CreateCourseStreamBidirectional"
	CourseService_ListCourses_FullMethodName                     = "/pb.CourseService/ListCourses"
	CourseService_GetCourse_FullMethodName                       = "/pb.CourseService/GetCourse"
	CourseService_ListCoursesStream_FullMethodName               = "/pb.CourseService/ListCoursesStream"
)

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error)
	CreateCourseStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirectionalClient, error)
	ListCourses(ctx context.Context, in *CourseBlank, opts ...grpc.CallOption) (*CourseList, error)
	GetCourse(ctx context.Context, in *CourseGetRequest, opts ...grpc.CallOption) (*Course, error)
	ListCoursesStream(ctx context.Context, in *CourseBlank, opts ...grpc.CallOption) (CourseService_ListCoursesStreamClient, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Course)
	err := c.cc.Invoke(ctx, CourseService_CreateCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) CreateCourseStream(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[0], CourseService_CreateCourseStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamClient{ClientStream: stream}
	return x, nil
}

type CourseService_CreateCourseStreamClient interface {
	Send(*CreateCourseRequest) error
	CloseAndRecv() (*CourseList, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamClient) CloseAndRecv() (*CourseList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CourseList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseServiceClient) CreateCourseStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (CourseService_CreateCourseStreamBidirectionalClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[1], CourseService_CreateCourseStreamBidirectional_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceCreateCourseStreamBidirectionalClient{ClientStream: stream}
	return x, nil
}

type CourseService_CreateCourseStreamBidirectionalClient interface {
	Send(*CreateCourseRequest) error
	Recv() (*Course, error)
	grpc.ClientStream
}

type courseServiceCreateCourseStreamBidirectionalClient struct {
	grpc.ClientStream
}

func (x *courseServiceCreateCourseStreamBidirectionalClient) Send(m *CreateCourseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirectionalClient) Recv() (*Course, error) {
	m := new(Course)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *courseServiceClient) ListCourses(ctx context.Context, in *CourseBlank, opts ...grpc.CallOption) (*CourseList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourseList)
	err := c.cc.Invoke(ctx, CourseService_ListCourses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourse(ctx context.Context, in *CourseGetRequest, opts ...grpc.CallOption) (*Course, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Course)
	err := c.cc.Invoke(ctx, CourseService_GetCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) ListCoursesStream(ctx context.Context, in *CourseBlank, opts ...grpc.CallOption) (CourseService_ListCoursesStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CourseService_ServiceDesc.Streams[2], CourseService_ListCoursesStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &courseServiceListCoursesStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CourseService_ListCoursesStreamClient interface {
	Recv() (*Course, error)
	grpc.ClientStream
}

type courseServiceListCoursesStreamClient struct {
	grpc.ClientStream
}

func (x *courseServiceListCoursesStreamClient) Recv() (*Course, error) {
	m := new(Course)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	CreateCourse(context.Context, *CreateCourseRequest) (*Course, error)
	CreateCourseStream(CourseService_CreateCourseStreamServer) error
	CreateCourseStreamBidirectional(CourseService_CreateCourseStreamBidirectionalServer) error
	ListCourses(context.Context, *CourseBlank) (*CourseList, error)
	GetCourse(context.Context, *CourseGetRequest) (*Course, error)
	ListCoursesStream(*CourseBlank, CourseService_ListCoursesStreamServer) error
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStream(CourseService_CreateCourseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStream not implemented")
}
func (UnimplementedCourseServiceServer) CreateCourseStreamBidirectional(CourseService_CreateCourseStreamBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCourseStreamBidirectional not implemented")
}
func (UnimplementedCourseServiceServer) ListCourses(context.Context, *CourseBlank) (*CourseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedCourseServiceServer) GetCourse(context.Context, *CourseGetRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServiceServer) ListCoursesStream(*CourseBlank, CourseService_ListCoursesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCoursesStream not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseService_CreateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_CreateCourseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStream(&courseServiceCreateCourseStreamServer{ServerStream: stream})
}

type CourseService_CreateCourseStreamServer interface {
	SendAndClose(*CourseList) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamServer) SendAndClose(m *CourseList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CourseService_CreateCourseStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CourseServiceServer).CreateCourseStreamBidirectional(&courseServiceCreateCourseStreamBidirectionalServer{ServerStream: stream})
}

type CourseService_CreateCourseStreamBidirectionalServer interface {
	Send(*Course) error
	Recv() (*CreateCourseRequest, error)
	grpc.ServerStream
}

type courseServiceCreateCourseStreamBidirectionalServer struct {
	grpc.ServerStream
}

func (x *courseServiceCreateCourseStreamBidirectionalServer) Send(m *Course) error {
	return x.ServerStream.SendMsg(m)
}

func (x *courseServiceCreateCourseStreamBidirectionalServer) Recv() (*CreateCourseRequest, error) {
	m := new(CreateCourseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CourseService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseBlank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseService_ListCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).ListCourses(ctx, req.(*CourseBlank))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseService_GetCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourse(ctx, req.(*CourseGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_ListCoursesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CourseBlank)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CourseServiceServer).ListCoursesStream(m, &courseServiceListCoursesStreamServer{ServerStream: stream})
}

type CourseService_ListCoursesStreamServer interface {
	Send(*Course) error
	grpc.ServerStream
}

type courseServiceListCoursesStreamServer struct {
	grpc.ServerStream
}

func (x *courseServiceListCoursesStreamServer) Send(m *Course) error {
	return x.ServerStream.SendMsg(m)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _CourseService_ListCourses_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _CourseService_GetCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCourseStream",
			Handler:       _CourseService_CreateCourseStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateCourseStreamBidirectional",
			Handler:       _CourseService_CreateCourseStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListCoursesStream",
			Handler:       _CourseService_ListCoursesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/course/v1/course.proto",
}