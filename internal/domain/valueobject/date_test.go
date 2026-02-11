package valueobject

import (
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	tests := []struct {
		name    string
		date    time.Time
		wantErr bool
	}{
		{
			name:    "Valid date",
			date:    time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "Zero time",
			date:    time.Time{},
			wantErr: true,
		},
		{
			name:    "Past date",
			date:    time.Date(2020, 6, 1, 12, 0, 0, 0, time.UTC),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, err := NewDate(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if date == nil && !tt.wantErr {
				t.Error("NewDate() returned nil date")
			}
			if date != nil && !date.Value().Equal(tt.date) {
				t.Errorf("NewDate() = %v, want %v", date.Value(), tt.date)
			}
		})
	}
}

func TestDateEquals(t *testing.T) {
	date1, _ := NewDate(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	date2, _ := NewDate(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	date3, _ := NewDate(time.Date(2026, 3, 20, 0, 0, 0, 0, time.UTC))

	if !date1.Equals(date2) {
		t.Error("date1 should equal date2")
	}

	if date1.Equals(date3) {
		t.Error("date1 should not equal date3")
	}

	if date1.Equals(nil) {
		t.Error("date1 should not equal nil")
	}
}

func TestDateIsBefore(t *testing.T) {
	date1, _ := NewDate(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	date2, _ := NewDate(time.Date(2026, 3, 20, 0, 0, 0, 0, time.UTC))

	if !date1.IsBefore(date2) {
		t.Error("date1 should be before date2")
	}

	if date2.IsBefore(date1) {
		t.Error("date2 should not be before date1")
	}

	if date1.IsBefore(nil) {
		t.Error("date1.IsBefore(nil) should return false")
	}
}

func TestDateIsAfter(t *testing.T) {
	date1, _ := NewDate(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	date2, _ := NewDate(time.Date(2026, 3, 20, 0, 0, 0, 0, time.UTC))

	if !date2.IsAfter(date1) {
		t.Error("date2 should be after date1")
	}

	if date1.IsAfter(date2) {
		t.Error("date1 should not be after date2")
	}

	if date1.IsAfter(nil) {
		t.Error("date1.IsAfter(nil) should return false")
	}
}

func TestDateString(t *testing.T) {
	date, _ := NewDate(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))

	expected := "2026-01-15"
	if date.String() != expected {
		t.Errorf("Date.String() = %v, want %v", date.String(), expected)
	}
}
