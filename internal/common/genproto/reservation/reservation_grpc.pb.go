// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reservation

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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	RentBook(ctx context.Context, in *CreateReservation, opts ...grpc.CallOption) (*ReservationId, error)
	ReturnBook(ctx context.Context, in *CreateReservation, opts ...grpc.CallOption) (*ReservationId, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) RentBook(ctx context.Context, in *CreateReservation, opts ...grpc.CallOption) (*ReservationId, error) {
	out := new(ReservationId)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/RentBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ReturnBook(ctx context.Context, in *CreateReservation, opts ...grpc.CallOption) (*ReservationId, error) {
	out := new(ReservationId)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/ReturnBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	RentBook(context.Context, *CreateReservation) (*ReservationId, error)
	ReturnBook(context.Context, *CreateReservation) (*ReservationId, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) RentBook(context.Context, *CreateReservation) (*ReservationId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RentBook not implemented")
}
func (UnimplementedReservationServiceServer) ReturnBook(context.Context, *CreateReservation) (*ReservationId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnBook not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_RentBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).RentBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/RentBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).RentBook(ctx, req.(*CreateReservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ReturnBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ReturnBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/ReturnBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ReturnBook(ctx, req.(*CreateReservation))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RentBook",
			Handler:    _ReservationService_RentBook_Handler,
		},
		{
			MethodName: "ReturnBook",
			Handler:    _ReservationService_ReturnBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation.proto",
}
