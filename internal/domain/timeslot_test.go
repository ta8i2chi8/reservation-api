package domain

import (
	"testing"
	"time"
)

func TestNewTimeSlot(t *testing.T) {
	date := time.Now()

	tests := []struct {
		name      string
		date      time.Time
		startTime string
		endTime   string
		capacity  int
		wantErr   bool
	}{
		{
			name:      "Valid time slot",
			date:      date,
			startTime: "09:00",
			endTime:   "10:00",
			capacity:  10,
			wantErr:   false,
		},
		{
			name:      "Invalid capacity",
			date:      date,
			startTime: "09:00",
			endTime:   "10:00",
			capacity:  0,
			wantErr:   true,
		},
		{
			name:      "Start after end",
			date:      date,
			startTime: "10:00",
			endTime:   "09:00",
			capacity:  10,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, err := NewTimeSlot(tt.date, tt.startTime, tt.endTime, tt.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimeSlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if ts == nil && !tt.wantErr {
				t.Error("NewTimeSlot() returned nil time slot")
			}
		})
	}
}
