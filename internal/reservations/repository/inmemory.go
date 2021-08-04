package repository

import "time"

type ReservationInmemoryRepo struct {
	reservations map[string]*ReservationEntity
}

func NewReservationInmemoryRepo() ReservationRepository {
	return &ReservationInmemoryRepo{
		reservations: map[string]*ReservationEntity{},
	}
}

func (r *ReservationInmemoryRepo) CreateReservation(entity *ReservationEntity) error {
	if _, ok := r.reservations[entity.ReservationId]; ok {
		return ReservationDuplicateError
	}

	r.reservations[entity.ReservationId] = entity

	return nil
}

func (r *ReservationInmemoryRepo) UpdateReturnAt(id string, returnAt time.Time) error {
	if _, ok := r.reservations[id]; !ok {
		return ReservationNotFound
	}

	r.reservations[id].ReturnAt = &returnAt

	return nil
}

func (r *ReservationInmemoryRepo) GetReservationList(filter ReservationFilter) (array []ReservationEntity) {
	array = []ReservationEntity{}

	for _, r := range r.reservations {
		array = append(array, *r)
	}

	return array
}
