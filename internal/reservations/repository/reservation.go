package repository

import (
	"errors"
	"time"
)

var ReservationDuplicateError = errors.New("reservation with such id already exists")
var ReservationError = errors.New("error during reservation")
var ReservationNotFound = errors.New("reservation not found")

type ReservationEntity struct {
	ReservationId string        `json:"string"`
	UserId        string        `json:"user_id"`
	BookId        string        `json:"book_id"`
	CreatedAt     time.Time     `json:"created_at"`
	ReturnAt      *time.Time    `json:"return_at,omitempty"`
	Duration      time.Duration `json:"duration"`
}

// ReservationFilter TODO implement filters
type ReservationFilter struct{}

type ReservationRepository interface {
	CreateReservation(entity *ReservationEntity) error
	UpdateReturnAt(id string, returnAt time.Time) error
	GetReservationList(filter ReservationFilter) []ReservationEntity
}
