package domain

import (
	"time"
)

// TimeSlot 時間枠値オブジェクト
type TimeSlot struct {
	Date      time.Time `json:"date"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Capacity  int       `json:"capacity"`
}

// NewTimeSlot 新規時間枠を作成
func NewTimeSlot(date time.Time, startTime, endTime string, capacity int) (*TimeSlot, error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	// 時間文字列をパースして比較
	start, err := time.Parse("15:04", startTime)
	if err != nil {
		return nil, ErrInvalidTimeRange
	}
	end, err := time.Parse("15:04", endTime)
	if err != nil {
		return nil, ErrInvalidTimeRange
	}

	if start.After(end) || start.Equal(end) {
		return nil, ErrInvalidTimeRange
	}

	return &TimeSlot{
		Date:      date,
		StartTime: startTime,
		EndTime:   endTime,
		Capacity:  capacity,
	}, nil
}

// IsAvailable 利用可能かチェック
func (ts *TimeSlot) IsAvailable(reservedCount int) bool {
	return ts.Capacity > reservedCount
}
