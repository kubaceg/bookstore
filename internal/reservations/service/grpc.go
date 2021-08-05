package service

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/kubaceg/bookstore/internal/common/eventbus"
	bookMessage "github.com/kubaceg/bookstore/internal/common/eventbus/message/book"
	"github.com/kubaceg/bookstore/internal/common/eventbus/message/reservations"
	"github.com/kubaceg/bookstore/internal/common/genproto/book"
	"github.com/kubaceg/bookstore/internal/common/genproto/reservation"
	"github.com/kubaceg/bookstore/internal/common/log"
	"github.com/kubaceg/bookstore/internal/common/uuid"
	"github.com/kubaceg/bookstore/internal/reservations/repository"
)

type ReservationsGrpcService struct {
	reservation.UnimplementedReservationServiceServer

	repository    repository.ReservationRepository
	bookClient    book.BookServiceClient
	logger        log.Logger
	uuidGenerator uuid.Generator
	publisher     eventbus.Publisher
}

func NewReservationsGrpcService(r repository.ReservationRepository, b book.BookServiceClient, l log.Logger, u uuid.Generator, p eventbus.Publisher) reservation.ReservationServiceServer {
	return &ReservationsGrpcService{repository: r, bookClient: b, logger: l, uuidGenerator: u, publisher: p}
}

func (r *ReservationsGrpcService) RentBook(ctx context.Context, in *reservation.CreateReservation) (*reservation.ReservationId, error) {
	status, err := r.bookClient.ReserveBook(ctx, &book.BookId{Id: in.BookId})

	if err != nil || (status != nil && !status.State) {
		e := fmt.Errorf("error during rentBook %s", err)
		r.logger.Error(ctx, e)

		return nil, e
	}

	reservationId, err := r.uuidGenerator.Generate()
	if err != nil || (status != nil && !status.State) {
		e := fmt.Errorf("error during rentBook %s", err)
		r.logger.Error(ctx, e)

		return nil, e
	}

	entity := repository.ReservationEntity{
		ReservationId: reservationId,
		UserId:        in.UserId,
		BookId:        in.BookId,
		CreatedAt:     time.Now(),
		ReturnAt:      nil,
		Duration:      in.Duration.AsDuration(),
	}

	err = r.repository.CreateReservation(ctx, &entity)
	if err != nil || (status != nil && !status.State) {
		e := fmt.Errorf("error during rentBook %s", err)
		r.logger.Error(ctx, e)

		event := reservations.Error{BookId: in.BookId}
		_ = r.publisher.Publish(&event)

		return nil, e
	}

	event := reservations.Created{
		ReservationId: reservationId,
		BookId:        in.BookId,
	}

	_ = r.publisher.Publish(&event)

	return &reservation.ReservationId{ReservationId: reservationId}, nil
}

func (r *ReservationsGrpcService) ReturnBook(ctx context.Context, reservationId *reservation.ReservationId) (*reservation.ReservationId, error) {
	entity, err := r.repository.GetReservation(ctx, reservationId.ReservationId)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	entity.ReturnAt = &now

	err = r.repository.UpdateReservation(ctx, *entity)
	if err != nil {
		return nil, err
	}

	event := bookMessage.Returned{
		ReservationId: reservationId.ReservationId,
		BookId:        entity.BookId,
	}

	_ = r.publisher.Publish(&event)

	return reservationId, nil
}

func (r *ReservationsGrpcService) GetReservationList(ctx context.Context, _ *emptypb.Empty) (*reservation.ReservationList, error) {
	elements := r.repository.GetReservationList(ctx)

	array := make([]*reservation.Reservation, len(elements))

	for i, e := range elements {
		res := reservation.Reservation{
			ReservationId: e.ReservationId,
			UserId:        e.UserId,
			BookId:        e.BookId,
			CreatedAt:     &timestamppb.Timestamp{Seconds: e.CreatedAt.Unix()},
			Duration:      &durationpb.Duration{Seconds: int64(e.Duration.Seconds())},
			Overdue:       e.GetOverdue(),
		}

		array[i] = &res
	}

	return &reservation.ReservationList{Reservation: array}, nil
}
