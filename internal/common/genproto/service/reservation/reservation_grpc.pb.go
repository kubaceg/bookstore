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
	CreateReservation(ctx context.Context, in *Reservation, opts ...grpc.CallOption) (*ReservationId, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *Reservation, opts ...grpc.CallOption) (*ReservationId, error) {
	out := new(ReservationId)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	CreateReservation(context.Context, *Reservation) (*ReservationId, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *Reservation) (*ReservationId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
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

func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*Reservation))
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
			MethodName: "CreateReservation",
			Handler:    _ReservationService_CreateReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation/reservation.proto",
}
