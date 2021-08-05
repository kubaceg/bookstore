package repository

import (
	"context"
	"errors"
	"time"
)

var ReservationDuplicateError = errors.New("reservation with such id already exists")
var ReservationNotFound = errors.New("reservation not found")

type ReservationEntity struct {
	ReservationId string        `json:"string"`
	UserId        string        `json:"user_id"`
	BookId        string        `json:"book_id"`
	CreatedAt     time.Time     `json:"created_at"`
	ReturnAt      *time.Time    `json:"return_at,omitempty"`
	Duration      time.Duration `json:"duration"`
}

func (r ReservationEntity) GetOverdue() bool {
	if r.ReturnAt == nil && time.Now().After(r.CreatedAt.Add(r.Duration)) {
		return true
	}

	return false
}

type ReservationRepository interface {
	CreateReservation(ctx context.Context, entity *ReservationEntity) error
	GetReservation(ctx context.Context, reservationId string) (*ReservationEntity, error)
	UpdateReservation(ctx context.Context, entity ReservationEntity) error
	GetReservationList(ctx context.Context) []ReservationEntity
}
