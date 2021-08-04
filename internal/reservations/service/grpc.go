package service

import (
	"context"

	"google.golang.org/grpc"

	"github.com/kubaceg/bookstore/internal/common/genproto/reservation"
	"github.com/kubaceg/bookstore/internal/reservations/repository"
)

type ReservationsGrpcService struct {
	reservation.UnimplementedReservationServiceServer

	repository repository.ReservationRepository
}

func NewReservationsGrpcService(r repository.ReservationRepository) reservation.UnsafeReservationServiceServer {
	return &ReservationsGrpcService{repository: r}
}

func (r ReservationsGrpcService) RentBook(ctx context.Context, in *reservation.CreateReservation, opts ...grpc.CallOption) (*reservation.ReservationId, error) {
	panic("implement me")
}

func (r ReservationsGrpcService) ReturnBook(ctx context.Context, in *reservation.CreateReservation, opts ...grpc.CallOption) (*reservation.ReservationId, error) {
	panic("implement me")
}

func (r ReservationsGrpcService) GetReservationList(ctx context.Context, in *reservation.ReservationFilter, opts ...grpc.CallOption) (*reservation.ReservationList, error) {
	panic("implement me")
}
