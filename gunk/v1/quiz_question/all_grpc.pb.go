// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package qzquestionpb

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

// QuizQuestionServiceClient is the client API for QuizQuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizQuestionServiceClient interface {
	CreateQuizQuestion(ctx context.Context, in *CreateQuizQuestionRequest, opts ...grpc.CallOption) (*CreateQuizQuestionResponse, error)
	EditQuizQuestion(ctx context.Context, in *EditQuizQuestionRequest, opts ...grpc.CallOption) (*EditQuizQuestionResponse, error)
	DeleteQuizQuestion(ctx context.Context, in *DeleteQuizQuestionRequest, opts ...grpc.CallOption) (*DeleteQuizQuestionResponse, error)
	ListQuizQuestion(ctx context.Context, in *ListQuizQuestionRequest, opts ...grpc.CallOption) (*ListQuizQuestionResponse, error)
	UpdateQuizQuestion(ctx context.Context, in *UpdateQuizQuestionRequest, opts ...grpc.CallOption) (*UpdateQuizQuestionResponse, error)
}

type quizQuestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizQuestionServiceClient(cc grpc.ClientConnInterface) QuizQuestionServiceClient {
	return &quizQuestionServiceClient{cc}
}

func (c *quizQuestionServiceClient) CreateQuizQuestion(ctx context.Context, in *CreateQuizQuestionRequest, opts ...grpc.CallOption) (*CreateQuizQuestionResponse, error) {
	out := new(CreateQuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/qzquestionpb.QuizQuestionService/CreateQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizQuestionServiceClient) EditQuizQuestion(ctx context.Context, in *EditQuizQuestionRequest, opts ...grpc.CallOption) (*EditQuizQuestionResponse, error) {
	out := new(EditQuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/qzquestionpb.QuizQuestionService/EditQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizQuestionServiceClient) DeleteQuizQuestion(ctx context.Context, in *DeleteQuizQuestionRequest, opts ...grpc.CallOption) (*DeleteQuizQuestionResponse, error) {
	out := new(DeleteQuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/qzquestionpb.QuizQuestionService/DeleteQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizQuestionServiceClient) ListQuizQuestion(ctx context.Context, in *ListQuizQuestionRequest, opts ...grpc.CallOption) (*ListQuizQuestionResponse, error) {
	out := new(ListQuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/qzquestionpb.QuizQuestionService/ListQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizQuestionServiceClient) UpdateQuizQuestion(ctx context.Context, in *UpdateQuizQuestionRequest, opts ...grpc.CallOption) (*UpdateQuizQuestionResponse, error) {
	out := new(UpdateQuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/qzquestionpb.QuizQuestionService/UpdateQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizQuestionServiceServer is the server API for QuizQuestionService service.
// All implementations must embed UnimplementedQuizQuestionServiceServer
// for forward compatibility
type QuizQuestionServiceServer interface {
	CreateQuizQuestion(context.Context, *CreateQuizQuestionRequest) (*CreateQuizQuestionResponse, error)
	EditQuizQuestion(context.Context, *EditQuizQuestionRequest) (*EditQuizQuestionResponse, error)
	DeleteQuizQuestion(context.Context, *DeleteQuizQuestionRequest) (*DeleteQuizQuestionResponse, error)
	ListQuizQuestion(context.Context, *ListQuizQuestionRequest) (*ListQuizQuestionResponse, error)
	UpdateQuizQuestion(context.Context, *UpdateQuizQuestionRequest) (*UpdateQuizQuestionResponse, error)
	mustEmbedUnimplementedQuizQuestionServiceServer()
}

// UnimplementedQuizQuestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuizQuestionServiceServer struct {
}

func (UnimplementedQuizQuestionServiceServer) CreateQuizQuestion(context.Context, *CreateQuizQuestionRequest) (*CreateQuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuizQuestion not implemented")
}
func (UnimplementedQuizQuestionServiceServer) EditQuizQuestion(context.Context, *EditQuizQuestionRequest) (*EditQuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditQuizQuestion not implemented")
}
func (UnimplementedQuizQuestionServiceServer) DeleteQuizQuestion(context.Context, *DeleteQuizQuestionRequest) (*DeleteQuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuizQuestion not implemented")
}
func (UnimplementedQuizQuestionServiceServer) ListQuizQuestion(context.Context, *ListQuizQuestionRequest) (*ListQuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuizQuestion not implemented")
}
func (UnimplementedQuizQuestionServiceServer) UpdateQuizQuestion(context.Context, *UpdateQuizQuestionRequest) (*UpdateQuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuizQuestion not implemented")
}
func (UnimplementedQuizQuestionServiceServer) mustEmbedUnimplementedQuizQuestionServiceServer() {}

// UnsafeQuizQuestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizQuestionServiceServer will
// result in compilation errors.
type UnsafeQuizQuestionServiceServer interface {
	mustEmbedUnimplementedQuizQuestionServiceServer()
}

func RegisterQuizQuestionServiceServer(s grpc.ServiceRegistrar, srv QuizQuestionServiceServer) {
	s.RegisterService(&QuizQuestionService_ServiceDesc, srv)
}

func _QuizQuestionService_CreateQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuizQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizQuestionServiceServer).CreateQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qzquestionpb.QuizQuestionService/CreateQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizQuestionServiceServer).CreateQuizQuestion(ctx, req.(*CreateQuizQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizQuestionService_EditQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditQuizQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizQuestionServiceServer).EditQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qzquestionpb.QuizQuestionService/EditQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizQuestionServiceServer).EditQuizQuestion(ctx, req.(*EditQuizQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizQuestionService_DeleteQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuizQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizQuestionServiceServer).DeleteQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qzquestionpb.QuizQuestionService/DeleteQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizQuestionServiceServer).DeleteQuizQuestion(ctx, req.(*DeleteQuizQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizQuestionService_ListQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuizQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizQuestionServiceServer).ListQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qzquestionpb.QuizQuestionService/ListQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizQuestionServiceServer).ListQuizQuestion(ctx, req.(*ListQuizQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuizQuestionService_UpdateQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuizQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizQuestionServiceServer).UpdateQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qzquestionpb.QuizQuestionService/UpdateQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizQuestionServiceServer).UpdateQuizQuestion(ctx, req.(*UpdateQuizQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizQuestionService_ServiceDesc is the grpc.ServiceDesc for QuizQuestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizQuestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qzquestionpb.QuizQuestionService",
	HandlerType: (*QuizQuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuizQuestion",
			Handler:    _QuizQuestionService_CreateQuizQuestion_Handler,
		},
		{
			MethodName: "EditQuizQuestion",
			Handler:    _QuizQuestionService_EditQuizQuestion_Handler,
		},
		{
			MethodName: "DeleteQuizQuestion",
			Handler:    _QuizQuestionService_DeleteQuizQuestion_Handler,
		},
		{
			MethodName: "ListQuizQuestion",
			Handler:    _QuizQuestionService_ListQuizQuestion_Handler,
		},
		{
			MethodName: "UpdateQuizQuestion",
			Handler:    _QuizQuestionService_UpdateQuizQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quiz/gunk/v1/quiz_question/all.proto",
}
