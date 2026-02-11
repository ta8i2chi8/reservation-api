package domain

import (
	"testing"
	"time"
)

func TestNewReservation(t *testing.T) {
	date := time.Now()
	ts, _ := NewTimeSlot(date, "09:00", "10:00", 10)

	tests := []struct {
		name     string
		userID   uint
		timeSlot *TimeSlot
		wantErr  bool
	}{
		{
			name:     "Valid reservation",
			userID:   1,
			timeSlot: ts,
			wantErr:  false,
		},
		{
			name:     "Invalid user ID",
			userID:   0,
			timeSlot: ts,
			wantErr:  true,
		},
		{
			name:     "Nil time slot",
			userID:   1,
			timeSlot: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservation, err := NewReservation(tt.userID, tt.timeSlot)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reservation == nil && !tt.wantErr {
				t.Error("NewReservation() returned nil reservation")
			}
		})
	}
}
