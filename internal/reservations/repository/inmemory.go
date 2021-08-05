package repository

import (
	"context"
)

type ReservationInmemoryRepo struct {
	reservations map[string]*ReservationEntity
}

func NewReservationInmemoryRepository() ReservationRepository {
	return &ReservationInmemoryRepo{
		reservations: map[string]*ReservationEntity{},
	}
}

func (r *ReservationInmemoryRepo) CreateReservation(_ context.Context, entity *ReservationEntity) error {
	if _, ok := r.reservations[entity.ReservationId]; ok {
		return ReservationDuplicateError
	}

	r.reservations[entity.ReservationId] = entity

	return nil
}

func (r *ReservationInmemoryRepo) GetReservation(_ context.Context, reservationId string) (*ReservationEntity, error) {
	if _, ok := r.reservations[reservationId]; !ok {
		return nil, ReservationNotFound
	}

	return r.reservations[reservationId], nil
}

func (r *ReservationInmemoryRepo) UpdateReservation(_ context.Context, entity ReservationEntity) error {
	if _, ok := r.reservations[entity.ReservationId]; !ok {
		return ReservationNotFound
	}

	r.reservations[entity.ReservationId] = &entity

	return nil
}

func (r *ReservationInmemoryRepo) GetReservationList(_ context.Context) (array []ReservationEntity) {
	array = []ReservationEntity{}

	for _, r := range r.reservations {
		array = append(array, *r)
	}

	return array
}
