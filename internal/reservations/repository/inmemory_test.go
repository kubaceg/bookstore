package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReservationInmemoryRepo_CreateReservation(t *testing.T) {
	tests := []struct {
		name         string
		reservations map[string]*ReservationEntity
		entity       *ReservationEntity
		wantErr      bool
	}{
		{
			name:         "Add reservation to empty repository",
			reservations: map[string]*ReservationEntity{},
			entity:       &ReservationEntity{ReservationId: "1234"},
			wantErr:      false,
		},
		{
			name:         "Add already existing reservation",
			reservations: map[string]*ReservationEntity{"1234": {ReservationId: "1234"}},
			entity:       &ReservationEntity{ReservationId: "1234"},
			wantErr:      true,
		},
		{
			name:         "Add new reservation",
			reservations: map[string]*ReservationEntity{"1234": {ReservationId: "1234"}},
			entity:       &ReservationEntity{ReservationId: "1235"},
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReservationInmemoryRepo{
				reservations: tt.reservations,
			}
			if err := r.CreateReservation(tt.entity); (err != nil) != tt.wantErr {
				t.Errorf("CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReservationInmemoryRepo_GetReservationList(t *testing.T) {
	elements := []ReservationEntity{
		{
			ReservationId: "1",
		},
		{
			ReservationId: "2",
		},
		{
			ReservationId: "3",
		},
	}

	tests := []struct {
		name         string
		reservations map[string]*ReservationEntity
		filter       ReservationFilter
		want         []ReservationEntity
	}{
		{
			name:         "Get empty array",
			reservations: map[string]*ReservationEntity{},
			filter:       ReservationFilter{},
			want:         []ReservationEntity{},
		},
		{
			name: "Get one element array",
			reservations: map[string]*ReservationEntity{
				"2": &elements[1],
			},
			filter: ReservationFilter{},
			want: []ReservationEntity{
				elements[1],
			},
		},
		{
			name: "Get multiple elements array",
			reservations: map[string]*ReservationEntity{
				"1": &elements[0],
				"2": &elements[1],
				"3": &elements[2],
			},
			filter: ReservationFilter{},
			want:   elements,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReservationInmemoryRepo{
				reservations: tt.reservations,
			}
			got := r.GetReservationList(tt.filter)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestReservationInmemoryRepo_UpdateReturnAt(t *testing.T) {
	date := time.Date(2021, time.August, 04, 21, 0, 0, 0, time.UTC)
	returnDate := time.Date(2021, time.August, 20, 21, 0, 0, 0, time.UTC)
	elements := []ReservationEntity{
		{
			ReservationId: "1",
			UserId:        "",
			BookId:        "",
			CreatedAt:     date,
			ReturnAt:      nil,
			Duration:      30,
		},
		{
			ReservationId: "2",
			UserId:        "",
			BookId:        "",
			CreatedAt:     date,
			ReturnAt:      &returnDate,
			Duration:      30,
		},
	}

	tests := []struct {
		name             string
		reservations     map[string]*ReservationEntity
		id               string
		returnAt         time.Time
		wantReservations map[string]ReservationEntity
		wantErr          error
	}{
		{
			name:             "Reservation not found",
			reservations:     map[string]*ReservationEntity{},
			id:               "1",
			returnAt:         returnDate,
			wantReservations: map[string]ReservationEntity{},
			wantErr:          ReservationNotFound,
		},
		{
			name: "Update reservation time",
			reservations: map[string]*ReservationEntity{
				"1": &elements[0],
				"2": &elements[1],
			},
			id:       "1",
			returnAt: returnDate,
			wantReservations: map[string]ReservationEntity{
				"1": {
					ReservationId: "1",
					UserId:        "",
					BookId:        "",
					CreatedAt:     date,
					ReturnAt:      &returnDate,
					Duration:      30,
				},
				"2": elements[1],
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ReservationInmemoryRepo{
				reservations: tt.reservations,
			}

			err := r.UpdateReturnAt(tt.id, tt.returnAt)

			if err != tt.wantErr {
				t.Errorf("UpdateReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for id, expected := range tt.wantReservations {
				if !reflect.DeepEqual(expected, *r.reservations[id]) {
					t.Errorf("UpdateReservation() expected %v, got %v", expected, *r.reservations[id])
					return
				}
			}
		})
	}
}
