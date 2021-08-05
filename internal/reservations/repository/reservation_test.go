package repository

import (
	"testing"
	"time"
)

func getDuration(s string) time.Duration {
	d, _ := time.ParseDuration(s)

	return d
}

func TestReservationEntity_GetOverdue(t *testing.T) {
	date := time.Now()

	type fields struct {
		CreatedAt time.Time
		ReturnAt  *time.Time
		Duration  time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "in time",
			fields: fields{
				CreatedAt: time.Now().AddDate(0, 0, -10),
				Duration:  getDuration("1000h"),
			},
			want: false,
		},
		{
			name: "over time",
			fields: fields{
				CreatedAt: time.Now().AddDate(0, 0, -70),
				Duration:  getDuration("1000h"),
			},
			want: true,
		},
		{
			name: "reservation ended",
			fields: fields{
				CreatedAt: time.Now().AddDate(0, 0, -30),
				ReturnAt:  &date,
				Duration:  getDuration("1000h"),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReservationEntity{
				CreatedAt: tt.fields.CreatedAt,
				ReturnAt:  tt.fields.ReturnAt,
				Duration:  tt.fields.Duration,
			}
			if got := r.GetOverdue(); got != tt.want {
				t.Errorf("GetOverdue() = %v, want %v", got, tt.want)
			}
		})
	}
}
